use crate::OperationComplete;
use crate::OperationReturn;
use crate::ReadReturn;
use crate::SectorVec;
use crate::StatusCode;
use serde::Serialize;
use serde_big_array::*;
use std::io::Error;
use tokio::io::AsyncWrite;

big_array! { BigArray; }

use crate::transfer::*;

pub async fn serialize_client_response(
    op: OperationComplete,
    writer: &mut (dyn AsyncWrite + Send + Unpin),
    hmac_key: &[u8; 32],
) -> Result<(), Error> {
    match op.op_return {
        OperationReturn::Read(ReadReturn { read_data }) => {
            let header = RecvHeader {
                padding: [0; 2],
                magic_value: op.status_code as u8,
                msg_type: ClientMsgType::Read as u8 + RESPONSE_SHIFT,
            };

            match op.status_code {
                StatusCode::Ok => {
                    serialize_message(
                        writer,
                        &header,
                        &ReadResponse {
                            request_number: op.request_identifier,
                            content: to_u8_4096(&read_data.unwrap()),
                        },
                        hmac_key,
                    )
                    .await
                }
                _ => serialize_message(writer, &header, &op.request_identifier, hmac_key).await,
            }
        }
        OperationReturn::Write => {
            serialize_message(
                writer,
                &RecvHeader {
                    padding: [0; 2],
                    magic_value: op.status_code as u8,
                    msg_type: ClientMsgType::Write as u8 + RESPONSE_SHIFT,
                },
                &op.request_identifier,
                hmac_key,
            )
            .await
        }
    }
}

#[derive(Serialize)]
struct ReadResponse {
    request_number: u64,
    #[serde(with = "BigArray")]
    content: [u8; 4096],
}

fn to_u8_4096(v: &SectorVec) -> [u8; 4096] {
    let mut res = [0; 4096];
    for i in 0..4096 {
        res[i] = v.0[i];
    }

    res
}
