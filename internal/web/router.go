package web

import (
	"log"
	auth "octodome/internal/auth/application"
	authinfra "octodome/internal/auth/infrastructure"
	authpres "octodome/internal/auth/presentation"
	equipment "octodome/internal/equipment/application"
	eqinfra "octodome/internal/equipment/infrastructure"
	eqpres "octodome/internal/equipment/presentation"
	user "octodome/internal/user/application"
	userdom "octodome/internal/user/domain"
	userinfra "octodome/internal/user/infrastructure"
	userpres "octodome/internal/user/presentation"
	"octodome/internal/web/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer() {
	r := gin.Default()

	dsn := "host=localhost user=user password=pass123 dbname=octome_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	userinfra.Migrate(db)
	eqinfra.Migrate(db)

	userRepo := userinfra.NewPgUserRepository(db)

	routes.RegisterUserRoutes(r, createUserController(userRepo))
	routes.RegisterAuthRoutes(r, createAuthController(userRepo))
	routes.RegisterEquipmentRoutes(r, createEquipmentController(db))

	r.Run(":8989")
}

func createUserController(userRepo userdom.UserRepository) *userpres.UserController {
	userHandler := user.NewUserHandler(userRepo)

	return userpres.NewUserController(userHandler)
}

func createAuthController(userRepo auth.AuthRepository) *authpres.AuthController {
	tokenGenerator := authinfra.NewJwtTokenGenerator()
	authHandler := auth.NewAuthHandler(userRepo, tokenGenerator)

	return authpres.NewAuthController(authHandler)
}

func createEquipmentController(db *gorm.DB) *eqpres.EquipmentController {
	eqRepo := eqinfra.NewPgEquipmentRepository(db)

	eqHandler := equipment.NewEquipmentHandler()
	eqTypeHandler := equipment.NewEquipmentTypeHandler(eqRepo)

	return eqpres.NewEquipmentController(eqHandler, eqTypeHandler)
}
