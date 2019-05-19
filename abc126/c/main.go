package main

import (
	"bufio"
	"fmt"
	"math"
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
	n_, _ := strconv.Atoi(input[0])
	k_, _ := strconv.Atoi(input[1])
	n := float64(n_)
	k := float64(k_)

	result := 0.0
	for i := 1; i <= n_; i++ {
		times := math.Log2(k / float64(i))
		var tmp float64
		if times < 0.0 || int(k) == i {
			tmp = 1.0 / n
		} else {
			tmp = 1.0 / n * math.Pow(0.5, float64(math.Ceil(times)))
		}
		result += tmp
		//fmt.Printf("%d %f %f\n", i, tmp, result)
	}

	fmt.Println(result)
}
