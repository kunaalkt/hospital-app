package main

import (
    "go-hospital-app/config"
    "go-hospital-app/models"
    "go-hospital-app/routes"
    "go-hospital-app/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    db := config.InitDB()
    models.AutoMigrate(db)
    router := routes.SetupRouter(db)

    router.Static("/static", "./static")
    router.StaticFile("/", "./static/index.html")
    router.StaticFile("/receptionist.html", "./static/receptionist.html")
    router.StaticFile("/doctor.html", "./static/doctor.html")

    router.GET("/receptionist", middleware.AuthMiddleware(), middleware.RoleMiddleware("receptionist"), func(c *gin.Context) {
        c.File("./static/receptionist.html")
    })
    router.GET("/doctor", middleware.AuthMiddleware(), middleware.RoleMiddleware("doctor"), func(c *gin.Context) {
        c.File("./static/doctor.html")
    })

    router.Run(":8080")
}