package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	age := 32
	fmt.Println(age)
	fmt.Println(Sunday, Monday)
	return
}
