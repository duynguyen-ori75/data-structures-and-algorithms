# Splay Tree implementation in C

Based on exceptional Splay Tree in www.geeksforgeeks.org.

New features:
- The number of child nodes on both sub-tree of any node (`SplayNode*->countLeft` & `SplayNode*->countRight`)

## Definition

A splay tree is a binary search tree with the additional property that recently accessed elements are quick to access again. Like self-balancing binary search trees, a splay tree performs basic operations such as insertion, look-up and removal in O(log n) amortized time. For many sequences of non-random operations, splay trees perform better than other search trees, even performing better than O(log n) for sufficiently non-random patterns, all without requiring advance knowledge of the pattern. The splay tree was invented by Daniel Sleator and Robert Tarjan in 1985.

![Splay tree](http://lcm.csa.iisc.ernet.in/dsa/img199.gif)

## Compile command

```shell
make
```

## Requirements

- [googletest](https://github.com/google/googletest)
- [googlebenchmark](https://github.com/google/benchmark)

## References
- https://www.geeksforgeeks.org/splay-tree-set-1-insert/
- https://en.wikipedia.org/wiki/Splay_tree