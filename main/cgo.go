package main

/*
#include <stdio.h>

void printthing() {
printf("foo\n");
}
*/
import "C"

func main() {
	C.printthing()
	C.ACFunction()
}
