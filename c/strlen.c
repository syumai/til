#include <stdio.h>

int mystrlen(char *);

int main() {
  char *s = "Hello, world!";
  printf("%s, len: %d\n", s, mystrlen(s));
  printf("%s, len: %d\n", "", mystrlen(""));
  printf("%s, len: %d\n", "abcde", mystrlen("abcde"));
}

int mystrlen(char *s) {
  char *sp = s;
  while(*(s++));
  return s - sp - 1;
}
