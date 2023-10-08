#ifndef INC_2_CLIENT_UTIL_H_
#define INC_2_CLIENT_UTIL_H_

#include "../Util/util.h"
#include "arpa/inet.h"

#include <iostream>
#include <sstream>
#include <stdexcept>
#include <string>

class OptionParser {
  options::ParsedOptions parsed_options;
  static std::string get_game_server(int argc, char **argv) {
    if (argc < 2) {
      throw std::runtime_error("Wrong invocation!");
    }

    return argv[1];
  }

  std::string const &get_or_default(std::string const &key, std::string const &def) const {
    if (parsed_options.count(key) > 0) {
      return parsed_options.at(key);
    }

    return def;
  }

public:
  OptionParser(int argc, char **argv) {
    std::string server = get_game_server(argc, argv);

    options::Options options;
    options['n'] = "player_name";
    options['p'] = "server_port";
    options['i'] = "gui_server";
    options['r'] = "gui_port";

    parsed_options = options::parse(argc, argv, options, 2);

    parsed_options["game_server"] = server;
  }

  uint16_t get_server_port() const {
    return string_to_port(get_or_default("server_port", "2021"));
  }

  uint16_t get_gui_port() const {
    return string_to_port(get_or_default("gui_port", "20210"));
  }

  std::string get_player_name() const {
    std::string player_name = get_or_default("player_name", "");
    if (!valid_username(player_name) && !player_name.empty()) {
      error_exit("Unsupported player name");
    }
    return player_name;
  }

  std::string get_game_server() const {
    return parsed_options.at("game_server");
  }

  std::string get_gui_server() const {
    return get_or_default("gui_server", "localhost");
  }
};

#endif // INC_2_CLIENT_UTIL_H_
