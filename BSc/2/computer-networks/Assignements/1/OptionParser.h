#ifndef INC_1__OPTIONPARSER_H_
#define INC_1__OPTIONPARSER_H_

#include <stdexcept>
#include <filesystem>

#include "Server.h"

class OptionParser {
  static constexpr Server::port_t DEFAULT_PORT = 8080;

  std::filesystem::path files_directory;
  std::filesystem::path correlated_servers_file_path;
  Server::port_t port;
 public:
  OptionParser(int argc, char **argv) {
    if (argc < 3) {
      throw std::invalid_argument(
          "serwer <nazwa-katalogu-z-plikami> <plik-z-serwerami-skorelowanymi> [<numer-portu-serwera>]");
    }
    files_directory = std::filesystem::canonical(argv[1]);
    correlated_servers_file_path = std::filesystem::canonical(argv[2]);
    if (argc == 3) {
      port = DEFAULT_PORT;
    } else {
      port = strtol(argv[3], nullptr, 10);
    }
  }

  [[nodiscard]] std::string getFilesDirectory() const {
    return files_directory;
  }

  [[nodiscard]] std::string getCorrelatedServersFilePath() const {
    return correlated_servers_file_path;
  }

  [[nodiscard]] Server::port_t getPort() const {
    return port;
  }

};

#endif //INC_1__OPTIONPARSER_H_
