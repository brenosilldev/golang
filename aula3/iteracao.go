package aula3


import (
	"fmt"
	"strings"
)

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < 5; i++ {
		repeticoes += strings.ToUpper(caractere) + "\n"
	}
	return repeticoes
}

func main() {
	fmt.Println(Repetir("a"))
}
