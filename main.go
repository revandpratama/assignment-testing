package main

import "fmt"

type Person struct {
	Name string
	Age  string
}

type Student struct {
	Person
	StudentID   string
	Class       string
	TeacherName string
}

type Teacher struct {
	Person
	TeacherID   string
	Salary      int
	ActiveMonth int
}

var StudentList = []*Student{
	{
		Person: Person{
			Name: "Esa",
			Age:  "18",
		},
		StudentID:   "1001",
		Class:       "XII",
		TeacherName: "Budi",
	},
	{
		Person: Person{
			Name: "Dian",
			Age:  "19",
		},
		StudentID:   "1002",
		Class:       "XII",
		TeacherName: "Budi",
	},
	{
		Person: Person{
			Name: "Luna",
			Age:  "17",
		},
		StudentID:   "1003",
		Class:       "XI",
		TeacherName: "Diki",
	},
	{
		Person: Person{
			Name: "Gian",
			Age:  "17",
		},
		StudentID:   "1004",
		Class:       "XI",
		TeacherName: "Diki",
	},
}

func main() {
	teacher1 := Teacher{
		Person:      Person{"Budi", "35"},
		TeacherID:   "12121212",
		Salary:      12000,
		ActiveMonth: 24,
	}
	teacher2 := Teacher{
		Person:      Person{"Diki", "45"},
		TeacherID:   "13131313",
		Salary:      15000,
		ActiveMonth: 36,
	}

	fmt.Println(teacher1.ShowStudent())
	fmt.Println(teacher2.ShowStudent())

	for _, s := range StudentList {
		fmt.Println(s.GetInfo())
	}
}

func (s *Student) GetInfo() map[string]string {
	return map[string]string{
		"Name":        s.Name,
		"Age":         s.Age,
		"StudentID":   s.StudentID,
		"Class":       s.Class,
		"TeacherName": s.TeacherName,
	}
}


func (t *Teacher) ShowStudent() []Student {
	var res []Student

	for _, s := range StudentList {
		if s.TeacherName == t.Name {
			res = append(res, *s)
		}
	}
	return res
}

func (t *Teacher) TotalSalary() int {
	return t.Salary * t.ActiveMonth
}
func (t *Teacher) TotalSalaryu() int {
	return t.Salary * t.ActiveMonth
}
