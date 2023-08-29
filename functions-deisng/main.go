package main

import "fmt"

// func (r reciever) identifier(p parameter) (returns) { code }
//everything is pass by value

func main() {
	foo()
	bar("YO")
	s := aloha("NOO")
	fmt.Println(s)
	s1, n := DogYears("jake", 16)
	fmt.Println(s1, n)
	n1 := []int{1, 2, 3, 4, 5}
	x := sum(n1...)
	fmt.Println(x)

}

// no parameters, no returns
func foo() {
	fmt.Println("I am from foo")
}

// one parameter, no returns
func bar(s string) {
	fmt.Println("my name is ", s)
}

// one parameter, one return
func aloha(s string) string {
	return fmt.Sprint("My name is ", s) //fmt.Sprint prints to a string
}

// 2 parametrs, 2 returns
func DogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, "Todd is this old in dog years"), age
}

// VARIADIC PARAMETERS (... parameters)
// VARIADIC PARAMETERS NEED TO BE THE FINAL INCOMING PARAMETER
func sum(ii ...int) int {
	fmt.Println(ii) //so it creates a range of slice of int
	fmt.Printf("%T\n", ii)
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
