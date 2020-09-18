extern "C" {
#include "rwlock.h"
#include "pthread.h"
}
#include <gtest/gtest.h>

struct Counter {
  SimpleRWLocker  *lk;
  int             counter;
};

void *atomicAdd(void *arg) {
  auto *data = (Counter*) arg;
  SimpleWLock(data->lk);
  data->counter ++;
  SimpleWUnlock(data->lk);
  pthread_exit(NULL);
}

void *atomicRead(void *arg) {
  auto *data = (Counter*) arg;
  SimpleRLock(data->lk);
  auto _ = data->counter;
  SimpleRUnlock(data->lk);
  pthread_exit(NULL);
}

TEST(SimpleReadWriteLockerTest, Basic) {
  auto cnt = new Counter{new SimpleRWLocker(), 10};
  pthread_t tid[100];

  // create 50 read threads and 50 write threads
  for (int idx = 0; idx < 100; idx ++) {
    if (idx % 2 == 0) {
      ASSERT_EQ(pthread_create(&(tid[idx]), NULL, &atomicAdd, cnt), 0);
    } else {
      ASSERT_EQ(pthread_create(&(tid[idx]), NULL, &atomicRead, cnt), 0);
    }
  }

  // join all threads
  for (int idx = 0; idx < 100; idx ++) pthread_join(tid[idx], NULL);

  // expect atomic counter to be correct
  EXPECT_EQ(cnt->counter, 60);
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}