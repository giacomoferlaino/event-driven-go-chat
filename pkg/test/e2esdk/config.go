package e2esdk

import "github.com/gin-gonic/gin"

type Config struct {
	Router       *gin.Engine
	KeycloakData *KeycloakData
	KeycloakUrl  string
}
