package controllers

import (
    "github.com/gofiber/fiber/v2"
)

func GoogleAuth(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "Google OAuth not implemented yet"})
}
