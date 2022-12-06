package main

/*
#include <stdio.h>

void myhello(int i){
	printf("Hello C:%d\n",i);
}

*/
import "C"
import "fmt"

func main() {
	C.myhello(C.int(12))
	fmt.Println("hellp C")
}
