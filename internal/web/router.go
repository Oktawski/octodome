package web

import (
	"log"
	"net/http"
	auth "octodome/internal/auth/application"
	authinfra "octodome/internal/auth/infrastructure"
	authpres "octodome/internal/auth/presentation"
	eqhandler "octodome/internal/equipment/application/handler"
	"octodome/internal/equipment/application/handler/eqtypehandler"
	eqdom "octodome/internal/equipment/domain"
	eqinfra "octodome/internal/equipment/infrastructure"
	eqpres "octodome/internal/equipment/presentation"
	userhandler "octodome/internal/user/application/user_handler"
	userdom "octodome/internal/user/domain"
	userinfra "octodome/internal/user/infrastructure"
	userpres "octodome/internal/user/presentation"
	"octodome/internal/web/routes"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer() {
	r := chi.NewRouter()

	dsn := "host=localhost user=sa password=pass123 dbname=octodome_db port=5432 sslmode=disable"
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

	http.ListenAndServe(":8989", r)
}

func createUserController(userRepo userdom.UserRepository) *userpres.UserController {
	userCreateHandler := userhandler.NewUserCreateHandler(userRepo)
	userGetByID := userhandler.NewUserGetByIDHandler(userRepo)

	return userpres.NewUserController(userCreateHandler, userGetByID)
}

func createAuthController(userRepo auth.AuthRepository) *authpres.AuthController {
	tokenGenerator := authinfra.NewJwtTokenGenerator()
	authHandler := auth.NewAuthenticateHandler(userRepo, tokenGenerator)

	return authpres.NewAuthController(authHandler)
}

func createEquipmentController(db *gorm.DB) *eqpres.EquipmentController {
	eqTypeRepo := eqinfra.NewPgEquipmentTypeRepository(db)
	eqTypeValidator := eqdom.NewEquipmentTypeValidator(eqTypeRepo)

	eqHandler := eqhandler.NewEquipmentHandler()
	eqTypeCreateHandler := eqtypehandler.NewCreateHandler(eqTypeValidator, eqTypeRepo)
	eqTypeDeleteHandler := eqtypehandler.NewDeleteHandler(eqTypeValidator, eqTypeRepo)
	eqTypeGetByIDHandler := eqtypehandler.NewGetByIDHandler(eqTypeValidator, eqTypeRepo)
	eqTypeGetListHandler := eqtypehandler.NewGetListHandler(eqTypeValidator, eqTypeRepo)

	return eqpres.NewEquipmentController(
		eqHandler,
		eqTypeCreateHandler,
		eqTypeDeleteHandler,
		eqTypeGetByIDHandler,
		eqTypeGetListHandler)
}
