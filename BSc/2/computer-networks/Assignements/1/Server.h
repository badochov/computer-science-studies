#ifndef INC_1_SERVER_H
#define INC_1_SERVER_H

#include "HeaderFields.h"
#include "error.h"
#include "utils.h"

#include <string>
#include <sstream>
#include <utility>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <functional>
#include <vector>
#include <regex>
#include <iostream>
#include <csignal>

struct Request {

  struct RequestData {
    const std::string method;
    const std::string target;
    const std::string http_version;

    static RequestData parse_request_data(std::string const &request_line);

    friend std::ostream &operator<<(std::ostream &os, RequestData const &request);

   private:

    inline static const std::vector<std::string> valid_methods = {"GET", "HEAD"};
    inline static const std::vector<std::string> valid_http_versions = {"HTTP/1.1"};

    template<class T>
    static bool in_vector(std::vector<T> const &v, T const &el) {
      return std::find(v.begin(), v.end(), el) != v.end();
    }

    static void validate_method(std::string const &token) {
      if (!in_vector(valid_methods, token)) {
        throw HTTPError(ResponseCode::NOT_IMPLEMENTED, "Method not implemented");
      }
    }

    static void validate_target(std::string const &token);

    static void validate_http_version(std::string const &token) {
      if (!in_vector(valid_http_versions, token)) {
        throw HTTPError(ResponseCode::INCORRECT_REQUEST, "Invalid HTTP version");
      }
    }

    static std::smatch get_request_line_parts(std::string const &request_line) {
      // 8192 is longest acceptable by server length of token in request line.
      std::regex re("([^ ]{1,8192}) ([^ ]{1,8192}) ([^ ]{1,8192})");
      std::smatch match;

      std::regex_match(request_line, match, re);

      return match;
    }
  };

  const HeaderFields headers;
  const RequestData request_data;

  Request(RequestData _request_data,
          HeaderFields headers)
      : headers(std::move(headers)), request_data(std::move(_request_data)) {}

  friend std::ostream &operator<<(std::ostream &os, Request const &request) {
    os << request.request_data << request.headers;

    return os;
  }
};

class DataStream {
  static constexpr size_t BUFFER_SIZE = 4096;

  char buffer_helper[BUFFER_SIZE]{};
  int fd;
  bool eof = false;

  std::stringstream out_stream;

  std::string buffer;

  ssize_t read_to_buffer();

  static size_t find_eol(std::string const &s) {
    return s.find(EOL);
  }

  void clear_out_buffer() {
    out_stream.clear();
    out_stream.str("");
  }

 public:
  explicit DataStream(int file_descriptor) : fd(file_descriptor) {
  }

  void change_socket(int new_socket) noexcept {
    fd = new_socket;
  }

  std::string read_line();

  [[nodiscard]] bool is_eof() const noexcept {
    return eof;
  }

  void clear();

  [[nodiscard]]std::stringstream &get_out_stream() {
    return out_stream;
  }

  void send_response();
};

class Server {
 public:
  using port_t = uint16_t;
 private:
  static constexpr int QUEUE_LENGTH = 5;

  int msg_socket{};
  int server_socket;
  struct sockaddr_in server_address{};
  struct sockaddr_in client_address{};
  socklen_t client_address_len{};

  const port_t port;
  bool connected = false;
  bool should_close = false;
  DataStream data_stream;

  HeaderFields get_header_fields();

  int init_socket();

  void turn_on_server();

  void turn_off_server() const;

  void wait_for_connection();

  [[nodiscard]] Request _get_request();

  void _respond(ResponseCode status_code,
                std::string const &reason,
                HeaderFields const &header_fields,
                std::string const &message);

  bool should_be_closed(Request const &request) {
    return request.headers.is_field_set_to("Connection", "close") || data_stream.is_eof();
  }

  static void validate_headers(HeaderFields const &header_fields) {
    std::string content_length = "Content-Length";
    std::string connection = "Connection";

    if (header_fields.has(content_length) && header_fields[content_length] != "0") {
      throw HTTPError(ResponseCode::INCORRECT_REQUEST, "Sending content is not supported");
    }

    if (header_fields.has(content_length) && header_fields[content_length] != "0") {
      throw HTTPError(ResponseCode::INCORRECT_REQUEST, "Sending content is not supported");
    }

  }
 public:
  explicit Server(port_t _port) : server_socket(init_socket()), port(_port), data_stream(server_socket) {
    turn_on_server();
  }

  ~Server() {
    if (should_close) {
      close_connection();
    }
    turn_off_server();
  }

  void close_connection();

  [[nodiscard]] Request get_request() {
    if (!connected) {
      wait_for_connection();
    }

    Request request = _get_request();
    should_close = should_be_closed(request);

    return request;
  }

  void respond(ResponseCode status_code,
               std::string const &reason,
               HeaderFields &header_fields,
               std::string const &message) {
    if (should_close) {
      header_fields.add("Connection", "close");
    }

    _respond(status_code, reason, header_fields, message);

    if (should_close) {
      should_close = false;
      close_connection();
    }
  }
};

#endif //INC_1_SERVER_H
