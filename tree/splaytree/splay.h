typedef struct SplayNode SplayNode;

struct SplayNode {
  i64 key;
  i64 countLeft, countRight;
  SplayNode *left, *right;
};

SplayNode *NewSplayNode(i64 key);
SplayNode *SplaySearch(SplayNode*, i64 key);
SplayNode *SplayInsert(SplayNode*, i64 key);
SplayNode *SplayDelete(SplayNode*, i64 key);