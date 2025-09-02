package web

import (
	"log"
	"net/http"
	auth "octodome/internal/auth/application"
	authinfra "octodome/internal/auth/infrastructure"
	authpres "octodome/internal/auth/presentation"
	"octodome/internal/equipment/application/handler/equipment"
	equipmenttype "octodome/internal/equipment/application/handler/equipment_type_handler"
	equipmentdom "octodome/internal/equipment/domain/equipment"
	eqtypedom "octodome/internal/equipment/domain/equipment_type"
	eqinfra "octodome/internal/equipment/infrastructure"
	eqpres "octodome/internal/equipment/presentation"
	userhandler "octodome/internal/user/application/handler"
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
	routes.RegisterEquipmentRoutes(r, createEquipmentController(db), createEquipmentTypeController(db))

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

func createEquipmentTypeController(db *gorm.DB) *eqpres.EquipmentTypeController {
	eqTypeRepo := eqinfra.NewPgEquipmentTypeRepository(db)
	eqTypeValidator := eqtypedom.NewEquipmentTypeValidator(eqTypeRepo)

	eqTypeCreateHandler := equipmenttype.NewCreateHandler(eqTypeValidator, eqTypeRepo)
	eqTypeUpdateHandler := equipmenttype.NewUpdateHandler(eqTypeValidator, eqTypeRepo)
	eqTypeDeleteHandler := equipmenttype.NewDeleteHandler(eqTypeValidator, eqTypeRepo)
	eqTypeGetByIDHandler := equipmenttype.NewGetByIDHandler(eqTypeValidator, eqTypeRepo)
	eqTypeGetListHandler := equipmenttype.NewGetListHandler(eqTypeValidator, eqTypeRepo)

	return eqpres.NewEquipmentTypeController(
		eqTypeCreateHandler,
		eqTypeUpdateHandler,
		eqTypeDeleteHandler,
		eqTypeGetByIDHandler,
		eqTypeGetListHandler)
}

func createEquipmentController(db *gorm.DB) *eqpres.EquipmentController {
	eqRepo := eqinfra.NewPgEquipmentRepository(db)
	equipmentValidator := equipmentdom.NewEquipmentValidator(eqRepo)

	eqCreateHandler := equipment.NewCreateHandler(equipmentValidator, eqRepo)
	eqUpdateHandler := equipment.NewUpdateHandler(equipmentValidator, eqRepo)
	eqDeleteHandler := equipment.NewDeleteHandler(equipmentValidator, eqRepo)
	eqGetByIDHandler := equipment.NewGetByIDHandler(eqRepo)
	eqGetListHandler := equipment.NewGetListHandler(eqRepo)

	return eqpres.NewEquipmentController(
		eqCreateHandler,
		eqUpdateHandler,
		eqDeleteHandler,
		eqGetByIDHandler,
		eqGetListHandler)
}
