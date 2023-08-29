package main

import "fmt"

type person struct {
	first string
	last  string
	ages  []int
}

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

type person1 struct {
	fav []string
}

type secret struct {
	person
	yo int
}

func main() {

	/*p1 := secret{
		person: person{
			first: "Jake",
			last:  "OK",
			age:   24,
		},
		yo: 30,
	}*/
	p1 := person{
		first: "J",
		last:  "R",
		ages:  []int{24, 25, 26, 27, 28},
	}
	p2 := person{
		first: "M",
		last:  "J",
		ages:  []int{28, 29, 30, 31, 32},
	}

	fmt.Println(p1)

	//fmt.Println(p1.first, p1.last, p1.age, p1.person)
	fmt.Printf("%T \n", p1)

	fmt.Println("------")
	xs := make([]person, 1, 1)
	fmt.Println(xs)
	//xs[0].age = 24
	xs[0].first = "Jake"
	xs[0].last = "reo"
	fmt.Println(xs)

	/*p2 := person1{
		fav: []string{"test1", "test2"},
	}
	fmt.Println(p2)
	for _, v := range p2.fav {
		fmt.Println(xs[0].first, "last is : ", v)
	}*/

	fmt.Println("=======")
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for k, v := range m {
		for _, v2 := range v.ages {
			fmt.Println(v2)
		}
		fmt.Println(k, v)
	}

	fmt.Println("------ embedded struct --------")
	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "ford",
		model: "fusion",
		doors: 4,
		color: "white",
	}
	v2 := vehicle{
		//inner type promotoion
		//engine is being promoted to vehicle
		engine: engine{
			electric: true,
		},
		make: "ford",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	//creating an anonymous struct
	fmt.Println("oooooooANON-STRUCToooooo")
	x := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "JT",
		friends: map[string]int{
			"J": 14,
			"A": 13,
		},
		favDrinks: []string{
			"Yo",
			"Hello",
		},
	}

	for k, v := range x.friends {
		fmt.Println(k, v)
	}

}
