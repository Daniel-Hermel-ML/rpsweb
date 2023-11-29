package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir  = "template/"
	templateBase = templateDir + "base.html"
)

type Player struct {
	Nome string
}

var player Player

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "game.html", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jogar")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Erro ao renderizar Página", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
