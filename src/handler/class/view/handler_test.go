package view_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/application/query"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/class/view"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/class/view/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	name         = "julián"
	title        = "title"
	classID      = "class1"
	creationDate = "2023-08-27"
	readTime     = 1.5
)

func TestHandlerWhenGetSuccessfulShouldReturnOK(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	gin.SetMode(gin.TestMode)
	expectedQuery := *query.NewView(classID)
	content := []string{"Content line 1", "Content line 2"}
	expectedClass := *domain.NewClass(
		classID,
		title,
		creationDate,
		content,
		readTime,
	)
	expectedResponse := contract.Response{
		ClassID:      classID,
		Title:        title,
		CreationDate: creationDate,
		Content:      content,
		ReadTime:     readTime,
	}
	mapperMock := new(MapperMock)
	useCaseMock := new(UseCaseMock)

	mapperMock.On("RequestToQuery", mock.Anything).Return(
		expectedQuery, nil,
	).Once()
	useCaseMock.On("Execute", expectedQuery).Return(
		expectedClass,
		nil,
	).Once()
	mapperMock.On("DomainToResponse", expectedClass).Return(
		expectedResponse,
	).Once()
	handler := view.NewHandler(mapperMock, useCaseMock)
	handler.Handler(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	mapperMock.AssertExpectations(t)
	useCaseMock.AssertExpectations(t)
}

func TestHandlerWhenRequestEmptyShouldReturnBadRequest(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	gin.SetMode(gin.TestMode)
	mapperMock := new(MapperMock)
	mapperMock.On("RequestToQuery", mock.Anything).Return(
		query.View{}, errors.New("empty request"),
	).Once()

	handler := view.NewHandler(mapperMock, nil)
	handler.Handler(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	mapperMock.AssertExpectations(t)
}

func TestHandlerWhenUseCaseFailShouldReturnInternalError(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	gin.SetMode(gin.TestMode)
	expectedQuery := *query.NewView(classID)
	mapperMock := new(MapperMock)
	useCaseMock := new(UseCaseMock)

	mapperMock.On("RequestToQuery", mock.Anything).Return(
		expectedQuery, nil,
	).Once()
	useCaseMock.On("Execute", expectedQuery).Return(
		domain.Class{},
		errors.New("usecase fail"),
	).Once()
	handler := view.NewHandler(mapperMock, useCaseMock)
	handler.Handler(ctx)

	assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	mapperMock.AssertExpectations(t)
	useCaseMock.AssertExpectations(t)
}
