CC=gcc
CXX=g++
RM=rm -f
CPPFLAGS=-O2 -Wall -Wextra -fdiagnostics-color=auto -g -std=c++17


all: client server

# Util

UTIL_PATH=Util

UTIL_SRCS=$(UTIL_PATH)/util.cpp
UTIL_OBJS=$(UTIL_PATH)/util.o

$(UTIL_PATH)/util.o: $(UTIL_PATH)/util.h $(UTIL_PATH)/util.cpp


# Server

SERVER_PATH=Server

SERVER_SRCS=$(SERVER_PATH)/Server.cpp $(SERVER_PATH)/main.cpp
SERVER_OBJS=$(subst .cpp,.o,$(SERVER_SRCS))

SERVER_NAME=screen-worms-server

server: $(SERVER_OBJS) $(UTIL_OBJS)
	$(CXX) $(CPPFLAGS) -o $(SERVER_NAME) $(SERVER_OBJS) $(UTIL_OBJS)

$(SERVER_PATH)/Server.o: $(SERVER_PATH)/Server.cpp $(SERVER_PATH)/Server.h

$(SERVER_PATH)/main.o: $(SERVER_PATH)/main.cpp


# Client

CLIENT_PATH=Client

CLIENT_SRCS=$(CLIENT_PATH)/Client.cpp $(CLIENT_PATH)/main.cpp
CLIENT_OBJS=$(subst .cpp,.o,$(CLIENT_SRCS))

CLIENT_NAME=screen-worms-client

client: $(CLIENT_OBJS) $(UTIL_OBJS)
	$(CXX) $(CPPFLAGS) -o $(CLIENT_NAME) $(CLIENT_OBJS) $(UTIL_OBJS)

$(CLIENT_PATH)/Server.o: $(CLIENT_PATH)/Server.cpp $(CLIENT_PATH)/Server.h

$(CLIENT_PATH)/main.o: $(CLIENT_PATH)/main.cpp


buildclean:
	$(RM) $(SERVER_OBJS) $(CLIENT_OBJS) $(UTIL_OBJS)

clean:
	make buildclean
	$(RM) $(CLIENT_NAME) $(SERVER_NAME)


.PHONY: clean server client all


