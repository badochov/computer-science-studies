#ifndef INC_1__UTILS_H_
#define INC_1__UTILS_H_

#include <string>
#include <exception>
#include <utility>

static const std::string EOL = "\r\n";

enum class ResponseCode {
  OK = 200, REDIRECT = 302, INCORRECT_REQUEST = 400, RESOURCE_NOT_FOUND = 404, SERVER_ERROR = 500, NOT_IMPLEMENTED = 501
};

#endif //INC_1__UTILS_H_
