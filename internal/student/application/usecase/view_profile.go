package usecase

import (
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
)

type RepositoryViewProfile interface {
	GetProfileByEmail(email string) (domain.Profile, []domain.Class, error)
}

type ViewUseCase struct {
	repositoryViewProfile RepositoryViewProfile
}

func NewViewUseCase(repositoryViewProfile RepositoryViewProfile) *ViewUseCase {
	return &ViewUseCase{repositoryViewProfile: repositoryViewProfile}
}

func (uc ViewUseCase) Execute(email string) (domain.Profile, []domain.Class, error) {

	domainProfile, classesDone, err := uc.repositoryViewProfile.GetProfileByEmail(email)
	if err != nil {
		return domain.Profile{}, []domain.Class{}, err
	}

	return domainProfile, classesDone, nil
}
