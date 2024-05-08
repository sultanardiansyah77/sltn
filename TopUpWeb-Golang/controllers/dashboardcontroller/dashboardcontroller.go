package dashboardcontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"topup-game/entities"
	"topup-game/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/dashboard/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Riwayat(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/dashboard/riwayat.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/dashboard/contact.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func RiwayatInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Mengambil nilai InvoiceNumber dari formulir HTTP
		InvoiceNumberStr := r.FormValue("invoiceNumber")
		InvoiceNumber, err := strconv.Atoi(InvoiceNumberStr)
		if err != nil {
			http.Error(w, "Invalid InvoiceNumber", http.StatusBadRequest)
			return
		}

		// Memanggil fungsi GetInvoiceWithProductByID dengan nilai InvoiceNumber yang sudah diubah
		result, err := models.GetInvoiceWithProductByID(InvoiceNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Menyiapkan dan mengeksekusi template HTML
		tmpl, err := template.ParseFiles("views/dashboard/riwayat-invoice.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Menyiapkan struktur data baru yang sesuai dengan template HTML
		data := struct {
			Result       *entities.InvoiceWithProduct
			InvoiceExist bool
		}{
			Result:       result,
			InvoiceExist: result != nil,
		}

		tmpl.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
