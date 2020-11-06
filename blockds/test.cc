#include <cstdlib>
#include <iostream>
#include <unordered_set>

#include <gtest/gtest.h>

#include "sorted_array.hh"
#include "slotted_page.hh"

#define MODULO 1000000
#define DS_SIZE 10000
#define NO_OPERATIONS 1000

TEST(SortedArray, Basic) {
  std::unordered_set<int> keys;
  auto sArray = SortedArray(DS_SIZE);

  for (int idx = 0; idx < NO_OPERATIONS; idx ++) {
    int key = std::rand() % MODULO;
    int value = std::rand() % MODULO;
    keys.insert(key);

    EXPECT_TRUE(sArray.Insert(key, value));
    auto result = sArray.Search(key);
    EXPECT_TRUE(result.first);
    EXPECT_EQ(result.second, value);
  }

  for(auto key: keys) {
    EXPECT_TRUE(sArray.Remove(key));
  }
  EXPECT_FALSE(sArray.Remove(123456));
}

TEST(SlottedPage, Basic) {
  std::unordered_set<int> keys;
  auto sPage = SlottedPage(DS_SIZE);

  for (int idx = 0; idx < NO_OPERATIONS; idx ++) {
    int key = std::rand() % MODULO;
    int value = std::rand() % MODULO;
    keys.insert(key);

    EXPECT_TRUE(sPage.Insert(key, value));
    auto result = sPage.Search(key);
    EXPECT_TRUE(result.first);
    EXPECT_EQ(result.second, value);
  }
  for(auto key: keys) {
    EXPECT_TRUE(sPage.Remove(key));
  }
  EXPECT_FALSE(sPage.Remove(123456));
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}