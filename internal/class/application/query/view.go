package query

type View struct {
	classID string
}

func NewView(classID string) *View {
	return &View{classID: classID}
}

func (v View) ClassID() string {
	return v.classID
}
