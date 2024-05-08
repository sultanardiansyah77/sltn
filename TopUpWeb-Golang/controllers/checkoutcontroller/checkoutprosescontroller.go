package checkoutcontroller

import (
	"html/template"
	"net/http"
	"topup-game/models"
)

func CheckoutProses(w http.ResponseWriter, r *http.Request) {
	latestInvoiceWithProduct, err := models.GetLatestInvoiceWithProduct()
	if err != nil {
		http.Error(w, "Error getting latest invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	templatePath := "views/checkout/checkoutproses.html"
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, latestInvoiceWithProduct)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
