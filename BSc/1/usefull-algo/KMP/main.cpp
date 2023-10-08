#include <iostream>
#include "KMP.h"
using namespace std;
int main() {
  string s, p;
  cin >> s >> p;
  for(int i:prefixTable(p)){
    cout<<i<<" ";
  }
  cout<<endl;
  auto a = KMP(s, p);
  for (int i : a) {
    cout << i << " ";
  }
  return 0;
}