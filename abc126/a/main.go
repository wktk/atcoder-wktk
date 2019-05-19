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
	k, _ := strconv.Atoi(input[1])
	str := []byte(line())
	char := strings.ToLower(string(str[k-1]))
	str[k-1] = []byte(char)[0]
	fmt.Println(string(str))
}
