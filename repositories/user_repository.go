package repositories

import (
    "dheepssupreme/go-microservis.git/config"
    "dheepssupreme/go-microservis.git/models"
)

func FindAllUsers() ([]models.User, error) {
    var users []models.User
    result := config.DB.Find(&users)
    return users, result.Error
}

func CreateUser(user *models.User) error {
    result := config.DB.Create(user)
    return result.Error
}
