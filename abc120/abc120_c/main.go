package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func line() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	var zero float64
	input := line()
	x := float64(len(input))
	for i := 0; i < len(input); i++ {
		if string(input[i]) == "0" {
			zero++
		}
	}
	min := math.Min(zero, x-zero)
	fmt.Printf("%d", int(min*2))
}
