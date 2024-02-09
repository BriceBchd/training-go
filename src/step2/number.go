package main

import (
	"fmt"
	"math"
	"strconv"
)

var integer8 int8 = 127
var integer16 int16 = 32767
var integer32 int32 = 2147483647
var integer64 int64 = 9223372036854775807

var floats32 float32 = 2147483647
var floats64 float64 = 9223372036854775807

const e = 2.71828
const Avogadro = 6.02214129e23
const Planck = 6.62606957e-34

var featureFlag bool = true

var defaultInt int
var defaultFloat32 float32
var defaultFloat64 float64
var defaultBool bool
var defaultString string

var rune = 'G'

func main() {
	fmt.Println(integer16, integer32)
	fmt.Println(rune)
	fmt.Println(math.MaxFloat32, math.MaxFloat64)
	fmt.Println(Avogadro + Planck)
	fmt.Println(featureFlag)
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)
	
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
