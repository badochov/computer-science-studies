use hmac::{Hmac, Mac, NewMac};
use rustls::RootCertStore;
use sha2::Sha256;
use std::convert::TryInto;
use std::io::{Read, Write};
use std::sync::Arc;

// Create a type alias:
type HmacSha256 = Hmac<Sha256>;
// You can add here other imports from std or crates listed in Cargo.toml.

pub struct SecureClient<L: Read + Write> {
    // Add here any fields you need.
    client: rustls::StreamOwned<rustls::ClientConnection, L>,
    hmac_key: HmacSha256,
}

pub struct SecureServer<L: Read + Write> {
    // Add here any fields you need.
    server: rustls::StreamOwned<rustls::ServerConnection, L>,
    hmac_key: HmacSha256,
}

struct Message {
    length: u32,
    content: Vec<u8>,
    hmac_tag: [u8; 32],
}

impl Message {
    fn serialize(&self) -> Vec<u8> {
        let mut bytes = Vec::new();
        bytes.extend_from_slice(&Message::serialize_length(self.length));
        bytes.extend(&self.content);
        bytes.extend_from_slice(&self.hmac_tag);

        return bytes;
    }

    fn deserialize_length(bytes: [u8; 4]) -> u32 {
        return ((bytes[0] as u32) << 24)
            + ((bytes[1] as u32) << 16)
            + ((bytes[2] as u32) << 8)
            + ((bytes[3] as u32) << 0);
    }

    fn serialize_length(length: u32) -> [u8; 4] {
        let mut bytes = [0u8; 4];
        for i in 0..4 {
            bytes[i] = ((length >> ((3 - i) * 8)) & 0xff) as u8;
        }
        return bytes;
    }
}

impl<L: Read + Write> SecureClient<L> {
    /// Creates a new instance of SecureClient.
    ///
    /// SecureClient communicates with SecureServer via `link`.
    /// The messages include a HMAC tag calculated using `hmac_key`.
    /// A certificate of SecureServer is signed by `root_cert`.
    pub fn new(link: L, hmac_key: &[u8], root_cert: &str) -> Self {
        // Create an empty store for root certificates:
        let mut root_store = RootCertStore::empty();

        // Add to the store the root certificate of the server:
        root_store
            .add_parsable_certificates(&rustls_pemfile::certs(&mut root_cert.as_bytes()).unwrap());

        // Create a TLS configuration for the client:
        let client_config = rustls::ClientConfig::builder()
            .with_safe_defaults()
            .with_root_certificates(root_store)
            .with_no_client_auth();

        // Create a TLS connection using the configuration prepared above.
        // "localhost" is the name of the remote server:
        let connection =
            rustls::ClientConnection::new(Arc::new(client_config), "localhost".try_into().unwrap())
                .unwrap();

        // Wrap the TCP stream in TLS:
        let conn = rustls::StreamOwned::new(connection, link);

        SecureClient {
            client: conn,
            hmac_key: HmacSha256::new_from_slice(hmac_key).unwrap(),
        }
    }

    /// Sends the data to the server. The sent message follows the
    /// format specified in the description of the assignment.
    pub fn send_msg(&mut self, data: Vec<u8>) {
        println!("Send");
        let msg = self.data_to_message(data);
        self.client.write_all(msg.serialize().as_slice()).unwrap();
    }

    fn data_to_message(&self, data: Vec<u8>) -> Message {
        let tag = self.calc_hmac_tag(&data);
        return Message {
            length: data.len().try_into().unwrap(),
            content: data,
            hmac_tag: tag,
        };
    }

    fn calc_hmac_tag(&self, data: &Vec<u8>) -> [u8; 32] {
        let mut hmac = self.hmac_key.clone();
        hmac.update(data.as_slice());

        return hmac.finalize().into_bytes().try_into().unwrap();
    }
}

impl<L: Read + Write> SecureServer<L> {
    /// Creates a new instance of SecureServer.
    ///
    /// SecureServer receives messages from SecureClients via `link`.
    /// HMAC tags of the messages are verified against `hmac_key`.
    /// The private key of the SecureServer's certificate is `server_private_key`,
    /// and the full certificate chain is `server_full_chain`.
    pub fn new(
        link: L,
        hmac_key: &[u8],
        server_private_key: &str,
        server_full_chain: &str,
    ) -> Self {
        // Load the certificate chain for the server:
        let certs = rustls_pemfile::certs(&mut server_full_chain.as_bytes())
            .unwrap()
            .iter()
            .map(|v| rustls::Certificate(v.clone()))
            .collect();

        // Load the private key for the server (for simplicity, we assume there is
        // provided one valid key and it is a RSA private key):
        let private_key = rustls::PrivateKey(
            rustls_pemfile::rsa_private_keys(&mut server_private_key.as_bytes())
                .unwrap()
                .first()
                .unwrap()
                .to_vec(),
        );

        // Create a TLS configuration for the server:
        let server_config = rustls::ServerConfig::builder()
            .with_safe_defaults()
            .with_no_client_auth()
            .with_single_cert(certs, private_key)
            .unwrap();

        // Create a TLS connection using the configuration prepared above:
        let connection = rustls::ServerConnection::new(Arc::new(server_config)).unwrap();

        // Wrap the TCP stream in TLS:
        let conn = rustls::StreamOwned::new(connection, link);

        SecureServer {
            server: conn,
            hmac_key: HmacSha256::new_from_slice(hmac_key).unwrap(),
        }
    }

    /// Receives the next incoming message and returns the message's content
    /// (i.e., without the message size and without the HMAC tag) if the
    /// message's HMAC tag is correct. Otherwise returns `SecureServerError`.
    pub fn recv_message(&mut self) -> Result<Vec<u8>, SecureServerError> {
        println!("Recv");
        let length = self.read_content_length().unwrap();
        let content = self.read_content(length).unwrap();
        let hmac_tag = self.read_hmac_tag().unwrap();
        println!("Vcer");
        if !self.verify_hmac_tag(&content, hmac_tag) {
            return Err(SecureServerError::InvalidHmac);
        }
        return Ok(content);
    }

    fn read_content_length(&mut self) -> Result<u32, ()> {
        let mut length_buffer = [0u8; 4];
        match self.server.read_exact(&mut length_buffer) {
            Ok(_) => Ok(Message::deserialize_length(length_buffer)),
            Err(_) => Err(()),
        }
    }

    fn read_hmac_tag(&mut self) -> Result<[u8; 32], ()> {
        let mut hmac_tag = [0u8; 32];
        match self.server.read_exact(&mut hmac_tag) {
            Ok(_) => Ok(hmac_tag),
            Err(_) => Err(()),
        }
    }

    fn read_content(&mut self, length: u32) -> Result<Vec<u8>, ()> {
        let mut length_buffer = Vec::new();
        length_buffer.resize(length.try_into().unwrap(), 0);
        match self.server.read_exact(&mut length_buffer) {
            Ok(_) => Ok(length_buffer),
            Err(_) => Err(()),
        }
    }

    fn verify_hmac_tag(&self, content: &Vec<u8>, hmac_tag: [u8; 32]) -> bool {
        let mut mac = self.hmac_key.clone();
        // Calculate MAC for the data (one can provide it in multiple portions):
        mac.update(content.as_slice());

        // Verify the tag:
        mac.verify(&hmac_tag).is_ok()
    }
}

#[derive(Copy, Clone, Eq, PartialEq, Hash, Debug)]
pub enum SecureServerError {
    /// The HMAC tag of a message is invalid.
    InvalidHmac,
}

// You can add any private types, structs, consts, functions, methods, etc., you need.
