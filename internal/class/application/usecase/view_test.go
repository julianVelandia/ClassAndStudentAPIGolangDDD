package usecase_test

import (
	"errors"
	"testing"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/application/query"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/application/usecase"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/domain"
	"github.com/stretchr/testify/assert"
)

const (
	classID      = "class_1"
	title        = "title"
	creationDate = "26/08/2023"
	content      = "content"
	readTime     = 5.5
)

func TestViewExecuteWhenRepositoryResponseOkShouldReturnOK(t *testing.T) {
	classDomain := *domain.NewClass(
		classID,
		title,
		creationDate,
		[]string{content, content},
		readTime,
	)

	qry := *query.NewView(
		classID,
	)
	repositoryViewMock := new(RepositoryViewClassMock)

	repositoryViewMock.On("GetClassByClassID", classID).Return(classDomain, nil).Once()
	getUseCase := usecase.NewViewUseCase(
		repositoryViewMock,
	)
	result, errResult := getUseCase.Execute(qry)

	assert.NoError(t, errResult)
	assert.Equal(t, result, classDomain)
	repositoryViewMock.AssertExpectations(t)
}

func TestViewExecuteWhenRepositoryViewFailShouldReturnError(t *testing.T) {
	qry := *query.NewView(
		classID,
	)
	repositoryViewMock := new(RepositoryViewClassMock)

	repositoryViewMock.On("GetClassByClassID", classID).Return(
		domain.Class{},
		errors.New(""),
	).Once()
	getUseCase := usecase.NewViewUseCase(
		repositoryViewMock,
	)
	result, errResult := getUseCase.Execute(qry)

	assert.Error(t, errResult)
	assert.Equal(t, result, domain.Class{})
	repositoryViewMock.AssertExpectations(t)
}
