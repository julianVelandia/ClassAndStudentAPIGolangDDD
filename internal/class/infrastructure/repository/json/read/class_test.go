package read_test

import (
	"errors"
	"testing"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/infrastructure/repository/json/dto"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/infrastructure/repository/json/read"
	"github.com/stretchr/testify/assert"
)

const (
	title         = "title Class 1"
	creationDate  = "08/26/2023"
	readTime      = 5.5
	classID       = "classid1"
	name          = "Julián"
	filenameClass = "../../../../../../dbtest/dev/Classes.json"
)

func TestGetClassByClassIDWhenSuccessShouldReturnProfile(t *testing.T) {
	mockMapper := new(MapperMock)
	content := []string{"content1", "content2"}
	expectedClassDTO := dto.Class{
		Title:        title,
		CreationDate: creationDate,
		Content:      content,
		ReadTime:     readTime,
	}
	expectedClass := *domain.NewClass(
		classID,
		title,
		creationDate,
		content,
		readTime,
	)
	mockMapper.On("DTOClassToDomain", expectedClassDTO).Return(expectedClass, nil)
	repository := read.NewClassRepositoryRead(mockMapper, filenameClass)
	result, err := repository.GetClassByClassID(classID)

	assert.NoError(t, err)
	assert.Equal(t, expectedClass, result)
	mockMapper.AssertExpectations(t)
}

func TestGetClassByClassIDWhenReadFailShouldReturnError(t *testing.T) {
	mockMapper := new(MapperMock)

	repository := read.NewClassRepositoryRead(nil, "")
	result, err := repository.GetClassByClassID(classID)

	assert.Error(t, err)
	assert.Equal(t, result, domain.Class{})
	mockMapper.AssertExpectations(t)
}

func TestGetClassByClassIDWhenMapperNotFoundShouldReturnEmpty(t *testing.T) {
	mockMapper := new(MapperMock)
	content := []string{"content1", "content2"}
	expectedClassDTO := dto.Class{
		Title:        title,
		CreationDate: creationDate,
		Content:      content,
		ReadTime:     readTime,
	}

	mockMapper.On("DTOClassToDomain", expectedClassDTO).Return(domain.Class{}, errors.New(""))
	repository := read.NewClassRepositoryRead(mockMapper, filenameClass)
	result, err := repository.GetClassByClassID(classID)

	assert.NoError(t, err)
	assert.Equal(t, result, domain.Class{})
	mockMapper.AssertExpectations(t)
}
