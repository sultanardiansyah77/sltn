package gamecontroller

import (
	"html/template"
	"net/http"
	"topup-game/entities"
	"topup-game/models"
)

func Freefire(w http.ResponseWriter, r *http.Request) {
	product, err := models.GetProductFF()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("views/game/ff.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, struct{ Products []entities.Product }{Products: product})
}
