package main

import "fmt"

func main() {
	defer foo() //when main ends, foo runs
	bar()
}

//func (r reciever) identifier( p parameter(s)) (r return(s)) {code}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
