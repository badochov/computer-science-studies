use async_channel::Sender;
use executor::{Handler, ModuleRef, System};
use std::convert::TryFrom;
use std::future::Future;
use std::pin::Pin;
use uuid::Uuid;

#[derive(Copy, Clone, Eq, PartialEq, Hash, Ord, PartialOrd, Debug)]
pub(crate) enum ProductType {
    Electronics,
    Toys,
    Books,
}

#[derive(Clone)]
pub(crate) struct StoreMsg {
    sender: ModuleRef<CyberStore2047>,
    content: StoreMsgContent,
}

#[derive(Clone, Debug)]
pub(crate) enum StoreMsgContent {
    /// Transaction Manager initiates voting for the transaction.
    RequestVote(Transaction),
    /// If every process is ok with transaction, TM issues commit.
    Commit,
    /// System-wide abort.
    Abort,
}

#[derive(Clone)]
pub(crate) struct NodeMsg {
    sender: ModuleRef<Node>,
    content: NodeMsgContent,
}

#[derive(Clone, Debug)]
pub(crate) enum NodeMsgContent {
    /// Process replies to TM whether it can/cannot commit the transaction.
    RequestVoteResponse(TwoPhaseResult),
    /// Process acknowledges to TM committing/aborting the transaction.
    FinalizationAck,
}

pub(crate) struct TransactionMessage {
    /// Request to change price.
    pub(crate) transaction: Transaction,

    /// Called after 2PC completes (i.e., the transaction was decided to be
    /// committed/aborted by CyberStore2047). This must be called after responses
    /// from all processes acknowledging commit or abort are collected.
    pub(crate) completed_callback:
        Box<dyn FnOnce(TwoPhaseResult) -> Pin<Box<dyn Future<Output = ()> + Send>> + Send>,
}

#[derive(Copy, Clone, Debug, Eq, PartialEq)]
pub(crate) enum TwoPhaseResult {
    Ok,
    Abort,
}

#[derive(Copy, Clone)]
pub(crate) struct Product {
    pub(crate) identifier: Uuid,
    pub(crate) pr_type: ProductType,
    pub(crate) price: u64,
}

#[derive(Copy, Clone, Debug)]
pub(crate) struct Transaction {
    pub(crate) pr_type: ProductType,
    pub(crate) shift: i32,
}

pub(crate) struct ProductPriceQuery {
    pub(crate) product_ident: Uuid,
    pub(crate) result_sender: Sender<ProductPrice>,
}

pub(crate) struct ProductPrice(pub(crate) Option<u64>);

/// Message which disables a node. Used for testing.
pub(crate) struct Disable;

/// Register and initialize a CyberStore2047 module.
pub(crate) async fn register_store(
    system: &mut System,
    store: CyberStore2047,
) -> ModuleRef<CyberStore2047> {
    let module = system.register_module(store).await;

    module.send(module.clone()).await;

    module
}

/// Register and initialize a Node module.
pub(crate) async fn register_node(system: &mut System, node: Node) -> ModuleRef<Node> {
    let module = system.register_module(node).await;

    module.send(module.clone()).await;

    module
}

/// CyberStore2047.
/// This structure serves as TM.
// Add any fields you need.
pub(crate) struct CyberStore2047 {
    nodes: Vec<ModuleRef<Node>>,
    response_count: usize,
    result: TwoPhaseResult,
    module_ref: Option<ModuleRef<CyberStore2047>>,
    callback:
        Option<Box<dyn FnOnce(TwoPhaseResult) -> Pin<Box<dyn Future<Output = ()> + Send>> + Send>>,
}

impl CyberStore2047 {
    pub(crate) fn new(nodes: Vec<ModuleRef<Node>>) -> Self {
        CyberStore2047 {
            nodes,
            module_ref: None,
            result: TwoPhaseResult::Ok,
            response_count: 0,
            callback: None,
        }
    }
}

/// Node of CyberStore2047.
/// This structure serves as a process of the distributed system.
// Add any fields you need.
pub(crate) struct Node {
    products: Vec<Product>,
    pending_transaction: Option<Transaction>,
    enabled: bool,
    module_ref: Option<ModuleRef<Node>>,
}

impl Node {
    pub(crate) fn new(products: Vec<Product>) -> Self {
        Self {
            products,
            pending_transaction: None,
            enabled: true,
            module_ref: None,
        }
    }
}

#[async_trait::async_trait]
impl Handler<NodeMsg> for CyberStore2047 {
    async fn handle(&mut self, msg: NodeMsg) {
        match msg.content {
            NodeMsgContent::FinalizationAck => {
                self.response_count += 1;
                if self.response_count == self.nodes.len() {
                    self.response_count = 0;
                    let callback = std::mem::replace(&mut self.callback, None);
                    (callback.unwrap())(self.result).await;
                }
            }
            NodeMsgContent::RequestVoteResponse(resp) => {
                if matches!(resp, TwoPhaseResult::Abort) {
                    self.result = resp;
                }
                self.response_count += 1;
                if self.response_count == self.nodes.len() {
                    self.response_count = 0;
                    let content = match self.result {
                        TwoPhaseResult::Abort => StoreMsgContent::Abort,
                        TwoPhaseResult::Ok => StoreMsgContent::Commit,
                    };

                    for node in self.nodes.iter() {
                        node.send(StoreMsg {
                            sender: self.module_ref.clone().unwrap(),
                            content: content.clone(),
                        })
                        .await
                    }
                }
            }
        }
    }
}

#[async_trait::async_trait]
impl Handler<StoreMsg> for Node {
    async fn handle(&mut self, msg: StoreMsg) {
        if self.enabled {
            let content = match msg.content {
                StoreMsgContent::RequestVote(transaction) => {
                    self.pending_transaction = Some(transaction);
                    let mut wrong_values = self
                        .products
                        .iter()
                        .filter(|p| p.pr_type == transaction.pr_type)
                        .filter(|p| {
                            match u64::try_from(p.price as i128 + transaction.shift as i128) {
                                Ok(v) => v == 0,
                                Err(_) => true,
                            }
                        });

                    match wrong_values.next() {
                        None => NodeMsgContent::RequestVoteResponse(TwoPhaseResult::Ok),
                        Some(_) => NodeMsgContent::RequestVoteResponse(TwoPhaseResult::Abort),
                    }
                }
                StoreMsgContent::Commit => {
                    let transaction = self.pending_transaction.unwrap();

                    self.products
                        .iter_mut()
                        .filter(|p| p.pr_type == transaction.pr_type)
                        .for_each(|p| {
                            p.price = (p.price as i128 + transaction.pr_type as i128) as u64
                        });
                    self.pending_transaction = None;
                    NodeMsgContent::FinalizationAck
                }
                StoreMsgContent::Abort => {
                    self.pending_transaction = None;
                    NodeMsgContent::FinalizationAck
                }
            };

            msg.sender
                .send(NodeMsg {
                    sender: self.module_ref.clone().unwrap(),
                    content,
                })
                .await;
        }
    }
}

#[async_trait::async_trait]
impl Handler<ProductPriceQuery> for Node {
    async fn handle(&mut self, msg: ProductPriceQuery) {
        if self.enabled {
            let mut iter = self.products.iter();
            let elem = iter.find(|&el| el.identifier == msg.product_ident);
            let product = match elem {
                None => ProductPrice(None),
                Some(p) => ProductPrice(Some(p.price)),
            };

            msg.result_sender.send(product).await.unwrap();
        }
    }
}

#[async_trait::async_trait]
impl Handler<Disable> for Node {
    async fn handle(&mut self, _msg: Disable) {
        self.enabled = false;
    }
}

#[async_trait::async_trait]
impl Handler<ModuleRef<CyberStore2047>> for CyberStore2047 {
    async fn handle(&mut self, module_ref: ModuleRef<CyberStore2047>) {
        self.module_ref = Some(module_ref);
    }
}

#[async_trait::async_trait]
impl Handler<ModuleRef<Node>> for Node {
    async fn handle(&mut self, module_ref: ModuleRef<Node>) {
        self.module_ref = Some(module_ref);
    }
}

#[async_trait::async_trait]
impl Handler<TransactionMessage> for CyberStore2047 {
    async fn handle(&mut self, msg: TransactionMessage) {
        let transaction = msg.transaction.clone();
        self.response_count = 0;
        self.result = TwoPhaseResult::Ok;
        self.callback = Some(msg.completed_callback);
        let sender = self.module_ref.clone().unwrap();

        for node in self.nodes.iter() {
            node.send(StoreMsg {
                sender: sender.clone(),
                content: StoreMsgContent::RequestVote(transaction),
            })
            .await
        }
    }
}
