#include <gtest/gtest.h>

extern "C" {
#include "splay.h"
}

void preOrder(SplayNode *root) {
  if (root != NULL) {
    printf("Node [key=%i]\n", *(int*)(root->key));
    preOrder(root->left);
    preOrder(root->right);
  }
}

int intComparison(SplayNode *lhs, SplayNode *rhs) {
  int leftKey   = *(int*)(lhs->key);
  int rightKey  = *(int*)(rhs->key);
  return leftKey - rightKey;
}

TEST(SplayTree, SplayInsertAndSearch) {
  SplayTree *tree = NewSplayTree(intComparison);
  SplayNode tmp, *ret;
  int data[100];

  for(int i = 2; i <= 200; i += 2) {
    data[i / 2 - 1] = i;
    tmp.key = &data[i / 2 - 1];
    SplayInsert(tree, &tmp);
  }

  for(int i = 1; i <= 200; i ++) {
    tmp.key = &i;
    ret = SplaySearchGreater(tree, &tmp);
    int expectedKey = (i % 2 == 1) ? i + 1 : i;
    ASSERT_EQ(ret, tree->root);
    ASSERT_EQ(*(int*)(ret->key), expectedKey);
  }

  SplayDestroy(tree);
}

TEST(SplayTree, SplayRemove) {
  SplayTree *tree = NewSplayTree(intComparison);
  SplayNode tmp, *ret;
  int data[100];

  for(int i = 2; i <= 200; i += 2) {
    data[i / 2 - 1] = i;
    tmp.key = &data[i / 2 - 1];
    SplayInsert(tree, &tmp);
  }

  for(int i = 1; i <= 200; i ++) {
    tmp.key = &i;
    SplayDelete(tree, &tmp);
    if (i % 2 == 1) {
      ASSERT_NE(tree->root, nullptr);
    } else if (i < 200) {
      ASSERT_NE(tree->root, nullptr);
      ret = SplaySearchGreater(tree, &tmp);
      ASSERT_EQ(ret, tree->root);
      ASSERT_EQ(*(int*)(ret->key), i + 2);
    } else {
      ASSERT_EQ(tree->root, nullptr);
    }
  }

  SplayDestroy(tree);
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}