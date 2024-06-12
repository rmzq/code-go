package main

import (
	"fmt"
	"golang-second-project/subject"
	"golang-second-project/user"
)

func main() {
	english := subject.NewEnglish("Bahasa Inggris")
	math := subject.NewMath("Matematika")

	student := user.NewStudent(english, math)
	student.SetName("Jason")

	fmt.Println(student.GetName())

}
