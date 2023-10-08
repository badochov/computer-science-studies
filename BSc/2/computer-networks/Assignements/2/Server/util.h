#ifndef INC_2_SERVER_UTIL_H_
#define INC_2_SERVER_UTIL_H_

#ifndef INC_2_CLIENT_UTIL_H_
#define INC_2_CLIENT_UTIL_H_

#include "../Util/util.h"
#include <arpa/inet.h>
#include <ctime>

#include <cmath>
#include <iostream>
#include <stdexcept>
#include <utility>
#include <vector>

using pixel_t = std::pair<int64_t, int64_t>;

struct Worm {
  double x = 0.;
  double y = 0.;
  uint32_t direction = 0;
  uint8_t turn_direction = 0;
  bool is_dead = true;

  uint8_t number = 0;

  [[nodiscard]] pixel_t get_current_pixel() const {
    return {std::floor(x), std::floor(y)};
  }

  void move() {
    x += std::cos(to_radians(direction));
    y += std::sin(to_radians(direction));
  }

  static double to_radians(double degrees) {
    return degrees * (M_PI / 180);
  }
};

struct PlayerData {
  uint8_t number;

  bool ready;

  uint64_t session_id;
  uint8_t turn_direction;
  bool has_worm;

  types::address_t address;

  std::string player_name;

  types::timestamp_t last_message_timestamp;

  PlayerData(communication::client_message_t const &client_message,
             types::address_t const &_address)
      : number(0), ready(false), session_id(client_message.session_id),
        turn_direction(client_message.turn_direction), has_worm(false), address(_address),
        player_name(client_message.player_name), last_message_timestamp(get_current_timestamp()){};

  [[nodiscard]] bool correct_session_id(uint64_t received_session_id) const {
    return received_session_id == session_id;
  }
};

struct SpectatorData {
  types::address_t address;
  types::timestamp_t last_message_timestamp;
};

class Board {
public:
  uint32_t width;
  uint32_t height;

  Board(uint32_t _height, uint32_t _width)
      : width(_width), height(_height), board(width, std::vector<uint8_t>(height, 0)){};

  [[nodiscard]] bool is_in(pixel_t const &pixel) const {
    return pixel.first < width && pixel.first >= 0 && pixel.second < height && pixel.second >= 0;
  }

  [[nodiscard]] bool is_taken(pixel_t const &pixel) const {
    return board[pixel.first][pixel.second] != 0;
  }

  void place(pixel_t const &pixel, uint8_t player) {
    board[pixel.first][pixel.second] = player + 1;
  }

  void clear() {
    board = std::vector<std::vector<uint8_t>>(width, std::vector<uint8_t>(height, 0));
  }

private:
  std::vector<std::vector<uint8_t>> board;
};

class OptionParser {
  static constexpr uint32_t MAX_DIMENSION = 5000;
  static constexpr uint32_t MAX_TURNING_SPEED = 359;
  static constexpr uint32_t MAX_ROUNDS_PER_SECOND = 250;
  options::ParsedOptions parsed_options;
  static uint32_t string_to_u32(std::string const &s) {
    errno = 0;
    char *end_ptr;

    long long num = strtoll(s.c_str(), &end_ptr, 10);
    if (errno || end_ptr != s.c_str() + s.size() || num < 0 ||
        num > std::numeric_limits<uint32_t>::max()) {
      error_exit("Number out of uin32_t range!");
    }

    return num;
  }

  std::string const &get_or_default(std::string const &key, std::string const &def) const {
    if (parsed_options.count(key) > 0) {
      return parsed_options.at(key);
    }

    return def;
  }

public:
  OptionParser(int argc, char **argv) {
    options::Options options;
    options['p'] = "port";
    options['s'] = "seed";
    options['t'] = "turning_speed";
    options['v'] = "round_per_second";
    options['w'] = "width";
    options['h'] = "height";

    parsed_options = options::parse(argc, argv, options, 1);
  }

  uint16_t get_server_port() const {
    return string_to_port(get_or_default("port", "2021"));
  }

  uint32_t get_seed() const {
    if (parsed_options.count("seed") > 0) {
      return string_to_u32(parsed_options.at("seed"));
    }
    return time(nullptr);
  }

  uint32_t get_turning_speed() const {
    uint32_t turning_speed = string_to_u32(get_or_default("turning_speed", "6"));
    if (turning_speed > MAX_TURNING_SPEED) {
      error_exit("Turning speed should be in [0, 360)");
    }
    return turning_speed;
  }

  uint32_t get_rounds_per_second() const {
    uint32_t rounds_per_second = string_to_u32(get_or_default("round_per_second", "50"));
    if (rounds_per_second == 0 || rounds_per_second > MAX_ROUNDS_PER_SECOND) {
      error_exit("Rounds per seconds to high or 0");
    }
    return rounds_per_second;
  }

  uint32_t get_width() const {
    uint32_t width = string_to_u32(get_or_default("width", "640"));
    if (width == 0 || width > MAX_DIMENSION) {
      error_exit("Incorrect width");
    }
    return width;
  }

  uint32_t get_height() const {
    uint32_t height = string_to_u32(get_or_default("height", "480"));
    if (height == 0 || height > MAX_DIMENSION) {
      error_exit("Incorrect height");
    }
    return height;
  }
};

class Random {
  uint64_t next_value;

  static constexpr uint64_t COEF = 279410273;
  static constexpr uint64_t MOD = 4294967291;

public:
  explicit Random(uint32_t seed) : next_value(seed) {
  }

  [[nodiscard]] uint32_t random_u32() {
    uint32_t ret = next_value;

    next_value = (next_value * COEF) % MOD;

    return ret;
  }
};

#endif // INC_2_CLIENT_UTIL_H_

#endif // INC_2_SERVER_UTIL_H_
