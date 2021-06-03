package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sing3demons/golanh-api/config"
	"github.com/sing3demons/golanh-api/controller"
	"github.com/sing3demons/golanh-api/middleware"
	"github.com/sing3demons/golanh-api/repository"
	"github.com/sing3demons/golanh-api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/v1/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("profile", userController.Profile)
		userRoutes.PUT("profile", userController.Update)
	}

	r.Run(":" + os.Getenv("PORT"))
}
