package alph_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/skiarn/practice-go/coincidence/alph"
)

func TestCoincidence(t *testing.T) {
	fmt.Println("Coincidence\nor\nNot")
	fmt.Println("If")
	fmt.Println(alph.Alph)
	fmt.Println("Equals")
	var nrs []int
	for i, _ := range alph.Alph {
		nrs = append(nrs, i+1)
	}
	fmt.Println(nrs)

	fmt.Println("Then")

	alph.Calculate(strings.NewReader("KNOWLEDGE"), os.Stdout)

	alph.Calculate(strings.NewReader("HARDWORK"), os.Stdout)
	fmt.Println("both are importatnt, but fall just short of 100%")
	fmt.Println("But")
	alph.Calculate(strings.NewReader("ATTITUDE"), os.Stdout)

}
