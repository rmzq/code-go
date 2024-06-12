package subject

import "golang-second-project/model"

type math struct {
	subjectName string
}

func (m math) GetSubjectName() string {
	return m.subjectName
}

func NewMath(subjectName string) model.Subject {
	return math{subjectName: subjectName}
}
