package routes

import (
    "github.com/gorilla/mux"
    "dheepssupreme/go-microservis.git/controllers"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
    router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
    return router
}
