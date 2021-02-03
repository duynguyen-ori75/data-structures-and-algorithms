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
  // update count
  x->countLeft = y->countRight;
  y->countRight = x->countLeft + x->countRight + 1;
  return y;
}

/**
 * @brief  A utility function to left rotate subtree rooted with x
 */
SplayNode *leftRotate(SplayNode *x) {
  SplayNode *y = x->right;
  x->right = y->left;
  y->left = x;
  // update count
  x->countRight = y->countLeft;
  y->countLeft = x->countLeft + x->countRight + 1;
  return y;
}

/**
 * @brief      Construct a new Splay node
 */
SplayNode *newSplayNode(void *key) {
  SplayNode* node = (SplayNode*)malloc(sizeof(SplayNode));
  node->key = key;
  node->countLeft = node->countRight = 0;
  node->left = node->right = NULL;
  return (node);
}

/**
 * @brief  This function brings the key at root if key is present in tree.
 *   If key is not present, then it brings the last accessed item at root.
 *   This function modifies the tree and returns the new root
 */
SplayNode *splay(SplayNode *root, SplayNode *lookUp, int (*nodeCmp)(SplayNode*, SplayNode*)) {
  if (root == NULL || nodeCmp(root, lookUp) == 0) return root;

  if (nodeCmp(root, lookUp) > 0) {
    if (root->left == NULL) return root;

    if (nodeCmp(root->left, lookUp) > 0) {
      root->left->left = splay(root->left->left, lookUp, nodeCmp);
      root = rightRotate(root);
    }
    else if (nodeCmp(root->left, lookUp) < 0) {
      root->left->right = splay(root->left->right, lookUp, nodeCmp);
      if (root->left->right != NULL)
        root->left = leftRotate(root->left);
    }

    return (root->left == NULL) ? root: rightRotate(root);
  }

  if (root->right == NULL) return root;

  if (nodeCmp(root->right, lookUp) > 0) {
    root->right->left = splay(root->right->left, lookUp, nodeCmp);
    if (root->right->left != NULL)
      root->right = rightRotate(root->right);
  }
  else if (nodeCmp(root->right, lookUp) < 0) {
    root->right->right = splay(root->right->right, lookUp, nodeCmp);
    root = leftRotate(root);
  }

  return (root->right == NULL) ? root: leftRotate(root);
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
    // update count
    newNode->countRight = pTmp->countRight + 1;
    newNode->countLeft = pTmp->countLeft;
    pTmp->countLeft = 0;
  } else {
    newNode->left = pTmp;
    newNode->right = pTmp->right;
    pTmp->right = NULL;
    // update count
    newNode->countLeft = pTmp->countLeft + 1;
    newNode->countRight = pTmp->countRight;
    pTmp->countRight = 0;
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
    pTmp->countRight = temp->countRight;
  }
  free(temp);
  tree->root = pTmp;
}

SplayNode *SplaySearchAtPosition(SplayTree *tree, int position) {
  if (!tree->root || tree->root->countLeft + tree->root->countRight + 1 < position) return NULL;

  SplayNode *cur = tree->root;
  SplayNode *ret = NULL;

  while (cur) {
    if (cur->countLeft + 1 == position) {
      ret = cur;
      break;
    } else if (cur->countLeft + 1 >= position) {
      cur = cur->left;
    } else {
      position -= cur->countLeft + 1;
      cur = cur->right;
    }
  }

  if (!ret) return ret;
  tree->root = splay(tree->root, ret, tree->nodeCmp);
  return tree->root;
}