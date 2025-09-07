package web

import (
	"log"
	"net/http"
	auth "octodome/internal/auth/application"
	authpres "octodome/internal/auth/authhttp"
	authinfra "octodome/internal/auth/infrastructure"
	eqhdl "octodome/internal/equipment/application/handler/equipment"
	eqtypehdl "octodome/internal/equipment/application/handler/equipmenttype"
	equipmentdom "octodome/internal/equipment/domain/equipment"
	eqtypedom "octodome/internal/equipment/domain/equipmenttype"
	eqinfra "octodome/internal/equipment/infrastructure"
	eqrepo "octodome/internal/equipment/infrastructure/repository"
	eqpres "octodome/internal/equipment/presentation/equipment"
	eqtypepres "octodome/internal/equipment/presentation/equipmenttype"
	userhdl "octodome/internal/user/application/handler"
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

func createUserController(userRepo userdom.Repository) *userpres.UserController {
	userCreateHandler := userhdl.NewUserCreateHandler(userRepo)
	userGetByID := userhdl.NewUserGetByIDHandler(userRepo)

	return userpres.NewUserController(userCreateHandler, userGetByID)
}

func createAuthController(userRepo auth.AuthRepository) *authpres.AuthController {
	tokenGenerator := authinfra.NewJwtTokenGenerator()
	authHandler := auth.NewAuthenticateHandler(userRepo, tokenGenerator)

	return authpres.NewAuthController(authHandler)
}

func createEquipmentTypeController(db *gorm.DB) *eqtypepres.EquipmentTypeController {
	eqTypeRepo := eqrepo.NewPgEquipmentTypeRepository(db)
	eqTypeValidator := eqtypedom.NewEquipmentTypeValidator(eqTypeRepo)

	create := eqtypehdl.NewCreateHandler(eqTypeValidator, eqTypeRepo)
	update := eqtypehdl.NewUpdateHandler(eqTypeValidator, eqTypeRepo)
	delete := eqtypehdl.NewDeleteHandler(eqTypeValidator, eqTypeRepo)
	getByID := eqtypehdl.NewGetByIDHandler(eqTypeValidator, eqTypeRepo)
	getList := eqtypehdl.NewGetListHandler(eqTypeValidator, eqTypeRepo)

	return eqtypepres.NewEquipmentTypeController(
		create,
		update,
		delete,
		getByID,
		getList)
}

func createEquipmentController(db *gorm.DB) *eqpres.EquipmentController {
	eqRepo := eqrepo.NewPgEquipmentRepository(db)
	equipmentValidator := equipmentdom.NewValidator(eqRepo)

	create := eqhdl.NewCreateHandler(equipmentValidator, eqRepo)
	update := eqhdl.NewUpdateHandler(equipmentValidator, eqRepo)
	delete := eqhdl.NewDeleteHandler(equipmentValidator, eqRepo)
	getByID := eqhdl.NewGetByIDHandler(eqRepo)
	getList := eqhdl.NewGetListHandler(eqRepo)

	return eqpres.NewEquipmentController(
		create,
		update,
		delete,
		getByID,
		getList,
	)
}
