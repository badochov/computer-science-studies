#include "KMP.h"
#include <iostream>
std::vector<int> prefixTable(const std::string &s) {
  std::vector<int> table(s.size(), 0);
  int j = 0;
  for (std::string::size_type i = 1; i < s.size(); i++) {
    while (s[j]!=s[i] && j > 0) {
      j = s[j - 1];
    }
    if (s[j]==s[i]) {
      j++;
    }
    table[i] = j;
  }
  return table;
}

std::vector<int> KMP(const std::string &s, const std::string &pattern) {
  std::vector<int> table = prefixTable(pattern);
  std::vector<int> matches;
  int j = 0;
  std::string::size_type i = 0;
  while (i <= s.size() - pattern.size()) {
    j = table[j];
    while (j < pattern.size() && pattern[j]==s[i + j]) {
      j++;
    }
    if (j==pattern.size()) {
      matches.push_back(i);
    }
    if (j==0) {
      i++;
    } else {
      i += j - table[j - 1];
      j--;
    }
    std::cout<<i<<" "<<j<<std::endl;
  }

  return matches;
}
