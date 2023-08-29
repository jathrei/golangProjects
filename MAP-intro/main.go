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
	an["pp"] = 69
	fmt.Println(an["pp"])

	fmt.Println("-------------")

	for k, v := range an {
		fmt.Println(k, v)
	}
	fmt.Println("-------------")

	for _, v := range an {
		fmt.Println(v)
	}

	//for range with a slice
	xi := []int{41, 42, 43, 44}
	fmt.Println("-----------slice")
	//i for slice because its an index, not a key
	for i, v := range xi {
		fmt.Println(i, v)
	}

	//deleting an elemnet from a map
	fmt.Println(an)
	delete(an, "pp")
	fmt.Println(an)

	//comma, ok idiom
	/*v, ok := an["Luke"]
	if ok {
		fmt.Println("the value exists", v)
	} else {
		fmt.Println(" does not exist")
	}*/

	//comma, ok
	if v, ok := an["Dee"]; ok {
		fmt.Println("the value exists", v)
	} else {
		fmt.Println("the value does not exist")
	}

	//comma, ok w/ a not ok checker
	if v, ok := an["Dee"]; !ok {
		fmt.Println("the value does not esist", v)
	} else {
		fmt.Println("the value does exist")
	}
}
