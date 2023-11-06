package mapper_test

import (
	"testing"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/infrastructure/repository/json/dto"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/infrastructure/repository/json/read/mapper"
	"github.com/stretchr/testify/assert"
)

const (
	classID      = "123"
	title        = "Test Class"
	creationDate = "2023-08-27"
	readTime     = 1.5
)

func TestMapper_DTOClassToDomain(t *testing.T) {
	mapperRead := mapper.Mapper{}
	content := []string{"Content line 1", "Content line 2"}
	classDTO := dto.Class{
		ClassID:      classID,
		Title:        title,
		CreationDate: creationDate,
		Content:      content,
		ReadTime:     readTime,
	}

	domainClass := mapperRead.DTOClassToDomain(classDTO)

	assert.Equal(t, classID, domainClass.ClassID())
	assert.Equal(t, title, domainClass.Title())
	assert.Equal(t, creationDate, domainClass.CreationDate())
	assert.Equal(t, content, domainClass.Content())
	assert.Equal(t, readTime, domainClass.ReadTime())
}
