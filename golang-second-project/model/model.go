package model

type Student interface {
	GetName() string
	SetName(name string)

	GetInfo() string
}

type Teacher interface {
	GetName() string
	SetName(name string)
}

type Subject interface {
	GetSubjectName() string
}
