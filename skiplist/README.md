# Skip list

## Definition:

A skiplist is an ordered data structure providing expected O(Log(n)) lookup, insertion, and deletion complexity. It provides this level of efficiency without the need for complex tree balancing or page splitting like that required by Btrees, redblack trees, or AVL trees. As a result, it’s a much simpler and more concise data structure to implement.

![Visualization](https://upload.wikimedia.org/wikipedia/commons/thumb/8/86/Skip_list.svg/800px-Skip_list.svg.png "Skip list")

## Implementation:
- SkipListPointers: Every node in Skip list is pointer, even nodes that have the same key (nodes in the same column)
- SkipList: Only one node per key, and each node maintains its own next pointers

## TODO:
- Finish SkipList implementation
- Benchmark two versions

## References:
- https://en.wikipedia.org/wiki/Skip_list
- https://www.cl.cam.ac.uk/teaching/2005/Algorithms/skiplists.pdf
- [MemSQL blog about Skip list](https://www.memsql.com/blog/what-is-skiplist-why-skiplist-index-for-memsql/#:~:text=MemSQL's%20skiplist%20employs%20a%20novel,particular%20node%2C%20of%20the%20skiplist.)
- https://www.cs.cmu.edu/~ckingsf/bioinfo-lectures/skiplists.pdf