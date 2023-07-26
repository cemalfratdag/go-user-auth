package cmd

import (
	"cfd/myapp/internal/core/domain/entity"
	"cfd/myapp/internal/core/service"
	"cfd/myapp/internal/server"
	"cfd/myapp/pkg/db/postgres"
	authjwt "cfd/myapp/pkg/jwt"
	"cfd/myapp/pkg/viper"
	"github.com/gin-gonic/gin"
	"log"
)

func Main() {
	config, err := viper.LoadConfig()
	if err != nil {
		log.Fatalf("Config error %s", err)
	}

	db, err := postgres.NewDatabase(config.Db)
	if err != nil {
		log.Fatalf("DB error %s", err)
	}

	err = db.Migrate(&entity.User{})
	if err != nil {
		log.Fatalf("DB migration error %s", err)
	}

	jwtToken := authjwt.NewJwtToken(config.Jwt)

	authService := service.NewAuthService(db.UserRepository, jwtToken)
	userService := service.NewUserService(db.UserRepository, db.RoleRepository)
	profileService := service.NewProfileService(db.UserRepository)
	roleService := service.NewRoleService(db.RoleRepository)

	router := gin.Default()

	webServer, err := server.NewWebServer(authService, userService, profileService, roleService, config, jwtToken)
	if err != nil {
		log.Fatalf("Error initializing web server: %s", err)
	}

	webServer.SetupRoutes(router)
}
