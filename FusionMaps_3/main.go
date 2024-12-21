package main

import (
	"fmt"
)

var ROMAN_NUM = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func main() {
	value := romainToInteger("XXI")
	fmt.Println(value) 
	fmt.Println(romainToInteger("XIV"))   
	fmt.Println(romainToInteger("IX")) 
	fmt.Println(romainToInteger("XL"))   
	fmt.Println(romainToInteger("C"))      
	fmt.Println(romainToInteger("MCMXCIV"))
}

func romainToInteger(s string) int {
	sum := 0
	n := len(s)

	for i := 0; i < n; i++ {
		if i < n-1 && ROMAN_NUM[s[i]] < ROMAN_NUM[s[i+1]] {
			sum -= ROMAN_NUM[s[i]]
		} else {
			sum += ROMAN_NUM[s[i]]
		}
	}

	return sum
}
