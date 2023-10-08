#include <iostream>

#include "OptionParser.h"
#include "FileGetter.h"
#include "CorrelatedServersHandler.h"
#include "Server.h"
#include "error.h"

void handle_http_error(HTTPError const &http_error, Server &server) {
  HeaderFields header_fields;
  if (http_error.get_response_conde() != ResponseCode::RESOURCE_NOT_FOUND) {
    header_fields.add("Connection", "close");
  }

  server.respond(http_error.get_response_conde(), http_error.get_reason(), header_fields, "");
}

int main(int argc, char **argv) {
  try {
    OptionParser options(argc, argv);
    FileGetter file_getter(options.getFilesDirectory());
    CorrelatedServersHandler correlated_servers(options.getCorrelatedServersFilePath());
    Server server(options.getPort());

    while (true) {
      try {
        try {
          Request request = server.get_request();
          std::string path = request.request_data.target;

          File file = file_getter.get_file(path);
          HeaderFields header_fields;

          header_fields.add_file_data(file);

          std::string file_content;

          if (request.request_data.method == "GET") {
            file_content = file.file_content;
          }

          server.respond(ResponseCode::OK, "File Found", header_fields, file_content);
        }
        catch (std::bad_alloc &) {
          throw HTTPError(ResponseCode::SERVER_ERROR, "Out of memory");
        }
        catch (FileNotFoundException &exception) {
          correlated_servers.redirect(exception.path, server);
        }
        catch (ClientDisconnected &) {
          server.close_connection();
        }
      }
      catch (HTTPError &http_error) {
        handle_http_error(http_error, server);
      }
    }
  }
  catch (std::exception &exception) {
    std::cerr << exception.what();

    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}