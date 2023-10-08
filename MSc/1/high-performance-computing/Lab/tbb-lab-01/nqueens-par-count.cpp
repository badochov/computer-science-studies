// Solves the n-quees puzzle on an n x x checkerboard.
//
// This sequential implementation is to be extended with TBB to get a
// parallel implementation.
//
// HPC course, MIM UW
// Krzysztof Rzadca, LGPL

#include "tbb/tbb.h"
#include <cmath>
#include <iostream>
#include <list>
#include <vector>

// Indexed by column. Value is the row where the queen is placed,
// or -1 if no queen.
typedef std::vector<int> Board;

void pretty_print(const Board &board) {
  for (int row = 0; row < (int)board.size(); row++) {
    for (const auto &loc : board) {
      if (loc == row)
        std::cout << "*";
      else
        std::cout << " ";
    }
    std::cout << std::endl;
  }
}

// Checks the location of queen in column 'col' against queens in cols [0, col).
bool check_col(Board &board, int col_prop) {
  int row_prop = board[col_prop];
  int col_queen = 0;
  for (auto i = board.begin(); (i != board.end()) && (col_queen < col_prop);
       ++i, ++col_queen) {
    int row_queen = *i;
    if (row_queen == row_prop) {
      return false;
    }
    if (abs(row_prop - row_queen) == col_prop - col_queen) {
      return false;
    }
  }
  return true;
}

void initialize(Board &board, int size) {
  board.reserve(size);
  for (int col = 0; col < size; ++col)
    board.push_back(-1);
}

int grain = 2;

// Solves starting from a partially-filled board (up to column col).
void recursive_solve(Board &partial_board, int col,
                     std::atomic<size_t> &solutions) {
  // std::cout << "rec solve col " << col << std::endl;
  // pretty_print(b_partial);

  int b_size = partial_board.size();
  if (col == b_size) {
    solutions++;
  } else {
    if (col < grain) {
      tbb::parallel_for(
          tbb::blocked_range<long>(0, b_size),
          [&partial_board, col, &solutions](const tbb::blocked_range<long> &r) {
            for (long tested_row = r.begin(); tested_row != r.end();
                 ++tested_row) {
              Board board = partial_board;
              board[col] = tested_row;
              if (check_col(board, col))
                recursive_solve(board, col + 1, solutions);
            }
          });
    } else {
      for (int tested_row = 0; tested_row < b_size; tested_row++) {
        partial_board[col] = tested_row;
        if (check_col(partial_board, col))
          recursive_solve(partial_board, col + 1, solutions);
      }
    }
  }
}

int main(int argc, char **argv) {
  if (argc == 2) {
    grain = atoi(argv[1]);
  }

  const int board_size = 13;
  Board board{};
  initialize(board, board_size);
  std::atomic<size_t> solutions{};

  tbb::tick_count seq_start_time = tbb::tick_count::now();
  recursive_solve(board, 0, solutions);
  tbb::tick_count seq_end_time = tbb::tick_count::now();
  double seq_time = (seq_end_time - seq_start_time).seconds();
  std::cout << "par time: " << seq_time << "[s]" << std::endl;

  std::cout << "solution count: " << solutions << std::endl;
  // for (const auto& sol : solutions) {
  //     pretty_print(sol);
  //     std::cout << std::endl;
  // }
}
