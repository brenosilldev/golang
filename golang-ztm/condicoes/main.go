package main

import (
	"fmt"
)

// func accessGranted() {
// 	fmt.Println("Access granted")
// }

// func accessDenied() {
// 	fmt.Println("Access denied")
// }

// const (
// 	Monday    = 0
// 	Tuesday   = 1
// 	Wednesday = 2
// 	Thursday  = 3
// 	Friday    = 4
// 	Saturday  = 5
// 	Sunday    = 6
// )

// const (
// 	Admin      = 10
// 	Manager    = 20
// 	Contractor = 30
// 	Member     = 40
// 	Guest      = 0
// )

// func weekDay(day int) bool {
// 	return day <= 4
// }

// func main() {

// 	today, role := Tuesday, Guest

// 	if role == Admin || role == Manager {
// 		accessGranted()
// 	} else if role == Contractor && !weekDay(today) {
// 		accessGranted()
// 	} else if role == Member && weekDay(today) {
// 		accessGranted()

// 	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
// 		accessGranted()
// 	} else {
// 		accessDenied()
// 	}

// }

// func VerificaValor(valor string ) {

// 	switch valor {
// 		case "verde":
// 			fmt.Println("Verde")
// 		case "amarelo":
// 			fmt.Println("Amarelo")
// 		case "vermelho":
// 			fmt.Println("Vermelho")
// 		default:
// 			fmt.Println("Cor desconhecida")
// 	}
// }

// func main() {
// 	VerificaValor("dede")
// }

func main() {
	switch age := 19; {
	case age == 0:
		fmt.Println("newborn")
	case age >=  1 && age <= 3:
		fmt.Println("toddler")
	case age >= 4 && age <= 12:
		fmt.Println("child")
	case age >= 13 && age <= 17:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")

	}

}
