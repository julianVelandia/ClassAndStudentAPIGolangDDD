package usecase

import (
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
)

type RepositoryUpdateProfile interface {
	UpdateClassesDoneInUserProfile(cmd command.UpdateClassesDone) error
}

type UpdateProfileUseCase struct {
	repositoryUpdateProfile RepositoryUpdateProfile
}

func NewUpdateProfileUseCase(repositoryUpdateProfile RepositoryUpdateProfile) *UpdateProfileUseCase {
	return &UpdateProfileUseCase{repositoryUpdateProfile: repositoryUpdateProfile}
}

func (uc UpdateProfileUseCase) Execute(cmd command.UpdateClassesDone) error {

	err := uc.repositoryUpdateProfile.UpdateClassesDoneInUserProfile(cmd)
	if err != nil {
		return err
	}

	return nil
}
