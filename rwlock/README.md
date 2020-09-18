# Simple RWLock implementation in C

Inspired by this excellent blog: https://eli.thegreenplace.net/2019/implementing-reader-writer-locks/

## Compile command

```shell
g++ rwlock_test.cc -lpthread -lgtest -o output
./output
```

## Requirements

- [googletest](https://github.com/google/googletest)