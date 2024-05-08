package main

import (
	"log"
	"net/http"
	"topup-game/config"
	"topup-game/controllers/checkoutcontroller"
	"topup-game/controllers/dashboardcontroller"
	"topup-game/controllers/gamecontroller"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))).ServeHTTP(w, r)
	})

	http.HandleFunc("/", dashboardcontroller.Home)
	http.HandleFunc("/contact", dashboardcontroller.Contact)
	http.HandleFunc("/riwayat", dashboardcontroller.Riwayat)
	http.HandleFunc("/riwayat-invoice", dashboardcontroller.RiwayatInvoiceHandler)

	http.HandleFunc("/freefire", gamecontroller.Freefire)
	http.HandleFunc("/mlbb", gamecontroller.Mlbb)

	http.HandleFunc("/checkout", checkoutcontroller.Checkout)
	http.HandleFunc("/checkoutproses", checkoutcontroller.CheckoutProses)

	server := &http.Server{Addr: ":8080"}

	log.Println("Server running on port 8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server error:", err)
	}

	log.Println("Server stopped")
}
