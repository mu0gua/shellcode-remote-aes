package run

/*
#include <stdio.h>
#include <string.h>

void Wcall(char *sc)
{
  void (*fp)(void) = (void (*)(void))sc;
  printf("Length: %d\n",strlen(sc));
  fp();
}
*/
import "unsafe"
import "C"

func LinuxRun(sc []byte) {
	C.Wcall((*C.char)(unsafe.Pointer(&sc[0])))
}
