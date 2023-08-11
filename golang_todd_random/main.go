package main

import (
	"fmt"

	"github.com/jathrei/puppy"
)

//comment to cehck git status
// random note - gotta config email in order for git to show up on ghub

func main() {
	fmt.Println("Yooooo")
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1, s2)
}
