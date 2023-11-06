package mapper_test

import (
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/command"
	"testing"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/update/mapper"
	"github.com/stretchr/testify/assert"
)

const (
	email   = "test@email.com"
	classID = "123"
	title   = "Updated Class"
)

func TestMapper_CommandToDTOClass(t *testing.T) {
	mapperWrite := mapper.Mapper{}
	updateCmd := command.NewUpdateClassesDone(
		email,
		classID,
		title,
	)

	dtoClass := mapperWrite.CommandToDTOClass(*updateCmd)

	assert.Equal(t, classID, dtoClass.ClassID)
	assert.Equal(t, title, dtoClass.Title)
}
