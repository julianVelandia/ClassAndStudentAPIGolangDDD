package domain

type Class struct {
	classID string
	title   string
}

func NewClass(classID string, title string) *Class {
	return &Class{classID: classID, title: title}
}

func (c Class) ClassID() string {
	return c.classID
}

func (c Class) Title() string {
	return c.title
}
