package subject

import "golang-second-project/model"

type english struct {
	subjectName string
}

func (e english) GetSubjectName() string {
	return e.subjectName
}

func NewEnglish(subjectName string) model.Subject {
	return english{subjectName: subjectName}
}
