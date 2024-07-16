package env

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
