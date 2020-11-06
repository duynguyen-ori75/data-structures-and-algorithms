#include <algorithm>
#include <vector>

typedef std::pair<int, int> KeyValue;

class SortedArray {
 private:
  std::vector<KeyValue> data_;
  int size_;

  std::vector<KeyValue>::iterator lookUp(int key) {
    return std::lower_bound(this->data_.begin(), this->data_.end(),
                             std::make_pair(key, -1),
                             [](const KeyValue& lhs, const KeyValue& rhs) {
                               return lhs.first < rhs.first;
                             });
  }

 public:
  SortedArray(int size) : size_(size) {}
  std::pair<bool, int> Search(int key) {
    auto it = this->lookUp(key);
    if (it == this->data_.end() || it->first != key)
      return std::make_pair(false, -1);
    return std::make_pair(true, it->second);
  }
  bool Insert(int key, int value) {
    auto it = this->lookUp(key);
    if (it != this->data_.end() && it->first == key) {
      it->second = value;
      return true;
    }
    if (this->data_.size() >= this->size_) return false;
    this->data_.insert(it, std::make_pair(key, value));
    return true;
  }
  bool Remove(int key) {
    auto it = this->lookUp(key);
    if (it == this->data_.end() || it->first != key) return false;
    this->data_.erase(it);
    return true;
  }
};