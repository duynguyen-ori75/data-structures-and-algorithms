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
pkg: github.com/duynguyen-ori75/playground/concurrency
BenchmarkSingleLockQueue-8           	      22	  47692904 ns/op
BenchmarkTwoLockQueue-8              	      27	  43281536 ns/op
BenchmarkLockFreeQueue-8             	      13	  77679139 ns/op
BenchmarkNormalStack_4_threads-8     	     222	   5400529 ns/op
BenchmarkLockFreeStack_2_threads-8   	     177	   6791986 ns/op
BenchmarkLockFreeStack_4_threads-8   	     132	   9193350 ns/op
BenchmarkLockFreeStack_8_threads-8   	      74	  14258238 ns/op
PASS
ok  	github.com/duynguyen-ori75/playground/concurrency	10.302s
```

## References:
- https://en.wikipedia.org/wiki/Non-blocking_algorithm
- https://www.cs.cmu.edu/~410-s05/lectures/L31_LockFree.pdf
- https://preshing.com/20120612/an-introduction-to-lock-free-programming/
- http://15418.courses.cs.cmu.edu/spring2013/article/46
- https://www.cs.rochester.edu/~scott/papers/1996_PODC_queues.pdf