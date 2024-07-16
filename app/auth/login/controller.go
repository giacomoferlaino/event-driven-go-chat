package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewController() *Controller {
	return &Controller{}
}

type Controller struct{}

func (c *Controller) Login(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Logged in")
}
