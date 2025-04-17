package main

// func main() {
// 	// type person struct {
// 	// 	name string
// 	// 	age  int
// 	// }

// 	// p1 := person{name: "Alice", age: 30}
// 	// p2 := person{name: "Bob", age: 25}

// 	// fmt.Println(p1)
// 	// fmt.Println(p2)

// 	// data := person{"breno", 10}

// 	// fmt.Println(data.name)
// 	// fmt.Println(data.age)

// 	person := struct {
// 		name string
// 		age  int
// 	}{
// 		"Hello",
// 		12,
// 	}

// 	fmt.Println(person)
// 	// println(person.age)

// }

// type Passeger struct {
// 	Name         string
// 	TicketNumber int
// 	Boarded      bool
// }

// type Bus struct {
// 	FrontSeat Passeger
// }

// func main() {
// 	breno := Passeger{"Breno", 1, true}
// 	fmt.Println(breno)
// 	rafael := Passeger{"Rafael", 2, false}
// 	fmt.Println(rafael)

// 	var (
// 		bill = Passeger{"Bill", 3, true}
// 		john = Passeger{"John", 4, false}
// 	)

// 	fmt.Println(bill)
// 	fmt.Println(john)

// 	var heidi Passeger
// 	heidi.Name = "Heidi"
// 	heidi.TicketNumber = 5
// 	fmt.Println(heidi)

// 	breno.Boarded = true
// 	fmt.Println(breno)
// 	rafael.Boarded = true

// 	if rafael.Boarded {
// 		fmt.Println("rafael boarded")
// 	}

// 	if breno.Boarded {
// 		fmt.Println("breno boarded")
// 	}

// 	heidi.Boarded = true

// 	bus := Bus{heidi}
// 	fmt.Println(bus.FrontSeat)
// 	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

// }

type Retangulo struct {
	Comprimento int
	Largura     int
}

func area(r Retangulo) int {
	return r.Comprimento * r.Largura
}

func perimetro(r Retangulo) int {
	return (r.Comprimento * 2) + (r.Largura * 2)
}

func printInfo(r Retangulo) {
	println("Comprimento:", r.Comprimento)
	println("Largura:", r.Largura)
	println("Area:", area(r))
	println("Perimetro:", perimetro(r))
}

func main() {

	r := Retangulo{10, 20}
	printInfo(r)

	r.Comprimento *= 2
	r.Largura *= 2
	printInfo(r)

}
