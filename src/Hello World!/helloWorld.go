package main

import (
	"fmt"
	"os"
	//"strings"
)

func main () {
	t, sep := "", " "
	i := 0
	for i, j := range os.Args {
		fmt.Println(i)
		t += j + sep
	}
	fmt.Println(t)
	for { //An infinite loop
		i ++;
	}
}