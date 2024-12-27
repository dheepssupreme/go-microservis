package controllers

import (
	"dheepssupreme/go-microservis.git/config"
	"dheepssupreme/go-microservis.git/models"
	"encoding/json"
	"github.com/gorilla/mux" // Tambahkan import ini
	"net/http"
	"strconv"
)

// GetUsers mendapatkan semua user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser membuat user baru
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUserByID mendapatkan user berdasarkan ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari parameter URL
	id := mux.Vars(r)["id"]

	// Convert ID ke integer
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Cari user berdasarkan ID
	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Kirim response dengan data user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari parameter URL
	id := mux.Vars(r)["id"]

	// Convert ID ke integer
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Cari user berdasarkan ID dan hapus
	var user models.User
	result := config.DB.Delete(&user, userID)
	if result.Error != nil {
		http.Error(w, "User not found or failed to delete", http.StatusNotFound)
		return
	}

	// Kirim response sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // No content, karena tidak ada data yang dikembalikan
}
