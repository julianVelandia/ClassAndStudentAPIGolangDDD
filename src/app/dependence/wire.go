package dependence

import (
	useCaseClass "github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/application/usecase"
	repositoryViewClass "github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/infrastructure/repository/json/read"
	mapperRepositoryViewClass "github.com/julianVelandia/EDteam/DDDyCQRS/internal/class/infrastructure/repository/json/read/mapper"
	useCaseStudent "github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/application/usecase"
	repositoryViewProfile "github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/read"
	mapperRepositoryViewProfile "github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/read/mapper"
	repositoryProfileUpdate "github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/update"
	repositoryUpdateProfile "github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/update"
	mapperRepositoryUpdateProfile "github.com/julianVelandia/EDteam/DDDyCQRS/internal/student/infrastructure/repository/json/update/mapper"
	handlerViewClass "github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/class/view"
	mapperViewClass "github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/class/view/mapper"
	handlerUpdateClassesDone "github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/updateclassesdone"
	mapperUpdateProfile "github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/updateclassesdone/mapper"
	handlerViewProfile "github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/view"
	mapperViewProfile "github.com/julianVelandia/EDteam/DDDyCQRS/src/handler/student/view/mapper"
)

type HandlerContainer struct {
	ViewProfileHandler       handlerViewProfile.Handler
	ViewClassHandler         handlerViewClass.Handler
	UpdateClassesDoneHandler handlerUpdateClassesDone.Handler
}

func NewWire() HandlerContainer {
	filenameProfile := "dbtest/prod/StudentsProfile.json"
	filenameClasses := "dbtest/prod/Classes.json"

	repositoryClassRead := repositoryViewClass.NewClassRepositoryRead(
		mapperRepositoryViewClass.Mapper{},
		filenameClasses,
	)
	repositoryProfileRead := repositoryViewProfile.NewProfileRepositoryRead(
		mapperRepositoryViewProfile.Mapper{},
		filenameProfile,
	)
	repositoryProfileUpdateClasses := repositoryProfileUpdate.NewClassRepositoryUpdate(
		mapperRepositoryUpdateProfile.Mapper{},
		filenameProfile,
	)

	return HandlerContainer{
		ViewClassHandler: newWireViewClassHandler(
			*repositoryClassRead,
		),
		ViewProfileHandler: newWireViewProfileHandler(
			*repositoryProfileRead,
		),
		UpdateClassesDoneHandler: newWireUpdateClassesDoneHandler(
			*repositoryProfileUpdateClasses,
		),
	}
}

func newWireViewClassHandler(
	repositoryViewClass repositoryViewClass.ClassRepositoryRead,
) handlerViewClass.Handler {

	useCaseView := useCaseClass.NewViewUseCase(
		repositoryViewClass,
	)
	return *handlerViewClass.NewHandler(
		mapperViewClass.Mapper{},
		useCaseView,
	)
}

func newWireViewProfileHandler(
	repositoryViewProfile repositoryViewProfile.ProfileRepositoryRead,
) handlerViewProfile.Handler {
	useCaseViewProfile := useCaseStudent.NewViewUseCase(
		repositoryViewProfile,
	)

	return *handlerViewProfile.NewHandler(
		mapperViewProfile.Mapper{},
		useCaseViewProfile,
	)
}

func newWireUpdateClassesDoneHandler(
	repositoryUpdateProfile repositoryUpdateProfile.ClassRepositoryUpdate,
) handlerUpdateClassesDone.Handler {
	useCaseUpdateProfile := useCaseStudent.NewUpdateProfileUseCase(
		repositoryUpdateProfile,
	)

	return *handlerUpdateClassesDone.NewHandler(
		mapperUpdateProfile.Mapper{},
		useCaseUpdateProfile,
	)
}
