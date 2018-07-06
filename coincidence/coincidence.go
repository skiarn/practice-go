package main

import (
	"fmt"
)

var alph = [...]string{"A", "B", "C", "D", "E", "F",
	"G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z"}

func main() {
	fmt.Println("Coincidence\nor\nNot")
	fmt.Println("If")
	fmt.Println(alph)
	fmt.Println("Equals")
	var nrs []int
	for i, _ := range alph {
		nrs = append(nrs, i+1)
	}
	fmt.Println(nrs)
	wisdom := func(k string) {
		fmt.Println(k)
		tot := 0
		for _, c := range k {
			tot += getIndex(string(c))
			fmt.Printf("%v ", getIndex(string(c)))
		}
		fmt.Printf("= %v%s\n", tot, "%")
	}
	fmt.Println("Then")
	wisdom("KNOWLEDGE")
	wisdom("HARDWORK")
	fmt.Println("both are importatnt, but fall just short of 100%")
	fmt.Println("But")
	wisdom("ATTITUDE")
}

func getIndex(char string) int {
	for i, c := range alph {
		if string(c) == char {
			return i + 1
		}

	}
	return -1
}
