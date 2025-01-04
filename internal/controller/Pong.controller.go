package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
	})
}
