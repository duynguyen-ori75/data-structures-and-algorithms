# B+Tree implementation

## Definition:

B+tree is an improvement of traditional B-tree, which is a self-balancing tree data structure that maintains sorted data and allows searches, sequential access, insertions, and deletions in logarithmic time. The B-tree generalizes the binary search tree, allowing for nodes with more than two children.

[Interactive visualization](https://www.cs.usfca.edu/~galles/visualization/BPlusTree.html)

## Implementation info:
- Single thread
- All keys should be unique
- Assume that all keys and values are `int`
- Does not support delete the right-most child (too lazy to handle this corner case)

## References:
- https://en.wikipedia.org/wiki/B%2B_tree
- https://www.cs.usfca.edu/~galles/visualization/BPlusTree.html
- https://www.geeksforgeeks.org/introduction-of-b-tree/
- http://pages.cs.wisc.edu/~paris/cs564-f15/lectures/lecture-12.pdf