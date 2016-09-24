package easy_fmt

import (
	"fmt"
	"testing"
)

type GenderType bool

const (
	Gender_Male   GenderType = false
	Gender_Female GenderType = true
)

type Person struct {
	Name   string
	Gender GenderType
	Age    int
}

type Teacher struct {
	*Person
	HomePage string
}

type Course struct {
	Name     string
	ID       string
	Teacher  *Teacher
	Students []*Student
}

type Student struct {
	*Person
	StudentId int64
}

var people = []*Person{
	&Person{
		Name:   "Jack",
		Gender: Gender_Male,
		Age:    19,
	},
	&Person{
		Name:   "Mark",
		Gender: Gender_Male,
		Age:    21,
	},
	&Person{
		Name:   "Marry",
		Gender: Gender_Female,
		Age:    20,
	},
	&Person{
		Name:   "Justin",
		Gender: Gender_Male,
		Age:    35,
	},
	&Person{
		Name:   "Hank",
		Gender: Gender_Male,
		Age:    32,
	},
	nil,
}

var students = []*Student{
	&Student{
		Person:    people[0],
		StudentId: 101,
	},
	&Student{
		Person:    people[1],
		StudentId: 102,
	},
	&Student{
		Person:    people[2],
		StudentId: 103,
	},
}

var teachers = []*Teacher{
	&Teacher{
		Person:   people[3],
		HomePage: "www.google.com",
	},
	&Teacher{
		Person:   people[4],
		HomePage: "www.facebook.com",
	},
	&Teacher{
		Person:   people[5], // nil
		HomePage: "www.linkedin.com",
	},
	nil,
}

var courses = []*Course{
	&Course{
		Name:     "Math",
		ID:       "mat-01",
		Teacher:  teachers[0],
		Students: []*Student{students[0], students[1], nil},
	},
	&Course{
		Name:     "Literatue",
		ID:       "lit-01",
		Teacher:  teachers[1],
		Students: []*Student{students[1], students[2]},
	},
	&Course{
		Name:     "Football",
		ID:       "spt-01",
		Teacher:  nil,
		Students: []*Student{students[1], students[2]},
	},
}

func TestStructToString(t *testing.T) {
	fmt.Println(EasyFormat(courses))
}
