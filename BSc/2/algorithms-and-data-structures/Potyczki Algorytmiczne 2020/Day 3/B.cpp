#include <bits/stdc++.h>

#define FOR_HELPER(start, end, var_name) for(auto var_name = start; var_name < end; var_name++)
#define FOR_HELPER_1(start, end) FOR_HELPER(start, end, i)
#define FOR_CHOOSER(x, start, end, increment, FOR_MACRO, ...) FOR_MACRO

#define FOR(...) FOR_CHOOSER(,##__VA_ARGS__, FOR_HELPER(__VA_ARGS__), FOR_HELPER_1(__VA_ARGS__))
#define FOR2(start, end) FOR_HELPER(start, end, j)

#define REV_FOR_HELPER(start, end, var_name) for(auto var_name = start; var_name >= end; var_name--)
#define REV_FOR_HELPER_1(start, end) FOR_HELPER(start, end, i)
#define REV_FOR_HELPER_2(start) FOR_HELPER(start, 0, i)

#define REV_FOR_CHOOSER(x, start, end, increment, FOR_MACRO, ...) FOR_MACRO

#define REV_FOR(...) REV_FOR_CHOOSER(,##__VA_ARGS__, REV_FOR_HELPER(__VA_ARGS__), REV_FOR_HELPER_1(__VA_ARGS__), REV_FOR_HELPER_2(__VA_ARGS__))
#define REV_FOR2(start, end) REV_HELPER(start, end, j)

#define END_LINE '\n'

using namespace std;

template<class t = int>
[[maybe_unused]] t input_unsigned_number() {
    t num = 0;
    int c = getchar_unlocked();
    while (c < '0') c = getchar_unlocked();
    while (c >= '0') {
        num = num * 10 + c - '0';
        c = getchar_unlocked();
    }
    return num;
}

void setupIO() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
}

constexpr int MOD = 1000000000 + 7;
constexpr long long max_c = 1000000000000000000;
constexpr int tree_size = 64 * 2;

int tree_to[tree_size];
int tree_after[tree_size];
long long longest_range_after[tree_size];
long long longest_range_before[tree_size];


void updateRangeUtil(int si, int ss, int se, int us,
                     int ue, int diff)
{
    if (lazy[si] != 0)
    {
        tree[si] += (se-ss+1)*lazy[si];

        if (ss != se)
        {
            lazy[si*2 + 1]   += lazy[si];
            lazy[si*2 + 2]   += lazy[si];
        }
        lazy[si] = 0;
    }

    if (ss>se || ss>ue || se<us)
        return ;

    if (ss>=us && se<=ue)
    {
        tree[si] += (se-ss+1)*diff;

        if (ss != se)
        {
            lazy[si*2 + 1]   += diff;
            lazy[si*2 + 2]   += diff;
        }
        return;
    }

    int mid = (ss+se)/2;
    updateRangeUtil(si*2+1, ss, mid, us, ue, diff);
    updateRangeUtil(si*2+2, mid+1, se, us, ue, diff);

    tree[si] = tree[si*2+1] + tree[si*2+2];
}

void updateRange(int n, int us, int ue, int diff)
{
    updateRangeUtil(0, 0, n-1, us, ue, diff);
}

int getSumUtil(int ss, int se, int qs, int qe, int si)
{
    if (lazy[si] != 0)
    {
        tree[si] += (se-ss+1)*lazy[si];

        if (ss != se)
        {
            lazy[si*2+1] += lazy[si];
            lazy[si*2+2] += lazy[si];
        }
        lazy[si] = 0;
    }

    if (ss>se || ss>qe || se<qs)
        return 0;

    if (ss>=qs && se<=qe)
        return tree[si];
    int mid = (ss + se)/2;
    return getSumUtil(ss, mid, qs, qe, 2*si+1) +
           getSumUtil(mid+1, se, qs, qe, 2*si+2);
}
int getSum(int n, int qs, int qe)
{
    if (qs < 0 || qe > n-1 || qs > qe)
    {
        printf("Invalid Input");
        return -1;
    }

    return getSumUtil(0, n-1, qs, qe, 0);
}

pair<long long, long long> get_outlying_blasting_mine(long long pos, long long range){

}


long long blasts_till(long long pos){
    get_till()
}

long long get_blasts_before_including(long long before, long long pos){

}

void add_blasts(long long pos, long long range, long long right, long long count){
    int right_most = max(pos+range, right);

}


long long dp(const vector<pair<long long, long long>> &mines) {
    vector<long long> res;
    res.reserve(mines.size());
    for (const auto &[pos, range] :mines) {
        const auto[left, right] = get_outlying_blasting_mine(pos, range);
        long long all_prev_blasts = blasts_till(left - 1);
        long long blasts_before_pos_including_curr = get_blasts_before_including(left - 1, pos);
        long long new_blasts = all_prev_blasts - blasts_before_pos_including_curr;
        add_blasts(pos, range, right, new_blasts);

        res.push_back(new_blasts);
    }

    return blasts_till(max_c +1);
}

int main() {
    setupIO();
    int n = input_unsigned_number();
    vector<pair<long long, long long>> mines;
    mines.reserve(n);
    FOR(0, n)mines.emplace_back(input_unsigned_number(), input_unsigned_number();


    cout << dp(mines);;
    return 0;
}