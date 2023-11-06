package mapper

import (
	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
)

type Mapper struct{}

func (m Mapper) RequestToCommand(ctx *gin.Context) command.UpdateClassesDone {

	return *command.NewUpdateClassesDone(
		ctx.Param("email"),
		ctx.Param("class_id"),
		ctx.Param("title"),
	)
}
