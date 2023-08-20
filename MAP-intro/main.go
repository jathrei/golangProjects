package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":  42,
		"Henry": 16,
		"Pag":   14,
	}

	fmt.Println("The age of Henry was :", am["Henry"])
	fmt.Println(am)

	fmt.Println()
	an := make(map[string]int)
	an["Luke"] = 4
	an["doggy"] = 5
	fmt.Println(an)
	fmt.Println(len(an), len(am))
}
