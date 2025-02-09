package controllers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"go-fiber-auth/config"
	"go-fiber-auth/models"
	"github.com/golang-jwt/jwt/v4"
	"os"
    "fmt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func generateJWT(email string, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"email":  email,
		"admin":  isAdmin,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func RegisterUser(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := hashPassword(data["password"].(string))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	isAdmin := false
	if val, ok := data["is_admin"].(bool); ok {
		// Pastikan hanya admin yang bisa membuat user admin baru
		claims, ok := c.Locals("user").(*jwt.Token)
		if !ok || !claims.Claims.(jwt.MapClaims)["admin"].(bool) {
			return c.Status(403).JSON(fiber.Map{"error": "Only admin can create another admin"})
		}
		isAdmin = val
	}

	user := models.User{Name: data["name"].(string), Email: data["email"].(string), Password: hashedPassword, IsAdmin: isAdmin}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.JSON(fiber.Map{"id": user.ID, "name": user.Name, "email": user.Email, "is_admin": user.IsAdmin})
}

func SetAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	claims, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	claimsMap := claims.Claims.(jwt.MapClaims)
	if !claimsMap["admin"].(bool) {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
	}

	user.IsAdmin = true
	config.DB.Save(&user)

	return c.JSON(fiber.Map{"message": "User is now an admin", "user": user})
}



func LoginUser(c *fiber.Ctx) error {
	var data LoginRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := config.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if !checkPasswordHash(data.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid password"})
	}

	token, err := generateJWT(user.Email, user.IsAdmin)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	config.DB.Find(&users)
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	claims, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - Token not found"})
	}

	claimsMap := claims.Claims.(jwt.MapClaims)
	if claimsMap["email"] != user.Email && !claimsMap["admin"].(bool) {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
	}

	// Cetak claims untuk debugging
	fmt.Println("JWT Claims:", claimsMap)

	var data UpdateUserRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	user.Name = data.Name
	user.Email = data.Email
	config.DB.Save(&user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	claims, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	claimsMap := claims.Claims.(jwt.MapClaims)
	if !claimsMap["admin"].(bool) {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
	}

	config.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
