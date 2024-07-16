package httpserver

import (
	"chat/internal/app/auth/login"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	loginController := login.NewController()
	router.POST(login.Path, loginController.Login)

	return router
}
