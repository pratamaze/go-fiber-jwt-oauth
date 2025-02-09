package config

import (

	"go-fiber-auth/models"
	"golang.org/x/crypto/bcrypt"
)


func SeedAdminUser() {
	var user models.User
	result := DB.Where("email = ?", "admin@example.com").First(&user)
	if result.Error != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := models.User{
			Name:     "Super Admin",
			Email:    "admin@example.com",
			Password: string(hashedPassword),
			IsAdmin:  true,
		}
		DB.Create(&admin)
	}
}

