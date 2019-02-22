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
	a, _ := strconv.Atoi(line())
	l2 := strings.Split(line(), " ")
	b, _ := strconv.Atoi(l2[0])
	c, _ := strconv.Atoi(l2[1])
	num := a + b + c
	l3 := line()
	fmt.Printf("%d %s\n", num, l3)
}
