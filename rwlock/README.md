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
------------------------------------------------------------
Benchmark                  Time             CPU   Iterations
------------------------------------------------------------
BM_SimpleRWLocker      34817 ns        28443 ns        24878
BM_RWLocker            34082 ns        28239 ns        24629
```

### 10 threads

```shell
------------------------------------------------------------
Benchmark                  Time             CPU   Iterations
------------------------------------------------------------
BM_SimpleRWLocker     106091 ns       101752 ns         6861
BM_RWLocker           109094 ns       104258 ns         6910
```

### 1000 threads

```shell
------------------------------------------------------------
Benchmark                  Time             CPU   Iterations
------------------------------------------------------------
BM_SimpleRWLocker   13561469 ns     13512624 ns           52
BM_RWLocker         13624746 ns     13619034 ns           52
```