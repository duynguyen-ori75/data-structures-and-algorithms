#include <gtest/gtest.h>

extern "C" {
#include "splay.h"
}

void preOrder(SplayNode *root) {
  if (root != NULL) {
    printf("%i (%i, %i)\n", root->key, root->countLeft, root->countRight);
    preOrder(root->left);
    preOrder(root->right);
  }
}

TEST(SplayTree, SplayInsert) {
  SplayNode *root = NewSplayNode(1);
  for(int i = 2; i <= 100; i ++)
    root = SplayInsert(root, i);
  for(int i = 1; i <= 100; i ++) {
    root = SplaySearch(root, i);
    ASSERT_EQ(root->key, i);
    ASSERT_EQ(root->countLeft, i - 1);
    ASSERT_EQ(root->countRight, 100 - i);
  }
}

TEST(SplayTree, SplayRemove) {
  SplayNode *root = NewSplayNode(1);
  for(int i = 2; i <= 100; i ++)
    root = SplayInsert(root, i);
  for(int i = 1; i <= 100; i ++) {
    root = SplayDelete(root, i);
    if (root != NULL) {
      ASSERT_NE(root->key, i);
      root = SplaySearch(root, i);
      ASSERT_EQ(root->key, i + 1);
      ASSERT_EQ(root->countLeft, 0);
      ASSERT_EQ(root->countRight, 99 - i);
    }
  }
  ASSERT_TRUE(root == NULL);
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}