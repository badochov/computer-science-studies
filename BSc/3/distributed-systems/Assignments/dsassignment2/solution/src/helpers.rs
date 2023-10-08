pub async fn get_socket_addr((addr, port): &(String, u16)) -> std::net::SocketAddr {
    tokio::net::lookup_host(format!("{}:{}", addr, port))
        .await
        .unwrap()
        .next()
        .unwrap()
}
