package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
    sum := sum(os.Args[1], os.Args[2])
    fmt.Println("Sum:", sum)

	sum, mul := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum: ", sum)
	fmt.Println("Mul: ", mul)

    firstName := "John"
    updateName(&firstName)
    fmt.Println(firstName)
}

func sum(number1 string, number2 string) (result int) {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    result = int1 + int2
	return
}

func calc(number1 string, number2 string) (sum int, mul int) {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    sum = int1 + int2
    mul = int1 * int2
    return
}

func updateName(name *string) {
    *name = "David"
}