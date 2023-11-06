package mapper

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/application/query"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/domain"
	"github.com/stretchr/testify/assert"
)

const (
	email        = "test@email.com"
	classID      = "123"
	title        = "Test Class"
	creationDate = "2023-08-27"
	readTime     = 1.5
)

func TestMapper_DomainToResponse(t *testing.T) {
	mapperClass := Mapper{}
	content := []string{"Content line 1", "Content line 2"}
	domainClass := domain.NewClass(classID, title, creationDate, content, readTime)

	response := mapperClass.DomainToResponse(*domainClass)

	assert.Equal(t, classID, response.ClassID)
	assert.Equal(t, title, response.Title)
	assert.Equal(t, creationDate, response.CreationDate)
	assert.Equal(t, content, response.Content)
	assert.Equal(t, readTime, response.ReadTime)
}

func TestMapper_ValidRequestToQuery(t *testing.T) {
	ctx := setupTestContext(email, classID, title)
	mapperClass := Mapper{}

	view, err := mapperClass.RequestToQuery(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, view)
}

func TestEmptyClassIDRequestToQuery(t *testing.T) {
	ctx := setupTestContext("user@example.com", ":", "Some Title")
	mapperClass := Mapper{}

	view, err := mapperClass.RequestToQuery(ctx)
	assert.Error(t, err)
	assert.Equal(t, "request empty", err.Error())
	assert.Equal(t, query.View{}, view)
}

func setupTestContext(email, classID, title string) *gin.Context {
	ctx := &gin.Context{}
	ctx.Params = append(ctx.Params, gin.Param{Key: "class_id", Value: classID})
	return ctx
}
