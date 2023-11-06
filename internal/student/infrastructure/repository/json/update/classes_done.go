package update

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/dto"
)

type Mapper interface {
	CommandToDTOClass(cmd command.UpdateClassesDone) dto.Class
}

type ClassRepositoryUpdate struct {
	mapper          Mapper
	filenameProfile string
}

func NewClassRepositoryUpdate(mapper Mapper, filenameProfile string) *ClassRepositoryUpdate {
	return &ClassRepositoryUpdate{mapper: mapper, filenameProfile: filenameProfile}
}

func (r ClassRepositoryUpdate) UpdateClassesDoneInUserProfile(cmd command.UpdateClassesDone) error {
	data, err := os.ReadFile(r.filenameProfile)
	if err != nil {
		return err
	}

	profiles := make(map[string]dto.Profile)
	err = json.Unmarshal(data, &profiles)
	if err != nil {
		return err
	}

	profile, found := profiles[cmd.Email()]
	if !found {
		return errors.New("profile not Found")
	}

	newClass := r.mapper.CommandToDTOClass(cmd)
	profile.ClassesDone = append(profile.ClassesDone, newClass)
	profiles[cmd.Email()] = profile

	updatedData, err := json.MarshalIndent(profiles, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filenameProfile, updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
