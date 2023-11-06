package view_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/query"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/view"
	"github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/view/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	email      = "test@email.com"
	name       = "julián"
	classID1   = "class1"
	classID2   = "class2"
	className1 = "className1"
	className2 = "className2"
)

func TestHandlerWhenGetSuccessfulShouldReturnOK(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	gin.SetMode(gin.TestMode)
	expectedQuery := *query.NewView(email)
	expectedProfile := *domain.NewProfile(email, name)
	expectedClasses := []domain.Class{
		*domain.NewClass(classID1, className1),
		*domain.NewClass(classID2, className2),
	}
	expectedResponse := contract.Response{
		Email: email,
		Name:  name,
		ClassesDone: []contract.Class{
			{classID1,
				className1,
			},
			{classID2,
				className2,
			},
		},
	}
	mapperMock := new(MapperMock)
	useCaseMock := new(UseCaseMock)

	mapperMock.On("RequestToQuery", mock.Anything).Return(
		expectedQuery, nil,
	).Once()
	useCaseMock.On("Execute", expectedQuery.Email()).Return(
		expectedProfile,
		expectedClasses,
		nil,
	).Once()
	mapperMock.On("DomainToResponse", expectedProfile, expectedClasses).Return(
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
	expectedQuery := *query.NewView(email)
	mapperMock := new(MapperMock)
	useCaseMock := new(UseCaseMock)

	mapperMock.On("RequestToQuery", mock.Anything).Return(
		expectedQuery, nil,
	).Once()
	useCaseMock.On("Execute", expectedQuery.Email()).Return(
		domain.Profile{},
		[]domain.Class{},
		errors.New("usecase fail"),
	).Once()

	handler := view.NewHandler(mapperMock, useCaseMock)
	handler.Handler(ctx)

	assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	mapperMock.AssertExpectations(t)
	useCaseMock.AssertExpectations(t)
}
