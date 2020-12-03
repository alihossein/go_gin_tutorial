package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowMainPage(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK , "index.html" , gin.H{})
}
