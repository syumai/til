#include <stdio.h>

char *mystrcat(char *s, char *t);

int main() {
  char s1[100] = "Hello";
  char *t1 = " world!";
  printf("%s\n", mystrcat(s1, t1));
  char s2[100] = "";
  char *t2 = "Hello world, with empty string!";
  printf("%s\n", mystrcat(s2, t2));
}

char *mystrcat(char *s, char *t) {
  char *sp = s;
  for (; *s; ++s);
  while ((*s++ = *t++));
  return sp;
}
