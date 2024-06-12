package user

import (
	"fmt"
	"golang-second-project/model"
)

type student struct {
	name string

	english model.Subject
	math    model.Subject
}

func (s student) GetName() string {
	return s.name
}

func (s *student) SetName(name string) {
	s.name = name
}

func (s student) GetInfo() string {
	return fmt.Sprintf("%s, %s, %s", s.english.GetSubjectName(), s.math.GetSubjectName(), s.name)
}

func NewStudent(english, math model.Subject) model.Student {
	return &student{
		english: english,
		math:    math,
	}
}
