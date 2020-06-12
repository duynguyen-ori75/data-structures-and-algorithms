# Lock-free stack benchmark

## Benchmark info

- Lenovo Thinkpad T490s
- CPU Core I5 - Ram 8Gb
- Cache information (in bytes)
- Push 80000 times and pop 10000 times

## Benchmark result


```
goos: linux
goarch: amd64
BenchmarkNormalStack-8               	     205	   5789611 ns/op
BenchmarkLockFreeStack_2_threads-8   	     178	   6743428 ns/op
BenchmarkLockFreeStack_4_threads-8   	     141	   8547525 ns/op
BenchmarkLockFreeStack_8_threads-8   	      98	  13555513 ns/op
PASS
ok  	_/home/duynguyen/Workplace/learning/src/lockfree	7.056s
```