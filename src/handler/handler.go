package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	Handler(ginCTX *gin.Context)
}
