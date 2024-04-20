package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/revandpratama/assignment-testing"
)

var _ = Describe("Main Test", func() {
	Describe("Test Student GetInfo Behaviour", func ()  {
		It("Should return student data", func() {
			//arrange
			var StudentList = []main.Student{
				{
					Person: main.Person{
						Name: "Esa",
						Age:  "18",
					},
					StudentID:   "1001",
					Class:       "XII",
					TeacherName: "Budi",
				},
				{
					Person: main.Person{
						Name: "Gian",
						Age:  "17",
					},
					StudentID:   "1004",
					Class:       "XI",
					TeacherName: "Diki",
				},
			}
		
			s1 := StudentList[0]
			s2 := StudentList[1]

			//act
			expected1 := map[string]string{
				"Name":        "Esa",
				"Age":         "18",
				"StudentID":   "1001",
				"Class":       "XII",
				"TeacherName": "Budi",
			}
			expected2 := map[string]string{
				"Name":        "Gian",
				"Age":         "17",
				"StudentID":   "1004",
				"Class":       "XI",
				"TeacherName": "Diki",
			}

			test1 := s1.GetInfo()
			test2 := s2.GetInfo()
			//assert

			Expect(test1).To(Equal(expected1))
			Expect(test2).To(Equal(expected2))
		})
	})

	Describe("Test Teacher Total Salary Calculation", func ()  {
		It("Should return correct total salary of teacher", func() {
			//arrange
			teach1 := main.Teacher{
				Person:      main.Person{"Angga", "12"},
				TeacherID:   "12121212",
				Salary:      13000,
				ActiveMonth: 24,
			}
		
			//act
			totalSalary := teach1.TotalSalary()
			expected := 312000

			//assert
			Expect(totalSalary).To(Equal(expected))
		})
	})

	Describe("Test Sort Student Based off Teacher", func() {
		It("Should return student with the TeacherName of the teacher object", func() {
			//assert
			student1 := main.Student{
				Person: main.Person{
					Name: "Esa",
					Age:  "18",
				},
				StudentID:   "1001",
				Class:       "XII",
				TeacherName: "Budi",
			}
		
			student2 := main.Student{
				Person: main.Person{
					Name: "Luna",
					Age:  "17",
				},
				StudentID:   "1003",
				Class:       "XI",
				TeacherName: "Diki",
			}
			teacher1 := main.Teacher{
				Person:      main.Person{"Budi", "35"},
				TeacherID:   "12121212",
				Salary:      12000,
				ActiveMonth: 24,
			}
			teacher2 := main.Teacher{
				Person:      main.Person{"Diki", "45"},
				TeacherID:   "13131313",
				Salary:      15000,
				ActiveMonth: 36,
			}

			//act
			test1 := teacher1.ShowStudent()
			test2 := teacher2.ShowStudent()

			//assert

			Expect(test1[0]).To(Equal(student1))
			Expect(test2[0]).To(Equal(student2))
		})
	})
})
