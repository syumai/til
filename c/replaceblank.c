#include <stdio.h>

int main() {
  int c = getchar();
  while(c != EOF) {
    switch(c) {
      case '\\':
        printf("\\");
        break;
      case '\t':
        printf("\\t");
        break;
      case '\b':
        printf("\\b");
        break;
      default:
        putchar(c);
        break;
    }
    c = getchar();
  }
}
