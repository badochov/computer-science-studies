#ifndef INC_2_SERVER_SERVER_H_
#define INC_2_SERVER_SERVER_H_
#include "../Util/util.h"
#include "util.h"
#include <map>
#include <optional>
#include <string>

class Server {
  static constexpr uint64_t TIMEOUT = 2e6;
  static constexpr size_t REQUIRED_PLAYERS = constants::REQUIRED_PLAYERS;

  uint8_t active_worms{};
  uint32_t game_id{};
  bool game_ongoing;

  uint16_t port;

  communication::Connection connection;

  std::map<std::string, PlayerData> players;
  std::vector<SpectatorData> spectators;
  std::vector<Worm> worms;

  std::vector<communication::event_t> events;

  Board board;

  Random random;
  uint32_t turning_speed;
  uint32_t rounds_per_second;

  int server_fd{};

  [[nodiscard]] std::string get_player_names_string() const;
  [[nodiscard]] communication::event_t generate_new_game() const;
  [[nodiscard]] communication::event_t generate_player_eliminated(uint32_t player_number) const;
  [[nodiscard]] communication::event_t generate_pixel(Worm const &worm) const;
  [[nodiscard]] communication::event_t generate_game_over() const;

  static communication::client_message_t parse_client_message(std::string const &message);

  [[nodiscard]] std::string create_events_message(uint32_t next_event) const;

  void clear_after_previous_round();

  void start_new_game();

  void start_new_game_handle_player(PlayerData &player_data);

  void perform_round();

  void perform_round_handle_worm(Worm &worm);

  void eliminate_worm(Worm &worm);

  void take_pixel(Worm &worm);

  void end_game();

  void handle_message(std::string const &message, types::address_t const &address);

  void handle_client_message(communication::client_message_t const &client_message,
                             types::address_t const &address);

  [[nodiscard]] std::string get_player_name(types::address_t const &address) const;

  [[nodiscard]] bool
  is_player_name_unique(communication::client_message_t const &client_message) const;

  void handle_connected_player_message(PlayerData &player,
                                       communication::client_message_t const &client_message);

  void handle_not_connected_player_message(communication::client_message_t const &client_message,
                                           types::address_t const &address);

  void disconnect_player(PlayerData &player);

  void connect_player(communication::client_message_t const &client_message,
                      types::address_t const &address);

  void send_player_events(PlayerData const &player, uint32_t next_event);

  void send_spectator_events(SpectatorData const &spectator_data, uint32_t next_event);

  void send_socket_events(types::address_t const &client_address, uint32_t next_event) const;

  void add_game_clock();

  void open_server();

  void handle_poll_fd_event(struct pollfd &poll_fd);

  void try_to_start_game();

  void add_event(communication::event_t const &event);

  void reconnect_player(PlayerData &player, communication::client_message_t const &client_message);

  Worm &get_player_worm(PlayerData const &player);

public:
  explicit Server(OptionParser const &option_parser)
      : game_ongoing(false), port(option_parser.get_server_port()),
        board(option_parser.get_height(), option_parser.get_width()),
        random(option_parser.get_seed()), turning_speed(option_parser.get_turning_speed()),
        rounds_per_second(option_parser.get_rounds_per_second()), server_fd(-1){};

  [[noreturn]] void run();

  [[nodiscard]] bool game_can_be_started() const;
  void handle_spectator_connection(const communication::client_message_t &message,
                                   const types::address_t &variant);
};

#endif // INC_2_SERVER_SERVER_H_
