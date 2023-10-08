#ifndef INC_1__ERROR_H_
#define INC_1__ERROR_H_

#include "utils.h"

class HTTPError : public std::exception {
  ResponseCode response_code;
  std::string reason;
 public:
  HTTPError(ResponseCode _response_code, std::string _reason)
      : response_code(_response_code), reason(std::move(_reason)) {}

  [[nodiscard]] char const *what() const noexcept override {
    return reason.c_str();
  }

  [[nodiscard]] ResponseCode get_response_conde() const {
    return response_code;
  }
  [[nodiscard]] std::string const &get_reason() const {
    return reason;
  }
};

class ClientDisconnected : public std::exception {
 public:

  [[nodiscard]] char const *what() const noexcept override {
    return "Client has disconnected!";
  }
};

#endif //INC_1__ERROR_H_
