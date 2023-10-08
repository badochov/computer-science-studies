CC=gcc
CXX=g++
RM=rm -f
CPPFLAGS=-Wall -Wextra -fdiagnostics-color=auto -g -std=c++17

SRCS=main.cpp Server.cpp HeaderFields.cpp
OBJS=$(subst .cpp,.o,$(SRCS))

NAME=serwer

all: serwer

serwer: $(OBJS)
	echo flags $(CPPFLAGS)
	$(CXX) $(CPPFLAGS) -o $(NAME) $(OBJS) -lstdc++fs
	make buildclean

Server.o: Server.cpp Server.h

HeaderFields.o: HeaderFields.cpp HeaderFields.h

main.o: main.cpp

.PHONY: clean serwer all

buildclean:
	$(RM) $(OBJS)

clean:
	make buildclean
	$(RM) $(NAME)