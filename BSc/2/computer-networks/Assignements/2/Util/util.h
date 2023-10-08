#ifndef INC_2_UTIL_UTIL_H_
#define INC_2_UTIL_UTIL_H_

#include <arpa/inet.h>
#include <functional>
#include <poll.h>
#include <string>
#include <unordered_map>
#include <utility>
#include <variant>
#include <vector>

#include <fstream>
#include <iostream>

namespace helpers {
  constexpr uint64_t x(int a) {
    uint64_t one = 1;
    return one << a;
  }
} // namespace helpers

namespace types {
  using timestamp_t = uint64_t;
  using address_t = struct sockaddr_storage;
} // namespace types

namespace constants {
  constexpr unsigned int MAX_PLAYER_NAME_LENGTH = 20;
  constexpr unsigned int MAX_PLAYERS = 25;

  constexpr unsigned int MAX_SERVER_MESSAGE_LENGTH = 550;

  constexpr uint64_t CRC32_DIVISOR = 0xEDB88320;

  constexpr uint32_t U64_SIZE = sizeof(uint64_t);
  constexpr uint32_t U32_SIZE = sizeof(uint32_t);
  constexpr uint32_t CHAR_SIZE = sizeof(char);
  constexpr uint32_t REQUIRED_PLAYERS = 2;
} // namespace constants

bool valid_username(std::string const &username);

void error_exit(std::string const &reason);

[[nodiscard]] uint16_t string_to_port(std::string const &port_str);

uint32_t network_data_to_u32(unsigned char const *c_str);

uint64_t network_data_to_u64(unsigned char const *c_str);

std::string u32_to_network_data(uint32_t u32);
std::string u64_to_network_data(uint64_t u64);

types::timestamp_t get_current_timestamp();

bool same_address(types::address_t const &address1, types::address_t const &address2);

namespace options {
  using Options = std::unordered_map<char, std::string>; // Option option character: name.
  using ParsedOptions =
      std::unordered_map<std::string, std::string>; // Option name : option character.

  ParsedOptions parse(int argc, char **argv, Options const &options, int start);
} // namespace options

namespace communication {

  enum DIRECTION { FORWARD, RIGHT, LEFT };

  enum EventType { NEW_GAME, PIXEL, PLAYER_ELIMINATED, GAME_OVER };

  constexpr size_t EVENT_TYPE_COUNT = 4;

  struct new_game_data_t {
    static constexpr uint32_t MAX_PLAYER_NAMES_SIZE =
        constants::MAX_PLAYERS * (constants::MAX_PLAYER_NAME_LENGTH + 1) - 1;

    static constexpr uint32_t BASE_SIZE = constants::U32_SIZE + constants::U32_SIZE;

    uint32_t maxx;
    uint32_t maxy;
    char player_names[MAX_PLAYER_NAMES_SIZE + 1];
    std::string::size_type players_len;

    [[nodiscard]] std::string to_string() const;
    [[nodiscard]] uint32_t size() const;
  };

  struct pixel_data_t {
    uint8_t player_number;
    uint32_t x;
    uint32_t y;

    [[nodiscard]] std::string to_string() const;
    static uint32_t size();
  };

  struct event_t {
    static constexpr uint32_t BASE_EVENT_LENGTH = constants::U32_SIZE + constants::CHAR_SIZE;

    uint32_t len = BASE_EVENT_LENGTH;
    uint32_t event_no;
    uint8_t event_type;
    struct event_data_t {
      union {
        new_game_data_t new_game_data;
        pixel_data_t pixel_data;
        uint8_t player_number;
      };
      [[nodiscard]] std::string to_string(uint32_t type) const;
    } event_data;
    uint32_t crc32;

    [[nodiscard]] std::string to_string() const;
  };

  struct server_message_t {
    uint32_t game_id;
    std::vector<event_t> events;
  };

  struct client_message_t {
    uint64_t session_id;
    uint8_t turn_direction;
    uint32_t next_expected_event_no;
    std::string player_name;

    [[nodiscard]] std::string to_string() const;

    static constexpr uint32_t BASE_SIZE =
        constants::U64_SIZE + constants::CHAR_SIZE + constants::U32_SIZE;

    static constexpr uint32_t MAX_SIZE = BASE_SIZE + constants::MAX_PLAYER_NAME_LENGTH;
  };

  uint32_t calculate_crc32(unsigned char const *data, size_t len, uint64_t reversed_key);

  uint32_t calculate_crc32(unsigned char const *data, size_t len);

  bool check_crc32(unsigned char const *data, size_t len, uint32_t crc);

  class Connection {
  public:
    using callback_t = std::function<void(struct pollfd &)>;

    void poll();

    void add_fd(struct pollfd const &poll_fd, callback_t const &callback);

    struct pollfd &get_poll_fd_by_fd(int fd);

  private:
    std::vector<struct pollfd> fds_vector;
    std::vector<callback_t> callbacks;

    struct pollfd *get_fds();
  };
} // namespace communication

std::ostream &operator<<(std::ostream &os, communication::client_message_t const &client_message);

#endif // INC_2_UTIL_UTIL_H_
