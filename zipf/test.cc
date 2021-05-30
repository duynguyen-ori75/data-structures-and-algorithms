#include "zipf.h"

#include <stdio.h>

int main(int argc, char **argv) {
  ZipfGenerator *p = ZipfGeneratorNew(0, 100);
  for(int i = 0; i < 100; i++) {
    printf("Generate time %d: %d\n", i, ZipfGenerate(p));
  }
}