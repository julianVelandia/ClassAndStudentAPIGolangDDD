package mapper

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/query"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/view/contract"
	"github.com/stretchr/testify/assert"
)

const (
	email = "test@example.com"
	name  = "My Name"
)

func TestMapper_DomainToResponse(t *testing.T) {
	domainProfile := domain.NewProfile(email, name)
	domainDoneArray := []domain.Class{
		*domain.NewClass("id1", "Clase 1"),
		*domain.NewClass("id2", "Clase 2"),
	}
	responseClassesDone := []contract.Class{
		{ClassID: "id1", Title: "Clase 1"},
		{ClassID: "id2", Title: "Clase 2"},
	}
	mapper := Mapper{}
	expectedResponse := contract.Response{
		Email:       email,
		Name:        name,
		ClassesDone: responseClassesDone,
	}

	response := mapper.DomainToResponse(*domainProfile, domainDoneArray)

	assert.Equal(t, expectedResponse, response)
}

func TestMapper_ValidRequestToQuery(t *testing.T) {
	ctx := setupTestContext(email)
	mapperClass := Mapper{}

	view, err := mapperClass.RequestToQuery(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, view)
}

func TestMapper_EmptyEmailRequestToQuery(t *testing.T) {
	ctx := setupTestContext(":")
	mapperClass := Mapper{}

	view, err := mapperClass.RequestToQuery(ctx)
	assert.Error(t, err)
	assert.Equal(t, "request empty", err.Error())
	assert.Equal(t, query.View{}, view)
}

func setupTestContext(email string) *gin.Context {
	ctx := &gin.Context{}
	ctx.Params = append(ctx.Params, gin.Param{Key: "email", Value: email})
	return ctx
}
