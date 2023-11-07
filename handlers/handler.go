package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Erro ao analizar a Página", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Erro ao renderizar Página", http.StatusInternalServerError)
		return
	}
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
