package main

import (
    "log"
    "net/http"
    "dheepssupreme/go-microservis.git/config"
    "dheepssupreme/go-microservis.git/routes"
)

func main() {
    // Initialize the database
    config.InitDB()

    // Setup routes
    router := routes.SetupRoutes()

    // Start the server
    log.Println("User Service running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
