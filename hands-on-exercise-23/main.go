package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 5; i++ {
		x := rand.Intn(250)
		fmt.Println("The value of x is: ", x)
		if x <= 100 {
			fmt.Println("between 0 and 100")
		} else if x <= 200 {
			fmt.Println("between 101 and 200")
		} else {
			fmt.Println("between 201 and 250")
		}
	}

	//doing it as a switch statement
	for i := 0; i < 5; i++ {
		x := rand.Intn(250)
		switch {
		case x <= 100:
			fmt.Println("between 0 and 100")
		case x > 100 && x <= 200:
			fmt.Println("between 101 and 200")
		case x > 200:
			fmt.Println("between 201 and 250")
		}
	}

	fmt.Println()

	fmt.Println("THIS IS START OF PART 27")
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Println("x is ", x, " y is ", y)

	var c bool = false

	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than 4")
		c = true
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than 6")
		c = true
	}
	if x >= 4 && x <= 6 {
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
		c = true
	}
	if y != 5 {
		fmt.Println("y is not 5")
		c = true
	}
	if !c {
		fmt.Println("none of the previous cases were met")
	}

	//not doing it via switch statement
	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("none of the previous cases were met")
	}

	//hands on num 29
	fmt.Println()
	fmt.Println("START OF NUM 29: print 0 to 99")
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println("PART 30")

	for i := 0; i < 42; i++ {
		z := rand.Intn(5)
		fmt.Printf("%v: ", z)
		switch z {
		case 0:
			fmt.Println("z is 0")
		case 1:
			fmt.Println("z is 1")
		case 2:
			fmt.Println("z is 2")
		case 3:
			fmt.Println("z is 3")
		case 4:
			fmt.Println("z is 4")
		}
	}
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		fmt.Println("newwww line")
		i++
	}

	for {
		fmt.Println("hey + ", i)
		if i == 20 {
			break
		}
		i++
	}
	for j := 0; j < 20; j++ {
		//print all odd numbers
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
	fmt.Printf("\n NESTED LOOP \n")
	//nested loop
	for k := 0; k < 5; k++ {
		for p := 0; p < 5; p++ {
			fmt.Printf("k is %v and p is %v \n", k, p)
		}
		fmt.Println()
	}

	xi := []int{40, 42, 42, 43, 44, 45, 46, 47}
	for k, v := range xi {
		fmt.Println("the value of k is: ", k)
		fmt.Println("v value is v: ", v)
	}
}

func init() {
	fmt.Println("This is where intiialization for my program occures")
}
