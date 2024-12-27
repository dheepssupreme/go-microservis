package controllers

import (
    "encoding/json"
    "net/http"
    "dheepssupreme/go-microservis.git/models"
    "dheepssupreme/go-microservis.git/repositories"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := repositories.FindAllUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := repositories.CreateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
