#include <unistd.h>
#define	BUFSIZ  64

int main() {
  int n;
  char buf[BUFSIZ];
  while((n = read(STDIN_FILENO, buf, BUFSIZ)))
    write(STDOUT_FILENO, buf, n);
}
