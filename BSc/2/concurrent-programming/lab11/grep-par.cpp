#include <iostream>
#include <fstream>
#include <locale>
#include <string>
#include <list>
#include <codecvt>
#include <future>
#include <cmath>

constexpr size_t POOL_SIZE = 3;

unsigned int grep(std::string filename, std::wstring word) {
    std::locale loc("pl_PL.UTF-8");
    std::wfstream file(filename);
    file.imbue(loc);
    std::wstring line;
    unsigned int count = 0;
    while (getline(file, line)) {
        for (auto pos = line.find(word, 0);
             pos != std::string::npos;
             pos = line.find(word, pos + 1))
            count++;
    }
    return count;
}

void fun(std::list<std::string> &&filelist, const std::wstring &word, std::promise<int> &count) {
    unsigned int locSum = 0;

    for (const auto &file : filelist) {
        locSum += grep(file, word);
    }
    count.set_value(locSum);
}

int main() {
    std::ios::sync_with_stdio(false);
    std::locale loc("pl_PL.UTF-8");
    std::wcout.imbue(loc);
    std::wcin.imbue(loc);

    std::wstring word;
    std::getline(std::wcin, word);

    std::wstring s_file_count;
    std::getline(std::wcin, s_file_count);
    int file_count = std::stoi(s_file_count);

    std::list<std::string> filenames{};

    std::wstring_convert<std::codecvt_utf8<wchar_t>, wchar_t> converter;

    for (int file_num = 0; file_num < file_count; file_num++) {
        std::wstring w_filename;
        std::getline(std::wcin, w_filename);
        std::string s_filename = converter.to_bytes(w_filename);
        filenames.push_back(s_filename);
    }

    int count = 0;

    std::list<std::thread> filethreads{};
    std::list<std::promise<int>> promiselist{};
    std::list<std::future<int>> futurelist{};

    std::list<std::string> threadFiles[POOL_SIZE];
    unsigned int size = ceil((double) filenames.size() / POOL_SIZE);
    for (auto &files : threadFiles) {
        for (int j = 0; j < size && !filenames.empty(); j++) {
            files.push_back(filenames.front());
            filenames.pop_front();
        }
        promiselist.emplace_back();
        std::promise<int> &promise = promiselist.back();
        futurelist.push_back(promise.get_future());
        filethreads.emplace_back([&files, word, &promise] { fun(std::move(files), word, promise); });
    }

    for (auto &future : futurelist) {
        count += future.get();
    }
    for (auto &thread : filethreads) {
        thread.join();
    }


    std::wcout << count << std::endl;
}
