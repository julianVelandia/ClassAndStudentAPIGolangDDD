package contract

type Request struct {
	ClassID string `json:"class_id" binding:"required"`
}
