use std::time::{Duration, SystemTime};

use async_channel::unbounded;
use executor::System;
use ntest::timeout;
use std::collections::HashSet;
use uuid::Uuid;

use assignment_3_solution::*;
use assignment_3_test_utils::*;

#[tokio::test]
#[timeout(1000)]
async fn added_server_participates_in_consensus() {
    // given
    let mut system = System::new().await;
    let leader_id = Uuid::new_v4();
    let follower_id = Uuid::new_v4();
    let added_id = Uuid::new_v4();
    let processes = vec![leader_id, follower_id];
    let sender = ExecutorSender::default();
    let first_log_entry_timestamp = SystemTime::now();
    let leader = Raft::new(
        &mut system,
        make_config(leader_id, Duration::from_millis(100), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    let follower = Raft::new(
        &mut system,
        make_config(follower_id, Duration::from_millis(300), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(leader_id, Box::new(leader.clone())).await;
    sender.insert(follower_id, Box::new(follower.clone())).await;
    let (result_sender, result_receiver) = unbounded();
    tokio::time::sleep(Duration::from_millis(150)).await;

    // when
    let client_id = register_client(&leader, &result_sender, &result_receiver).await;

    let added = Raft::new(
        &mut system,
        make_config(added_id, Duration::from_millis(300), processes),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(added_id, Box::new(added.clone())).await;

    leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::AddServer {
                new_server: added_id,
            },
        })
        .await;
    let add_result = result_receiver.recv().await.unwrap();
    sender.break_link(follower_id, leader_id).await;

    leader
        .send(ClientRequest {
            reply_to: result_sender,
            content: ClientRequestContent::Command {
                command: vec![1, 2, 3, 4],
                client_id,
                sequence_num: 0,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;
    let command_result = result_receiver.recv().await.unwrap();

    // then
    assert_eq!(
        add_result,
        ClientRequestResponse::AddServerResponse(AddServerResponseArgs {
            new_server: added_id,
            content: AddServerResponseContent::ServerAdded,
        })
    );
    assert_eq!(
        command_result,
        ClientRequestResponse::CommandResponse(CommandResponseArgs {
            client_id,
            sequence_num: 0,
            content: CommandResponseContent::CommandApplied {
                output: vec![1, 2, 3, 4]
            }
        })
    );

    system.shutdown().await;
}

#[tokio::test]
#[timeout(1000)]
async fn added_server_is_caught_up_with_snapshot() {
    // given
    let mut system = System::new().await;
    let leader_id = Uuid::new_v4();
    let follower_id = Uuid::new_v4();
    let added_id = Uuid::new_v4();
    let processes = vec![leader_id, follower_id];
    let sender = ExecutorSender::default();
    let first_log_entry_timestamp = SystemTime::now();
    let leader = Raft::new(
        &mut system,
        make_config(leader_id, Duration::from_millis(100), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    let follower = Raft::new(
        &mut system,
        make_config(follower_id, Duration::from_millis(300), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(leader_id, Box::new(leader.clone())).await;
    sender.insert(follower_id, Box::new(follower.clone())).await;
    let (result_sender, result_receiver) = unbounded();
    tokio::time::sleep(Duration::from_millis(150)).await;

    // when
    let client_id = register_client(&leader, &result_sender, &result_receiver).await;

    let added = Raft::new(
        &mut system,
        make_config(added_id, Duration::from_millis(300), processes),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(added_id, Box::new(added.clone())).await;
    leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::Command {
                command: vec![0],
                client_id,
                sequence_num: 0,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;
    result_receiver.recv().await.unwrap();
    leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::Command {
                command: vec![1],
                client_id,
                sequence_num: 1,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;
    result_receiver.recv().await.unwrap();
    leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::Snapshot,
        })
        .await;
    assert_eq!(
        result_receiver.recv().await.unwrap(),
        ClientRequestResponse::SnapshotResponse(SnapshotResponseArgs {
            content: SnapshotResponseContent::SnapshotCreated {
                last_included_index: 3
            }
        })
    );

    leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::AddServer {
                new_server: added_id,
            },
        })
        .await;
    let add_result = result_receiver.recv().await.unwrap();
    sender.break_link(follower_id, leader_id).await;

    leader
        .send(ClientRequest {
            reply_to: result_sender,
            content: ClientRequestContent::Command {
                command: vec![1, 2, 3, 4],
                client_id,
                sequence_num: 2,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;
    let command_result = result_receiver.recv().await.unwrap();

    // then
    assert_eq!(
        add_result,
        ClientRequestResponse::AddServerResponse(AddServerResponseArgs {
            new_server: added_id,
            content: AddServerResponseContent::ServerAdded,
        })
    );
    assert_eq!(
        command_result,
        ClientRequestResponse::CommandResponse(CommandResponseArgs {
            client_id,
            sequence_num: 2,
            content: CommandResponseContent::CommandApplied {
                output: vec![1, 2, 3, 4]
            }
        })
    );

    system.shutdown().await;
}

#[tokio::test]
#[timeout(1000)]
async fn many_added_and_then_removed() {
    let n = 6;
    // given
    let mut system = System::new().await;
    let leader_id = Uuid::new_v4();
    let follower_id = Uuid::new_v4();

    let processes = vec![leader_id, follower_id];
    let sender = ExecutorSender::default();
    let first_log_entry_timestamp = SystemTime::now();
    let leader = Raft::new(
        &mut system,
        make_config(leader_id, Duration::from_millis(100), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    let follower = Raft::new(
        &mut system,
        make_config(follower_id, Duration::from_millis(300), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(leader_id, Box::new(leader.clone())).await;
    sender.insert(follower_id, Box::new(follower.clone())).await;
    let (result_sender, result_receiver) = unbounded();
    tokio::time::sleep(Duration::from_millis(150)).await;

    // when
    let client_id = register_client(&leader, &result_sender, &result_receiver).await;

    for i in 0..n {
        let id = Uuid::from_u128(i as u128);
        let added = Raft::new(
            &mut system,
            make_config(id, Duration::from_millis(300), processes.clone()),
            first_log_entry_timestamp,
            Box::new(IdentityMachine),
            Box::new(RamStorage::default()),
            Box::new(sender.clone()),
        )
        .await;
        sender.insert(id, Box::new(added.clone())).await;
    }

    for i in 0..n {
        let id = Uuid::from_u128(i as u128);
        leader
            .send(ClientRequest {
                reply_to: result_sender.clone(),
                content: ClientRequestContent::AddServer { new_server: id },
            })
            .await;
    }
    let mut r = HashSet::new();
    let mut l = HashSet::new();
    for i in 0..n {
        l.insert(Uuid::from_u128(i as u128));

        let add_result = result_receiver.recv().await.unwrap();
        match add_result {
            ClientRequestResponse::AddServerResponse(AddServerResponseArgs {
                new_server,
                content: AddServerResponseContent::ServerAdded,
            }) => {
                r.insert(new_server);
            }
            _ => {
                assert!(false);
            }
        };
    }
    assert!(l == r);

    leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::Command {
                command: vec![1, 2, 3, 4],
                client_id,
                sequence_num: 0,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;
    let command_result = result_receiver.recv().await.unwrap();

    // then

    assert_eq!(
        command_result,
        ClientRequestResponse::CommandResponse(CommandResponseArgs {
            client_id,
            sequence_num: 0,
            content: CommandResponseContent::CommandApplied {
                output: vec![1, 2, 3, 4]
            }
        })
    );

    //removing all that were added
    for i in 0..n {
        let id = Uuid::from_u128(i as u128);
        leader
            .send(ClientRequest {
                reply_to: result_sender.clone(),
                content: ClientRequestContent::RemoveServer { old_server: id },
            })
            .await;
    }
    let mut r2 = HashSet::new();
    let mut l2 = HashSet::new();
    for i in 0..n {
        l2.insert(Uuid::from_u128(i as u128));

        let add_result = result_receiver.recv().await.unwrap();
        match add_result {
            ClientRequestResponse::RemoveServerResponse(RemoveServerResponseArgs {
                old_server,
                content: RemoveServerResponseContent::ServerRemoved,
            }) => {
                r2.insert(old_server);
            }
            _ => {
                assert!(false);
            }
        };
    }
    assert!(l2 == r2);

    //can still make progress

    leader
        .send(ClientRequest {
            reply_to: result_sender,
            content: ClientRequestContent::Command {
                command: vec![5, 6, 7, 8],
                client_id,
                sequence_num: 2,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;
    let command_result2 = result_receiver.recv().await.unwrap();

    // then

    assert_eq!(
        command_result2,
        ClientRequestResponse::CommandResponse(CommandResponseArgs {
            client_id,
            sequence_num: 2,
            content: CommandResponseContent::CommandApplied {
                output: vec![5, 6, 7, 8]
            }
        })
    );

    system.shutdown().await;
}

#[tokio::test]
#[timeout(1000)]
async fn server_is_caught_up() {
    let n: u8 = 6;
    // given
    let mut system = System::new().await;
    let leader_id = Uuid::new_v4();
    let follower_id = Uuid::new_v4();
    let added_id = Uuid::new_v4();
    let processes = vec![leader_id, follower_id];
    let sender = ExecutorSender::default();
    let first_log_entry_timestamp = SystemTime::now();
    let leader = Raft::new(
        &mut system,
        make_config(leader_id, Duration::from_millis(100), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    let follower = Raft::new(
        &mut system,
        make_config(follower_id, Duration::from_millis(300), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(leader_id, Box::new(leader.clone())).await;
    sender.insert(follower_id, Box::new(follower.clone())).await;
    let (result_sender, result_receiver) = unbounded();
    tokio::time::sleep(Duration::from_millis(150)).await;

    // when
    let client_id = register_client(&leader, &result_sender, &result_receiver).await;

    let added = Raft::new(
        &mut system,
        make_config(added_id, Duration::from_millis(300), processes),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(added_id, Box::new(added.clone())).await;
    for i in 0..n {
        leader
            .send(ClientRequest {
                reply_to: result_sender.clone(),
                content: ClientRequestContent::Command {
                    command: vec![i],
                    client_id,
                    sequence_num: i as u64,
                    lowest_sequence_num_without_response: 0,
                },
            })
            .await;
    }
    for _ in 0..n {
        result_receiver.recv().await.unwrap();
    }
    //wait for commands to take effect
    tokio::time::sleep(Duration::from_millis(150)).await;

    leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::AddServer {
                new_server: added_id,
            },
        })
        .await;
    //wait for server to catch up
    tokio::time::sleep(Duration::from_millis(100)).await;
    let add_result = result_receiver.recv().await.unwrap();
    sender.break_link(follower_id, leader_id).await;

    leader
        .send(ClientRequest {
            reply_to: result_sender,
            content: ClientRequestContent::Command {
                command: vec![n],
                client_id,
                sequence_num: n as u64,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;
    let command_result = result_receiver.recv().await.unwrap();

    // then
    assert_eq!(
        add_result,
        ClientRequestResponse::AddServerResponse(AddServerResponseArgs {
            new_server: added_id,
            content: AddServerResponseContent::ServerAdded,
        })
    );
    assert_eq!(
        command_result,
        ClientRequestResponse::CommandResponse(CommandResponseArgs {
            client_id,
            sequence_num: n as u64,
            content: CommandResponseContent::CommandApplied { output: vec![n] }
        })
    );

    system.shutdown().await;
}

#[tokio::test]
#[timeout(1000)]
async fn removed_server_does_not_participate_in_consensus() {
    // given
    let mut system = System::new().await;
    let leader_id = Uuid::new_v4();
    let follower_id = Uuid::new_v4();
    let removed_id = Uuid::new_v4();
    let processes = vec![leader_id, follower_id, removed_id];
    let sender = ExecutorSender::default();
    let first_log_entry_timestamp = SystemTime::now();
    let leader = Raft::new(
        &mut system,
        make_config(leader_id, Duration::from_millis(100), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    let follower = Raft::new(
        &mut system,
        make_config(follower_id, Duration::from_millis(300), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    let removed = Raft::new(
        &mut system,
        make_config(removed_id, Duration::from_millis(300), processes.clone()),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(leader_id, Box::new(leader.clone())).await;
    sender.insert(follower_id, Box::new(follower.clone())).await;
    sender.insert(removed_id, Box::new(removed.clone())).await;
    let (result_sender, result_receiver) = unbounded();
    tokio::time::sleep(Duration::from_millis(150)).await;

    // when
    let client_id = register_client(&leader, &result_sender, &result_receiver).await;

    leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::RemoveServer {
                old_server: removed_id,
            },
        })
        .await;
    let remove_result = result_receiver.recv().await.unwrap();
    for process_id in processes {
        sender.break_link(removed_id, process_id).await;
    }
    sender.break_link(follower_id, leader_id).await;

    leader
        .send(ClientRequest {
            reply_to: result_sender,
            content: ClientRequestContent::Command {
                command: vec![1, 2, 3, 4],
                client_id,
                sequence_num: 0,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;
    tokio::time::sleep(Duration::from_millis(700)).await;

    // then
    assert_eq!(
        remove_result,
        ClientRequestResponse::RemoveServerResponse(RemoveServerResponseArgs {
            old_server: removed_id,
            content: RemoveServerResponseContent::ServerRemoved,
        })
    );
    assert!(result_receiver.is_empty());

    //system.shutdown().await;
}

#[tokio::test]
#[timeout(500)]
async fn messages_outside_cluster_configuration_are_accepted() {
    // given
    let mut system = System::new().await;
    let follower_id = Uuid::new_v4();
    let outside_ids = vec![Uuid::new_v4(), Uuid::new_v4()];
    let processes = vec![follower_id];
    let (tx, rx) = unbounded();
    let first_log_entry_timestamp = SystemTime::now();
    let follower = Raft::new(
        &mut system,
        make_config(follower_id, Duration::from_millis(100), processes),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(RamSender { tx }),
    )
    .await;

    // when
    follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: outside_ids[0],
                term: 1,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0,
            }),
        })
        .await;
    let client_id = Uuid::from_u128(1);
    follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: outside_ids[1],
                term: 2,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                prev_log_index: 0,
                prev_log_term: 0,
                entries: vec![
                    LogEntry {
                        term: 2,
                        timestamp: SystemTime::now(),
                        content: LogEntryContent::RegisterClient,
                    },
                    LogEntry {
                        term: 2,
                        timestamp: SystemTime::now(),
                        content: LogEntryContent::Command {
                            data: vec![1, 2, 3, 4],
                            client_id,
                            sequence_num: 0,
                            lowest_sequence_num_without_response: 0,
                        },
                    },
                ],
                leader_commit: 0,
            }),
        })
        .await;

    // then
    assert_eq!(
        rx.recv().await.unwrap(),
        RaftMessage {
            header: RaftMessageHeader {
                source: follower_id,
                term: 1,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true
            })
        }
    );
    assert_eq!(
        rx.recv().await.unwrap(),
        RaftMessage {
            header: RaftMessageHeader {
                source: follower_id,
                term: 2,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: true,
                last_log_index: 2
            })
        }
    );

    system.shutdown().await;
}
