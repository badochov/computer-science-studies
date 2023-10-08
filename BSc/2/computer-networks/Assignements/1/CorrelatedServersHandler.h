#ifndef INC_1__CORRELATEDSERVERSHANDLER_H_
#define INC_1__CORRELATEDSERVERSHANDLER_H_

#include "Server.h"

#include <utility>
#include <fstream>

class CorrelatedServersHandler {
  const std::string path;
  struct CorrelatedServer {
    const std::string ip;
    const std::string port;

    CorrelatedServer(std::string ipv4, std::string _port) : ip(std::move(ipv4)), port(std::move(_port)) {};
  };

  [[nodiscard]] CorrelatedServer get_file_correlated_server(std::string const &file_path) const {
    std::ifstream correlated_servers_file_stream(path);

    std::string current_file_path;
    std::string ip;
    std::string port;

    while (correlated_servers_file_stream >> current_file_path >> ip >> port) {
      if (current_file_path == file_path) {
        return CorrelatedServer(ip, port);
      }
    }

    throw HTTPError(ResponseCode::RESOURCE_NOT_FOUND, "Not Found");
  }

  static std::string make_redirection_address(CorrelatedServer const &server,
                                              std::string const &path) noexcept {
    return "http://" + server.ip + ":" + server.port + path;
  }

 public:
  explicit CorrelatedServersHandler(std::string lookup_path) : path(std::move(
      lookup_path)) {}

  void redirect(std::string const &file_path, Server &server) const {
    CorrelatedServer correlated_server = get_file_correlated_server(file_path);
    std::string redirection_address = make_redirection_address(correlated_server, file_path);

    HeaderFields header_fields;
    header_fields.add("Location", redirection_address);

    server.respond(ResponseCode::REDIRECT, "Redirect", header_fields, "");
  }
};

#endif //INC_1__CORRELATEDSERVERSHANDLER_H_
