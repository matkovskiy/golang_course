package structs

type Student struct {
	Name     string
	Surname  string
	Subjects []Subject
}

type Subject struct {
	Name       string `json:"name"`
	Techical   bool   `json:"techical"`
	Complexity int    `json:"complexity"`
}

type Subjects struct {
	Subjects []Subject `json:"subjects"`
}

func (student *Student) Learn_subject(subject Subject) {
	student.Subjects = append(student.Subjects, subject)
}
