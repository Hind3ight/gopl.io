// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++ // 找到key对应的v，进行计数
		// NOTE: ignoring potential errors from input.Err()
		for line, n := range counts {
			if n > 0 {
				fmt.Printf("key值为%s的map存储了第%d次值\n", line, n)
			}
		}
	}
}

//!-
