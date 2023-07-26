package server

import (
	"cfd/myapp/config"
	"cfd/myapp/internal/common"
	"cfd/myapp/internal/core/domain/enum"
	"cfd/myapp/internal/core/service"
	handler "cfd/myapp/internal/handlers"
	"cfd/myapp/internal/helper"
	authjwt "cfd/myapp/pkg/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type WebServer struct {
	authService    *service.AuthService
	userService    *service.UserService
	profileService *service.ProfileService
	roleService    *service.RoleService
	jwtToken       authjwt.JwtToken
	authMiddleware gin.HandlerFunc
	config         config.Config
}

func NewWebServer(authService *service.AuthService, userService *service.UserService, profileService *service.ProfileService, roleService *service.RoleService, config config.Config, jwtToken authjwt.JwtToken) (*WebServer, error) {
	server := &WebServer{
		authService:    authService,
		userService:    userService,
		profileService: profileService,
		roleService:    roleService,
		jwtToken:       jwtToken,
		config:         config,
	}
	return server, nil
}

func (ws *WebServer) SetupRoutes(router *gin.Engine) {
	authHandler := handler.NewAuthHandler(*ws.authService)
	userHandler := handler.NewUserHandler(*ws.userService)
	profileHandler := handler.NewProfileHandler(*ws.profileService)
	roleHandler := handler.NewRoleHandler(*ws.roleService)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.Signup)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/verify-user", authHandler.VerifyUser)
		authGroup.POST("/refresh-token", authHandler.RefreshToken)
		authGroup.POST("/refresh-verification-code", authHandler.RefreshVerificationCode)
	}

	usersGroup := router.Group("/user")
	{
		usersGroup.Use(ws.AuthMiddleware(), ws.PermissionMiddleware())
		usersGroup.POST("", userHandler.CreateUser)
		usersGroup.GET("/:id", userHandler.GetUserByID)
		usersGroup.PUT("/:id", userHandler.UpdateUser)
		usersGroup.DELETE("/:id", userHandler.DeleteUser)
	}

	profileGroup := router.Group("/profile")
	{
		profileGroup.POST("/forgot-password", profileHandler.ForgotPassword)
		profileGroup.PUT("/reset-password", profileHandler.ResetPassword)

		profileGroup.Use(ws.AuthMiddleware())
		profileGroup.PUT("/change-password", profileHandler.ChangePassword)
		profileGroup.GET("/view", profileHandler.ViewProfile)
		profileGroup.PUT("/update", profileHandler.UpdateProfile)
		profileGroup.DELETE("/delete", profileHandler.DeleteProfile)
	}

	roleGroup := router.Group("/role")
	{
		roleGroup.Use(ws.AuthMiddleware(), ws.PermissionMiddleware())
		roleGroup.POST("", roleHandler.CreateRole)
		roleGroup.GET("/:id", roleHandler.GetRoleByID)
		roleGroup.PUT("/:id", roleHandler.UpdateRole)
		roleGroup.DELETE("/:id", roleHandler.DeleteRole)
	}

	err := ws.Run(router)
	if err != nil {
		log.Fatal(err)
	}
}

func (ws *WebServer) Run(router *gin.Engine) error {
	log.Println("App is starting on port " + ws.config.Server.Port)
	return router.Run(":" + ws.config.Server.Port)
}

func (ws *WebServer) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ResponseError("Authorization header missing"))
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		tokenUser, err := ws.jwtToken.ValidateAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ResponseError(err.Error()))
			return
		}

		c.Set("userID", tokenUser.ID)
		c.Next()
	}
}

func (ws *WebServer) PermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusInternalServerError, helper.ResponseError(common.ErrMissingContext.Error()))
			return
		}

		roles, err := ws.userService.GetUserRoleByID(userID.(int))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, helper.ResponseError(err.Error()))
			return
		}

		hasPermission := false
		for _, role := range roles {
			roleCode := enum.Role(role.Code)
			if roleCode == enum.Admin {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.AbortWithStatusJSON(http.StatusForbidden, helper.ResponseError(common.ErrPermission.Error()))
			return
		}

		c.Next()
	}
}
