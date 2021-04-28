#include <pthread.h>
#include <stdbool.h>

struct SimpleRWLocker {
  pthread_mutex_t lock_;
  int             count_;
};

void SimpleRLock(struct SimpleRWLocker *locker) {
  pthread_mutex_lock(&(locker->lock_));
  locker->count_ ++;
  pthread_mutex_unlock(&(locker->lock_));
}

void SimpleRUnlock(struct SimpleRWLocker *locker) {
  pthread_mutex_lock(&(locker->lock_));
  locker->count_ --;
  pthread_mutex_unlock(&(locker->lock_));
}

void SimpleWLock(struct SimpleRWLocker *locker) {
  while (true) {
    pthread_mutex_lock(&(locker->lock_));
    if (locker->count_ > 0) {
      pthread_mutex_unlock(&(locker->lock_));
    } else {
      break;
    }
  }
}

void SimpleWUnlock(struct SimpleRWLocker *locker) {
  pthread_mutex_unlock(&(locker->lock_));
}