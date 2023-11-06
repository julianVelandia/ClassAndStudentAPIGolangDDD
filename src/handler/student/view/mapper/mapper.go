package mapper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/query"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/view/contract"
	"strings"
)

type Mapper struct{}

func (m Mapper) DomainToResponse(profile domain.Profile, classesDone []domain.Class) contract.Response {
	classesResponse := make([]contract.Class, len(classesDone))
	for i := range classesDone {
		classesResponse[i] = contract.Class{
			ClassID: classesDone[i].ClassID(),
			Title:   classesDone[i].Title(),
		}
	}

	return contract.Response{
		Email:       profile.Email(),
		Name:        profile.Name(),
		ClassesDone: classesResponse,
	}
}

func (m Mapper) RequestToQuery(ctx *gin.Context) (query.View, error) {
	request := contract.Request{
		Email: ctx.Param("email"),
	}

	if strings.HasPrefix(request.Email, ":") {
		return query.View{}, errors.New("request empty")
	}

	return *query.NewView(
		request.Email,
	), nil
}
