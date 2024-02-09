package main

import (
	"fmt"

	"github.com/BriceBchd/calculator"
)

func main() {
    fmt.Println("Hello World!")

    total := calculator.Sum(3, 5)
    fmt.Println(total)
    fmt.Println("Version: " + calculator.Version)
}

const HTTPStatusOK = 200

const (
    StatusOK = 0
    StatusConnectionReset = 1
    StatusOtherError = 2
)
