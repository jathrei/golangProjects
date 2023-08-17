package main

import "fmt"

func main() {
	x := [5]int{0, 1, 1, 3, 4}
	for i, v := range x {
		fmt.Printf("index: %v, value %v of type %T\n", i, v, v)
	}

	y := []int{42, 43, 44, 45, 46, 47}
	fmt.Println(y[:5])
	y = append(y, 48, 49, 50, 51)
	fmt.Println(y[5:])
	fmt.Println(y[2:7])
	fmt.Println(y[1:6])
	z := []int{52, 53, 54, 55, 56, 57, 58, 59, 60}
	y = append(y, z...)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	y = append(y[:10])
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	states := make([]string, 0, 50)
	//initial cap and len
	fmt.Println(len(states))
	fmt.Println(cap(states))
	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(states))
	fmt.Println(cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Printf("index: %v   state: %v \n", i+1, states[i])
	}
}
