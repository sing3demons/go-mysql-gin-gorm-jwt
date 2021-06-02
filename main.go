package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sing3demons/golanh-api/config"
	"github.com/sing3demons/golanh-api/controller"
	"github.com/sing3demons/golanh-api/repository"
	"github.com/sing3demons/golanh-api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	jwtService     service.JWTService        = service.NewJWTService()
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
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
