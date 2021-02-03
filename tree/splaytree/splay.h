typedef struct SplayNode SplayNode;
typedef struct SplayTree SplayTree;

struct SplayNode {
  void *key;
  int countLeft, countRight;
  SplayNode *left, *right;
};

struct SplayTree {
  SplayNode *root;

  int (*nodeCmp)(SplayNode*, SplayNode*);
};

SplayTree *NewSplayTree(int (*cmpFunction)(SplayNode*, SplayNode*));
int SplayTreeIsEmpty(SplayTree*);
void SplayDestroy(SplayTree*);
void SplayInsert(SplayTree*, SplayNode *item);
void SplayDelete(SplayTree*, SplayNode *item);
SplayNode *SplaySearchGreater(SplayTree*, SplayNode *lookUp);
SplayNode *SplaySearchAtPosition(SplayTree*, int position);