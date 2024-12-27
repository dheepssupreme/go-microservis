package main

import (
	"dheepssupreme/go-microservis.git/config"
	"dheepssupreme/go-microservis.git/routes"
	"net/http"
	"log"
)

func main() {
	// Inisialisasi database
	config.InitDB()

	// Setup routes
	router := routes.SetupRoutes()

	// Serve static files (CSS, JS, images, etc.)
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Serve HTML template
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/users.html")
	})

	// Jalankan server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
