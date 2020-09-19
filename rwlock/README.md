# Simple RWLock implementation in C

Inspired by this excellent blog: https://eli.thegreenplace.net/2019/implementing-reader-writer-locks/

## Implementation details

- Two versions: `SimpleRWLocker` and `RWLocker`
- `SimpleRWLocker` is implemented using a simple `pthread_mutex_t`
- `RWLocker` use several techniques to achieve higher performance
  - two conditional variables - `reader_` & `writer_` - to notify the threads to start the execution
  - a bool `writer_entered_` to notify all threads that a write thread is ready to run
    - no read thread can execute after the toggle
- Testing is done in C++

## Compile command

```shell
make
```

## Requirements

- [googletest](https://github.com/google/googletest)