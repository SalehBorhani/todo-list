package requestparam

import "time"

type Request struct {
	Command           string
	CreateTaskRequest CreateTaskRequest
}

type CreateTaskRequest struct {
	Title      string
	DueDate    time.Time
	CategoryID int
}
