#ifndef INC_1__FILEGETTER_H_
#define INC_1__FILEGETTER_H_

#include "error.h"

#include <iostream>
#include <filesystem>
#include <utility>
#include <stdexcept>
#include <exception>
#include <fstream>

struct File {
  const std::string path;
  const size_t file_size;
  const std::string file_content;
  const std::string mime_type = "application/octet-stream";

  File(std::string _path, std::string _content)
      : path(std::move(_path)), file_size(_content.size()), file_content(std::move(_content)) {
  }
};

class FileNotFoundException : public std::exception {
 public:
  const std::string path;

  explicit FileNotFoundException(std::string _path) : path(std::move(_path)) {

  }

  [[nodiscard]] const char *what() const noexcept override {
    return "FileNotFound";
  }
};

class FileGetter {
  std::string dir;
 public:
  explicit FileGetter(std::string _dir) : dir(std::move(_dir)) {
    if (!std::filesystem::is_directory(dir)) {
      throw std::invalid_argument("Pierwszy argument musi być katalogiem!");
    }
  }

  [[nodiscard]] File get_file(std::string const &file_path) const {
    std::string full_path = dir + file_path;

    char *abs_path_c = canonicalize_file_name(full_path.c_str());

    if (abs_path_c == nullptr) {
      throw FileNotFoundException(file_path);
    }

    std::string abs_path = abs_path_c;
    free(abs_path_c);

    if (abs_path.find(dir + "/") != 0) {
      throw HTTPError(ResponseCode::RESOURCE_NOT_FOUND, "Not Found");
    }

    std::ifstream file_stream(abs_path);
    if (std::filesystem::is_directory(abs_path) || file_stream.fail()) {
      throw FileNotFoundException(file_path);
    }

    std::stringstream ss;
    ss << file_stream.rdbuf();

    std::string file_content = ss.str();

    return File(file_path, file_content);
  }

};

#endif //INC_1__FILEGETTER_H_
