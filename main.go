package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sing3demons/golanh-api/config"
	"github.com/sing3demons/golanh-api/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run(":" + os.Getenv("PORT"))
}
