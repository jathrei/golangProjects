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
}
