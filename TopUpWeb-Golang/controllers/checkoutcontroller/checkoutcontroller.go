package checkoutcontroller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"topup-game/models"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Your existing code to retrieve the latest invoice
		latestInvoiceWithProduct, err := models.GetLatestInvoiceWithProduct()
		if err != nil {
			http.Error(w, "Error getting latest invoice: "+err.Error(), http.StatusInternalServerError)
			return
		}

		templatePath := "views/checkout/checkout.html"
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

	if r.Method == http.MethodPost {
		// Retrieve invoice ID from the form or request parameters
		invoiceID, err := strconv.Atoi(r.FormValue("invoice_id"))
		if err != nil {
			http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
			return
		}

		// Update the status to "Proses"
		if ok := models.UpdateInvoiceStatus(invoiceID, "Proses"); !ok {
			log.Println("Error: Unable to update invoice status")
			http.Error(w, "Unable to update invoice status", http.StatusInternalServerError)
			return
		}

		// Redirect to a success page or perform other actions as needed
		http.Redirect(w, r, "/checkoutproses", http.StatusSeeOther)
	}
}
