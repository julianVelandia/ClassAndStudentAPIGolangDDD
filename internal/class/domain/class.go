package domain

type Class struct {
	classID      string
	title        string
	creationDate string
	content      []string
	readTime     float64
}

func (c Class) ClassID() string {
	return c.classID
}

func (c Class) Title() string {
	return c.title
}

func (c Class) CreationDate() string {
	return c.creationDate
}

func (c Class) Content() []string {
	return c.content
}

func (c Class) ReadTime() float64 {
	return c.readTime
}

func NewClass(classID string, title string, creationDate string, content []string, readTime float64) *Class {
	return &Class{classID: classID, title: title, creationDate: creationDate, content: content, readTime: readTime}
}
