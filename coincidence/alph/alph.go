package alph

import (
	"fmt"
	"io"
)

//Alph is english alphabet
var Alph = [...]string{"A", "B", "C", "D", "E", "F",
	"G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z"}

//Calculate calculates the value of a input and writes to outut.
func Calculate(in io.Reader, out io.Writer) error {
	tot := 0
	b := make([]byte, 1)
	for {
		n, err := in.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		c := b[:n]
		i := Index(string(c))
		out.Write([]byte(fmt.Sprintf("%v ", i)))
		tot += i

	}
	_, err := out.Write([]byte(fmt.Sprintf("= %v%s\n", tot, "%")))
	return err
}

//Index returns index of char in alphabet, starting from 1 - 26.
func Index(char string) int {
	for i, c := range Alph {
		if string(c) == char {
			return i + 1
		}

	}
	return -1
}
