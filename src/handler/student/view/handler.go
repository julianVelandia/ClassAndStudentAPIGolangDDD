package view

import (
	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/query"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/view/contract"
	"net/http"
)

type Mapper interface {
	RequestToQuery(ctx *gin.Context) (query.View, error)
	DomainToResponse(profile domain.Profile, classesDone []domain.Class) contract.Response
}

type UseCase interface {
	Execute(email string) (domain.Profile, []domain.Class, error)
}

type Handler struct {
	mapper  Mapper
	useCase UseCase
}

func NewHandler(mapper Mapper, useCase UseCase) *Handler {
	return &Handler{mapper: mapper, useCase: useCase}
}

func (h Handler) Handler(ginCTX *gin.Context) {
	qry, errBinding := h.mapper.RequestToQuery(ginCTX)
	if errBinding != nil {
		ginCTX.JSON(http.StatusBadRequest, nil)
		return
	}

	domainProfile, classesDone, errorUseCase := h.useCase.Execute(qry.Email())
	if errorUseCase != nil {
		ginCTX.JSON(http.StatusInternalServerError, domainProfile)
		return
	}

	response := h.mapper.DomainToResponse(domainProfile, classesDone)
	ginCTX.JSON(http.StatusOK, response)
}
