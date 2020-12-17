# Simple RWLock implementation in C

Inspired by this excellent blog: https://eli.thegreenplace.net/2019/implementing-reader-writer-locks/

## Implementation details

- Two versions: `SimpleRWLocker` and `RWLocker`
- `SimpleRWLocker` is implemented using a simple `pthread_mutex_t`
- `RWLocker` use several techniques
  - two conditional variables - `reader_` & `writer_` - to notify the threads to start the execution
  - a bool `writer_entered_` to notify all threads that a write thread is ready to run
    - no read thread can execute after the toggle
- Test and benchmark are done in C++

## Compile command

```shell
make
```

## Requirements

- [googletest](https://github.com/google/googletest)
- [googlebenchmark](https://github.com/google/benchmark)

## Benchmark

### 4 threads

```shell
---------------------------------------------------------------------------
Benchmark                                 Time             CPU   Iterations
---------------------------------------------------------------------------
BM_SimpleRWLocker_ReadIntensive       34481 ns        28212 ns        25228
BM_SimpleRWLocker_WriteIntensive      33950 ns        27629 ns        24695
BM_RWLocker_ReadIntensive             33799 ns        27663 ns        25346
BM_RWLocker_WriteIntensive            33785 ns        27656 ns        25382
```

### 10 threads

```shell
---------------------------------------------------------------------------
Benchmark                                 Time             CPU   Iterations
---------------------------------------------------------------------------
BM_SimpleRWLocker_ReadIntensive      106568 ns       101374 ns         6983
BM_SimpleRWLocker_WriteIntensive     107133 ns       101907 ns         6916
BM_RWLocker_ReadIntensive            105410 ns       100439 ns         6321
BM_RWLocker_WriteIntensive           103995 ns        98949 ns         6995
```

### 1000 threads

```shell
---------------------------------------------------------------------------
Benchmark                                 Time             CPU   Iterations
---------------------------------------------------------------------------
BM_SimpleRWLocker_ReadIntensive    13251754 ns     13147669 ns           53
BM_SimpleRWLocker_WriteIntensive   13382452 ns     13332255 ns           53
BM_RWLocker_ReadIntensive          13029691 ns     13029422 ns           53
BM_RWLocker_WriteIntensive         13042133 ns     13017316 ns           53
```