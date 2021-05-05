#include <stdio.h>


int increment(void);
int decrement(void);
int chcnt(int diff);

int main() {
  increment();
  increment();
  increment();
  decrement();
  int result = decrement();
  printf("%d\n", result);
}

int increment(void) {
  return chcnt(1);
}

int decrement(void) {
  return chcnt(-1);;
}

int chcnt(int diff) {
  static int count = 0;
  count += diff;
  return count;
}

