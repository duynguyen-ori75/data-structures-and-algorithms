# Lock-free stack benchmark

## Information

There are a lof of online resources about this topic, and one of them is [Introduction to Lock-free Programming](https://preshing.com/20120612/an-introduction-to-lock-free-programming/).

To sum up, do not attempt to implement Lock-free data structures, the performance is probably not as good as traditional mutex, unless you have a deep knowledge about this area.

## Benchmark info:

- Lenovo Thinkpad T490s
- CPU Core I5 - Ram 8Gb
- Cache information (in bytes)
- Push 80000 times and pop 10000 times

## Benchmark result:

```
goos: linux
goarch: amd64
pkg: github.com/duynguyen-ori75/playground/lockfree
BenchmarkSingleLockQueue_8_threads-8   	      25	  47138609 ns/op
BenchmarkTwoLockQueue_8_threads-8      	      26	  45014106 ns/op
BenchmarkNormalStack_4_threads-8       	     220	   5412046 ns/op
BenchmarkLockFreeStack_2_threads-8     	     176	   6926597 ns/op
BenchmarkLockFreeStack_4_threads-8     	     126	  10744846 ns/op
BenchmarkLockFreeStack_8_threads-8     	      93	  14817868 ns/op
PASS
ok  	github.com/duynguyen-ori75/playground/lockfree	9.817s
```

## References:
- https://en.wikipedia.org/wiki/Non-blocking_algorithm
- https://www.cs.cmu.edu/~410-s05/lectures/L31_LockFree.pdf
- https://preshing.com/20120612/an-introduction-to-lock-free-programming/
- http://15418.courses.cs.cmu.edu/spring2013/article/46