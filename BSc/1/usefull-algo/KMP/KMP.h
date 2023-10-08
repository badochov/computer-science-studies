#ifndef KMP__KMP_H_
#define KMP__KMP_H_
#include<vector>
#include<string>
std::vector<int> prefixTable(const std::string &s);
std::vector<int> KMP(const std::string &s, const std::string &pattern);
#endif //KMP__KMP_H_
