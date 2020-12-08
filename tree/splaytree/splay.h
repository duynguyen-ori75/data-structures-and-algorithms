typedef struct SplayNode SplayNode;

struct SplayNode {
  int key;
  int countLeft, countRight;
  SplayNode *left, *right;
};

SplayNode *NewSplayNode(int key);
SplayNode *SplaySearch(SplayNode*, int key);
SplayNode *SplayInsert(SplayNode*, int key);
SplayNode *SplayDelete(SplayNode*, int key);
SplayNode *SplayLeftmostNode(SplayNode*);
SplayNode *SplayRightmostNode(SplayNode*);
SplayNode *SplayMoveNext(SplayNode*, int currentKey);
SplayNode *SplayMovePrevious(SplayNode*, int currentKey);