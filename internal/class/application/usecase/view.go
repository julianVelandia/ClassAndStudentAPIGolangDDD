package usecase

import (
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/application/query"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/domain"
)

type RepositoryViewClass interface {
	GetClassByClassID(classID string) (domain.Class, error)
}

type ViewUseCase struct {
	repositoryViewClass RepositoryViewClass
}

func NewViewUseCase(repositoryViewClass RepositoryViewClass) *ViewUseCase {
	return &ViewUseCase{repositoryViewClass: repositoryViewClass}
}

func (uc ViewUseCase) Execute(qry query.View) (domain.Class, error) {
	domainClass, err := uc.repositoryViewClass.GetClassByClassID(qry.ClassID())
	if err != nil {
		return domain.Class{}, err
	}

	return domainClass, nil
}
