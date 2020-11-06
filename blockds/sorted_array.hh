#include <algorithm>
#include <vector>

class SortedArray {
 private:
  std::vector<int> data_;
  int maxSize_, currentSize_;

  std::pair<int, int> cellPayload(int index) {
    return std::make_pair(this->data_[index * 2], this->data_[index * 2 + 1]);
  }

  int lookUp(int key) {
    int low = 0;
    int high = currentSize_ - 1;
    while (low <= high) {
      auto mid = (low + high) / 2;
      auto cell = cellPayload(mid);
      if (cell.first < key) low = mid + 1;
      else high = mid - 1;
    }
    return low;
  }

 public:
  SortedArray(int size) : maxSize_(size), currentSize_(0) { data_.resize(size * 2); }
  std::pair<bool, int> Search(int key) {
    auto index = this->lookUp(key);
    if (index >= this->currentSize_ || cellPayload(index).first != key)
      return std::make_pair(false, -1);
    return std::make_pair(true, cellPayload(index).second);
  }
  bool Insert(int key, int value) {
    auto index = this->lookUp(key);
    if (index < this->currentSize_ && cellPayload(index).first == key) {
      this->data_[index * 2 + 1] = value;
      return true;
    }
    if (this->currentSize_ >= this->maxSize_) return false;
    for (int idx = this->currentSize_; idx > index; idx --) {
      this->data_[idx * 2] = this->data_[(idx - 1) * 2];
      this->data_[idx * 2 + 1] = this->data_[(idx - 1) * 2 + 1];
    }
    this->data_[index * 2] = key;
    this->data_[index * 2 + 1] = value;
    this->currentSize_ ++;
    return true;
  }
  bool Remove(int key) {
    auto index = this->lookUp(key);
    if (index >= this->currentSize_ || cellPayload(index).first != key) return false;
    for (int idx = index; idx < this->currentSize_ - 1; idx ++) {
      this->data_[idx * 2] = this->data_[(idx + 1) * 2];
      this->data_[idx * 2 + 1] = this->data_[(idx + 1) * 2];
    }
    this->currentSize_ --;
    this->data_[this->currentSize_ * 2] = 0;
    this->data_[this->currentSize_ * 2 + 1] = 0;
    return true;
  }
};