package main

import (
	"task5/config"
	"task5/controllers"
	"task5/middleware"
	"task5/repository"
	"task5/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                    = config.SetupDatabaseConnection()
	userRepository  repository.UserRepository   = repository.NewUserRepository(db)
	photoRepository repository.PhotoRepository  = repository.NewPhotoRepository(db)
	jwtService      service.JWTService          = service.NewJWTService()
	authService     service.AuthService         = service.NewAuthService(userRepository)
	userService     service.UserService         = service.NewUserService(userRepository)
	photoService    service.PhotoService        = service.NewPhotoService(photoRepository)
	authController  controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
	userController  controllers.UserController  = controllers.NewUserController(userService, jwtService)
	photoController controllers.PhotoController = controllers.NewPhotoController(photoService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("users")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRoutes := r.Group("users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
	photoRoutes := r.Group("photos", middleware.AuthorizeJWT(jwtService))
	{
		photoRoutes.GET("/", photoController.All)
		photoRoutes.POST("/", photoController.Insert)
		photoRoutes.PUT("/:id", photoController.Update)
		photoRoutes.DELETE("/:id", photoController.Delete)
	}
	r.Run()
}
