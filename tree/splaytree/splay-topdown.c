#include "splay.h"

#include <assert.h>
#include <stdlib.h>
#include <stdio.h>

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
 * @brief      Construct a new Splay node
 */
SplayNode *newSplayNode(void *key) {
  SplayNode* node = (SplayNode*)malloc(sizeof(SplayNode));
  node->key = key;
  node->left = node->right = NULL;
  return (node);
}

/**
 * @brief  This function brings the key at root if key is present in tree.
 *   If key is not present, then it brings the last accessed item at root.
 *   This function modifies the tree and returns the new root
 */
SplayNode *splay(SplayNode *root, SplayNode *lookUp, int (*nodeCmp)(SplayNode*, SplayNode*)) {
  if (!root) return root;
  SplayNode N;
  N.left = N.right = NULL;
  SplayNode *leftT, *rightT;
  leftT = rightT = &N;

  for (;;) {
    int cmp = nodeCmp(root, lookUp);
    if (cmp == 0) break;

    if (cmp > 0) {
      if (!root->left) break;

      if (nodeCmp(root->left, lookUp) > 0) {
        root = rightRotate(root);
        if (!root->left) break;
      }
      rightT->left = root;
      rightT = root;
      root = root->left;
    } else {
      if (!root->right) break;

      if (nodeCmp(root->right, lookUp) < 0) {
        root = leftRotate(root);
        if(!root->right) break;
      }

      leftT->right = root;
			leftT = root;
			root = root->right;
    }
  }

  leftT->right = root->left;
  rightT->left = root->right;

  root->left = N.right;
  root->right = N.left;
  return root;
}

void destroySplayNode(SplayNode *cur) {
  if (!cur) return;
  destroySplayNode(cur->left);
  destroySplayNode(cur->right);
  free(cur);
}

/**
 * @brief    Below is the implementation of all public functions
 */
SplayTree *NewSplayTree(int (*cmpFunction)(SplayNode*, SplayNode*)) {
  SplayTree *ret = (SplayTree*)malloc(sizeof(SplayTree));
  ret->root = NULL;
  ret->nodeCmp = cmpFunction;
  return ret;
}

void SplayDestroy(SplayTree *tree) {
  destroySplayNode(tree->root);
  free(tree);
}

int SplayTreeIsEmpty(SplayTree *tree) {
  return (tree->root) ? 0 : 1;
}

SplayNode *SplaySearchGreater(SplayTree *tree, SplayNode *lookUp) {
  SplayNode *cur = tree->root;
  SplayNode *ret = NULL;

  while (cur) {
    if (tree->nodeCmp(cur, lookUp) == 0) {
      ret = cur;
      break;
    } else if (tree->nodeCmp(cur, lookUp) > 0) {
      ret = cur;
      cur = cur->left;
    } else cur = cur->right;
  }

  if (!ret) return ret;
  tree->root = splay(tree->root, ret, tree->nodeCmp);
  return tree->root;
}

void SplayInsert(SplayTree *tree, SplayNode *item) {
  if (tree->root == NULL) {
    tree->root = newSplayNode(item->key);
    return;
  }

  SplayNode *pTmp = SplaySearchGreater(tree, item);
  if (pTmp && tree->nodeCmp(pTmp, item) == 0) return;

  pTmp = tree->root;
  SplayNode *newNode = newSplayNode(item->key);
  if (tree->nodeCmp(pTmp, newNode) > 0) {
    newNode->right = pTmp;
    newNode->left = pTmp->left;
    pTmp->left = NULL;
  } else {
    newNode->left = pTmp;
    newNode->right = pTmp->right;
    pTmp->right = NULL;
  }

  tree->root = newNode;
}

void SplayDelete(SplayTree *tree, SplayNode *item) {
  if (!tree->root) return;

  SplayNode *pTmp = SplaySearchGreater(tree, item);
  if (pTmp && tree->nodeCmp(pTmp, item) != 0) return;

  // come here means that pTmp should be splay-ed to the root
  assert(pTmp == tree->root);

  SplayNode *temp;
  if (!pTmp->left) {
    temp = pTmp;
    pTmp = pTmp->right;
  } else {
    temp = pTmp;
    pTmp = splay(pTmp->left, item, tree->nodeCmp);
    pTmp->right = temp->right;
  }
  free(temp);
  tree->root = pTmp;
}