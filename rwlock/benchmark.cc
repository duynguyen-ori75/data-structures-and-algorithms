#include "test_infras.hh"

#include <benchmark/benchmark.h>

static void BM_SimpleRWLocker_ReadIntensive(benchmark::State& state) {
  for (auto _ : state) {
    auto cnt = std::make_unique<Counter_SimpleRWLocker>(10);
    pthread_t tid[NO_THREADS];

    for (int idx = 0; idx < NO_THREADS; idx ++) {
      if (idx % 10 == 0) {
        pthread_create(&(tid[idx]), NULL, &(atomicRead<Counter_SimpleRWLocker>), cnt.get());
      } else {
        pthread_create(&(tid[idx]), NULL, &(atomicAdd<Counter_SimpleRWLocker>), cnt.get());
      }
    }

    for (int idx = 0; idx < NO_THREADS; idx ++) pthread_join(tid[idx], NULL);
  }
}

static void BM_RWLocker_ReadIntensive(benchmark::State& state) {
  for (auto _ : state) {
    auto cnt = std::make_unique<Counter_RWLocker>(10);
    pthread_t tid[NO_THREADS];

    for (int idx = 0; idx < NO_THREADS; idx ++) {
      if (idx % 10 == 0) {
        pthread_create(&(tid[idx]), NULL, &(atomicRead<Counter_RWLocker>), cnt.get());
      } else {
        pthread_create(&(tid[idx]), NULL, &(atomicAdd<Counter_RWLocker>), cnt.get());
      }
    }

    for (int idx = 0; idx < NO_THREADS; idx ++) pthread_join(tid[idx], NULL);
  }
}

static void BM_SimpleRWLocker_WriteIntensive(benchmark::State& state) {
  for (auto _ : state) {
    auto cnt = std::make_unique<Counter_SimpleRWLocker>(10);
    pthread_t tid[NO_THREADS];

    for (int idx = 0; idx < NO_THREADS; idx ++) {
      if (idx % 10 == 0) {
        pthread_create(&(tid[idx]), NULL, &(atomicAdd<Counter_SimpleRWLocker>), cnt.get());
      } else {
        pthread_create(&(tid[idx]), NULL, &(atomicRead<Counter_SimpleRWLocker>), cnt.get());
      }
    }

    for (int idx = 0; idx < NO_THREADS; idx ++) pthread_join(tid[idx], NULL);
  }
}

static void BM_RWLocker_WriteIntensive(benchmark::State& state) {
  for (auto _ : state) {
    auto cnt = std::make_unique<Counter_RWLocker>(10);
    pthread_t tid[NO_THREADS];

    for (int idx = 0; idx < NO_THREADS; idx ++) {
      if (idx % 10 == 0) {
        pthread_create(&(tid[idx]), NULL, &(atomicAdd<Counter_RWLocker>), cnt.get());
      } else {
        pthread_create(&(tid[idx]), NULL, &(atomicRead<Counter_RWLocker>), cnt.get());
      }
    }

    for (int idx = 0; idx < NO_THREADS; idx ++) pthread_join(tid[idx], NULL);
  }
}

BENCHMARK(BM_SimpleRWLocker_ReadIntensive);
BENCHMARK(BM_SimpleRWLocker_WriteIntensive);
BENCHMARK(BM_RWLocker_ReadIntensive);
BENCHMARK(BM_RWLocker_WriteIntensive);

BENCHMARK_MAIN();