#include <iostream>
#include <vector>

class SlottedPage {
 private:
  std::vector<int> data_;
  int maxSize_, currentSize_, payloadOffset_;

  std::pair<int, int> cellPayload(int offset) {
    return std::make_pair(this->data_[offset], this->data_[offset + 1]);
  }

  int lookUp(int key) {
    int low = 0;
    int high = currentSize_ - 1;
    while (low <= high) {
      auto mid = (low + high) / 2;
      auto cell = cellPayload(this->data_[mid]);
      if (cell.first < key)
        low = mid + 1;
      else
        high = mid - 1;
    }
    return low;
  }

 public:
  SlottedPage(int size)
      : maxSize_(size + 10), currentSize_(0), payloadOffset_(maxSize_ * 3 - 2) {
    data_.resize(maxSize_ * 3);
  }
  std::pair<bool, int> Search(int key) {
    auto index = this->lookUp(key);
    if (index >= this->currentSize_) return std::make_pair(false, -1);
    return std::make_pair(true, cellPayload(this->data_[index]).second);
  }
  bool Insert(int key, int value) {
    auto index = this->lookUp(key);
    if (index < this->currentSize_ &&
        cellPayload(this->data_[index]).first == key) {
      this->data_[this->data_[index] + 1] = value;
      return true;
    }
    if (this->currentSize_ >= this->maxSize_) return false;
    // inserting
    currentSize_++;
    this->payloadOffset_ -= 2;
    this->data_[this->payloadOffset_] = key;
    this->data_[this->payloadOffset_ + 1] = value;
    // shift-right elements
    for (int idx = currentSize_; idx > index; idx--)
      this->data_[idx] = this->data_[idx - 1];
    this->data_[index] = this->payloadOffset_;
    return true;
  }
  bool Remove(int key) {
    auto index = this->lookUp(key);
    if (index >= this->currentSize_ ||
        cellPayload(this->data_[index]).first != key)
      return false;
    // shift-left elements
    int deletedOffset = this->data_[index];
    this->currentSize_--;
    for (int idx = index; idx < this->currentSize_; idx++)
      this->data_[idx] = this->data_[idx + 1];
    this->data_[currentSize_] = 0;
    // empty payload cell
    this->data_[deletedOffset] = this->data_[deletedOffset + 1] = 0;
    return true;
  }
};