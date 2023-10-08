#include <bits/stdc++.h>

#define FOR_HELPER(start, end, var_name) for(auto var_name = start; var_name < end; var_name++)
#define FOR_HELPER_1(start, end) FOR_HELPER(start, end, i)
#define FOR_CHOOSER(x, start, end, increment, FOR_MACRO, ...) FOR_MACRO

#define FOR(...) FOR_CHOOSER(,##__VA_ARGS__, FOR_HELPER(__VA_ARGS__), FOR_HELPER_1(__VA_ARGS__))
#define FOR2(start, end) FOR_HELPER(start, end, j)

#define REV_FOR_HELPER(start, end, var_name) for(auto var_name = start; var_name >= end; var_name--)
#define REV_FOR_HELPER_1(start, end) REV_FOR_HELPER(start, end, i)
#define REV_FOR_HELPER_2(start) REV_FOR_HELPER(start, 0, i)

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

// https://stackoverflow.com/a/32685618
struct pair_hash {
    std::size_t operator()(const std::pair<unsigned int, unsigned int> &p) const {
        return ((unsigned long long) p.first << 32u) + p.second;
    }
};

using island_coords_t = pair<unsigned int, unsigned int>;
using island_neighbours_coords_t = vector<island_coords_t>;
using islands_map_t = map<island_coords_t, island_neighbours_coords_t>;

islands_map_t get_neighbours(vector<island_coords_t> &islands) {
    islands_map_t neighbours;
    for (auto &island : islands) {
        island_neighbours_coords_t island_neighbours;
        neighbours.emplace(island, island_neighbours);
    }

    sort(islands.begin(), islands.end());
    for (int i = 0; i < islands.size() - 1; i++) {
        auto island = islands[i];
        auto r_island = islands[i + 1];

        neighbours.at(island).push_back(r_island);
    }
    for (auto i = islands.size() - 1; i > 0; i--) {
        auto island = islands[i];
        auto l_island = islands[i - 1];

        neighbours.at(island).push_back(l_island);
    }

    sort(islands.begin(), islands.end(), [](const island_coords_t &i1, const island_coords_t &i2) {
        if (i1.second == i2.second) {
            return i1.first > i2.first;
        }
        return i1.second > i2.second;
    });

    for (int i = 0; i < islands.size() - 1; i++) {
        auto island = islands[i];
        auto r_island = islands[i + 1];

        neighbours.at(island).push_back(r_island);
    }
    for (auto i = islands.size() - 1; i > 0; i--) {
        auto island = islands[i];
        auto l_island = islands[i - 1];

        neighbours.at(island).push_back(l_island);
    }

    return neighbours;

}

unsigned int calc_dist(const island_coords_t &is1, const island_coords_t &is2) {
    return min(abs((int) is1.first - (int) is2.first), abs((int) is1.second - (int) is2.second));
}

int get_dist(const islands_map_t &island_map, island_coords_t start, island_coords_t destination) {
    map<island_coords_t, bool> visited;
    using cont = pair<unsigned int, island_coords_t>;
    priority_queue<cont, vector<cont>,greater<>> pq;
    pq.emplace(0, start);
    while (!pq.empty()) {
        auto [dist, coords] = pq.top();
        pq.pop();
        if (visited.count(coords) != 0) {
            continue;
        }
        visited[coords] = true;
        if (coords == destination) {
            return dist;
        }
        for (auto &neighbor : island_map.at(coords)) {
            if (visited.count(neighbor) == 0) {
                unsigned int new_dist = dist + calc_dist(neighbor, coords);
                pq.emplace(new_dist, neighbor);
            }
        }
    }

    return -1;
}

int main() {
    setupIO();
    int n = input_unsigned_number();
    vector<island_coords_t> islands;
    islands.reserve(n);

    FOR(0, n) {
        int x = input_unsigned_number(), y = input_unsigned_number();
        islands.emplace_back(x, y);
    }
    island_coords_t start = islands.front();
    island_coords_t destination = islands.back();

    islands_map_t islands_with_neighbours = get_neighbours(islands);

    cout << get_dist(islands_with_neighbours, start, destination);

    return 0;
}