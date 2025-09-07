package web

import (
	"log"
	"net/http"
	auth "octodome/internal/auth/application"
	authpres "octodome/internal/auth/authhttp"
	authinfra "octodome/internal/auth/infrastructure"
	eq "octodome/internal/equipment/application/handler/equipment"
	eqtype "octodome/internal/equipment/application/handler/equipmenttype"
	eqdom "octodome/internal/equipment/domain/equipment"
	eqtypedom "octodome/internal/equipment/domain/equipmenttype"
	eqinfra "octodome/internal/equipment/infrastructure"
	eqhttp "octodome/internal/equipment/presentation/equipmenthttp"
	eqtypehttp "octodome/internal/equipment/presentation/equipmenttypehttp"
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

func createEquipmentTypeController(db *gorm.DB) *eqtypehttp.EquipmentTypeController {
	eqTypeRepo := eqinfra.NewPgEquipmentTypeRepository(db)
	eqTypeValidator := eqtypedom.NewEquipmentTypeValidator(eqTypeRepo)

	eqTypeCreateHandler := eqtype.NewCreateHandler(eqTypeValidator, eqTypeRepo)
	eqTypeUpdateHandler := eqtype.NewUpdateHandler(eqTypeValidator, eqTypeRepo)
	eqTypeDeleteHandler := eqtype.NewDeleteHandler(eqTypeValidator, eqTypeRepo)
	eqTypeGetByIDHandler := eqtype.NewGetByIDHandler(eqTypeValidator, eqTypeRepo)
	eqTypeGetListHandler := eqtype.NewGetListHandler(eqTypeValidator, eqTypeRepo)

	return eqtypehttp.NewEquipmentTypeController(
		eqTypeCreateHandler,
		eqTypeUpdateHandler,
		eqTypeDeleteHandler,
		eqTypeGetByIDHandler,
		eqTypeGetListHandler)
}

func createEquipmentController(db *gorm.DB) *eqhttp.EquipmentController {
	eqRepo := eqinfra.NewPgEquipmentRepository(db)
	equipmentValidator := eqdom.NewValidator(eqRepo)

	eqCreateHandler := eq.NewCreateHandler(equipmentValidator, eqRepo)
	eqUpdateHandler := eq.NewUpdateHandler(equipmentValidator, eqRepo)
	eqDeleteHandler := eq.NewDeleteHandler(equipmentValidator, eqRepo)
	eqGetByIDHandler := eq.NewGetByIDHandler(eqRepo)
	eqGetListHandler := eq.NewGetListHandler(eqRepo)

	return eqhttp.NewEquipmentController(
		eqCreateHandler,
		eqUpdateHandler,
		eqDeleteHandler,
		eqGetByIDHandler,
		eqGetListHandler)
}
