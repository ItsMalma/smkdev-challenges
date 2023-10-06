package main

import (
	"fmt"
	"os"
	"smkdev-challenges/lib"
)

func main() {
	inputNumber, err := lib.InputInt("input_number = ")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}

	for i := int64(1); i <= inputNumber; i++ {
		fmt.Printf("\nCurrent Number is : %v and the cube is %v\n", i, i*i*i)
	}
}
