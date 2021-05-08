#include <fcntl.h>
#define PERMS 0600

int main() {
  int fd;
  fd = creat("example.txt", PERMS);
}
