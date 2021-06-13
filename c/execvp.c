#include <unistd.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
  int ret;
  ret = execvp(argv[1], argv+1);
  if (ret == -1)
    perror("execvp");
  return 0;
}
