package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Página de início")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Criar novo Game")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jogo")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jogar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A cerca de")
}
