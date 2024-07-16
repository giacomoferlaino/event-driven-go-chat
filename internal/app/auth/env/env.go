package env

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

var (
	godotenvLoad = godotenv.Load
	osGetenv     = os.Getenv
	ginSetMode   = gin.SetMode
)

func Init() error {
	err := godotenvLoad(".env")
	if err != nil {
		return err
	}
	GIN_MODE := osGetenv("GIN_MODE")
	ginSetMode(GIN_MODE)
	return nil
}

func Port() string {
	port := osGetenv("PORT")
	if port == "" {
		return defaultPort
	}
	return port
}
