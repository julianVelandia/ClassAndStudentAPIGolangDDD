package command

type UpdateClassesDone struct {
	email   string
	classID string
	title   string
}

func NewUpdateClassesDone(email string, classID string, title string) *UpdateClassesDone {
	return &UpdateClassesDone{email: email, classID: classID, title: title}
}

func (u UpdateClassesDone) Title() string {
	return u.title
}

func (u UpdateClassesDone) Email() string {
	return u.email
}

func (u UpdateClassesDone) ClassID() string {
	return u.classID
}
