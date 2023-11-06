package mapper_test

import (
	"testing"

	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/dto"
	"github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/read/mapper"
	"github.com/stretchr/testify/assert"
)

const (
	email    = "test@email.com"
	name     = "Julián Velandia"
	classID1 = "123"
	title1   = "Class 1"
	classID2 = "456"
	title2   = "Class 2"
)

func TestMapper_DTOProfileToDomain(t *testing.T) {
	mapperRead := mapper.Mapper{}
	profile := dto.Profile{
		Email: email,
		Name:  name,
	}

	domainProfile := mapperRead.DTOProfileToDomain(email, profile)

	assert.Equal(t, email, domainProfile.Email())
	assert.Equal(t, name, domainProfile.Name())
}

func TestMapper_DTOClassesToDomain(t *testing.T) {
	mapperRead := mapper.Mapper{}
	classesDTO := []dto.Class{
		{ClassID: classID1, Title: title1},
		{ClassID: classID2, Title: title2},
	}

	domainClasses := mapperRead.DTOClassesToDomain(classesDTO)

	assert.Len(t, domainClasses, len(classesDTO))
	assert.Equal(t, classID1, domainClasses[0].ClassID())
	assert.Equal(t, title1, domainClasses[0].Title())
	assert.Equal(t, classID2, domainClasses[1].ClassID())
	assert.Equal(t, title2, domainClasses[1].Title())
}
