package read_test

import (
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/dto"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/read"
	"github.com/stretchr/testify/assert"
)

const (
	email           = "test1@email.com"
	name            = "Julián"
	classID         = "classid1"
	title           = "clase 1"
	filenameProfile = "../../../../../../dbtest/dev/StudentsProfile.json"
)

func TestGetProfileByEmailWhenSuccessShouldReturnProfile(t *testing.T) {
	mockMapper := new(MapperMock)
	classesDoneDTO := make([]dto.Class, 0)
	classesDoneDTO = append(classesDoneDTO, dto.Class{
		ClassID: classID,
		Title:   title,
	})
	expectedProfile := *domain.NewProfile("", name)
	expectedClassesDone := make([]domain.Class, 0)
	expectedClassesDone = append(expectedClassesDone,
		*domain.NewClass(classID, title))
	mockMapper.On("DTOProfileToDomain", email, mock.Anything).Return(expectedProfile, nil)
	mockMapper.On("DTOClassesToDomain", mock.Anything).Return(expectedClassesDone, nil)
	repository := read.NewProfileRepositoryRead(mockMapper, filenameProfile)
	resultProfile, resultClasses, err := repository.GetProfileByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, expectedProfile, resultProfile)
	assert.Equal(t, expectedClassesDone, resultClasses)
	mockMapper.AssertExpectations(t)
}

func TestGetProfileByEmailWhenReadFailShouldReturnError(t *testing.T) {
	mockMapper := new(MapperMock)

	repository := read.NewProfileRepositoryRead(nil, "")
	result, resultClasses, err := repository.GetProfileByEmail(email)

	assert.Error(t, err)
	assert.Equal(t, result, domain.Profile{})
	assert.Equal(t, resultClasses, []domain.Class{})
	mockMapper.AssertExpectations(t)
}
