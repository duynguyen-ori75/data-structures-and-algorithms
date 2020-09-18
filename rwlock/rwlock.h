#include <pthread.h>

struct SimpleRWLocker {
  pthread_mutex_t lock_;
  int             count_;
};

void SimpleRLock(SimpleRWLocker *locker) {
  pthread_mutex_lock(&(locker->lock_));
  locker->count_ ++;
  pthread_mutex_unlock(&(locker->lock_));
}

void SimpleRUnlock(SimpleRWLocker *locker) {
  pthread_mutex_lock(&(locker->lock_));
  locker->count_ --;
  pthread_mutex_unlock(&(locker->lock_));
}

void SimpleWLock(SimpleRWLocker *locker) {
  while (true) {
    pthread_mutex_lock(&(locker->lock_));
    if (locker->count_ > 0) {
      pthread_mutex_unlock(&(locker->lock_));
    } else {
      break;
    }
  }
}

void SimpleWUnlock(SimpleRWLocker *locker) {
  pthread_mutex_unlock(&(locker->lock_));
}