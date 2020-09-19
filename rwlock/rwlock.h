#include <pthread.h>

#define MAX_READERS 50

struct RWLocker {
  pthread_mutex_t lock_;
  pthread_cond_t  writer_;  // cond var to notify writer
  pthread_cond_t  reader_;  // cond var to notify readers
  int             reader_count_;  // the number of running readers
  bool            writer_entered_;  // check whether there is a running-or-registered writer
};

void ReadLock(RWLocker *locker) {
  pthread_mutex_lock(&(locker->lock_));
  /**
   * there are two cases:
   * - there is a registered writer -> we should prioritize its execution -> all future readers should wait for its broadcast
   * - the number of concurrent readers reaches its limit -> the current reader should wait for notification from other readers
   */
  if (locker->reader_count_ >= MAX_READERS || locker->writer_entered_) {
    pthread_cond_wait(&(locker->reader_), &(locker->lock_));
  }
  locker->reader_count_ ++;
  pthread_mutex_unlock(&(locker->lock_));
}

void ReadUnlock(RWLocker *locker) {
  pthread_mutex_lock(&(locker->lock_));
  locker->reader_count_ --;
  // one writer registered to be executed
  if (locker->writer_entered_) {
    // only notify the writer if the number of readers is 0
    if (locker->reader_count_ == 0) {
      pthread_cond_signal(&(locker->writer_));
    }
  } else { // no registered writer so far
    if (locker->reader_count_ == MAX_READERS - 1) {
      // signal 1 reader to be executed
      pthread_cond_signal(&(locker->reader_));
    }
  }
  pthread_mutex_unlock(&(locker->lock_));
}

void WriteLock(RWLocker *locker) {
  pthread_mutex_lock(&(locker->lock_));
  /**
   * there is (at least) one running writer -> the current writer have to wait for it to finish
   * however, the finished writer will only broadcast all readers for simplicity, as it will be
   * more complicated to notify all readers and writers.
   * Therefore, the current thread should wait for signal on locker->readers
   */
  if (locker->writer_entered_) {
    pthread_cond_wait(&(locker->reader_), &(locker->lock_));
  }
  /**
   * have to (re)announce that a writer has been registered for it execution
   * to maintain write preference
   */
  locker->writer_entered_ = true;
  // there is some running readers -> the writer have to wait for them to finish their execution
  if (locker->reader_count_ > 0) {
    pthread_cond_wait(&(locker->writer_), &(locker->lock_));
  }
  pthread_mutex_unlock(&(locker->lock_));
}

void WriteUnlock(RWLocker *locker) {
  pthread_mutex_lock(&(locker->lock_));
  // announce that there should be no registered writer at the moment
  locker->writer_entered_ = false;
  pthread_cond_broadcast(&(locker->reader_));
  pthread_mutex_unlock(&(locker->lock_));
}