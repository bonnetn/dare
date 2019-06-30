package entity

type UUID = string
type Version = int64
type TaskName = string
type Content = string

var NullTask Task

type Task struct {
	uuid    UUID
	version Version
	name    TaskName
	content Content
}

func NewTask(uuid UUID, version Version, name TaskName, content Content) Task {
	return Task{
		uuid:    uuid,
		version: version,
		name:    name,
		content: content,
	}
}

func (t Task) UUID() UUID {
	return t.uuid
}

func (t Task) Version() Version {
	return t.version
}

func (t Task) Name() TaskName {
	return t.name
}

func (t Task) Content() Content {
	return t.content
}
