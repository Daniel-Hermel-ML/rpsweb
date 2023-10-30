package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo")
	})

	port := ":8090"
	fmt.Printf("Servidor rodando em http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
