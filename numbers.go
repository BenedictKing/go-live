package main

import (
	"fmt"
	"strconv"
)

func main() {
	number8, err := strconv.Atoi(string8)
	if err != nil {
		panic(err)
	}
	sum := num10 + number8 + one
	fmt.Println(sum)

	for number := 20; number <= 100; number++ {
		if number%3 == 0 {
			fmt.Println(number)
		}
	}

	total := 0
	for _, score := range scores {
		total += score
	}
	average := float64(total) / float64(len(scores))
	fmt.Println(average)
}

// 1. Provide the sum of the following variables
// Hint: one of them is a string, so you will need to convert it before adding

var num10 = 10
var string8 = "8"
var one = 1

// 2. Write a loop that will print only numbers divisible by 3 between 20 - 100

// 3. Find the class average given the following test scores

var scores = []int{88, 84, 100, 92, 70, 76, 76, 84, 86, 98}
