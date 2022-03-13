// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	start := time.Now().Unix()
	for i := 0; i <= 100000; i++ {
		fmt.Println(strings.Join(os.Args[0:], " "))
	}
	end := time.Now().Unix()
	fmt.Printf("消耗的时间为%v", end-start)
}

//!-
