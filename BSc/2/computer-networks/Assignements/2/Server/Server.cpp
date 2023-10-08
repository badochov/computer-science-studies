#include "Server.h"
#include <arpa/inet.h>
#include <chrono>
#include <cstring>
#include <sstream>

#include <sys/timerfd.h>
#include <unistd.h>

#define MSG (msg + byte)

communication::client_message_t Server::parse_client_message(std::string const &message) {
  communication::client_message_t res;

  auto *msg = reinterpret_cast<const unsigned char *>(message.c_str());

  std::string::size_type byte = 0;

  res.session_id = network_data_to_u64(MSG);
  byte += 8;
  res.turn_direction = *MSG;
  byte += 1;
  res.next_expected_event_no = network_data_to_u32(MSG);
  byte += 4;
  res.player_name = message.substr(byte);

  return res;
}

std::string Server::create_events_message(uint32_t next_event) const {
  std::string::size_type length = 0;
  std::stringstream ss;

  ss << u32_to_network_data(game_id);
  length += 4;

  while (next_event < events.size()) {
    communication::event_t const &event = events[next_event];
    std::string event_str = event.to_string();
    length += event_str.size();
    if (length > constants::MAX_SERVER_MESSAGE_LENGTH) {
      break;
    }
    ss << event_str;

    ++next_event;
  }

  return ss.str();
}

void Server::start_new_game() {
  clear_after_previous_round();
  game_ongoing = true;

  game_id = random.random_u32();

  add_event(generate_new_game());
  for (auto &player : players) {
    start_new_game_handle_player(player.second);
  }
}

communication::event_t Server::generate_new_game() const {
  communication::event_t event;
  event.event_no = events.size();
  event.event_type = communication::NEW_GAME;

  communication::new_game_data_t new_game_data{};

  std::string players_string = get_player_names_string();

  new_game_data.maxx = board.width;
  new_game_data.maxy = board.height;
  memcpy(new_game_data.player_names, players_string.c_str(), players_string.size());
  new_game_data.players_len = players_string.size();

  event.event_data.new_game_data = new_game_data;
  event.len += new_game_data.size();

  return event;
}

communication::event_t Server::generate_player_eliminated(uint32_t player_number) const {
  communication::event_t event;
  event.event_no = events.size();
  event.event_type = communication::PLAYER_ELIMINATED;
  event.event_data.player_number = player_number;

  event.len += constants::CHAR_SIZE;

  return event;
}

communication::event_t Server::generate_pixel(Worm const &worm) const {
  communication::event_t event;
  event.event_no = events.size();
  event.event_type = communication::PIXEL;

  communication::pixel_data_t pixel_data{};
  pixel_data.player_number = worm.number;
  pixel_t pixel = worm.get_current_pixel();
  pixel_data.x = pixel.first;
  pixel_data.y = pixel.second;

  event.event_data.pixel_data = pixel_data;

  event.len += pixel_data.size();

  return event;
}

communication::event_t Server::generate_game_over() const {
  communication::event_t event;
  event.event_no = events.size();
  event.event_type = communication::GAME_OVER;

  return event;
}

std::string Server::get_player_names_string() const {
  std::stringstream ss;

  for (auto &player : players) {
    ss << player.first << '\0';
  }

  return ss.str();
}

void Server::start_new_game_handle_player(PlayerData &player) {
  uint8_t player_number = worms.size();
  worms.emplace_back();
  Worm &worm = worms.back();
  worm.x = double(random.random_u32() % board.width) + 0.5;
  worm.y = double(random.random_u32() % board.height) + 0.5;

  worm.direction = random.random_u32() % 360;
  worm.is_dead = false;
  worm.turn_direction = player.turn_direction;
  worm.number = player_number;

  player.number = player_number;
  player.ready = false;
  player.has_worm = true;

  pixel_t pixel = worm.get_current_pixel();

  if (board.is_taken(pixel)) {
    eliminate_worm(worm);
  } else {
    take_pixel(worm);
  }
}

void Server::perform_round() {
  types::timestamp_t timestamp = get_current_timestamp();

  for (auto it = players.begin(); it != players.end();) {
    PlayerData player_data = it->second;
    ++it;
    if (timestamp - player_data.last_message_timestamp > TIMEOUT) {
      disconnect_player(player_data);
    }
  }
  uint32_t ongoing = game_ongoing;
  try_to_start_game();

  for (auto &worm : worms) {
    if (ongoing) {
      perform_round_handle_worm(worm);
      ongoing = game_ongoing;
    } else {
      break;
    }
  }
}

void Server::perform_round_handle_worm(Worm &worm) {
  if (worm.is_dead) {
    return;
  }
  if (worm.turn_direction == communication::DIRECTION::RIGHT) {
    worm.direction = (360 + worm.direction + turning_speed) % 360;
  } else if (worm.turn_direction == communication::DIRECTION::LEFT) {
    worm.direction = (360 + worm.direction - turning_speed) % 360;
  }

  pixel_t previous = worm.get_current_pixel();
  worm.move();
  pixel_t current = worm.get_current_pixel();

  if (previous != current) {
    pixel_t pixel = worm.get_current_pixel();
    if (!board.is_in(pixel) || board.is_taken(pixel)) {
      eliminate_worm(worm);
    } else {
      take_pixel(worm);
    }
  }
}

void Server::eliminate_worm(Worm &worm) {
  add_event(generate_player_eliminated(worm.number));

  worm.is_dead = true;
  --active_worms;
  if (active_worms < 2) {
    end_game();
  }
}

void Server::take_pixel(Worm &worm) {
  pixel_t pixel = worm.get_current_pixel();
  board.place(pixel, worm.number);

  add_event(generate_pixel(worm));
}

void Server::clear_after_previous_round() {
  events.clear();
  board.clear();
  active_worms = players.size();
  worms.clear();
}

void Server::end_game() {
  game_ongoing = false;

  add_event(generate_game_over());
}

void Server::handle_client_message(communication::client_message_t const &client_message,
                                   types::address_t const &address) {
  std::string player_name = get_player_name(address);

  if (player_name.empty()) {
    handle_not_connected_player_message(client_message, address);
  } else {
    handle_connected_player_message(players.at(player_name), client_message);
  }
}

bool Server::is_player_name_unique(communication::client_message_t const &client_message) const {
  return players.count(client_message.player_name) == 0;
}

void Server::handle_connected_player_message(
    PlayerData &player, communication::client_message_t const &client_message) {
  if (!player.correct_session_id(client_message.session_id)) {
    reconnect_player(player, client_message);
  } else {
    player.turn_direction = client_message.turn_direction;
    if (player.has_worm) {
      get_player_worm(player).turn_direction = player.turn_direction;
    }
    if (!game_ongoing && client_message.turn_direction != communication::DIRECTION::FORWARD) {
      player.ready = true;
    }

    player.last_message_timestamp = get_current_timestamp();

    send_player_events(player, client_message.next_expected_event_no);
  }
}

void Server::handle_not_connected_player_message(
    communication::client_message_t const &client_message, types::address_t const &address) {
  if (client_message.player_name.empty()) {
    handle_spectator_connection(client_message, address);
  } else {
    if (players.size() < constants::MAX_PLAYERS && is_player_name_unique(client_message)) {
      connect_player(client_message, address);
    }
  }
}

void Server::disconnect_player(PlayerData &player) {
  players.erase(player.player_name);
}

void Server::connect_player(communication::client_message_t const &client_message,
                            types::address_t const &address) {
  if (valid_username(client_message.player_name)) {
    players.emplace(client_message.player_name, PlayerData{client_message, address});

    handle_connected_player_message(players.at(client_message.player_name), client_message);
  }
}

void Server::send_player_events(const PlayerData &player, uint32_t next_event) {
  send_socket_events(player.address, next_event);
}

void Server::handle_message(const std::string &message, const types::address_t &address) {
  communication::client_message_t client_message = parse_client_message(message);

  handle_client_message(client_message, address);
}

[[noreturn]] void Server::run() {
  add_game_clock();
  open_server();

  for (;;) {
    connection.poll();
  }
}

void Server::open_server() {
  struct sockaddr_in6 server_addr {};

  int fd = socket(AF_INET6, SOCK_DGRAM, IPPROTO_UDP);
  if (fd < 0) {
    error_exit("Error openning socket");
  }
  int on = 1;
  if (setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &on, sizeof(on)) < 0) {
    error_exit("Error openning socket");
  }

  std::memset(&server_addr, 0, sizeof(server_addr));
  server_addr.sin6_family = AF_INET6;
  server_addr.sin6_port = htons(port);

  server_addr.sin6_addr = in6addr_any;

  if (bind(fd, reinterpret_cast<const struct sockaddr *>(&server_addr), sizeof(server_addr)) < 0) {
    error_exit("Bind failed");
  }

  communication::Connection::callback_t callback = [this](struct pollfd &poll_fd) {
    handle_poll_fd_event(poll_fd);
  };

  connection.add_fd({.fd = fd, .events = POLLIN, .revents = 0}, callback);

  server_fd = fd;
}

void Server::add_game_clock() {
  int timer = timerfd_create(CLOCK_MONOTONIC, 0);
  struct timespec period {
    .tv_sec = 1 / rounds_per_second, .tv_nsec = long(1e9 / rounds_per_second) % long(1e9)
  };

  struct itimerspec timer_period {
    .it_interval = period, .it_value = period
  };
  if (timer < 0 || timerfd_settime(timer, 0, &timer_period, nullptr)) {
    error_exit("Could create game clock");
  }

  communication::Connection::callback_t callback = [this](struct pollfd &poll_fd) {
    poll_fd.revents = 0;
    uint64_t elapses;

    if (read(poll_fd.fd, &elapses, constants::U64_SIZE) != constants::U64_SIZE) {
      error_exit("Server clock error");
    }

    perform_round();
  };

  connection.add_fd({.fd = timer, .events = POLLIN, .revents = 0}, callback);
}

void Server::handle_poll_fd_event(struct pollfd &poll_fd) {
  types::address_t client_addr;
  // One byte longer to check if player name isn't too long.
  char buffer[communication::client_message_t::MAX_SIZE + 1];
  socklen_t len = sizeof(client_addr);

  ssize_t read_count = recvfrom(poll_fd.fd, buffer, communication::client_message_t::MAX_SIZE + 1,
                                0, reinterpret_cast<struct sockaddr *>(&client_addr), &len);

  if (read_count < 0) {
    std::cerr << "Read error" << std::endl;
    return;
  }
  if ((uint32_t)read_count < communication::client_message_t::BASE_SIZE) {
    std::cerr << "Message too short" << std::endl;
    return;
  }

  std::string message(buffer, read_count);

  handle_message(message, client_addr);
}

bool Server::game_can_be_started() const {
  if (players.size() >= REQUIRED_PLAYERS) {
    return std::all_of(players.begin(), players.end(),
                       [](auto const &player) { return player.second.ready; });
  }

  return false;
}

void Server::try_to_start_game() {
  if (!game_ongoing) {
    if (game_can_be_started()) {
      start_new_game();
    }
  }
}

std::string Server::get_player_name(const types::address_t &address) const {
  for (auto &player : players) {
    if (same_address(address, player.second.address)) {
      return player.first;
    }
  }

  return "";
}

void Server::add_event(communication::event_t const &event) {
  events.push_back(event);
  for (auto &player : players) {
    send_player_events(player.second, events.size() - 1);
  }
  for (auto &spectator : spectators) {
    send_spectator_events(spectator, events.size() - 1);
  }
}
void Server::reconnect_player(PlayerData &player,
                              communication::client_message_t const &client_message) {
  disconnect_player(player);
  connect_player(client_message, player.address);
}

void Server::send_spectator_events(SpectatorData const &spectator_data, uint32_t next_event) {
  send_socket_events(spectator_data.address, next_event);
}

void Server::send_socket_events(types::address_t const &client_address, uint32_t next_event) const {
  if (next_event < events.size()) {
    std::string message = create_events_message(next_event);

    if (sendto(server_fd, message.c_str(), message.size(), 0,
               reinterpret_cast<sockaddr const *>(&client_address), sizeof(client_address)) < 0) {
      std::cerr << "Sendto error" << std::endl;
    }
  }
}

void Server::handle_spectator_connection(const communication::client_message_t &message,
                                         const types::address_t &address) {
  for (auto &spectator : spectators) {
    if (same_address(spectator.address, address)) {
      spectator.last_message_timestamp = get_current_timestamp();
      return send_spectator_events(spectator, message.next_expected_event_no);
    }
  }

  spectators.push_back({.address = address, .last_message_timestamp = get_current_timestamp()});

  return send_spectator_events(spectators.back(), message.next_expected_event_no);
}

Worm &Server::get_player_worm(PlayerData const &player) {
  return worms.at(player.number);
}
