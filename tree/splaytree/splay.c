#include "splay.h"

/**
 * @brief  A utility function to right rotate subtree rooted with x
 */
SplayNode *rightRotate(SplayNode *x) {
  SplayNode *y = x->left;
  x->left = y->right;
  y->right = x;
  return y;
}

/**
 * @brief  A utility function to left rotate subtree rooted with x
 */
SplayNode *leftRotate(SplayNode *x) {
  SplayNode *y = x->right;
  x->right = y->left;
  y->left = x;
  return y;
}

/**
 * @brief  This function brings the key at root if key is present in tree.
 *   If key is not present, then it brings the last accessed item at root.
 *   This function modifies the tree and returns the new root
 */
SplayNode *splay(SplayNode *root, i64 key) {
  if (root == NULL || root->key == key) return root;

  if (root->key > key) {
    if (root->left == NULL) return root;

    if (root->left->key > key) {
      root->left->left = splay(root->left->left, key);
      root = rightRotate(root);
    }
    else if (root->left->key < key) {
      root->left->right = splay(root->left->right, key);
      if (root->left->right != NULL)
        root->left = leftRotate(root->left);
    }

    return (root->left == NULL)? root: rightRotate(root);
  }

  if (root->right == NULL) return root;

  if (root->right->key > key) {
    root->right->left = splay(root->right->left, key);
    if (root->right->left != NULL)
      root->right = rightRotate(root->right);
  }
  else if (root->right->key < key) {
    root->right->right = splay(root->right->right, key);
    root = leftRotate(root);
  }

  return (root->right == NULL)? root: leftRotate(root);
}

/**
 * @brief    Below is the implementation of all public functions
 */

SplayNode *NewSplayNode(i64 key) {
  SplayNode* node = (SplayNode*)malloc(sizeof(SplayNode));
  node->key = key;
  node->countLeft = node->countRight = 0;
  node->left = node->right = NULL;
  return (node);
}

SplayNode *SplaySearch(SplayNode *root, i64 key) {
  return splay(root, key);
}

SplayNode *SplayInsert(SplayNode *root, i64 key) {
  if (root == NULL) return NewSplayNode(key);

  root = splay(root, key);
  if (root->key == key) return root;

  SplayNode *newNode  = NewSplayNode(key);
  if (root->key > key) {
    newNode->right = root;
    newNode->left = root->left;
    root->left = NULL;
    // update count
    newNode->countRight = root->countRight + 1;
    newNode->countLeft = root->countLeft;
    root->countLeft = 0;
  } else {
    newNode->left = root;
    newNode->right = root->right;
    root->right = NULL;
    // update count
    newNode->countLeft = root->countLeft + 1;
    newNode->countRight = root->countRight;
    root->countRight = 0;
  }

  return newNode;
}

SplayNode *SplayDelete(SplayNode *root, i64 key) {
  if (!root) return NULL;

  root = splay(root, key);
  if (key != root->key) return root;

  SplayNode *temp;
  if (!root->left) {
    temp = root;
    root = root->right;
  } else {
    temp = root;
    root = splay(root->left, key);
    root->right = temp->right;
    root->countRight = temp->countRight;
  }
  free(temp);
  return root;
}