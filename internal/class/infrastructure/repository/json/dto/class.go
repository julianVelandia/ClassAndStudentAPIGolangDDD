package dto

type Class struct {
	ClassID      string   `json:"class_id"`
	Title        string   `json:"title"`
	CreationDate string   `json:"creation_date"`
	Content      []string `json:"content"`
	ReadTime     float64  `json:"read_time"`
}

type ClassStudent struct {
	ClassID string `json:"class_id"`
	Title   string `json:"title"`
}
