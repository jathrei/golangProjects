package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 40
	y := 5

	//if statement
	if x < 42 {
		fmt.Println("x is less than the meaning of life")
	} else {
		fmt.Println("x is greater than or equal to the meaning of life")
	}
	//logical operators
	if (x < 42) && (y < 42) {
		fmt.Println("x and y are less than the meaning of life")
	} else {
		fmt.Println("x and y are greater than or equal to the meaning of life")
	}
	//statement; statement idiom
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Println("z is greater than or equal to x")
	} else {
		fmt.Println("z is less than x")
	}

	//switch statements
	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQULAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("this is default case for x")
	}

	//for loop init; condition; post (after loop completes)
	for i := 0; i < 5; i++ {
		fmt.Println("count to five: ", i)
	}

	//for loop (condition) -- post is in the loop
	for y < 10 {
		fmt.Println(" y is value : ", y)
		y++
	}

	//for loop (every condition is inside loop, witha  break)
	for {
		fmt.Println("y is value :: ", y)
		if y > 20 {
			break
		}
		y++
	}

	//for range
	//create a slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice ", i, v)
	}
}
