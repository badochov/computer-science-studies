#include "Client.h"
#include "util.h"
#include <iostream>

int main(int argc, char **argv) {
  OptionParser option_parser(argc, argv);

  std::cout << "server: " << option_parser.get_game_server() << std::endl;
  std::cout << "server_port: " << option_parser.get_server_port() << std::endl;
  std::cout << "gui: " << option_parser.get_gui_server() << std::endl;
  std::cout << "gui_port: " << option_parser.get_gui_port() << std::endl;
  std::cout << "player: " << option_parser.get_player_name() << std::endl;

  Client client(option_parser);

  client.run();

  return 0;
}
