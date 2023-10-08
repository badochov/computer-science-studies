#ifndef INC_2_CLIENT_CLIENT_H_
#define INC_2_CLIENT_CLIENT_H_
#include "../Util/util.h"
#include "util.h"

#include <string>
#include <vector>

class Client {
  static constexpr long SERVER_MESSAGE_FREQUENCY_NS = 30e6;

  std::vector<std::string> player_names;

  communication::client_message_t data;

  std::string server_ip;
  uint16_t server_port;

  std::string gui_ip;
  uint16_t gui_port;

  uint32_t game_id;

  int server_fd{};

  int gui_fd{};

  uint32_t width{};
  uint32_t height{};

  communication::Connection connection;

  std::string gui_buffer_in;

  std::string gui_buffer_out;

  enum class GUI_MESSAGE { LEFT_KEY_DOWN, LEFT_KEY_UP, RIGHT_KEY_DOWN, RIGHT_KEY_UP, UNKNOWN };

  [[nodiscard]] static GUI_MESSAGE create_message_from_gui_type(std::string const &message);
  [[nodiscard]] std::string create_message_for_gui(communication::event_t const &event) const;

  [[nodiscard]] std::string const &player_number_to_name(uint8_t player_number) const;

  [[nodiscard]] static communication::event_t::event_data_t
  parse_event_data(unsigned char const *msg, size_t &byte, uint8_t event_type,
                   communication::event_t const &event, size_t len);

  [[nodiscard]] static communication::server_message_t
  parse_server_message(const std::string &message);

  [[nodiscard]] std::string create_message_for_server() const;

  void handle_server_raw_message(std::string const &message);

  void handle_server_message(communication::server_message_t const &server_message);

  void send_to_server();

  void send_server_message(std::string const &message) const;

  void connect_to_server();

  void connect_to_gui();

  void handle_raw_message_from_gui(std::string const &message);

  void handle_message_from_gui(GUI_MESSAGE gui_message_t);

  void send_gui_message(std::string const &message);

  void handle_new_game(communication::server_message_t const &message);

  void create_server_socket();

  void create_gui_socket();

  void create_timer();

  void handle_gui_poll_in();

  void handle_gui_poll_out();

  struct pollfd &get_gui_poll_fd();

  void set_player_names(communication::event_t const &event_data);

  [[nodiscard]] std::string get_player_names_string() const;

  void handle_event(communication::event_t const &event);

  void handle_new_game_event(communication::event_t const &event);

  void handle_game_over_event(communication::event_t const &event);

  void handle_pixel_event(communication::event_t const &event);

  void handle_player_elimination_event(communication::event_t const &event);

public:
  explicit Client(OptionParser const &option_parser);

  [[noreturn]] void run();
  [[nodiscard]] std::string _handle_new_game_message(const communication::event_t &event) const;
  [[nodiscard]] std::string _handle_pixel_message(const communication::event_t &event) const;
  [[nodiscard]] std::string
  _handle_player_eliminated_message(const communication::event_t &event) const;
};

#endif // INC_2_CLIENT_CLIENT_H_
