package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNum(r *bufio.Reader) int {
	fmt.Print("please enter a number:")
	line, _, _ := r.ReadLine()
	num, _ := strconv.Atoi(string(line))
	return num
}

func divide(n, d int) int {
	return n / d
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Numerator  ")
	numerator := getNum(r)
	fmt.Print("Denominator  ")
	denominator := getNum(r)
	result := divide(numerator, denominator)
	fmt.Printf("%d/%d = %d\n", numerator, denominator, result)
}
