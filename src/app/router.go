package app

import (
	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/app/dependence"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	handlers := dependence.NewWire()
	configureMappings(router, handlers)
	return router
}

func configureMappings(router *gin.Engine, handlers dependence.HandlerContainer) {
	// Student
	apiGroupStudent := router.Group("v1.0/student")
	apiGroupStudent.GET("/profile/:email", handlers.ViewProfileHandler.Handler)
	apiGroupStudent.PUT("/class/:class_id/email/:email/title/:title", handlers.UpdateClassesDoneHandler.Handler)

	// Classes
	apiGroupClasses := router.Group("v1.0/classes")
	apiGroupClasses.GET("/class/:class_id/email/:email/title/:title", handlers.ViewClassHandler.Handler)
}
