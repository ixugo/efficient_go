package main

/*
#include <stdio.h>

void SayHello(const char *s){
	puts(s);
}
*/
import "C"

func main() {
	s := C.CString("Hello world\n")
	C.SayHello(s)
}
