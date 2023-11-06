package updateclassesdone

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
)

type Mapper interface {
	RequestToCommand(ctx *gin.Context) command.UpdateClassesDone
}

type UseCase interface {
	Execute(cmd command.UpdateClassesDone) error
}

type Handler struct {
	mapper  Mapper
	useCase UseCase
}

func NewHandler(mapper Mapper, useCase UseCase) *Handler {
	return &Handler{mapper: mapper, useCase: useCase}
}

func (h Handler) Handler(ginCTX *gin.Context) {
	cmd := h.mapper.RequestToCommand(ginCTX)

	errorUseCase := h.useCase.Execute(cmd)
	if errorUseCase != nil {
		ginCTX.JSON(http.StatusInternalServerError, nil)
		return
	}

	ginCTX.JSON(http.StatusOK, nil)
}
