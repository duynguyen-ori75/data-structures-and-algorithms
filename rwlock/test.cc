#include "test_infras.hh"

#include <gtest/gtest.h>

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