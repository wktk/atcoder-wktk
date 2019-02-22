package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func line() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	input := strings.Split(line(), " ")
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	if a*b&0x1 == 0x1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
