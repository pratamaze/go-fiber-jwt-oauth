// package models

// import (
//     "gorm.io/gorm"
// )

// type User struct {
//     gorm.Model
//     Name      string `gorm:"not null"`
//     Email     string `gorm:"unique;not null"`
//     Password  string `gorm:"not null"`
// }

// // // GetUsers - Handler untuk mendapatkan daftar user
// // func GetUsers(c *fiber.Ctx) error {
// // 	return c.JSON(fiber.Map{
// // 		"message": "List of users",
// // 	})
// // }

// package models

// import (
// 	"time"
// )

// type User struct {
// 	ID        uint      `json:"id" gorm:"primaryKey"`
// 	Name      string    `json:"name"`
// 	Email     string    `json:"email" gorm:"unique"`
// 	Password  string    `json:"-"`
// 	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
// }


package models

import (
	"time"

)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"-"`
	IsAdmin   bool           `json:"is_admin"`
	DeletedAt *time.Time     `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
