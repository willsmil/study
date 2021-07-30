package data_structure

import (
	"fmt"
	"testing"
)

func TestFloatPrint(t *testing.T)  {
	a := 1.0005
	b := 2.0005
	c := "2.0005"
	fmt.Printf("a: %.3f, b: %.3f\n", a, b)
	fmt.Println([]byte(c))
}
