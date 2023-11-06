package usecase_test

import (
	"errors"
	"testing"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/usecase"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
	"github.com/stretchr/testify/assert"
)

const (
	email      = "test@email.com"
	classID1   = "class_1"
	classID2   = "class_2"
	name       = "julián"
	className1 = "name 1"
	className2 = "name 2"
)

func TestViewExecuteWhenRepositoryResponseOkShouldReturnOK(t *testing.T) {
	classesDomain := []domain.Class{
		*domain.NewClass(classID1, className1),
		*domain.NewClass(classID2, className2),
	}
	profileDomain := *domain.NewProfile(
		email,
		name,
	)
	repositoryViewProfile := new(RepositoryViewProfileMock)

	repositoryViewProfile.On("GetProfileByEmail", email).Return(profileDomain, classesDomain, nil).Once()
	getUseCase := usecase.NewViewUseCase(
		repositoryViewProfile,
	)
	resultProfile, resultClasses, errResult := getUseCase.Execute(email)

	assert.NoError(t, errResult)
	assert.Equal(t, resultProfile, profileDomain)
	assert.Equal(t, resultClasses, classesDomain)
	repositoryViewProfile.AssertExpectations(t)
}

func TestViewExecuteWhenGetProfileByEmailFailShouldReturnError(t *testing.T) {
	repositoryViewProfile := new(RepositoryViewProfileMock)

	repositoryViewProfile.On("GetProfileByEmail", email).Return(domain.Profile{}, []domain.Class{}, errors.New("")).Once()
	getUseCase := usecase.NewViewUseCase(
		repositoryViewProfile,
	)
	resultProfile, resultClasses, errResult := getUseCase.Execute(email)

	assert.Error(t, errResult)
	assert.Equal(t, resultProfile, domain.Profile{})
	assert.Equal(t, resultClasses, []domain.Class{})
	repositoryViewProfile.AssertExpectations(t)
}
