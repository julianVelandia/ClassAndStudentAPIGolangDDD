package mapper

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	email   = "test@email.com"
	classID = "123"
	title   = "title"
)

func TestMapper_RequestToCommand(t *testing.T) {
	ctx := setupTestContext(email, classID, title)
	mapperClass := Mapper{}

	cmd := mapperClass.RequestToCommand(ctx)
	assert.Equal(t, cmd.Email(), email)
	assert.Equal(t, cmd.ClassID(), classID)
	assert.Equal(t, cmd.Title(), title)
}

func setupTestContext(email, classID, title string) *gin.Context {
	ctx := &gin.Context{}
	ctx.Params = append(ctx.Params, gin.Param{Key: "email", Value: email})
	ctx.Params = append(ctx.Params, gin.Param{Key: "class_id", Value: classID})
	ctx.Params = append(ctx.Params, gin.Param{Key: "title", Value: title})
	return ctx
}
