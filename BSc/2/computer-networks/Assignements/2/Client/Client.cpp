#include "Client.h"
#include "../Util/util.h"

#include <cstring>
#include <fcntl.h>
#include <functional>
#include <netdb.h>
#include <netinet/in.h>
#include <sstream>
#include <sys/socket.h>
#include <sys/timerfd.h>
#include <unistd.h>

#define MSG (msg + byte)

namespace server {

  void check_size(size_t a, size_t b) {
    if (a > b) {
      error_exit("Malformed data");
    }
  }

  void increase_byte(size_t &byte, size_t amount, size_t len) {
    byte += amount;
    check_size(byte, len);
  }

  communication::event_t::event_data_t parse_new_game_data(const unsigned char *msg, size_t &byte,
                                                           communication::event_t const &event,
                                                           size_t len) {
    communication::event_t::event_data_t event_data{};

    event_data.new_game_data.maxx = network_data_to_u32(MSG);
    increase_byte(byte, 4, len);
    event_data.new_game_data.maxy = network_data_to_u32(MSG);
    increase_byte(byte, 4, len);

    event_data.new_game_data.players_len = event.len - communication::event_t::BASE_EVENT_LENGTH -
                                           communication::new_game_data_t::BASE_SIZE;

    check_size(byte + event_data.new_game_data.players_len, len);

    std::memcpy(event_data.new_game_data.player_names, reinterpret_cast<const char *>(MSG),
                event_data.new_game_data.players_len - 1);

    byte += event_data.new_game_data.players_len;

    return event_data;
  }
  communication::event_t::event_data_t parse_pixel_data(const unsigned char *msg, size_t &byte,
                                                        communication::event_t const &,
                                                        size_t len) {
    communication::event_t::event_data_t event_data{};

    event_data.pixel_data.player_number = *MSG;
    increase_byte(byte, 1, len);
    event_data.pixel_data.x = network_data_to_u32(MSG);
    increase_byte(byte, 4, len);
    event_data.pixel_data.y = network_data_to_u32(MSG);
    increase_byte(byte, 4, len);

    return event_data;
  }
  communication::event_t::event_data_t parse_player_eliminated_data(const unsigned char *msg,
                                                                    size_t &byte,
                                                                    communication::event_t const &,
                                                                    size_t len) {
    communication::event_t::event_data_t event_data{};

    event_data.player_number = *MSG;
    increase_byte(byte, 1, len);

    return event_data;
  }
  communication::event_t::event_data_t
  parse_game_over_data(const unsigned char *, size_t &, communication::event_t const &, size_t) {
    return communication::event_t::event_data_t{};
  }
} // namespace server

std::string Client::_handle_new_game_message(communication::event_t const &event) const {
  std::stringstream ss;
  ss << "NEW_GAME " << event.event_data.new_game_data.maxx << " "
     << event.event_data.new_game_data.maxy << " " << get_player_names_string() << '\n';

  return ss.str();
}
std::string Client::_handle_pixel_message(communication::event_t const &event) const {
  std::string player = player_number_to_name(event.event_data.pixel_data.player_number);

  std::stringstream ss;
  ss << "PIXEL " << event.event_data.pixel_data.x << " " << event.event_data.pixel_data.y << " "
     << player_number_to_name(event.event_data.pixel_data.player_number) << '\n';

  return ss.str();
}
std::string Client::_handle_player_eliminated_message(communication::event_t const &event) const {
  std::stringstream ss;
  ss << "PLAYER_ELIMINATED " << player_number_to_name(event.event_data.player_number) << '\n';

  return ss.str();
}

Client::GUI_MESSAGE Client::create_message_from_gui_type(std::string const &message) {
  if (message == "LEFT_KEY_DOWN") {
    return GUI_MESSAGE::LEFT_KEY_DOWN;
  }
  if (message == "LEFT_KEY_UP") {
    return GUI_MESSAGE::LEFT_KEY_UP;
  }
  if (message == "RIGHT_KEY_DOWN") {
    return GUI_MESSAGE::RIGHT_KEY_DOWN;
  }
  if (message == "RIGHT_KEY_UP") {
    return GUI_MESSAGE::RIGHT_KEY_UP;
  }
  std::cerr << "Unknown GUI message: '" << message << "'";

  return GUI_MESSAGE::UNKNOWN;
}

std::string Client::create_message_for_gui(communication::event_t const &event) const {
  switch (event.event_type) {
    case communication::NEW_GAME:
      return _handle_new_game_message(event);
    case communication::PIXEL:
      return _handle_pixel_message(event);
    case communication::PLAYER_ELIMINATED:
      return _handle_player_eliminated_message(event);
    default:
      throw std::runtime_error("Unexpected event type!");
  }
}

std::string const &Client::player_number_to_name(uint8_t player_number) const {
  if (player_number >= player_names.size()) {
    error_exit("Player out of list");
  }
  return player_names.at(player_number);
}

communication::server_message_t Client::parse_server_message(const std::string &message) {
  communication::server_message_t res;

  auto msg = reinterpret_cast<unsigned char const *>(message.c_str());
  size_t len = message.size();

  size_t byte = 0;

  res.game_id = network_data_to_u32(MSG);
  server::increase_byte(byte, 4, len);

  while (byte != len) {
    unsigned char const *start = MSG;
    communication::event_t event{};

    event.len = network_data_to_u32(MSG);

    server::increase_byte(byte, 4, len);

    event.event_no = network_data_to_u32(MSG);

    server::increase_byte(byte, 4, len);

    event.event_type = *MSG;

    server::increase_byte(byte, 1, len);

    event.event_data = parse_event_data(msg, byte, event.event_type, event, len);

    event.crc32 = network_data_to_u32(MSG);

    server::increase_byte(byte, 4, len);

    if (!communication::check_crc32(start, event.len + constants::U32_SIZE, event.crc32)) {
      std::cerr << "WRONG CRC" << std::endl;
      break;
    }
    //    debug << event.event_data.to_string(event.event_type) << std::endl;

    if ((event.event_no == 0) != (event.event_type == communication::EventType::NEW_GAME)) {
      error_exit("New game should be event_no 0");
    }
    if (event.len < communication::event_t::BASE_EVENT_LENGTH) {
      error_exit("Event len to short");
    }

    res.events.push_back(event);
  }

  return res;
}

communication::event_t::event_data_t Client::parse_event_data(const unsigned char *msg,
                                                              size_t &byte, uint8_t event_type,
                                                              communication::event_t const &event,
                                                              size_t len) {
  switch (event_type) {
    case communication::NEW_GAME:
      return server::parse_new_game_data(msg, byte, event, len);
    case communication::PIXEL:
      return server::parse_pixel_data(msg, byte, event, len);
    case communication::PLAYER_ELIMINATED:
      return server::parse_player_eliminated_data(msg, byte, event, len);
    case communication::GAME_OVER:
      return server::parse_game_over_data(msg, byte, event, len);
    default:
      server::increase_byte(byte, event.len - communication::event_t::BASE_EVENT_LENGTH, len);
      return communication::event_t::event_data_t{};
  }
}

std::string Client::create_message_for_server() const {
  return data.to_string();
}

void Client::handle_raw_message_from_gui(std::string const &message) {
  GUI_MESSAGE gui_message_t = create_message_from_gui_type(message);

  handle_message_from_gui(gui_message_t);
}

void Client::handle_message_from_gui(Client::GUI_MESSAGE gui_message_t) {
  switch (gui_message_t) {
    case GUI_MESSAGE::LEFT_KEY_UP:
      if (data.turn_direction == communication::DIRECTION::LEFT) {
        data.turn_direction = communication::DIRECTION::FORWARD;
      }
      break;
    case GUI_MESSAGE::LEFT_KEY_DOWN:
      data.turn_direction = communication::DIRECTION::LEFT;
      break;
    case GUI_MESSAGE::RIGHT_KEY_UP:
      if (data.turn_direction == communication::DIRECTION::RIGHT) {
        data.turn_direction = communication::DIRECTION::FORWARD;
      }
      break;
    case GUI_MESSAGE::RIGHT_KEY_DOWN:
      data.turn_direction = communication::DIRECTION::RIGHT;
      break;
    case GUI_MESSAGE::UNKNOWN:
      break;
  }
}

void Client::handle_server_raw_message(std::string const &message) {
  communication::server_message_t server_message = parse_server_message(message);

  handle_server_message(server_message);
}

void Client::handle_server_message(communication::server_message_t const &server_message) {
  if (game_id != server_message.game_id) {
    communication::event_t first = server_message.events.front();
    if (first.event_type == communication::EventType::NEW_GAME) {
      handle_new_game(server_message);
    } else {
      return;
    }
  }
  for (communication::event_t const &event : server_message.events) {
    if (event.event_no == data.next_expected_event_no) {
      handle_event(event);

      ++data.next_expected_event_no;
    } else {
      std::cerr << "UNEXPECTED EVENT: " << event.event_no << std::endl;
    }
  }
}

void Client::handle_new_game(communication::server_message_t const &message) {
  data.next_expected_event_no = 0;
  game_id = message.game_id;
}

void Client::send_to_server() {
  std::string message = create_message_for_server();

  send_server_message(message);
}

[[noreturn]] void Client::run() {
  create_timer();
  connect_to_gui();
  connect_to_server();

  for (;;) {
    connection.poll();
  }
}

Client::Client(OptionParser const &option_parser) {
  game_id = 0;

  data.player_name = option_parser.get_player_name();
  data.turn_direction = communication::DIRECTION::FORWARD;
  data.next_expected_event_no = 0;
  data.session_id = get_current_timestamp();

  server_ip = option_parser.get_game_server();
  server_port = option_parser.get_server_port();

  gui_ip = option_parser.get_gui_server();
  gui_port = option_parser.get_gui_port();
}

void Client::send_server_message(std::string const &message) const {
  ssize_t written_bytes = write(server_fd, message.c_str(), message.size());
  if (written_bytes < 0 || std::string::size_type(written_bytes) != message.size()) {
    error_exit("Error writing data to server");
  }
}

void Client::create_server_socket() {
  struct addrinfo *addr_info;
  struct addrinfo addr_hints {};

  memset(&addr_hints, 0, sizeof(struct addrinfo));
  addr_hints.ai_flags = 0;
  addr_hints.ai_family = AF_UNSPEC;
  addr_hints.ai_socktype = SOCK_DGRAM;
  addr_hints.ai_protocol = IPPROTO_UDP;

  if (getaddrinfo(server_ip.c_str(), std::to_string(server_port).c_str(), &addr_hints, &addr_info) <
      0) {
    error_exit("Wrong server address");
  }

  server_fd = socket(addr_info->ai_family, SOCK_DGRAM, IPPROTO_UDP);
  if (server_fd < 0) {
    error_exit("Error creating server socket");
  }

  if (connect(server_fd, addr_info->ai_addr, addr_info->ai_addrlen) < 0) {
    error_exit("Error connecting to server");
  }
  freeaddrinfo(addr_info);
}
void Client::create_gui_socket() {
  struct addrinfo *addr_info;
  struct addrinfo addr_hints {};

  memset(&addr_hints, 0, sizeof(struct addrinfo));
  addr_hints.ai_flags = 0;
  addr_hints.ai_family = AF_UNSPEC;
  addr_hints.ai_socktype = SOCK_STREAM;
  addr_hints.ai_protocol = IPPROTO_TCP;

  if (getaddrinfo(gui_ip.c_str(), std::to_string(gui_port).c_str(), &addr_hints, &addr_info) < 0) {
    error_exit("Wrong gui address");
  }

  gui_fd = socket(addr_info->ai_family, SOCK_STREAM, IPPROTO_TCP);
  if (gui_fd < 0) {
    error_exit("Error creating gui socket");
  }

  if (connect(gui_fd, addr_info->ai_addr, addr_info->ai_addrlen) < 0) {
    error_exit("Error connecting to gui");
  }

  if (fcntl(gui_fd, F_SETFL, O_NONBLOCK) < 0) {
    error_exit("Error creating gui socket");
  }

  freeaddrinfo(addr_info);
}

void Client::connect_to_server() {
  create_server_socket();

  communication::Connection::callback_t callback = [this](struct pollfd &poll_fd) {
    poll_fd.revents = 0;

    char server_data[constants::MAX_SERVER_MESSAGE_LENGTH];
    ssize_t len;
    if ((len = read(server_fd, server_data, constants::MAX_SERVER_MESSAGE_LENGTH)) < 0) {
      error_exit("Error reading server data");
    }

    handle_server_raw_message(std::string(server_data, len));
  };

  connection.add_fd({.fd = server_fd, .events = POLLIN, .revents = 0}, callback);
}

void Client::connect_to_gui() {
  create_gui_socket();

  communication::Connection::callback_t callback = [this](struct pollfd &poll_fd) {
    if (poll_fd.revents & POLLOUT) {
      handle_gui_poll_out();
      poll_fd.revents &= ~POLLOUT;
    }
    if (poll_fd.revents & POLLHUP) {
      error_exit("GUI hung up");
    }
    if (poll_fd.revents & POLLIN) {
      handle_gui_poll_in();
      poll_fd.revents &= ~POLLIN;
    }
    poll_fd.revents = 0;
  };

  connection.add_fd({.fd = gui_fd, .events = POLLIN, .revents = 0}, callback);
}

void Client::send_gui_message(std::string const &message) {
  gui_buffer_out += message;

  get_gui_poll_fd().events |= POLLOUT;
}

void Client::create_timer() {
  int timer = timerfd_create(CLOCK_MONOTONIC, 0);
  struct timespec period {
    .tv_sec = 0, .tv_nsec = SERVER_MESSAGE_FREQUENCY_NS
  };

  struct itimerspec timer_period {
    .it_interval = period, .it_value = period
  };
  if (timer < 0 || timerfd_settime(timer, 0, &timer_period, nullptr)) {
    error_exit("Could create game clock");
  }

  communication::Connection::callback_t callback = [this](struct pollfd &poll_fd) {
    uint64_t expires;
    if (read(poll_fd.fd, &expires, constants::U64_SIZE) != constants::U64_SIZE) {
      error_exit("Game clock error");
    }
    send_to_server();
  };

  connection.add_fd({.fd = timer, .events = POLLIN, .revents = 0}, callback);
}

void Client::handle_gui_poll_in() {
  constexpr size_t BUFFER_SIZE = 4096;
  char buffer[BUFFER_SIZE];

  ssize_t read_size = read(gui_fd, buffer, BUFFER_SIZE);

  if (read_size < 0) {
    error_exit("Error reading from GUI");
  }
  if (read_size == 0) {
    error_exit("GUI connection closed");
  }

  gui_buffer_in += std::string(buffer, read_size);

  std::string::size_type pos;
  std::string::size_type last_pos = 0;

  while ((pos = gui_buffer_in.find('\n', last_pos)) != std::string::npos) {
    std::string message = gui_buffer_in.substr(last_pos, pos - last_pos);
    handle_raw_message_from_gui(message);

    last_pos = pos + 1;
  }

  gui_buffer_in = gui_buffer_in.substr(last_pos);
}

void Client::handle_gui_poll_out() {
  if (!gui_buffer_out.empty()) {
    ssize_t write_size;
    if ((write_size = write(gui_fd, gui_buffer_out.c_str(), gui_buffer_out.size())) < 0) {
      error_exit("Error writing to gui");
    }
    gui_buffer_out = gui_buffer_out.substr(write_size);

    if (gui_buffer_in.empty()) {
      get_gui_poll_fd().events &= ~POLLOUT;
    }
  }
}

struct pollfd &Client::get_gui_poll_fd() {
  return connection.get_poll_fd_by_fd(gui_fd);
}

void Client::set_player_names(const communication::event_t &event_data) {
  player_names.clear();
  std::stringstream ss;
  ss << std::string(event_data.event_data.new_game_data.player_names,
                    event_data.event_data.new_game_data.players_len - 1);
  std::string name;
  std::string last_name;
  while (std::getline(ss, name, '\0')) {
    if (name <= last_name) {
      error_exit("Incorrect name order or duplicated names");
    }
    last_name = name;
    player_names.push_back(name);
  }

  if (player_names.size() < constants::REQUIRED_PLAYERS) {
    error_exit("Player list to short");
  }
}

std::string Client::get_player_names_string() const {
  std::stringstream ss;
  for (auto const &player : player_names) {
    ss << player << " ";
  }
  std::string res = ss.str();
  res.pop_back();
  return res;
}

void Client::handle_event(communication::event_t const &event) {
  switch (event.event_type) {
    case communication::EventType::NEW_GAME:
      handle_new_game_event(event);
      break;
    case communication::EventType::PLAYER_ELIMINATED:
      handle_player_elimination_event(event);
      break;
    case communication::EventType::PIXEL:
      handle_pixel_event(event);
      break;
    case communication::EventType::GAME_OVER:
      handle_game_over_event(event);
      break;
    default:
      break;
  }
}

void Client::handle_new_game_event(communication::event_t const &event) {
  set_player_names(event);
  width = event.event_data.new_game_data.maxx;
  height = event.event_data.new_game_data.maxy;

  std::string message = create_message_for_gui(event);

  send_gui_message(message);
}

void Client::handle_pixel_event(communication::event_t const &event) {
  if (event.event_data.pixel_data.x >= width || event.event_data.pixel_data.y >= height) {
    error_exit("Pixel out of map");
  }

  std::string message = create_message_for_gui(event);

  send_gui_message(message);
}

void Client::handle_player_elimination_event(communication::event_t const &event) {
  std::string message = create_message_for_gui(event);

  send_gui_message(message);
}

void Client::handle_game_over_event(communication::event_t const &) {
}
