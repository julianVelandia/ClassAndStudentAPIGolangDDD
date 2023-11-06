package read

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/domain"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/dto"
)

type Mapper interface {
	DTOProfileToDomain(email string, profile dto.Profile) domain.Profile
	DTOClassesToDomain(classes []dto.Class) []domain.Class
}

type ProfileRepositoryRead struct {
	mapper          Mapper
	filenameProfile string
}

func NewProfileRepositoryRead(mapper Mapper, filenameProfile string) *ProfileRepositoryRead {
	return &ProfileRepositoryRead{mapper: mapper, filenameProfile: filenameProfile}
}

func (r ProfileRepositoryRead) GetProfileByEmail(email string) (domain.Profile, []domain.Class, error) {
	data, err := os.ReadFile(r.filenameProfile)
	if err != nil {
		return domain.Profile{}, []domain.Class{}, err
	}

	profiles := make(map[string]dto.Profile)
	err = json.Unmarshal(data, &profiles)
	if err != nil {
		return domain.Profile{}, []domain.Class{}, err
	}

	foundProfileDTO, found := profiles[email]
	if !found {
		return domain.Profile{}, []domain.Class{}, fmt.Errorf("profile not found for email: %s", email)
	}

	foundProfile := r.mapper.DTOProfileToDomain(email, foundProfileDTO)

	classesDone := r.mapper.DTOClassesToDomain(foundProfileDTO.ClassesDone)
	return foundProfile, classesDone, nil
}
