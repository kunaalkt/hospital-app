package routes

import (
    "go-hospital-app/controllers"
    "go-hospital-app/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRouter(dbConn interface{}) *gin.Engine {
    r := gin.Default()

    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())

    receptionistRoutes := auth.Group("/receptionist")
    receptionistRoutes.Use(middleware.RoleMiddleware("receptionist"))
    {
        receptionistRoutes.POST("/patients", controllers.CreatePatient)
        receptionistRoutes.GET("/patients", controllers.GetPatients)
        receptionistRoutes.GET("/patients/:id", controllers.GetPatient)
        receptionistRoutes.PUT("/patients/:id", controllers.UpdatePatient)
        receptionistRoutes.DELETE("/patients/:id", controllers.DeletePatient)
    }

    doctorRoutes := auth.Group("/doctor")
    doctorRoutes.Use(middleware.RoleMiddleware("doctor"))
    {
        doctorRoutes.GET("/patients", controllers.GetPatients)
        doctorRoutes.GET("/patients/:id", controllers.GetPatient)
        doctorRoutes.PUT("/patients/:id", controllers.UpdatePatient)
    }

    return r
}