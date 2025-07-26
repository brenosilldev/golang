package main

// Generics

// T = define o tipo de dado que a função vai receber - Type constraint
// int | float64 = define os tipos de dados que a função vai receber
// map[string]T = define o tipo de dado que a função vai receber

type Number interface {
	~int | float64
}

type MyInt int

func soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// constraint comparable
func Compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]MyInt{
		"a": 100,
		"b": 2,
		"c": 3,
		"d": 100,
	}

	println(soma(m))
	println(Compara(1, 3.4))
}
