package config

import (

	"go-fiber-auth/models"
	"golang.org/x/crypto/bcrypt"
)

// func SeedAdmin(db *gorm.DB) {
// 	var user models.User
// 	result := db.Where("email = ?", "admin@example.com").First(&user)
// 	if result.RowsAffected == 0 {
// 		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
// 		admin := models.User{
// 			Name:     "Super Admin",
// 			Email:    "admin@example.com",
// 			Password: string(hashedPassword),
// 			IsAdmin:  true,
// 		}
// 		db.Create(&admin)
// 		log.Println("Admin user created: admin@example.com (password: admin123)")
// 	} else {
// 		log.Println("Admin user already exists")
// 	}
// }

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

