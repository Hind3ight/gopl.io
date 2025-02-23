// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for k, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Printf("第%v个元素为%v", k, s)
	}
}

//!-
