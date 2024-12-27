package main

import (
    "log"
    "net/http"
    "dheepssupreme/go-microservis.git/config"
    "dheepssupreme/go-microservis.git/routes"
)

func main() {
    
    config.InitDB()

    router := routes.SetupRoutes()


    log.Println("User Service running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
