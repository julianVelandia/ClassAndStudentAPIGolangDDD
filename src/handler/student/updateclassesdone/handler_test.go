package updateclassesdone_test

import (
	"errors"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/updateclassesdone"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	email   = "test@email.com"
	name    = "julián"
	classID = "class1"
	title   = "className1"
)

func TestHandlerWhenUpdateSuccessfulShouldReturnOK(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	gin.SetMode(gin.TestMode)
	cmd := *command.NewUpdateClassesDone(
		email, classID, title,
	)
	mapperMock := new(MapperMock)
	useCaseMock := new(UseCaseMock)

	mapperMock.On("RequestToCommand", mock.Anything).Return(
		cmd,
	).Once()
	useCaseMock.On("Execute", cmd).Return(
		nil,
	).Once()
	handler := updateclassesdone.NewHandler(mapperMock, useCaseMock)
	handler.Handler(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	mapperMock.AssertExpectations(t)
	useCaseMock.AssertExpectations(t)
}

func TestHandlerWhenUseCaseFailShouldReturnInternalError(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	gin.SetMode(gin.TestMode)
	cmd := *command.NewUpdateClassesDone(
		email, classID, title,
	)
	mapperMock := new(MapperMock)
	useCaseMock := new(UseCaseMock)

	mapperMock.On("RequestToCommand", mock.Anything).Return(
		cmd,
	).Once()
	useCaseMock.On("Execute", cmd).Return(
		errors.New("usecase fail"),
	).Once()
	handler := updateclassesdone.NewHandler(mapperMock, useCaseMock)
	handler.Handler(ctx)

	assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	mapperMock.AssertExpectations(t)
	useCaseMock.AssertExpectations(t)
}
