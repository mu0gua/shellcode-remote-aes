package run

/*
#include <stdio.h>
#include <string.h>

void Lcall(char *sc)
{
  void (*fp)(void) = (void (*)(void))sc;
  printf("Length: %d\n",strlen(sc));
  fp();
}
*/
import "unsafe"
import "C"

func LinuxRun(sc []byte) {
	C.Lcall((*C.char)(unsafe.Pointer(&sc[0])))
}
