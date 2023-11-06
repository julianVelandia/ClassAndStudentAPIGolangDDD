package mapper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/application/query"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/class/view/contract"
	"strings"
)

type Mapper struct{}

func (m Mapper) DomainToResponse(classes domain.Class) contract.Response {

	return contract.Response{
		ClassID:      classes.ClassID(),
		Title:        classes.Title(),
		CreationDate: classes.CreationDate(),
		Content:      classes.Content(),
		ReadTime:     classes.ReadTime(),
	}
}

func (m Mapper) RequestToQuery(ctx *gin.Context) (query.View, error) {
	request := contract.Request{
		ClassID: ctx.Param("class_id"),
	}

	if strings.HasPrefix(request.ClassID, ":") {
		return query.View{}, errors.New("request empty")
	}

	return *query.NewView(
		request.ClassID,
	), nil
}
