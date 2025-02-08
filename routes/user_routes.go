package routes

import (
    "github.com/gofiber/fiber/v2"
    "go-fiber-auth/controllers"
    "go-fiber-auth/middleware"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/register", controllers.RegisterUser)
    app.Post("/login", controllers.LoginUser)
    app.Get("/auth/google", controllers.GoogleAuth) // Tambahkan route Google Auth

    userRoutes := app.Group("/users")
    userRoutes.Use(middleware.AuthMiddleware()) // Tambahkan middleware JWT di sini
    userRoutes.Use(middleware.Protected) // Middleware untuk proteksi endpoint
    userRoutes.Get("/", controllers.GetUsers)
    userRoutes.Get("/:id", controllers.GetUserByID)
    userRoutes.Put("/:id", controllers.UpdateUser)
    userRoutes.Delete("/:id", controllers.DeleteUser)
    // userRoutes.Put("/:id/make-admin", controllers.MakeAdmin)
    userRoutes.Put("/:id/admin", controllers.SetAdmin)


}
