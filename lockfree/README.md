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
BenchmarkSingleLockQueue_4_threads-8   	      37	  29474997 ns/op
BenchmarkTwoLockQueue_4_threads-8      	      40	  30265734 ns/op
BenchmarkNormalStack_4_threads-8       	     207	   5692627 ns/op
BenchmarkLockFreeStack_2_threads-8     	     158	   7439368 ns/op
BenchmarkLockFreeStack_4_threads-8     	     100	  10112392 ns/op
BenchmarkLockFreeStack_8_threads-8     	      84	  15641695 ns/op
PASS
ok  	github.com/duynguyen-ori75/playground/lockfree	9.417s
```

## References:
- https://en.wikipedia.org/wiki/Non-blocking_algorithm
- https://www.cs.cmu.edu/~410-s05/lectures/L31_LockFree.pdf
- https://preshing.com/20120612/an-introduction-to-lock-free-programming/
- http://15418.courses.cs.cmu.edu/spring2013/article/46