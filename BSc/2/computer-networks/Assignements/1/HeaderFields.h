#ifndef INC_1_HEADERFIELDS_H
#define INC_1_HEADERFIELDS_H

#include "utils.h"
#include "FileGetter.h"
#include "error.h"

#include <unordered_map>
#include <string>
#include <regex>
#include <algorithm>

class HeaderFields {
  static constexpr size_t HEADER_LINE_MAX_LENGTH = 10000;
 public:
  using header_fields_map_t = std::unordered_map<std::string, std::string>;

  void add(std::string const &key, std::string const &value) {
    fields_map[to_lower(key)] = value;
  }

  std::string operator[](std::string const &key) const {
    return fields_map.at(to_lower(key));
  }

  [[nodiscard]] bool is_field_set_to(std::string const &key, std::string const &value) const {
    return has(key) && (*this)[key] == value;
  }

  bool has(std::string const &key) const {
    return fields_map.count(to_lower(key));
  }

  void add_field_from_header_line(std::string const &header_line) {
    if (header_line.size() > HEADER_LINE_MAX_LENGTH) {
      throw HTTPError(ResponseCode::INCORRECT_REQUEST, "Header too long");
    }
    std::smatch match = match_header_line(header_line);
    if (match.ready() && match.empty()) {
      throw HTTPError(ResponseCode::INCORRECT_REQUEST, "Malformed header");
    }
    std::string key = match[1];

    if (has(key)) {
      throw HTTPError(ResponseCode::INCORRECT_REQUEST, "Malformed header");
    }

    add(key, match[2]);

  }

  friend std::ostream &operator<<(std::ostream &os, HeaderFields const &header_fields) {
    for (auto const &[key, value] : header_fields.fields_map) {
      os << key << ": " << value << EOL;
    }

    return os;
  }

  void add_file_data(File const &file) {
    add("Content-length", std::to_string(file.file_size));
    add("Content-Type", file.mime_type);
  }

  void erase(std::string const &key) {
    fields_map.erase(key);
  }

 private:
  static std::string to_lower(std::string s) {
    std::transform(s.begin(), s.end(), s.begin(), [](unsigned char c) { return std::tolower(c); });
    return s;
  }

  static std::smatch match_header_line(std::string const &header_line) {

    std::regex re("([a-zA-Z0-9_-]+): *(.*?) *");
    std::smatch match;

    std::regex_match(header_line, match, re);

    return match;
  }

  header_fields_map_t fields_map;
};

#endif //INC_1_HEADERFIELDS_H
