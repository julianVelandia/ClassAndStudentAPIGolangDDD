package query

type View struct {
	email string
}

func (v View) Email() string {
	return v.email
}

func NewView(email string) *View {
	return &View{email: email}
}
