package mapper

import (
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
	dto2 "github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/dto"
)

type Mapper struct{}

func (m Mapper) DTOProfileToDomain(email string, profile dto2.Profile) domain.Profile {
	return *domain.NewProfile(
		email,
		profile.Name,
	)
}

func (m Mapper) DTOClassesToDomain(classes []dto2.Class) []domain.Class {
	domainClasses := make([]domain.Class, len(classes))
	for i := range classes {
		domainClasses[i] = *domain.NewClass(
			classes[i].ClassID,
			classes[i].Title,
		)
	}
	return domainClasses
}
