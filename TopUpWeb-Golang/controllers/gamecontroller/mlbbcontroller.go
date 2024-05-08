package gamecontroller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"topup-game/entities"
	"topup-game/models"
)

func Mlbb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		product, err := models.GetProductMlbb()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Generate CSRF token and pass it to the template
		csrfToken := generateCSRFToken()
		tmpl, err := template.ParseFiles("views/game/mlbb.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, struct {
			Products  []entities.InvoiceWithProduct
			CSRFToken string
		}{Products: product, CSRFToken: csrfToken})
	}

	if r.Method == "POST" {
		// Validate CSRF token
		if !validateCSRFToken(r) {
			http.Error(w, "Invalid CSRF token", http.StatusForbidden)
			return
		}

		var invoice entities.InvoiceWithProduct

		dataID, err := strconv.Atoi(r.FormValue("dataID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		productID, err := strconv.Atoi(r.FormValue("product_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		phoneNumber, err := strconv.ParseInt(r.FormValue("phone_number"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		invoice.DataID = dataID
		invoice.ProductID = productID
		invoice.Payment = r.FormValue("payment")
		invoice.PhoneNumber = phoneNumber
		invoice.Status = "BelumDibayar"
		invoice.Time = time.Now()

		if ok := models.CreateInvoice(&invoice); !ok {
			log.Println("Error: Unable to create invoice")
			http.Error(w, "Unable to create invoice", http.StatusInternalServerError)
			return
		}

		log.Println("Debug: Invoice created successfully")
		http.Redirect(w, r, "/checkout", http.StatusSeeOther)
	}
}

// Function to generate a CSRF token (replace with your own implementation)
func generateCSRFToken() string {
	// Implement your CSRF token generation logic here
	// Example: Generate a random token
	return "random_csrf_token"
}

// Function to validate CSRF token (replace with your own implementation)
func validateCSRFToken(r *http.Request) bool {
	// Implement your CSRF token validation logic here
	// Example: Compare the token in the request with the expected token
	expectedToken := "random_csrf_token"
	actualToken := r.FormValue("csrf_token")
	return actualToken == expectedToken
}
