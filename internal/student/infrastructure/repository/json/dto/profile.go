package dto

type Profile struct {
	Email       string  `json:"email"`
	Name        string  `json:"name"`
	ClassesDone []Class `json:"classes_done"`
}
