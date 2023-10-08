#include "util.h"

#include <arpa/inet.h>
#include <cerrno>
#include <chrono>
#include <cstring>
#include <iostream>
#include <limits>
#include <sstream>
#include <stdexcept>
#include <unistd.h>
// https://stackoverflow.com/a/4410728
#if defined(__linux__)
#include <endian.h>
#elif defined(__FreeBSD__) || defined(__NetBSD__)
#include <sys/endian.h>
#elif defined(__OpenBSD__)
#include <sys/types.h>
#define be16toh(x) betoh16(x)
#define be32toh(x) betoh32(x)
#define be64toh(x) betoh64(x)
#endif

namespace helpers {
  uint64_t network_data_to_u_helper(unsigned char const *c_str, size_t len) {
    uint64_t res = 0;

    for (size_t i = 0; i < len; ++i) {
      res <<= 8;
      res += c_str[i];
    }

    return res;
  }
} // namespace helpers

void error_exit(std::string const &reason) {
  std::cerr << reason;
  exit(EXIT_FAILURE);
}

uint32_t network_data_to_u32(unsigned char const *c_str) {
  return helpers::network_data_to_u_helper(c_str, 4);
}

uint64_t network_data_to_u64(unsigned char const *c_str) {
  return helpers::network_data_to_u_helper(c_str, 8);
}

uint16_t string_to_port(const std::string &port_str) {
  errno = 0;
  char *end_ptr;
  long port = strtol(port_str.c_str(), &end_ptr, 10);

  if (errno || end_ptr != port_str.c_str() + port_str.size() ||
      port > std::numeric_limits<uint16_t>::max() || port < 0) {
    error_exit("Invalid Port!");
  }
  return port;
}

types::timestamp_t get_current_timestamp() {
  std::chrono::time_point<std::chrono::system_clock> now = std::chrono::system_clock::now();
  auto duration = now.time_since_epoch();
  return std::chrono::duration_cast<std::chrono::microseconds>(duration).count();
}

bool same_address(types::address_t const &address1, types::address_t const &address2) {
  if (address1.ss_family != address2.ss_family) {
    return false;
  }
  if (address1.ss_family == AF_INET) {
    sockaddr_in sock_addr1{};
    sockaddr_in sock_addr2{};

    std::memcpy(&sock_addr1, &address1, sizeof(sock_addr1));
    std::memcpy(&sock_addr2, &address2, sizeof(sock_addr2));

    return sock_addr1.sin_port == sock_addr2.sin_port &&
           sock_addr1.sin_addr.s_addr == sock_addr2.sin_addr.s_addr;
  }

  sockaddr_in6 sock_addr1{};
  sockaddr_in6 sock_addr2{};

  std::memcpy(&sock_addr1, &address1, sizeof(sock_addr1));
  std::memcpy(&sock_addr2, &address2, sizeof(sock_addr2));

  return sock_addr1.sin6_port == sock_addr2.sin6_port &&
         std::memcmp(&sock_addr1.sin6_addr, &sock_addr2.sin6_addr, sizeof(sock_addr1.sin6_addr)) == 0;
}

std::string u32_to_network_data(uint32_t u32) {
  uint32_t network_u32 = htonl(u32);

  return std::string(reinterpret_cast<const char *>(&network_u32), constants::U32_SIZE);
}

std::string u64_to_network_data(uint64_t u64) {
  uint64_t network_u64 = htobe64(u64);

  return std::string(reinterpret_cast<const char *>(&network_u64), constants::U64_SIZE);
}

bool is_valid_username_character(unsigned char c) {
  return c > 32 && c < 127;
}

bool valid_username(std::string const &username) {
  if (username.empty() || username.size() > constants::MAX_PLAYER_NAME_LENGTH) {
    return false;
  }
  return std::all_of(username.begin(), username.end(), is_valid_username_character);
}

namespace options {

  static std::string get_options_string(Options const &options) {
    std::stringstream ss;
    ss << '+';
    for (auto const &option : options) {
      ss << option.first << ':';
    }

    return ss.str();
  }

  ParsedOptions parse(int argc, char **argv, Options const &options, int start) {
    ParsedOptions parsed_options;

    std::string options_string = get_options_string(options);
    optind = start;
    int opt;
    while ((opt = getopt(argc, argv, options_string.c_str())) != -1) {
      if (opt == '?') {
        error_exit("Wrong invocation!");
      }
      if (options.count(opt) > 0) {
        parsed_options[options.at(opt)] = optarg;
      } else {
        error_exit("Wrnog invocation!");
      }
    }

    if (optind != argc) {
      error_exit("Wrong invocation!");
    }

    return parsed_options;
  }
} // namespace options

std::string communication::client_message_t::to_string() const {
  char res[communication::client_message_t::MAX_SIZE + 1] = {0};
  char *ptr = res;

  uint64_t session_id_host = htobe64(session_id);
  uint32_t next_expected_event_no_host = htonl(next_expected_event_no);

  memcpy(ptr, &session_id_host, constants::U64_SIZE);
  ptr += constants::U64_SIZE;
  *ptr = turn_direction;
  ptr += constants::CHAR_SIZE;
  memcpy(ptr, &next_expected_event_no_host, constants::U32_SIZE);
  ptr += constants::U32_SIZE;
  strncpy(ptr, player_name.c_str(), constants::MAX_PLAYER_NAME_LENGTH);

  return std::string(res, BASE_SIZE + player_name.size());
}

std::ostream &operator<<(std::ostream &os, communication::client_message_t const &client_message) {
  os << "Session id:" << client_message.session_id << std::endl
     << "Turn direction:" << int(client_message.turn_direction) << std::endl
     << "Next expected event no:" << client_message.next_expected_event_no << std::endl
     << "Player name:" << client_message.player_name << std::endl;

  return os;
}

std::string communication::event_t::to_string() const {
  std::stringstream res;

  res << u32_to_network_data(len) << u32_to_network_data(event_no) << event_type
      << event_data.to_string(event_type);

  std::string data = res.str();

  data += u32_to_network_data(
      calculate_crc32(reinterpret_cast<const unsigned char *>(data.c_str()), data.size()));

  return data;
}
uint32_t communication::calculate_crc32(const unsigned char *data, size_t len,
                                        uint64_t reversed_key) {
  // Source: https://medium.com/@nxtchg/crc-32-without-lookup-tables-1682405aa048
  uint r = ~0;
  unsigned char const *end = data + len;

  while (data < end) {
    r ^= *data++;

    for (int i = 0; i < 8; i++) {
      uint t = ~((r & 1) - 1);
      r = (r >> 1) ^ (reversed_key & t);
    }
  }

  return ~r;
}

uint32_t communication::calculate_crc32(const unsigned char *data, size_t len) {
  return calculate_crc32(data, len, constants::CRC32_DIVISOR);
}

bool communication::check_crc32(const unsigned char *data, size_t len, uint32_t crc) {
  return calculate_crc32(data, len) == crc;
}

std::string communication::event_t::event_data_t::to_string(uint32_t type) const {
  switch (type) {
    case communication::NEW_GAME:
      return new_game_data.to_string();
    case communication::PIXEL:
      return pixel_data.to_string();
    case communication::PLAYER_ELIMINATED:
      return std::string(1, player_number);
    case communication::GAME_OVER:
      return "";
    default:
      throw std::runtime_error("Unexpected event type!");
  }
}

std::string communication::new_game_data_t::to_string() const {
  std::stringstream ss;

  ss << u32_to_network_data(maxx) << u32_to_network_data(maxy)
     << std::string(player_names, players_len);

  return ss.str();
}
uint32_t communication::new_game_data_t::size() const {
  return BASE_SIZE + players_len;
}

std::string communication::pixel_data_t::to_string() const {
  std::stringstream ss;

  ss << player_number << u32_to_network_data(x) << u32_to_network_data(y);

  return ss.str();
}

uint32_t communication::pixel_data_t::size() {
  return constants::U32_SIZE + constants::U32_SIZE + constants::CHAR_SIZE;
}

struct pollfd *communication::Connection::get_fds() {
  return fds_vector.data();
}

void communication::Connection::poll() {
  int ready = ::poll(get_fds(), fds_vector.size(), -1);
  if (ready < 0) {
    error_exit("Poll failure");
  }

  for (std::vector<struct pollfd>::size_type i = 0; i < fds_vector.size(); i++) {
    if (fds_vector[i].revents != 0) {
      callbacks[i](fds_vector[i]);
    }
  }
}

void communication::Connection::add_fd(struct pollfd const &poll_fd,
                                       communication::Connection::callback_t const &callback) {
  fds_vector.push_back(poll_fd);
  callbacks.push_back(callback);
}

struct pollfd &communication::Connection::get_poll_fd_by_fd(int fd) {
  for (auto &poll_fd : fds_vector) {
    if (poll_fd.fd == fd) {
      return poll_fd;
    }
  }
  throw std::runtime_error("Fd not in list");
}
