#pragma once

#include <utility>

class BlockDataStructure {
 public:
  bool Insert(int key, int value) = 0;
  bool Remove(int key) = 0;
  int Search(int key) = 0;
}