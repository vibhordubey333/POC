package grades

import "fmt"

type GradeType string
type Students []Student

var students Students


type Student struct{
	ID int
	FirstName string
	LastName string
	Grades []Grade
}

const(
	GradeTest = GradeType("Test")
	GradeHomework = GradeType("Homework")
	GradeQuiz 	= GradeType("Quiz")
)

type Grade struct{
	Title string
	Type GradeType
	Score float32
}

func (s Student) Average() float32{
	var result float32
	for _,grade := range s.Grades{
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

func (s Students) GetByID(id int) (*Student,error){
	for i := range s{
		if s[i].ID == id{
			return &s[i],nil
		}
	}
	return nil, fmt.Errorf("student with ID %v not found",id)
}