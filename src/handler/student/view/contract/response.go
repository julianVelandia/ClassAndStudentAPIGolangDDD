package contract

type Response struct {
	Email       string  `json:"email"`
	Name        string  `json:"name"`
	ClassesDone []Class `json:"classes_done"`
}

type Class struct {
	ClassID string `json:"class_id"`
	Title   string `json:"title"`
}
