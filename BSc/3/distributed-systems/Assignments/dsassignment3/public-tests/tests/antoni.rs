use std::time::{Duration, Instant, SystemTime};

use async_channel::unbounded;
use ntest::timeout;
use uuid::Uuid;

use executor::{ModuleRef, System};

use assignment_3_solution::*;
use assignment_3_test_utils::*;

#[tokio::test]
#[timeout(1000)]
async fn system_makes_progress_when_there_is_a_majority() {
    env_logger::try_init().unwrap_or(());
    // given
    let mut system = System::new().await;

    let entry_data = vec![1, 2, 3, 4, 5];
    let ident_leader = Uuid::new_v4();
    let ident_follower = Uuid::new_v4();
    let processes = vec![ident_leader, ident_follower, Uuid::new_v4()];
    let sender = ExecutorSender::default();

    let first_log_entry_timeout = SystemTime::now();
    let raft_leader = Raft::new(
        &mut system,
        make_config(ident_leader, Duration::from_millis(100), processes.clone()),
        first_log_entry_timeout,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    let raft_follower = Raft::new(
        &mut system,
        make_config(ident_follower, Duration::from_millis(300), processes),
        first_log_entry_timeout,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender
        .insert(ident_leader, Box::new(raft_leader.clone()))
        .await;
    sender.insert(ident_follower, Box::new(raft_follower)).await;
    tokio::time::sleep(Duration::from_millis(200)).await;

    let (result_sender, result_receiver) = unbounded();

    // when
    let client_id = register_client(&raft_leader, &result_sender, &result_receiver).await;

    raft_leader
        .send(ClientRequest {
            reply_to: result_sender,
            content: ClientRequestContent::Command {
                command: entry_data.clone(),
                client_id,
                sequence_num: 0,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;

    // then
    assert_eq!(
        entry_data,
        *unwrap_output(&result_receiver.recv().await.unwrap())
    );

    system.shutdown().await;
}

struct TestEnv {
    leader_id: Uuid,
    leader: ModuleRef<Raft>,
    follower_id: Uuid,
    follower: ModuleRef<Raft>,
    spy_id: Uuid,
    spy_receiver: async_channel::Receiver<RaftMessage>,
    system: System,
    sender: ExecutorSender,
}

impl TestEnv {
    async fn stabilize(&mut self) {
        tokio::time::sleep(Duration::from_millis(200)).await;
    }

    async fn finish(mut self) {
        self.system.shutdown().await;
    }

    async fn spy_get_no_prologue(&mut self) -> RaftMessage {
        let request = RaftMessage {
            header: RaftMessageHeader {
                source: self.leader_id,
                term: 1,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0,
            }),
        };

        let msg = self.spy_receiver.recv().await.unwrap();
        if msg != request {
            return msg;
        }

        loop {
            let msg = self.spy_receiver.recv().await.unwrap();
            let append = RaftMessage {
                header: RaftMessageHeader {
                    source: self.leader_id,
                    term: 1,
                },
                content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                    prev_log_index: 0,
                    prev_log_term: 0,
                    entries: vec![],
                    leader_commit: 0,
                }),
            };
            if msg != append {
                return msg;
            }
        }
    }

    fn clear_spy_receiver(&mut self) {
        while let Ok(x) = self.spy_receiver.try_recv() {
            log::debug!("TEST: dropping message {:?}", x);
            drop(x);
        }
    }
}

async fn prepare_cluster() -> TestEnv {
    env_logger::try_init().unwrap_or(());
    // given
    let mut system = System::new().await;

    let leader_id = Uuid::new_v4();
    let follower_id = Uuid::new_v4();
    let spy_id = Uuid::new_v4();
    let processes = vec![leader_id, follower_id, spy_id];
    let sender = ExecutorSender::default();

    let first_log_entry_timeout = SystemTime::now();
    let leader = Raft::new(
        &mut system,
        make_config(leader_id, Duration::from_millis(100), processes.clone()),
        first_log_entry_timeout,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    let follower = Raft::new(
        &mut system,
        make_config(follower_id, Duration::from_millis(300), processes),
        first_log_entry_timeout,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(leader_id, Box::new(leader.clone())).await;
    sender.insert(follower_id, Box::new(follower.clone())).await;
    let (spy_sender, spy_receiver) = unbounded();
    sender
        .insert(
            spy_id,
            Box::new(RaftSpy::new(&mut system, None, spy_sender).await),
        )
        .await;

    TestEnv {
        leader_id,
        leader,
        follower_id,
        follower,
        spy_id,
        spy_receiver,
        system,
        sender,
    }
}

#[tokio::test]
#[timeout(1000)]
async fn respond_to_old_term_leader_append_entries() {
    let mut env = prepare_cluster().await;

    env.stabilize().await;

    env.leader
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 0,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                entries: Vec::new(),
                leader_commit: 0,
                prev_log_index: 0,
                prev_log_term: 0,
            }),
        })
        .await;

    // then
    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.leader_id,
                term: 1
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: false,
                last_log_index: 0
            })
        },
        env.spy_get_no_prologue().await
    );

    env.finish().await;
}

#[tokio::test]
#[timeout(1000)]
async fn respond_to_old_term_follower_append_entries() {
    let mut env = prepare_cluster().await;

    env.stabilize().await;

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 0,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                entries: Vec::new(),
                leader_commit: 0,
                prev_log_index: 0,
                prev_log_term: 0,
            }),
        })
        .await;

    // then
    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 1
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: false,
                last_log_index: 0
            })
        },
        env.spy_get_no_prologue().await
    );

    env.finish().await;
}

#[tokio::test]
#[timeout(1000)]
async fn respond_to_old_term_request_vote() {
    let mut env = prepare_cluster().await;

    // when

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 1,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: false,
            }),
        })
        .await;

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 0,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0,
            }),
        })
        .await;

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 1,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0,
            }),
        })
        .await;

    // then
    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 1
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: false,
            })
        },
        env.spy_get_no_prologue().await
    );

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 1
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            })
        },
        env.spy_get_no_prologue().await
    );

    env.finish().await;
}

#[tokio::test]
#[timeout(1000)]
async fn leader_steps_down_on_term() {
    let mut env = prepare_cluster().await;

    env.stabilize().await;

    env.leader
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0,
            }),
        })
        .await;

    env.leader
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: false,
            }),
        })
        .await;

    env.leader
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 3,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0,
            }),
        })
        .await;

    // then
    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.leader_id,
                term: 3
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            })
        },
        env.spy_get_no_prologue().await
    );

    env.finish().await;
}

#[tokio::test]
#[timeout(1000)]
async fn append_entries_batch_size() {
    let mut env = prepare_cluster().await;

    env.stabilize().await;

    let (result_sender, result_receiver) = unbounded();

    let client_id = register_client(&env.leader, &result_sender, &result_receiver).await;

    let mut entry_datas = Vec::new();

    for i in 0..15 {
        let val = i as u8;
        let entry_data = vec![val, val + 1, 2 * val, val * val];
        entry_datas.push(entry_data.clone());

        env.leader
            .send(ClientRequest {
                reply_to: result_sender.clone(),
                content: ClientRequestContent::Command {
                    command: entry_data,
                    client_id,
                    sequence_num: i,
                    lowest_sequence_num_without_response: 0,
                },
            })
            .await;
    }

    tokio::time::sleep(Duration::from_millis(100)).await;

    env.clear_spy_receiver();

    let entries = extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.leader_id,
            term: 1,
        },
        0,
        0,
        16,
    );

    let mut idx = 0;

    for entry in &entries[1..] {
        let LogEntry {
            content,
            term,
            timestamp: _,
        } = entry.clone();
        assert_eq!(
            LogEntryContent::Command {
                data: entry_datas[idx].clone(),
                client_id,
                sequence_num: idx as u64,
                lowest_sequence_num_without_response: 0,
            },
            content
        );
        assert_eq!(term, 1);
        idx += 1;
    }

    assert_eq!(idx, 9);

    env.leader
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 1,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 10,
                success: true,
            }),
        })
        .await;

    while let Ok(x) = env.spy_receiver.try_recv() {
        drop(x);
    }

    let entries = extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.leader_id,
            term: 1,
        },
        10,
        1,
        16,
    );

    for entry in entries {
        let LogEntry {
            content,
            term,
            timestamp: _,
        } = entry;
        assert_eq!(
            LogEntryContent::Command {
                data: entry_datas[idx].clone(),
                client_id,
                sequence_num: idx as u64,
                lowest_sequence_num_without_response: 0,
            },
            content
        );
        assert_eq!(term, 1);
        idx += 1;
    }

    assert_eq!(idx, 15);

    env.finish().await;
}

fn extract_entries(
    msg: RaftMessage,
    header: RaftMessageHeader,
    prev_idx: usize,
    prev_term: u64,
    commit_idx: usize,
) -> Vec<LogEntry> {
    assert_eq!(msg.header, header);
    match msg.content {
        RaftMessageContent::AppendEntries(AppendEntriesArgs {
            prev_log_index,
            prev_log_term,
            entries,
            leader_commit,
        }) if prev_log_index == prev_idx
            && prev_log_term == prev_term
            && leader_commit == commit_idx =>
        {
            entries
        }
        content => {
            panic!(
                "content {:?} does not fit command with prev_idx {}, prev_term {}, commit {}",
                content, prev_idx, prev_term, commit_idx
            );
        }
    }
}

#[tokio::test]
#[timeout(1500)]
async fn follower_overwrites_bad_logs() {
    let mut env = prepare_cluster().await;

    env.sender.break_link(env.leader_id, env.spy_id).await;

    env.stabilize().await;

    let (result_sender, result_receiver) = unbounded();

    let client_id = register_client(&env.leader, &result_sender, &result_receiver).await;

    env.sender.break_link(env.follower_id, env.leader_id).await;

    env.leader
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

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 2
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 2,
                last_log_term: 1
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 2,
        },
        2,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: false,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 2,
        },
        1,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: true,
            }),
        })
        .await;

    let entries: Vec<_> = extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 2,
        },
        1,
        1,
        1,
    )
    .into_iter()
    .map(unwrap_log_command(1))
    .collect();

    assert_eq!(entries, vec![vec![1, 2, 3, 4]]);

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 100,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                leader_commit: 0,
                prev_log_index: 1,
                prev_log_term: 1,
                entries: vec![LogEntry {
                    term: 70,
                    timestamp: SystemTime::now(),
                    content: LogEntryContent::Command {
                        data: vec![123, 32, 11],
                        sequence_num: 123,
                        client_id,
                        lowest_sequence_num_without_response: 12,
                    },
                }],
            }),
        })
        .await;

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 100
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: true,
                last_log_index: 2
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 101
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 2,
                last_log_term: 70
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        2,
        70,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: false,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        1,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: true,
            }),
        })
        .await;

    let entries: Vec<_> = extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        1,
        1,
        1,
    )
    .into_iter()
    .map(unwrap_log_command(70))
    .collect();

    assert_eq!(entries, vec![vec![123, 32, 11]]);

    env.finish().await;
}

fn unwrap_log_command(term: u64) -> Box<dyn Fn(LogEntry) -> Vec<u8>> {
    Box::new(move |entry: LogEntry| {
        assert_eq!(entry.term, term);
        match entry.content {
            LogEntryContent::Command { data, .. } => data,
            content => {
                panic!("Not a command {:?}", content);
            }
        }
    })
}

#[tokio::test]
#[timeout(1500)]
async fn follower_deletes_trailing_logs() {
    let mut env = prepare_cluster().await;

    env.sender.break_link(env.leader_id, env.spy_id).await;

    env.stabilize().await;

    let (result_sender, result_receiver) = unbounded();

    let client_id = register_client(&env.leader, &result_sender, &result_receiver).await;

    env.sender.break_link(env.follower_id, env.leader_id).await;

    env.leader
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

    env.leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::Command {
                command: vec![5, 2, 9, 4],
                client_id,
                sequence_num: 1,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 2
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 3,
                last_log_term: 1
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 2,
        },
        3,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: false,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 2,
        },
        1,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: true,
            }),
        })
        .await;

    let entries: Vec<_> = extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 2,
        },
        1,
        1,
        1,
    )
    .into_iter()
    .map(unwrap_log_command(1))
    .collect();

    assert_eq!(entries, vec![vec![1, 2, 3, 4], vec![5, 2, 9, 4]]);

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 100,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                leader_commit: 0,
                prev_log_index: 1,
                prev_log_term: 1,
                entries: vec![LogEntry {
                    term: 70,
                    timestamp: SystemTime::now(),
                    content: LogEntryContent::Command {
                        data: vec![123, 32, 11],
                        sequence_num: 123,
                        client_id,
                        lowest_sequence_num_without_response: 12,
                    },
                }],
            }),
        })
        .await;

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 100
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: true,
                last_log_index: 2
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 101
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 2,
                last_log_term: 70
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        2,
        70,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: false,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        1,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: true,
            }),
        })
        .await;

    let entries: Vec<_> = extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        1,
        1,
        1,
    )
    .into_iter()
    .map(unwrap_log_command(70))
    .collect();

    assert_eq!(entries, vec![vec![123, 32, 11]]);

    env.finish().await;
}

#[tokio::test]
#[timeout(1500)]
async fn follower_doesnt_delete_on_mismatch() {
    let mut env = prepare_cluster().await;

    env.sender.break_link(env.leader_id, env.spy_id).await;

    env.stabilize().await;

    let (result_sender, result_receiver) = unbounded();

    let client_id = register_client(&env.leader, &result_sender, &result_receiver).await;

    env.sender.break_link(env.follower_id, env.leader_id).await;

    env.leader
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

    env.leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::Command {
                command: vec![5, 2, 9, 4],
                client_id,
                sequence_num: 1,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 2
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 3,
                last_log_term: 1
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 2,
        },
        3,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: false,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 2,
        },
        1,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: true,
            }),
        })
        .await;

    let entries: Vec<_> = extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 2,
        },
        1,
        1,
        1,
    )
    .into_iter()
    .map(unwrap_log_command(1))
    .collect();

    assert_eq!(entries, vec![vec![1, 2, 3, 4], vec![5, 2, 9, 4]]);

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 100,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                leader_commit: 0,
                prev_log_index: 2,
                prev_log_term: 18,
                entries: vec![],
            }),
        })
        .await;

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 100
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: false,
                last_log_index: 3
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 100,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                leader_commit: 0,
                prev_log_index: 3,
                prev_log_term: 18,
                entries: vec![],
            }),
        })
        .await;

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 100
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: false,
                last_log_index: 3
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 101
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 3,
                last_log_term: 1
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        3,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 3,
                success: false,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        2,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 3,
                success: false,
            }),
        })
        .await;

    extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        1,
        1,
        1,
    );

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 101,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                last_log_index: 1,
                success: true,
            }),
        })
        .await;

    let entries: Vec<_> = extract_entries(
        env.spy_receiver.recv().await.unwrap(),
        RaftMessageHeader {
            source: env.follower_id,
            term: 101,
        },
        1,
        1,
        1,
    )
    .into_iter()
    .map(unwrap_log_command(1))
    .collect();

    assert_eq!(entries, vec![vec![1, 2, 3, 4], vec![5, 2, 9, 4]]);

    env.finish().await;
}

#[tokio::test]
#[timeout(1500)]
async fn follower_revotes_the_same_candidate() {
    let env = prepare_cluster().await;

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                term: 5,
                source: env.spy_id,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_term: 0,
                last_log_index: 0,
            }),
        })
        .await;

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                term: 5,
                source: env.spy_id,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_term: 0,
                last_log_index: 0,
            }),
        })
        .await;

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                term: 6,
                source: env.leader_id,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_term: 0,
                last_log_index: 0,
            }),
        })
        .await;

    env.follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                term: 6,
                source: env.spy_id,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_term: 0,
                last_log_index: 0,
            }),
        })
        .await;

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 5
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 5
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 6
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: false
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.leader_id,
                term: 7
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    env.finish().await;
}

#[tokio::test]
#[timeout(2000)]
async fn candidate_converts_on_append_entries() {
    let env = prepare_cluster().await;

    env.sender.break_link(env.follower_id, env.leader_id).await;
    env.sender.break_link(env.follower_id, env.spy_id).await;

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.leader_id,
                term: 1
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.leader_id,
                term: 2
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    env.leader
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.spy_id,
                term: 2,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                leader_commit: 0,
                prev_log_index: 0,
                prev_log_term: 0,
                entries: vec![
                    LogEntry {
                        term: 1,
                        timestamp: SystemTime::now(),
                        content: LogEntryContent::RegisterClient,
                    },
                    LogEntry {
                        term: 1,
                        timestamp: SystemTime::now(),
                        content: LogEntryContent::RegisterClient,
                    },
                ],
            }),
        })
        .await;

    env.leader
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: env.follower_id,
                term: 2,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: true,
            }),
        })
        .await;

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.leader_id,
                term: 2
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: true,
                last_log_index: 2
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );

    assert_eq!(
        RaftMessage {
            header: RaftMessageHeader {
                source: env.leader_id,
                term: 3
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 2,
                last_log_term: 1
            })
        },
        env.spy_receiver.recv().await.unwrap()
    );
    env.finish().await;
}

#[tokio::test]
#[timeout(2000)]
async fn candidate_timeouts_properly() {
    let env = prepare_cluster().await;

    let start = Instant::now();

    env.sender.break_link(env.follower_id, env.leader_id).await;
    env.sender.break_link(env.follower_id, env.spy_id).await;

    let iterations = 15;
    for i in 0..iterations {
        assert_eq!(
            RaftMessage {
                header: RaftMessageHeader {
                    source: env.leader_id,
                    term: i + 1
                },
                content: RaftMessageContent::RequestVote(RequestVoteArgs {
                    last_log_index: 0,
                    last_log_term: 0
                })
            },
            env.spy_receiver.recv().await.unwrap()
        );
    }

    let elapsed = start.elapsed();

    println!("{:?}", elapsed);
    assert!(elapsed < Duration::from_millis(1550));
    assert!(elapsed > Duration::from_millis(1450));

    env.finish().await;
}

#[tokio::test]
#[timeout(1500)]
async fn system_does_not_make_progress_without_majority() {
    env_logger::try_init().unwrap_or(());
    // given
    let mut system = System::new().await;

    let entry_data = vec![1, 2, 3, 4, 5];
    let ident_leader = Uuid::new_v4();
    let ident_follower = Uuid::new_v4();
    let processes = vec![ident_leader, ident_follower, Uuid::new_v4()];
    let sender = ExecutorSender::default();

    let first_log_entry_timestamp = SystemTime::now();
    let raft_leader = Raft::new(
        &mut system,
        make_config(ident_leader, Duration::from_millis(100), processes.clone()),
        first_log_entry_timestamp,
        Box::new(DummyMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    let raft_follower = Raft::new(
        &mut system,
        make_config(
            ident_follower,
            Duration::from_millis(300),
            processes.clone(),
        ),
        first_log_entry_timestamp,
        Box::new(DummyMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender
        .insert(ident_leader, Box::new(raft_leader.clone()))
        .await;
    sender
        .insert(ident_follower, Box::new(raft_follower.clone()))
        .await;
    tokio::time::sleep(Duration::from_millis(200)).await;

    let (result_sender, result_receiver) = unbounded();

    // when
    let client_id = register_client(&raft_leader, &result_sender, &result_receiver).await;

    sender.break_link(ident_follower, ident_leader).await;

    raft_leader
        .send(ClientRequest {
            reply_to: result_sender,
            content: ClientRequestContent::Command {
                command: entry_data,
                client_id,
                sequence_num: 0,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;

    tokio::time::sleep(Duration::from_millis(1000)).await;

    // then
    assert!(result_receiver.is_empty());

    system.shutdown().await;
}

#[tokio::test]
#[timeout(1500)]
async fn follower_denies_vote_for_candidate_with_outdated_log() {
    env_logger::try_init().unwrap_or(());
    // given
    let mut system = System::new().await;

    let (tx, rx) = unbounded();
    let storage = RamStorage::default();
    let other_ident_1 = Uuid::new_v4();
    let other_ident_2 = Uuid::new_v4();
    let ident_follower = Uuid::new_v4();
    let first_log_entry_timestamp = SystemTime::now();
    let raft_follower = Raft::new(
        &mut system,
        make_config(
            ident_follower,
            Duration::from_millis(500),
            vec![other_ident_1, other_ident_2, ident_follower],
        ),
        first_log_entry_timestamp,
        Box::new(DummyMachine),
        Box::new(storage),
        Box::new(RamSender { tx }),
    )
    .await;
    let client_id = Uuid::from_u128(1);

    // when

    raft_follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                term: 2,
                source: other_ident_1,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                prev_log_index: 0,
                prev_log_term: 0,
                entries: vec![
                    LogEntry {
                        content: LogEntryContent::RegisterClient,
                        term: 2,
                        timestamp: SystemTime::now(),
                    },
                    LogEntry {
                        content: LogEntryContent::Command {
                            data: vec![1],
                            client_id,
                            sequence_num: 0,
                            lowest_sequence_num_without_response: 0,
                        },
                        term: 2,
                        timestamp: SystemTime::now(),
                    },
                ],
                leader_commit: 0,
            }),
        })
        .await;

    // Wait longer than election timeout so that the follower does not ignore the vote request
    tokio::time::sleep(Duration::from_millis(600)).await;

    // Older term of the last message.
    raft_follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                term: 4,
                source: other_ident_2,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 2,
                last_log_term: 1,
            }),
        })
        .await;

    // Shorter log in candidate.
    raft_follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                term: 5,
                source: other_ident_2,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 1,
                last_log_term: 2,
            }),
        })
        .await;

    // then
    assert_eq!(
        rx.recv().await.unwrap(),
        RaftMessage {
            header: RaftMessageHeader {
                source: ident_follower,
                term: 2,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: true,
                last_log_index: 2,
            })
        }
    );
    for _ in 0..2 {
        let msg = rx.recv().await.unwrap();
        log::debug!("Dropping message {:?}", msg);
        log::logger().flush();
    }
    log::debug!("TEST TEST");
    log::logger().flush();
    assert_eq!(
        rx.recv().await.unwrap(),
        RaftMessage {
            header: RaftMessageHeader {
                source: ident_follower,
                term: 4,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: false,
            })
        }
    );
    assert_eq!(
        rx.recv().await.unwrap(),
        RaftMessage {
            header: RaftMessageHeader {
                source: ident_follower,
                term: 5,
            },
            content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                vote_granted: false,
            })
        }
    );

    system.shutdown().await;
}

#[tokio::test]
#[timeout(500)]
async fn follower_rejects_inconsistent_append_entry() {
    env_logger::try_init().unwrap_or(());
    // given
    let mut system = System::new().await;

    let (tx, rx) = unbounded();
    let storage = RamStorage::default();
    let first_log_entry_timestamp = SystemTime::now();
    let other_ident = Uuid::new_v4();
    let ident_follower = Uuid::new_v4();
    let raft_follower = Raft::new(
        &mut system,
        make_config(
            ident_follower,
            Duration::from_secs(10),
            vec![other_ident, ident_follower],
        ),
        first_log_entry_timestamp,
        Box::new(DummyMachine),
        Box::new(storage),
        Box::new(RamSender { tx }),
    )
    .await;

    // when
    let client_id = Uuid::from_u128(1);

    raft_follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                term: 1,
                source: other_ident,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                prev_log_index: 0,
                prev_log_term: 0,
                entries: vec![
                    LogEntry {
                        content: LogEntryContent::RegisterClient,
                        term: 1,
                        timestamp: SystemTime::now(),
                    },
                    LogEntry {
                        content: LogEntryContent::Command {
                            data: vec![1, 2, 3, 4],
                            client_id,
                            sequence_num: 0,
                            lowest_sequence_num_without_response: 0,
                        },
                        term: 1,
                        timestamp: SystemTime::now(),
                    },
                ],
                leader_commit: 0,
            }),
        })
        .await;

    raft_follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                term: 2,
                source: other_ident,
            },
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                prev_log_index: 2,
                prev_log_term: 2,
                entries: vec![LogEntry {
                    content: LogEntryContent::Command {
                        data: vec![5, 6, 7, 8],
                        client_id,
                        sequence_num: 0,
                        lowest_sequence_num_without_response: 0,
                    },
                    term: 2,
                    timestamp: SystemTime::now(),
                }],
                leader_commit: 0,
            }),
        })
        .await;

    // then
    assert_eq!(
        rx.recv().await.unwrap(),
        RaftMessage {
            header: RaftMessageHeader {
                term: 1,
                source: ident_follower,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: true,
                last_log_index: 2
            })
        }
    );
    assert_eq!(
        rx.recv().await.unwrap(),
        RaftMessage {
            header: RaftMessageHeader {
                term: 2,
                source: ident_follower,
            },
            content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                success: false,
                last_log_index: 2
            })
        }
    );

    system.shutdown().await;
}

#[tokio::test]
#[timeout(500)]
async fn follower_redirects_to_leader() {
    env_logger::try_init().unwrap_or(());
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
        make_config(follower_id, Duration::from_millis(300), processes),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(leader_id, Box::new(leader.clone())).await;
    sender.insert(follower_id, Box::new(follower.clone())).await;
    tokio::time::sleep(Duration::from_millis(200)).await;

    // when
    let (follower_result_sender, follower_result_receiver) = unbounded();
    let (leader_result_sender, leader_result_receiver) = unbounded();

    let client_id = register_client(&leader, &leader_result_sender, &leader_result_receiver).await;

    follower
        .send(ClientRequest {
            reply_to: follower_result_sender,
            content: ClientRequestContent::Command {
                command: vec![1, 2, 3, 4],
                client_id,
                sequence_num: 1,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;
    leader
        .send(ClientRequest {
            reply_to: leader_result_sender,
            content: ClientRequestContent::Command {
                command: vec![5, 6, 7, 8],
                client_id,
                sequence_num: 0,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;

    // then
    assert_eq!(
        follower_result_receiver.recv().await.unwrap(),
        ClientRequestResponse::CommandResponse(CommandResponseArgs {
            client_id,
            sequence_num: 1,
            content: CommandResponseContent::NotLeader {
                leader_hint: Some(leader_id)
            }
        })
    );
    assert_eq!(
        leader_result_receiver.recv().await.unwrap(),
        ClientRequestResponse::CommandResponse(CommandResponseArgs {
            client_id,
            sequence_num: 0,
            content: CommandResponseContent::CommandApplied {
                output: vec![5, 6, 7, 8]
            },
        })
    );

    system.shutdown().await;
}

#[tokio::test]
#[timeout(500)]
async fn leader_steps_down_without_heartbeat_responses_from_majority() {
    env_logger::try_init().unwrap_or(());
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
        make_config(follower_id, Duration::from_millis(300), processes),
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

    let client_id = register_client(&leader, &result_sender, &result_receiver).await;

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

    let response_1 = result_receiver.recv().await.unwrap();

    sender.break_link(follower_id, leader_id).await;
    tokio::time::sleep(Duration::from_millis(200)).await;

    leader
        .send(ClientRequest {
            reply_to: result_sender.clone(),
            content: ClientRequestContent::Command {
                command: vec![5, 6, 7, 8],
                client_id,
                sequence_num: 1,
                lowest_sequence_num_without_response: 0,
            },
        })
        .await;

    let response_2 = result_receiver.recv().await.unwrap();

    assert_eq!(
        response_1,
        ClientRequestResponse::CommandResponse(CommandResponseArgs {
            client_id,
            sequence_num: 0,
            content: CommandResponseContent::CommandApplied {
                output: vec![1, 2, 3, 4]
            },
        })
    );
    assert_eq!(
        response_2,
        ClientRequestResponse::CommandResponse(CommandResponseArgs {
            client_id,
            sequence_num: 1,
            content: CommandResponseContent::NotLeader { leader_hint: None }
        })
    );

    system.shutdown().await;
}

#[tokio::test]
#[timeout(1000)]
async fn follower_ignores_request_vote_within_election_timeout_of_leader_heartbeat() {
    env_logger::try_init().unwrap_or(());
    // given
    let mut system = System::new().await;
    let leader_id = Uuid::new_v4();
    let follower_id = Uuid::new_v4();
    let spy_id = Uuid::new_v4();
    let processes = vec![leader_id, follower_id, spy_id];
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
        make_config(follower_id, Duration::from_millis(300), processes),
        first_log_entry_timestamp,
        Box::new(IdentityMachine),
        Box::new(RamStorage::default()),
        Box::new(sender.clone()),
    )
    .await;
    sender.insert(leader_id, Box::new(leader.clone())).await;
    sender.insert(follower_id, Box::new(follower.clone())).await;
    tokio::time::sleep(Duration::from_millis(150)).await;

    // when
    let (spy_sender, spy_receiver) = unbounded();
    sender
        .insert(
            spy_id,
            Box::new(RaftSpy::new(&mut system, None, spy_sender).await),
        )
        .await;

    follower
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: spy_id,
                term: 2,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0,
            }),
        })
        .await;
    leader
        .send(RaftMessage {
            header: RaftMessageHeader {
                source: spy_id,
                term: 2,
            },
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: 0,
                last_log_term: 0,
            }),
        })
        .await;
    tokio::time::sleep(Duration::from_millis(500)).await;

    // then
    while let Ok(msg) = spy_receiver.try_recv() {
        assert_eq!(
            msg,
            RaftMessage {
                header: RaftMessageHeader {
                    source: leader_id,
                    term: 1,
                },
                content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                    prev_log_index: 0,
                    prev_log_term: 0,
                    entries: vec![],
                    leader_commit: 0,
                })
            }
        );
    }

    system.shutdown().await;
}
