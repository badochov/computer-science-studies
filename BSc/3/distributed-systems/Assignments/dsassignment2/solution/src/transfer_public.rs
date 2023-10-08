use crate::SectorVec;
use crate::{ClientRegisterCommandContent, RegisterCommand, SystemRegisterCommandContent};

use std::io::Error;
use tokio::io::{AsyncRead, AsyncWrite};

use crate::transfer::*;

pub async fn deserialize_register_command(
    data: &mut (dyn AsyncRead + Send + Unpin),
    hmac_system_key: &[u8; 64],
    hmac_client_key: &[u8; 32],
) -> Result<(RegisterCommand, bool), Error> {
    loop {
        match deserialize_message(data, hmac_system_key, hmac_client_key).await? {
            (Message::Normal(msg), hmac_ok) => return Ok((msg, hmac_ok)),
        };
    }
}

pub async fn serialize_register_command(
    cmd: &RegisterCommand,
    writer: &mut (dyn AsyncWrite + Send + Unpin),
    hmac_key: &[u8],
) -> Result<(), Error> {
    match cmd {
        RegisterCommand::Client(cmd) => {
            let header = &RecvHeader {
                padding: [0; 2],
                magic_value: 0,
                msg_type: get_client_message_type(cmd),
            };
            match &cmd.content {
                ClientRegisterCommandContent::Read => {
                    serialize_message(
                        writer,
                        &header,
                        &client_message::ClientIn {
                            header: cmd.header,
                            content: client_message::ClientInRead {},
                        },
                        hmac_key,
                    )
                    .await
                }
                ClientRegisterCommandContent::Write { data } => {
                    serialize_message(
                        writer,
                        &header,
                        &client_message::ClientIn {
                            header: cmd.header,
                            content: client_message::ClientInWrite {
                                data: to_u8_4096(data),
                            },
                        },
                        hmac_key,
                    )
                    .await
                }
            }
        }
        RegisterCommand::System(cmd) => {
            let header = &RecvHeader {
                padding: [0; 2],
                magic_value: cmd.header.process_identifier,
                msg_type: get_system_message_type(cmd),
            };

            match &cmd.content {
                SystemRegisterCommandContent::Ack => {
                    serialize_message(
                        writer,
                        &header,
                        &system_message::SystemIn {
                            uuid: *cmd.header.msg_ident.as_bytes(),
                            read_ident: cmd.header.read_ident,
                            sector_index: cmd.header.sector_idx,
                            content: system_message::SystemInAck {},
                        },
                        hmac_key,
                    )
                    .await
                }
                SystemRegisterCommandContent::ReadProc => {
                    serialize_message(
                        writer,
                        &header,
                        &system_message::SystemIn {
                            uuid: *cmd.header.msg_ident.as_bytes(),
                            read_ident: cmd.header.read_ident,
                            sector_index: cmd.header.sector_idx,
                            content: system_message::SystemInReadProc {},
                        },
                        hmac_key,
                    )
                    .await
                }
                SystemRegisterCommandContent::WriteProc {
                    timestamp,
                    write_rank,
                    data_to_write,
                } => {
                    serialize_message(
                        writer,
                        &&header,
                        &system_message::SystemIn {
                            uuid: *cmd.header.msg_ident.as_bytes(),
                            read_ident: cmd.header.read_ident,
                            sector_index: cmd.header.sector_idx,
                            content: system_message::SystemInWriteProc {
                                timestamp: *timestamp,
                                write_rank: *write_rank,
                                padding: [0; 7],
                                data_to_write: to_u8_4096(data_to_write),
                            },
                        },
                        hmac_key,
                    )
                    .await
                }
                SystemRegisterCommandContent::Value {
                    timestamp,
                    write_rank,
                    sector_data,
                } => {
                    serialize_message(
                        writer,
                        &header,
                        &system_message::SystemIn {
                            uuid: *cmd.header.msg_ident.as_bytes(),
                            read_ident: cmd.header.read_ident,
                            sector_index: cmd.header.sector_idx,
                            content: system_message::SystemInValue {
                                timestamp: *timestamp,
                                write_rank: *write_rank,
                                padding: [0; 7],
                                data: to_u8_4096(sector_data),
                            },
                        },
                        hmac_key,
                    )
                    .await
                }
            }
        }
    }
}

fn to_u8_4096(v: &SectorVec) -> [u8; 4096] {
    let mut res = [0; 4096];
    for i in 0..4096 {
        res[i] = v.0[i];
    }

    res
}
