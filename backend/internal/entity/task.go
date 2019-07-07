package entity

type UUID = string
type TaskName = string

type Task struct {
	uuid UUID
	name TaskName
}

func NewTask(uuid UUID, name TaskName) Task {
	return Task{
		uuid: uuid,
		name: name,
	}
}

func (t Task) UUID() UUID {
	return t.uuid
}

func (t Task) Name() TaskName {
	return t.name
}
