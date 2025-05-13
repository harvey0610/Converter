package main

import (
	"currency-converter-go/api"
	"html/template"
	"log"
	"net/http"
)

// serveHome renders the main HTML page
func serveHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

func main() {
	// Serve static files from the "public" folder
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Route to render the frontend
	http.HandleFunc("/", serveHome)

	// API route for currency conversion
	http.HandleFunc("/api/convert", api.ConvertHandler)

	log.Println("✅ Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}
