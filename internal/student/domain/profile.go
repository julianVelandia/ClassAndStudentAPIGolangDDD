package domain

type Profile struct {
	email string
	name  string
}

func (p Profile) Email() string {
	return p.email
}

func (p Profile) Name() string {
	return p.name
}

func NewProfile(email string, name string) *Profile {
	return &Profile{email: email, name: name}
}
