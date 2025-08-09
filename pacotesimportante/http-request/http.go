package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	fmt.Println(string(r.Body.Read()))
	// })

	// http.Post("http://localhost:8080", "application/json", strings.NewReader("Hello, World!"))

	// http.ListenAndServe(":8080", nil)

	req, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		panic(err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da requisição:", err)
		panic(err)
	}

	fmt.Println(req.Header)
	fmt.Println(req.StatusCode)

	fmt.Println(string(res))

	// Defer é usado para fechar o corpo da requisição após o uso
	// Isso é importante para evitar vazamentos de recursos
	// Ele executa  por ultimo
	defer req.Body.Close()

}
