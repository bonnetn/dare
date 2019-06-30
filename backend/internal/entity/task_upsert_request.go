package entity

type TaskUpsertRequest interface {
	UUID() UUID
	Version() Version

	Name() TaskName
	Content() Content
	RequestID() string
}

func NewTaskUpsertRequest(uuid, name, content, requestID string, version int64) TaskUpsertRequest {
	return &taskUpsertRequest{
		Task: Task{
			uuid:    uuid,
			version: version,
			name:    name,
			content: content,
		},
		requestID: requestID,
	}
}

type taskUpsertRequest struct {
	Task
	requestID string
}

func (t taskUpsertRequest) RequestID() string {
	return t.requestID
}
