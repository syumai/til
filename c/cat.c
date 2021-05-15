#include <fcntl.h>
#include <unistd.h>
#include <stdlib.h>

#define	BUFSIZ  1024

int copy(int dst_fd, int src_fd);

int main(int argc, char *argv[]) {
  int fd, result;
  if(argc == 1) {
    result = copy(STDOUT_FILENO, STDIN_FILENO);
    if(result < 0) {
      return result;
    }
    return 0;
  }
  while(--argc > 0) {
    fd = open(*++argv, O_RDONLY);
    result = copy(STDOUT_FILENO, fd);
    if(result < 0) {
      return result;
    }
    result = close(fd);
    if(result < 0) {
      return result;
    }
  }
}

int copy(int dst_fd, int src_fd) {
  int n;
  char buf[BUFSIZ];
  while((n = read(src_fd, buf, BUFSIZ)) > 0)
    write(dst_fd, buf, n);
  if(n < 0) {
    return n;
  }
  return 0;
}

