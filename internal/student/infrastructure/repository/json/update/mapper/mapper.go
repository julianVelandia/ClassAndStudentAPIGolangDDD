package mapper

import (
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/dto"
)

type Mapper struct{}

func (m Mapper) CommandToDTOClass(cmd command.UpdateClassesDone) dto.Class {
	return dto.Class{
		ClassID: cmd.ClassID(),
		Title:   cmd.Title(),
	}
}
