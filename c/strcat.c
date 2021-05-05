#include <stdio.h>

char *mystrcat(char *s, char *t);

int main() {
  char s[100] = "Hello,";
  char *t = " world!";
  printf("%s\n", mystrcat(s, t));
}

char *mystrcat(char *s, char *t) {
  char *sp = s;
  while(*(++s));
  while((*(s++) = *(t++)));
  return sp;
}
