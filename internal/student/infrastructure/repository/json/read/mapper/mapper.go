package mapper

import (
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/dto"
)

type Mapper struct{}

func (m Mapper) DTOProfileToDomain(email string, profile dto.Profile) domain.Profile {
	return *domain.NewProfile(
		email,
		profile.Name,
	)
}

func (m Mapper) DTOClassesToDomain(classes []dto.Class) []domain.Class {
	domainClasses := make([]domain.Class, len(classes))
	for i := range classes {
		domainClasses[i] = *domain.NewClass(
			classes[i].ClassID,
			classes[i].Title,
		)
	}
	return domainClasses
}
