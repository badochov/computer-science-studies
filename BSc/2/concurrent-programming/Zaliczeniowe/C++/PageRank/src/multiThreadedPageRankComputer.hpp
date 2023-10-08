#ifndef SRC_MULTITHREADEDPAGERANKCOMPUTER_HPP_
#define SRC_MULTITHREADEDPAGERANKCOMPUTER_HPP_

#include <atomic>
#include <mutex>
#include <thread>

#include <unordered_map>
#include <unordered_set>
#include <vector>

#include "immutable/network.hpp"
#include "immutable/pageIdAndRank.hpp"
#include "immutable/pageRankComputer.hpp"

#include <chrono>
#include <condition_variable>
#include <cstdlib>
#include <iostream>
#include <mutex>
#include <thread>

class Barrier {
public:
    Barrier(int threadCount, std::function<void()> _callback)
        : resistance(threadCount)
        , org_resistance(threadCount)
        , callback(_callback)
    {
    }

    void reach(std::mutex& m)
    {
        std::unique_lock<std::mutex> lock(m, std::adopt_lock);
        --resistance;
        if (resistance == 0) {
            resistance = org_resistance;
            callback();
            cond.notify_all();
        } else {
            cond.wait(lock);
        }
    }

private:
    int resistance;
    int org_resistance;
    std::condition_variable cond;
    std::function<void()> callback;
};

template <typename T>
class MyElementsIterator {
    size_t id;
    size_t parts;
    T& it_holder;
    const typename T::iterator _begin;
    const typename T::iterator _end;

    typename T::iterator get_begin() const
    {
        size_t part_size = it_holder.size() / parts;
        typename T::iterator it = it_holder.begin();
        std::advance(it, part_size * id);
        return it;
    }

    typename T::iterator get_end() const
    {
        if (id == parts - 1) {
            return it_holder.end();
        }

        size_t part_size = it_holder.size() / parts;
        typename T::iterator it = it_holder.begin();
        std::advance(it, part_size * (id + 1));
        return it;
    }

public:
    MyElementsIterator(size_t _id, size_t _parts, T& _it_holder)
        : id(_id)
        , parts(_parts)
        , it_holder(_it_holder)
        , _begin(get_begin())
        , _end(get_end())
    {
    }

    typename T::iterator begin() const
    {
        return _begin;
    }

    typename T::iterator end() const
    {
        return _end;
    }
};

template <typename T>
class MyElementsConstIterator {
    size_t id;
    size_t parts;
    T const& it_holder;
    const typename T::const_iterator _begin;
    const typename T::const_iterator _end;

    typename T::const_iterator get_begin() const
    {
        size_t part_size = it_holder.size() / parts;
        typename T::const_iterator it = it_holder.begin();
        std::advance(it, part_size * id);
        return it;
    }

    typename T::const_iterator get_end() const
    {
        if (id == parts - 1) {
            return it_holder.end();
        }

        size_t part_size = it_holder.size() / parts;
        typename T::const_iterator it = it_holder.begin();
        std::advance(it, part_size * (id + 1));
        return it;
    }

public:
    MyElementsConstIterator(size_t _id, size_t _parts, T const& _it_holder)
        : id(_id)
        , parts(_parts)
        , it_holder(_it_holder)
        , _begin(get_begin())
        , _end(get_end())
    {
    }

    typename T::const_iterator begin() const
    {
        return _begin;
    }

    typename T::const_iterator end() const
    {
        return _end;
    }
};

class MultiThreadedPageRankComputer : public PageRankComputer {
public:
    MultiThreadedPageRankComputer(uint32_t numThreadsArg)
        : numThreads(numThreadsArg) {};

    std::vector<PageIdAndRank>
    computeForNetwork(Network const& network, double alpha, uint32_t iterations, double tolerance) const
    {
        std::unordered_map<PageId, PageRank, PageIdHash> pageHashMap;

        std::vector<PageIdAndRank> result;
        result.reserve(pageHashMap.size());
        std::vector<std::thread> threads;
        std::unordered_map<PageId, PageRank, PageIdHash> previousPageHashMap;

        double dangleSum = 0.;
        double difference = 0.;

        std::mutex mutex;

        Barrier dangleBarrier(numThreads, [&]() {
            dangleSum *= alpha;
            difference = 0;
        });

        Barrier diffBarrier(numThreads, [&]() {
            result.clear();
            for (auto iter : pageHashMap) {
                result.push_back(PageIdAndRank(iter.first, iter.second));
            }

            ASSERT(result.size() == network.getSize(),
                "Invalid result size=" << result.size() << ", for network" << network);

            previousPageHashMap = pageHashMap;
            dangleSum = 0;
        });

        threads.reserve(numThreads);

        for (size_t i = 0; i < numThreads; i++) {
            threads.push_back(std::thread([&, num = i]() {
                const auto myPagesIterator = MyElementsConstIterator<std::vector<Page>>(num, numThreads,
                    network.getPages());
                for (auto const& page : myPagesIterator) {
                    page.generateId(network.getGenerator());
                }
            }));
        }

        for (std::thread& t : threads) {
            t.join();
        }

        for (auto const& page : network.getPages()) {
            pageHashMap[page.getId()] = 1.0 / network.getSize();
        }

        std::unordered_map<PageId, uint32_t, PageIdHash> numLinks;
        for (auto page : network.getPages()) {
            numLinks[page.getId()] = page.getLinks().size();
        }

        std::unordered_set<PageId, PageIdHash> danglingNodes;
        for (auto page : network.getPages()) {
            if (page.getLinks().size() == 0) {
                danglingNodes.insert(page.getId());
            }
        }

        std::unordered_map<PageId, std::vector<PageId>, PageIdHash> edges;
        for (auto page : network.getPages()) {
            for (auto link : page.getLinks()) {
                edges[link].push_back(page.getId());
            }
        }
        previousPageHashMap = pageHashMap;

        threads.clear();

        for (size_t i = 0; i < numThreads; i++) {
            threads.push_back(std::thread([&, num = i]() {
                const auto myDanglingNodesIterator = MyElementsConstIterator<decltype(danglingNodes)>(num, numThreads,
                    danglingNodes);

                const auto myPageHashMapIterator = MyElementsIterator<decltype(pageHashMap)>(num, numThreads,
                    pageHashMap);

                for (uint32_t _i = 0; _i < iterations; ++_i) {
                    double threadSum = 0;

                    for (auto danglingNode : myDanglingNodesIterator) {
                        threadSum += previousPageHashMap[danglingNode];
                    }
                    mutex.lock();
                    dangleSum += threadSum;
                    dangleBarrier.reach(mutex);

                    double threadDiff = 0;
                    for (auto& pageMapElem : myPageHashMapIterator) {
                        PageId pageId = pageMapElem.first;

                        double danglingWeight = 1.0 / network.getSize();
                        pageMapElem.second = dangleSum * danglingWeight + (1.0 - alpha) / network.getSize();

                        if (edges.count(pageId) > 0) {
                            for (auto link : edges[pageId]) {
                                pageMapElem.second += alpha * previousPageHashMap[link] / numLinks[link];
                            }
                        }
                        threadDiff += std::abs(previousPageHashMap[pageId] - pageHashMap[pageId]);
                    }

                    mutex.lock();
                    difference += threadDiff;
                    diffBarrier.reach(mutex);

                    if (difference < tolerance) {
                        return;
                    }
                }

                ASSERT(false, "Not able to find result in iterations=" << iterations);
            }));
        }

        for (std::thread& t : threads) {
            t.join();
        }

        return result;
    }

    std::string getName() const
    {
        return "MultiThreadedPageRankComputer[" + std::to_string(this->numThreads) + "]";
    }

private:
    uint32_t numThreads;
};

#endif /* SRC_MULTITHREADEDPAGERANKCOMPUTER_HPP_ */
