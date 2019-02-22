package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func line() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	count := 0
	input := []byte(line())

	for index := 0; index < len(input); index++ {
		if string(input[index]) == "1" {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}
