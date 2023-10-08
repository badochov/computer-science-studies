use std::convert::TryInto;
use std::io::Error;

use bincode::Options;
use hmac::{Hmac, Mac, NewMac};
use serde::{Deserialize, Serialize};
use sha2::Sha256;
use tokio::io::{AsyncRead, AsyncReadExt, AsyncWrite, AsyncWriteExt};
use uuid::Uuid;

use crate::SystemCommandHeader;
use crate::{
    ClientCommandHeader, ClientRegisterCommand, ClientRegisterCommandContent, RegisterCommand,
    SectorVec, SystemRegisterCommand, SystemRegisterCommandContent, MAGIC_NUMBER,
};

pub async fn serialize_message<H: ?Sized, C: ?Sized>(
    writer: &mut (dyn AsyncWrite + Send + Unpin),
    header: &H,
    content: &C,
    hmac_key: &[u8],
) -> Result<(), Error>
where
    H: serde::Serialize,
    C: serde::Serialize,
{
    let mut msg = Vec::new();
    let header_bytes = serialize(header);
    let content_bytes = serialize(content);

    msg.reserve(MAGIC_NUMBER.len() + header_bytes.len() + content_bytes.len() + 32);

    msg.extend(&MAGIC_NUMBER);
    msg.extend(header_bytes);
    msg.extend(content_bytes);

    let mut hmac = HmacSha256::new_from_slice(hmac_key).unwrap();
    hmac.update(&msg);
    msg.extend(hmac.finalize().into_bytes());

    writer.write_all(&msg).await.map(|_| Ok(()))?
}

pub async fn deserialize_message(
    data: &mut (dyn AsyncRead + Send + Unpin),
    hmac_system_key: &[u8; 64],
    hmac_client_key: &[u8; 32],
) -> Result<(Message, bool), Error> {
    loop {
        let reader = MagicBytesReader::new(data);
        let header_reader = reader.read_to_magic_bytes().await?;

        let content_reader = header_reader.read_header().await?;
        let verifier: HmacVerifier;
        let msg: Message;

        match client_message::get(content_reader, hmac_client_key).await {
            GetResult::Correct(cmd_res) => {
                let (_verifier, cmd) = cmd_res?;
                verifier = _verifier;
                msg = Message::Normal(cmd);
            }
            GetResult::Wrong(r) => match system_message::get(r, hmac_system_key).await {
                GetResult::Correct(cmd_res) => {
                    let (_verifier, cmd) = cmd_res?;
                    verifier = _verifier;
                    msg = Message::Normal(cmd);
                }
                GetResult::Wrong(reader) => {
                    log::debug!("Encountered unknown message {:?}", reader.get_header());
                    continue;
                }
            },
        }

        let hmac_ok = verifier.verify_hmac().await?;

        return Ok((msg, hmac_ok));
    }
}

struct MagicBytesReader<'a> {
    reader: &'a mut (dyn AsyncRead + Send + Unpin),
}

struct HeaderReader<'a> {
    reader: &'a mut (dyn AsyncRead + Send + Unpin),
}

struct HmaclessContentReader<'a> {
    reader: &'a mut (dyn AsyncRead + Send + Unpin),
    header_bytes: Vec<u8>,
}

struct ContentReader<'a> {
    reader: &'a mut (dyn AsyncRead + Send + Unpin),
    hmac: HmacSha256,
}

struct HmacVerifier<'a> {
    reader: &'a mut (dyn AsyncRead + Send + Unpin),
    hmac: HmacSha256,
}

impl<'a> MagicBytesReader<'a> {
    pub fn new(reader: &'a mut (dyn AsyncRead + Send + Unpin)) -> Self {
        Self { reader }
    }

    pub async fn read_to_magic_bytes(self) -> Result<HeaderReader<'a>, Error> {
        let mut matching = 0;
        let mut buf = [0u8];
        while matching != MAGIC_NUMBER.len() {
            self.reader.read_exact(&mut buf).await?;

            if buf[0] == MAGIC_NUMBER[matching] {
                matching += 1;
            } else {
                matching = 0;
            }
        }

        Ok(HeaderReader {
            reader: self.reader,
        })
    }
}

impl<'a> HeaderReader<'a> {
    pub async fn read_header(self) -> Result<HmaclessContentReader<'a>, Error> {
        let mut header_bytes = Vec::new();
        header_bytes.resize(RecvHeader::size(), 0);

        self.reader.read_exact(&mut header_bytes).await?;

        Ok(HmaclessContentReader {
            reader: self.reader,
            header_bytes,
        })
    }
}

impl<'a> HmaclessContentReader<'a> {
    pub fn get_header(&self) -> RecvHeader {
        deserialize(&self.header_bytes)
    }

    pub fn set_hmac_key(self, hmac_key: &[u8]) -> ContentReader<'a> {
        let mut hmac = HmacSha256::new_from_slice(hmac_key).unwrap();
        hmac.update(&MAGIC_NUMBER);
        hmac.update(&self.header_bytes);

        ContentReader {
            reader: self.reader,
            hmac,
        }
    }
}

impl<'a> ContentReader<'a> {
    pub async fn read_raw_of_type<T>(mut self) -> Result<(HmacVerifier<'a>, Vec<u8>), Error>
    where
        T: Size,
    {
        let mut buf = Vec::new();
        buf.resize(T::size(), 0);

        self.reader.read_exact(&mut buf).await?;
        self.hmac.update(&buf);

        Ok((
            HmacVerifier {
                reader: self.reader,
                hmac: self.hmac,
            },
            buf,
        ))
    }
}

impl<'a> HmacVerifier<'a> {
    pub async fn verify_hmac(self) -> Result<bool, Error> {
        let mut hmac_tag = [0u8; 32];

        self.reader.read_exact(&mut hmac_tag).await?;

        Ok(self.hmac.verify(&hmac_tag).is_ok())
    }
}

#[derive(Debug)]
pub enum Message {
    Normal(RegisterCommand),
}

pub trait Size {
    fn size() -> usize {
        return 0;
    }
}

enum GetResult<'a> {
    Correct(Result<(HmacVerifier<'a>, RegisterCommand), Error>),
    Wrong(HmaclessContentReader<'a>),
}

pub mod client_message {
    use serde::{Deserialize, Serialize};
    use serde_big_array::*;

    pub use crate::transfer::*;

    big_array! { BigArray; }

    #[derive(Serialize, Deserialize)]
    pub struct ClientIn<T>
    where
        T: ToCmdContent,
        T: Size,
    {
        pub header: ClientCommandHeader,
        pub content: T,
    }

    pub trait ToCmdContent {
        fn to_cmd_content(self) -> ClientRegisterCommandContent;
    }

    impl<T> ClientIn<T>
    where
        T: ToCmdContent,
        T: Size,
    {
        fn to_cmd(self) -> ClientRegisterCommand {
            ClientRegisterCommand {
                header: self.header,
                content: self.content.to_cmd_content(),
            }
        }
    }

    impl<T> Size for ClientIn<T>
    where
        T: ToCmdContent,
        T: Size,
    {
        fn size() -> usize {
            return 8 + 8 + T::size();
        }
    }

    #[derive(Serialize, Deserialize)]
    pub struct ClientInWrite {
        #[serde(with = "BigArray")]
        pub data: [u8; 4096],
    }

    impl ToCmdContent for ClientInWrite {
        fn to_cmd_content(self) -> ClientRegisterCommandContent {
            return ClientRegisterCommandContent::Write {
                data: SectorVec(self.data.to_vec()),
            };
        }
    }

    impl Size for ClientInWrite {
        fn size() -> usize {
            return 4096;
        }
    }

    #[derive(Serialize, Deserialize)]
    pub struct ClientInRead {}

    impl ToCmdContent for ClientInRead {
        fn to_cmd_content(self) -> ClientRegisterCommandContent {
            return ClientRegisterCommandContent::Read;
        }
    }

    impl Size for ClientInRead {}

    macro_rules! cmd_of_type {
        ($content_reader:expr, $cmd_type:ty) => {{
            let (verifier, bytes) = $content_reader.read_raw_of_type::<$cmd_type>().await?;
            let msg: $cmd_type = deserialize(&bytes);

            Ok((verifier, RegisterCommand::Client(msg.to_cmd())))
        }};
    }

    pub(super) async fn get<'a>(
        reader: HmaclessContentReader<'a>,
        hmac_key: &[u8; 32],
    ) -> GetResult<'a> {
        let recv_header = reader.get_header();
        match recv_header.msg_type.try_into() {
            Ok(t) => GetResult::Correct(get_helper(reader, t, hmac_key).await),
            Err(_) => GetResult::Wrong(reader),
        }
    }

    async fn get_helper<'a>(
        reader: HmaclessContentReader<'a>,
        msg_type: ClientMsgType,
        hmac_key: &[u8; 32],
    ) -> Result<(HmacVerifier<'a>, RegisterCommand), Error> {
        let content_reader = reader.set_hmac_key(hmac_key);
        match msg_type {
            ClientMsgType::Write => {
                cmd_of_type!(content_reader, ClientIn<ClientInWrite>)
            }
            ClientMsgType::Read => {
                cmd_of_type!(content_reader, ClientIn<ClientInRead>)
            }
        }
    }
}

pub mod system_message {
    use serde::{Deserialize, Serialize};
    use serde_big_array::*;

    pub use crate::transfer::*;

    big_array! { BigArray; }

    #[derive(Serialize, Deserialize)]
    pub struct SystemIn<T>
    where
        T: ToCmdContent,
        T: Size,
    {
        pub uuid: [u8; 16],
        pub read_ident: u64,
        pub sector_index: u64,
        pub content: T,
    }

    pub trait ToCmdContent {
        fn to_cmd_content(self) -> SystemRegisterCommandContent;
    }

    impl<T> SystemIn<T>
    where
        T: ToCmdContent,
        T: Size,
    {
        fn to_cmd(self, process_identifier: u8) -> SystemRegisterCommand {
            SystemRegisterCommand {
                header: SystemCommandHeader {
                    process_identifier,
                    msg_ident: Uuid::from_slice(&self.uuid).unwrap(),
                    read_ident: self.read_ident,
                    sector_idx: self.sector_index,
                },
                content: self.content.to_cmd_content(),
            }
        }
    }

    impl<T> Size for SystemIn<T>
    where
        T: ToCmdContent,
        T: Size,
    {
        fn size() -> usize {
            return 16 + 8 + 8 + T::size();
        }
    }

    #[derive(Serialize, Deserialize)]
    pub struct SystemInValue {
        pub timestamp: u64,
        pub padding: [u8; 7],
        pub write_rank: u8,
        #[serde(with = "BigArray")]
        pub data: [u8; 4096],
    }

    impl ToCmdContent for SystemInValue {
        fn to_cmd_content(self) -> SystemRegisterCommandContent {
            return SystemRegisterCommandContent::Value {
                timestamp: self.timestamp,
                write_rank: self.write_rank,
                sector_data: SectorVec(self.data.to_vec()),
            };
        }
    }

    impl Size for SystemInValue {
        fn size() -> usize {
            return 8 + 7 + 1 + 4096;
        }
    }

    #[derive(Serialize, Deserialize)]
    pub struct SystemInWriteProc {
        pub timestamp: u64,
        pub padding: [u8; 7],
        pub write_rank: u8,
        #[serde(with = "BigArray")]
        pub data_to_write: [u8; 4096],
    }

    impl ToCmdContent for SystemInWriteProc {
        fn to_cmd_content(self) -> SystemRegisterCommandContent {
            return SystemRegisterCommandContent::WriteProc {
                timestamp: self.timestamp,
                write_rank: self.write_rank,
                data_to_write: SectorVec(self.data_to_write.to_vec()),
            };
        }
    }

    impl Size for SystemInWriteProc {
        fn size() -> usize {
            return 8 + 7 + 1 + 4096;
        }
    }

    #[derive(Serialize, Deserialize)]
    pub struct SystemInReadProc {}

    impl ToCmdContent for SystemInReadProc {
        fn to_cmd_content(self) -> SystemRegisterCommandContent {
            return SystemRegisterCommandContent::ReadProc;
        }
    }

    impl Size for SystemInReadProc {}

    #[derive(Serialize, Deserialize)]
    pub struct SystemInAck {}

    impl ToCmdContent for SystemInAck {
        fn to_cmd_content(self) -> SystemRegisterCommandContent {
            return SystemRegisterCommandContent::Ack;
        }
    }

    impl Size for SystemInAck {}

    macro_rules! cmd_of_type {
        ($content_reader:expr, $cmd_type:ty, $process_identifier: expr) => {{
            let (verifier, bytes) = $content_reader.read_raw_of_type::<$cmd_type>().await?;
            let msg: $cmd_type = deserialize(&bytes);

            Ok((
                verifier,
                RegisterCommand::System(msg.to_cmd($process_identifier)),
            ))
        }};
    }

    pub(super) async fn get<'a>(
        reader: HmaclessContentReader<'a>,
        hmac_key: &[u8; 64],
    ) -> GetResult<'a> {
        let header = reader.get_header();
        match header.msg_type.try_into() {
            Ok(t) => GetResult::Correct(get_helper(reader, t, header.magic_value, hmac_key).await),
            Err(_) => GetResult::Wrong(reader),
        }
    }

    async fn get_helper<'a>(
        reader: HmaclessContentReader<'a>,
        msg_type: SystemMsgType,
        process_identifier: u8,
        hmac_key: &[u8; 64],
    ) -> Result<(HmacVerifier<'a>, RegisterCommand), Error> {
        let content_reader = reader.set_hmac_key(hmac_key);
        match msg_type {
            SystemMsgType::ReadProc => {
                cmd_of_type!(
                    content_reader,
                    SystemIn<SystemInReadProc>,
                    process_identifier
                )
            }
            SystemMsgType::Value => {
                cmd_of_type!(content_reader, SystemIn<SystemInValue>, process_identifier)
            }
            SystemMsgType::WriteProc => {
                cmd_of_type!(
                    content_reader,
                    SystemIn<SystemInWriteProc>,
                    process_identifier
                )
            }
            SystemMsgType::Ack => {
                cmd_of_type!(content_reader, SystemIn<SystemInAck>, process_identifier)
            }
        }
    }
}

pub const RESPONSE_SHIFT: u8 = 0x40;

#[derive(Serialize, Deserialize)]
struct ResponseHeader {
    padding: [u8; 2],
    type_specific: u8,
    msg_type: u8,
}

#[derive(Clone)]
#[repr(u8)]
pub enum ClientMsgType {
    Read = 0x01,
    Write = 0x02,
}

impl TryInto<ClientMsgType> for u8 {
    fn try_into(self) -> Result<ClientMsgType, <Self as TryInto<ClientMsgType>>::Error> {
        for msg in [ClientMsgType::Read, ClientMsgType::Write] {
            if msg.clone() as u8 == self {
                return Ok(msg);
            }
        }

        Err(())
    }
    type Error = ();
}

#[repr(u8)]
#[derive(Clone)]
pub enum SystemMsgAckType {
    ReadProc = 0x43,
    Value = 0x44,
    WriteProc = 0x45,
    Ack = 0x46,
}

impl TryInto<SystemMsgAckType> for u8 {
    fn try_into(self) -> Result<SystemMsgAckType, <Self as TryInto<SystemMsgAckType>>::Error> {
        for msg in [
            SystemMsgAckType::ReadProc,
            SystemMsgAckType::Value,
            SystemMsgAckType::WriteProc,
            SystemMsgAckType::Ack,
        ] {
            if msg.clone() as u8 == self {
                return Ok(msg);
            }
        }

        Err(())
    }
    type Error = ();
}

#[repr(u8)]
#[derive(Clone, Debug)]
pub enum SystemMsgType {
    ReadProc = 0x03,
    Value = 0x04,
    WriteProc = 0x05,
    Ack = 0x06,
}

impl TryInto<SystemMsgType> for u8 {
    fn try_into(self) -> Result<SystemMsgType, <Self as TryInto<SystemMsgType>>::Error> {
        for msg in [
            SystemMsgType::ReadProc,
            SystemMsgType::Value,
            SystemMsgType::WriteProc,
            SystemMsgType::Ack,
        ] {
            if msg.clone() as u8 == self {
                return Ok(msg);
            }
        }

        Err(())
    }
    type Error = ();
}

#[derive(Deserialize, Serialize, Debug)]
pub struct RecvHeader {
    pub padding: [u8; 2],
    pub magic_value: u8,
    pub msg_type: u8,
}

impl Size for RecvHeader {
    fn size() -> usize {
        return 2 + 1 + 1;
    }
}

fn deserialize<'a, T>(bytes: &'a [u8]) -> T
where
    T: serde::de::Deserialize<'a>,
{
    bincode::options()
        .with_big_endian()
        .with_fixint_encoding()
        .allow_trailing_bytes()
        .deserialize(bytes)
        .unwrap()
}

type HmacSha256 = Hmac<Sha256>;

fn serialize<T: ?Sized>(data: &T) -> Vec<u8>
where
    T: serde::Serialize,
{
    bincode::options()
        .with_big_endian()
        .with_fixint_encoding()
        .allow_trailing_bytes()
        .serialize(data)
        .unwrap()
}

pub fn get_client_message_type(msg: &ClientRegisterCommand) -> u8 {
    match &msg.content {
        ClientRegisterCommandContent::Read => 0x01,
        ClientRegisterCommandContent::Write { .. } => 0x02,
    }
}

pub fn get_system_message_type(msg: &SystemRegisterCommand) -> u8 {
    match &msg.content {
        SystemRegisterCommandContent::ReadProc => 0x03,
        SystemRegisterCommandContent::Value { .. } => 0x04,
        SystemRegisterCommandContent::WriteProc { .. } => 0x05,
        SystemRegisterCommandContent::Ack => 0x06,
    }
}
