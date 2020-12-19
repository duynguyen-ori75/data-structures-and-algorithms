## Implementation of various block-based data structures

Including `sorted array` and `slotted page`

Benchmark result

```shell
Run on (8 X 3900 MHz CPU s)
CPU Caches:
  L1 Data 32 KiB (x4)
  L1 Instruction 32 KiB (x4)
  L2 Unified 256 KiB (x4)
  L3 Unified 6144 KiB (x1)
Load Average: 0.69, 0.82, 0.81
***WARNING*** CPU scaling is enabled, the benchmark real time measurements may be noisy and will incur extra overhead.
-----------------------------------------------------------------
Benchmark                       Time             CPU   Iterations
-----------------------------------------------------------------
BM_SortedArray_Insert   183208629 ns    183173992 ns            4
BM_SlottedPage_Insert   108501777 ns    108498235 ns            7
BM_SortedArray_Search     3497068 ns      3497010 ns          200
BM_SlottedPage_Search     3950770 ns      3950684 ns          177
BM_SortedArray_Generic  377638498 ns    377630396 ns            2
BM_SlottedPage_Generic  210670045 ns    210667168 ns            3
```

### Sorted array

![sorted array](https://media.geeksforgeeks.org/wp-content/cdn-uploads/Insert-Operation-in-Sorted-Array.png)

### Slotted page

![slotted page](https://media.vlpt.us/images/gwak2837/post/f3c9156a-7605-4dbd-86d5-c13c034a78d4/image.png)