#include <condition_variable>
#include <mutex>
#include <cstdlib>
#include <thread>
#include <chrono>
#include <iostream>

class Barrier {
public:
  Barrier(int ile) : resistance(ile) {}
  void reach() {
    std::unique_lock<std::mutex> lock(mutex);
    resistance--;
    if (resistance < 0) {
      return;
    } else if (resistance == 0) {
      cond.notify_all();
    } else
      cond.wait(lock, [this] { return this->resistance <= 0; });
  }

private:
  int resistance;
  std::mutex mutex;
  std::condition_variable cond;
};

Barrier barr(3);
std::mutex cout_mutex;
void f() {
    int time = std::rand() % 10;
    std::this_thread::sleep_for(std::chrono::seconds(time));
    auto id = std::this_thread::get_id();
    cout_mutex.lock();
    std::cout<<"Wątek "<<id<<" doszedł do bariery!"<<std::endl;
    cout_mutex.unlock();
    barr.reach();
    cout_mutex.lock();
    std::cout<<"Wątek "<<id<<" przeszedł bariery!"<<std::endl;
    cout_mutex.unlock();

}

int main()
{
    std::array<std::thread, 4> threads = {std::thread{f}, std::thread{f},
                                          std::thread{f}, std::thread{f}};
    for (auto& t : threads)
        t.join();
}