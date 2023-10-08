#include "Server.h"
#include "util.h"
#include <iostream>

int main(int argc, char **argv) {
  OptionParser option_parser(argc, argv);

  std::cout << "port: " << option_parser.get_server_port() << std::endl;
  std::cout << "seed: " << option_parser.get_seed() << std::endl;
  std::cout << "turning_speed: " << option_parser.get_turning_speed() << std::endl;
  std::cout << "round_per_second: " << option_parser.get_rounds_per_second() << std::endl;
  std::cout << "width: " << option_parser.get_width() << std::endl;
  std::cout << "height: " << option_parser.get_height() << std::endl;

  Server server(option_parser);

  server.run();

  return 0;
}
