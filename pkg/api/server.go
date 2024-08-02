package api

import (
	"chat/pkg/env"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(envFilePath string, router *gin.Engine) error {
	err := env.Init(".env")
	if err != nil {
		return fmt.Errorf("environment init error: %w", err)
	}

	port := env.Port()
	err = router.Run(":" + port)
	return fmt.Errorf("http server error: %w", err)
}
