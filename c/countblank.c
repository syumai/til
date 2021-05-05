#include <stdio.h>

int main() {
  int c, bc, tc, rc;
  c = getchar();
  while(c != EOF) {
    switch(c) {
      case ' ':
        ++bc;
        break;
      case '\t':
        ++tc;
        break;
      case '\n':
        ++rc;
        break;
    }
    c = getchar();
  }
  printf("blank: %d, tab: %d, return: %d\n", bc, tc, rc);
}
