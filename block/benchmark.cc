#include <benchmark/benchmark.h>

#include <unordered_set>

#include "slotted_page.hh"
#include "sorted_array.hh"

#define MODULO 1000000
#define DS_SIZE 20000
#define NO_OPERATIONS 10000

static void BM_SortedArray_Insert(benchmark::State& state) {
  for (auto _ : state) {
    auto sArray = SortedArray(DS_SIZE);

    for (int idx = 0; idx < NO_OPERATIONS; idx++) {
      int key = std::rand() % MODULO;
      int value = std::rand() % MODULO;

      sArray.Insert(key, value);
    }
  }
}

static void BM_SlottedPage_Insert(benchmark::State& state) {
  for (auto _ : state) {
    auto sPage = SlottedPage(DS_SIZE);

    for (int idx = 0; idx < NO_OPERATIONS; idx++) {
      int key = std::rand() % MODULO;
      int value = std::rand() % MODULO;

      sPage.Insert(key, value);
    }
  }
}

static void BM_SortedArray_Search(benchmark::State& state) {
  auto sArray = SortedArray(DS_SIZE);

  for (int idx = 0; idx < NO_OPERATIONS; idx++) {
    int key = std::rand() % MODULO;
    int value = std::rand() % MODULO;

    sArray.Insert(key, value);
  }

  for (auto _ : state) {
    for (int idx = 0; idx < NO_OPERATIONS; idx++) {
      int key = std::rand() % MODULO;
      sArray.Search(key);
    }
  }
}

static void BM_SlottedPage_Search(benchmark::State& state) {
  auto sPage = SlottedPage(DS_SIZE);

  for (int idx = 0; idx < NO_OPERATIONS; idx++) {
    int key = std::rand() % MODULO;
    int value = std::rand() % MODULO;

    sPage.Insert(key, value);
  }
  for (auto _ : state) {
    for (int idx = 0; idx < NO_OPERATIONS; idx++) {
      int key = std::rand() % MODULO;
      sPage.Search(key);
    }
  }
}

static void BM_SortedArray_Generic(benchmark::State& state) {
  std::unordered_set<int> keys;

  for (auto _ : state) {
    auto sArray = SortedArray(DS_SIZE);

    for (int idx = 0; idx < NO_OPERATIONS; idx++) {
      int key = std::rand() % MODULO;
      int value = std::rand() % MODULO;
      keys.insert(key);

      sArray.Insert(key, value);
      auto _ = sArray.Search(key);
    }

    for (auto key : keys) sArray.Remove(key);
    keys.clear();
  }
}

static void BM_SlottedPage_Generic(benchmark::State& state) {
  std::unordered_set<int> keys;

  for (auto _ : state) {
    auto sPage = SlottedPage(DS_SIZE);

    for (int idx = 0; idx < NO_OPERATIONS; idx++) {
      int key = std::rand() % MODULO;
      int value = std::rand() % MODULO;
      keys.insert(key);

      sPage.Insert(key, value);
      auto _ = sPage.Search(key);
    }
    for (auto key : keys) sPage.Remove(key);
    keys.clear();
  }
}

// Register the function as a benchmark
BENCHMARK(BM_SortedArray_Insert);
BENCHMARK(BM_SlottedPage_Insert);
BENCHMARK(BM_SortedArray_Search);
BENCHMARK(BM_SlottedPage_Search);
BENCHMARK(BM_SortedArray_Generic);
BENCHMARK(BM_SlottedPage_Generic);

BENCHMARK_MAIN();