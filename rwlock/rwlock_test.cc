extern "C" {
#include "simple_rwlock.h"
#include "rwlock.h"
#include "pthread.h"
}
#include <gtest/gtest.h>

#define NO_THREADS 1000

#define TEST_PACKAGE(locker_type, RLockFunc, RUnlockFunc, WLockFunc, WUnlockFunc) \
  class Counter_##locker_type { \
    public: \
    locker_type   *lk; \
    int           counter_; \
    Counter_##locker_type(int default_counter) : counter_(default_counter), lk(new locker_type()) {}\
    ~Counter_##locker_type() { delete lk; } \
    void RLock() { RLockFunc(lk); } \
    void RUnlock() { RUnlockFunc(lk); } \
    void WLock() { WLockFunc(lk); } \
    void WUnlock() { WUnlockFunc(lk); } \
  } \

TEST_PACKAGE(SimpleRWLocker, SimpleRLock, SimpleRUnlock, SimpleWLock, SimpleWUnlock);
TEST_PACKAGE(RWLocker, ReadLock, ReadUnlock, WriteLock, WriteUnlock);

template <class locker_T>
void *atomicAdd(void *arg) {
  auto *data = (locker_T*) arg;
  data->WLock();
  data->counter_ ++;
  data->WUnlock();
  pthread_exit(NULL);
}

template <class locker_T>
void *atomicRead(void *arg) {
  auto *data = (locker_T*) arg;
  data->RLock();
  auto _ = data->counter_;
  data->RUnlock();
  pthread_exit(NULL);
}


TEST(SimpleReadWriteLockerTest, Basic) {
  auto cnt = std::make_unique<Counter_SimpleRWLocker>(10);
  pthread_t tid[NO_THREADS];

  for (int idx = 0; idx < NO_THREADS; idx ++) {
    if (idx % 2 == 0) {
      ASSERT_EQ(pthread_create(&(tid[idx]), NULL, &(atomicRead<Counter_SimpleRWLocker>), cnt.get()), 0);
    } else {
      ASSERT_EQ(pthread_create(&(tid[idx]), NULL, &(atomicAdd<Counter_SimpleRWLocker>), cnt.get()), 0);
    }
  }

  for (int idx = 0; idx < NO_THREADS; idx ++) pthread_join(tid[idx], NULL);
  EXPECT_EQ(cnt->counter_, NO_THREADS / 2 + 10);
}

TEST(ReadWriteLockerTest, Basic) {
  auto cnt = std::make_unique<Counter_RWLocker>(10);
  pthread_t tid[NO_THREADS];

  for (int idx = 0; idx < NO_THREADS; idx ++) {
    if (idx % 2 == 0) {
      ASSERT_EQ(pthread_create(&(tid[idx]), NULL, &(atomicRead<Counter_RWLocker>), cnt.get()), 0);
    } else {
      ASSERT_EQ(pthread_create(&(tid[idx]), NULL, &(atomicAdd<Counter_RWLocker>), cnt.get()), 0);
    }
  }

  for (int idx = 0; idx < NO_THREADS; idx ++) pthread_join(tid[idx], NULL);
  EXPECT_EQ(cnt->counter_, NO_THREADS / 2 + 10);
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}