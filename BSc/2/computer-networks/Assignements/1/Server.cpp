#include "Server.h"

int Server::init_socket() {
  server_socket = socket(PF_INET, SOCK_STREAM, 0);
  if (server_socket < 0) {
    throw std::runtime_error("Error while calling socket function!");
  }

  return server_socket;
}

void Server::turn_on_server() {
  server_address.sin_family = AF_INET;
  server_address.sin_addr.s_addr = htonl(INADDR_ANY);
  server_address.sin_port = htons(port);

  if (bind(server_socket, (struct sockaddr *) &server_address, sizeof(server_address)) < 0) {
    throw std::runtime_error("Error while calling bind function!");
  }

  if (listen(server_socket, QUEUE_LENGTH) < 0) {
    throw std::runtime_error("Error while calling listen function!");
  }
}
void Server::turn_off_server() const {
  if (shutdown(server_socket, SHUT_RDWR) < 0) {
    throw std::runtime_error("Error while closing server!");
  }
  if (close(server_socket) < 0) {
    throw std::runtime_error("Error while closing server!");
  }
}

void Server::wait_for_connection() {
  client_address_len = sizeof(client_address);

  msg_socket = accept(server_socket, (struct sockaddr *) &client_address, &client_address_len);
  if (msg_socket < 0) {
    throw std::runtime_error("Error while calling accept function!");
  }

  data_stream.change_socket(msg_socket);
  connected = true;
}

Request Server::_get_request() {
  std::string request_line = data_stream.read_line();

  try {
    Request::RequestData request_data = Request::RequestData::parse_request_data(request_line);
    HeaderFields header_fields = get_header_fields();
    return Request(request_data, header_fields);
  }
  catch (HTTPError &http_err) {
    if (http_err.get_response_conde() != ResponseCode::RESOURCE_NOT_FOUND) {
      throw;
    }
    get_header_fields();
    throw;
  }
}

void Server::_respond(ResponseCode status_code,
                      std::string const &reason,
                      HeaderFields const &header_fields,
                      std::string const &message) {
  std::stringstream &out_stream = data_stream.get_out_stream();
  out_stream << "HTTP/1.1 " << static_cast<int>(status_code) << " " << reason << EOL;

  out_stream << header_fields << EOL;

  out_stream << message;

  data_stream.send_response();

  if (header_fields.is_field_set_to("Connection", "close")) {
    should_close = true;
  }
//  out_stream << message;
}

void Server::close_connection() {
  if (close(msg_socket) < 0) {
    throw std::runtime_error("Error while closing server!");
  }

  data_stream.clear();
  connected = false;

}

HeaderFields Server::get_header_fields() {
  HeaderFields header_fields;

  std::string header_line = data_stream.read_line();

  while (!header_line.empty()) {
    header_fields.add_field_from_header_line(header_line);

    header_line = data_stream.read_line();
  }

  validate_headers(header_fields);

  return header_fields;
}

void DataStream::send_response() {
  std::string response = out_stream.str();

  char const *c_str = response.c_str();
  ssize_t written = 0;
  ssize_t to_send = response.length();

  ssize_t curr_write;

  if (signal(SIGPIPE, SIG_IGN) == SIG_ERR) {
    throw std::runtime_error("Error while calling signal function!");
  }

  while (to_send > 0) {
    if ((curr_write = write(fd, c_str + written, to_send)) < 0) {
      if (errno == EPIPE || errno == ECONNRESET) {
        throw ClientDisconnected();
      }

      throw std::runtime_error("Error while sending the response!");
    }
    written += curr_write;
    to_send -= curr_write;
  }

  clear_out_buffer();
}

void DataStream::clear() {
  eof = false;
  buffer.clear();

  clear_out_buffer();
}

std::string DataStream::read_line() {
  size_t pos;

  while ((pos = find_eol(buffer)) == std::string::npos) {
    size_t bytes_read = read_to_buffer();

    if (bytes_read == 0) {
      eof = true;
      if (buffer.empty()) {
        return "";
      } else {
        throw HTTPError(ResponseCode::INCORRECT_REQUEST, "Malformed request");
      }
    }

  }

  std::string res = buffer.substr(0, pos);

  buffer = buffer.substr(pos + EOL.length());

  return res;
}

ssize_t DataStream::read_to_buffer() {
  ssize_t read_count;

  memset(buffer_helper, 0, BUFFER_SIZE);

  if (signal(SIGPIPE, SIG_IGN) == SIG_ERR) {
    throw std::runtime_error("Error while calling signal function!");
  }

  if ((read_count = read(fd, (void *) buffer_helper, BUFFER_SIZE - 1)) < 0) {
    if (errno == EPIPE || errno == ECONNRESET) {
      throw ClientDisconnected();
    }
    throw std::runtime_error("Error while calling read function!");
  }

  buffer.append(buffer_helper, read_count);

  return read_count;
}

void Request::RequestData::validate_target(const std::string &token) {
  if (token[0] != '/') {
    throw HTTPError(ResponseCode::INCORRECT_REQUEST, "Invalid target");
  }

  std::regex re(R"(\/[a-zA-Z0-9.\-\/]*)");
  std::smatch m;

  if (!std::regex_match(token, m, re)) {
    throw HTTPError(ResponseCode::RESOURCE_NOT_FOUND, "Not found");
  }
}

Request::RequestData Request::RequestData::parse_request_data(std::string const &request_line) {
  std::smatch match = get_request_line_parts(request_line);

  if (match.ready() && match.empty()) {
    throw HTTPError(ResponseCode::INCORRECT_REQUEST, "Wrong request");
  }

  std::string method = match[1];
  std::string target = match[2];
  std::string http_version = match[3];

  validate_method(method);
  validate_target(target);
  validate_http_version(http_version);

  return {.method=method, .target=target, .http_version=http_version};
}

std::ostream &operator<<(std::ostream &os, const Request::RequestData &request) {
  os << "Method: " << request.method << std::endl;
  os << "Target: " << request.target << std::endl;
  os << "HTTP version: " << request.http_version << std::endl;

  return os;
}
