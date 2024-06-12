package user

import "golang-second-project/model"

type teacher struct {
	name string
}

func (t teacher) GetName() string {
	return t.name
}

func (t *teacher) SetName(name string) {
	t.name = name
}

func NewTeacher(name string) model.Teacher {
	return &teacher{}
}
