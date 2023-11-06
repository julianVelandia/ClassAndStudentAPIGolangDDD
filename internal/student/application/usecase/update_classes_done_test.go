package usecase_test

import (
	"errors"
	"testing"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/usecase"
	"github.com/stretchr/testify/assert"
)

const (
	classID string = "classID"
	title   string = "title"
)

func TestUpdateExecuteWhenRepositoryResponseOkShouldReturnOK(t *testing.T) {
	cmd := *command.NewUpdateClassesDone(
		email, classID, title,
	)
	repositoryUpdateProfile := new(RepositoryUpdateProfileMock)

	repositoryUpdateProfile.On("UpdateClassesDoneInUserProfile", cmd).Return(nil).Once()
	updateUseCase := usecase.NewUpdateProfileUseCase(
		repositoryUpdateProfile,
	)
	errResult := updateUseCase.Execute(cmd)

	assert.NoError(t, errResult)
	repositoryUpdateProfile.AssertExpectations(t)
}

func TestUpdateExecuteWhenUpdateClassesDoneInUserProfileFailShouldReturnError(t *testing.T) {
	cmd := *command.NewUpdateClassesDone(
		email, classID, title,
	)
	repositoryUpdateProfile := new(RepositoryUpdateProfileMock)

	repositoryUpdateProfile.On("UpdateClassesDoneInUserProfile", cmd).Return(errors.New("")).Once()
	updateUseCase := usecase.NewUpdateProfileUseCase(
		repositoryUpdateProfile,
	)
	errResult := updateUseCase.Execute(cmd)
	assert.Error(t, errResult)
	repositoryUpdateProfile.AssertExpectations(t)
}
