package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) > 3 {
		fmt.Println("We only accept 2 numbers")
		os.Exit(1)
	}

	ai, err := processArgs(os.Args[1:])

	if err != nil {
		fmt.Printf("This is an error: %s\n", err)
		os.Exit(1)
	}

	a := &Adder{
		num1: ai[0],
		num2: ai[1],
	}

	fmt.Println(a.Calculate())
}

func processArgs(args []string) ([]int, error) {
	ai := []int{}

	for _, a := range args {
		as, err := strconv.Atoi(a)

		if err != nil {
			return nil, err
		}

		ai = append(ai, as)
	}

	return ai, nil
}

type Adder struct {
	num1 int
	num2 int
}

func (a *Adder) Calculate() int {
	return a.num1 + a.num2
}

func (a *Adder) Numbers(nums []int) error {
	if len(nums) > 2 {
		return errors.New("You can only add 2 numbers")
	}

	a.num1 = nums[0]
	a.num2 = nums[1]

	return nil
}
