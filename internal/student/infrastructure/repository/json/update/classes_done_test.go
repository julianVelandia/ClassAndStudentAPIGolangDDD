package update_test

import (
	"testing"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/dto"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/update"
	"github.com/stretchr/testify/assert"
)

const (
	email         = "test1@email.com"
	title         = "title Class 1"
	classID       = "classid1"
	name          = "Julián"
	filenameClass = "../../../../../../dbtest/dev/StudentsProfile.json"
)

func TestUpdateClassesByEmailWhenSuccessShouldReturnOk(t *testing.T) {
	mockMapper := new(MapperMock)
	updateCmd := *command.NewUpdateClassesDone(
		email,
		classID,
		title,
	)
	classDTO := dto.Class{
		ClassID: classID,
		Title:   title,
	}

	mockMapper.On("CommandToDTOClass", updateCmd).Return(classDTO)
	repository := update.NewClassRepositoryUpdate(mockMapper, filenameClass)
	err := repository.UpdateClassesDoneInUserProfile(updateCmd)

	assert.NoError(t, err)
	mockMapper.AssertExpectations(t)
}

func TestUpdateClassesByEmailWhenReadFailShouldReturnError(t *testing.T) {
	mockMapper := new(MapperMock)

	repository := update.NewClassRepositoryUpdate(nil, "")
	err := repository.UpdateClassesDoneInUserProfile(command.UpdateClassesDone{})

	assert.Error(t, err)
	mockMapper.AssertExpectations(t)
}
