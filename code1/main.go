package main

import (
	"CODE1/manager"
	"CODE1/student"
	"fmt"
)

func main() {

	m := manager.StudentManager{}

	for {

		fmt.Println()
		fmt.Println("===== Student Management System =====")
		fmt.Println("1. Print Student Details")
		fmt.Println("2. Add Student")
		fmt.Println("3. Delete Student")
		fmt.Println("4. Add Subject")
		fmt.Println("5. Update Year")
		fmt.Println("6. List All Students")
		fmt.Println("7. Count Students")
		fmt.Println("8. Exit")

		var choice int

		fmt.Print("Enter Choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var prn string

			fmt.Print("Enter PRN: ")
			fmt.Scan(&prn)

			if !m.FindAndPrint(prn) {
				fmt.Println("Student not found.")
			}

		case 2:

			var s student.Student

			fmt.Print("Enter PRN: ")
			fmt.Scan(&s.PRN)

			fmt.Print("Enter Name: ")
			fmt.Scan(&s.Name)

			fmt.Print("Enter Branch: ")
			fmt.Scan(&s.Branch)

			fmt.Print("Enter Year: ")
			fmt.Scan(&s.Year)

			fmt.Print("Enter Division: ")
			fmt.Scan(&s.Division)

			fmt.Print("Enter Roll Number: ")
			fmt.Scan(&s.RollNo)

			fmt.Print("Enter CGPA: ")
			fmt.Scan(&s.CGPA)

			var n int

			fmt.Print("Number of Subjects: ")
			fmt.Scan(&n)

			s.Subjects = make([]string, n)

			for i := 0; i < n; i++ {
				fmt.Printf("Subject %d: ", i+1)
				fmt.Scan(&s.Subjects[i])
			}

			fmt.Print("Enter Guide Name: ")
			fmt.Scan(&s.Guide)

			m.AddStudent(s)

			fmt.Println("Student Added Successfully.")

		case 3:

			var prn string

			fmt.Print("Enter PRN: ")
			fmt.Scan(&prn)

			if m.RemoveStudent(prn) {
				fmt.Println("Student Removed Successfully.")
			} else {
				fmt.Println("Student not found.")
			}

		case 4:

			var prn string
			var sub string

			fmt.Print("Enter PRN: ")
			fmt.Scan(&prn)

			fmt.Print("Enter Subject: ")
			fmt.Scan(&sub)

			if m.FindAndUpdateSubject(prn, sub) {
				fmt.Println("Subject Updated Successfully.")
			} else {
				fmt.Println("Student not found.")
			}

		case 5:

			var prn string
			var year int

			fmt.Print("Enter PRN: ")
			fmt.Scan(&prn)

			fmt.Print("Enter New Year: ")
			fmt.Scan(&year)

			if m.FindAndUpdateYear(prn, year) {
				fmt.Println("Year Updated Successfully.")
			} else {
				fmt.Println("Student not found.")
			}

		case 6:
			m.ListStudents()

		case 7:
			fmt.Println("Total Students:", m.CountStudents())

		case 8:
			return

		default:
			fmt.Println("Invalid Choice")
		}
	}
}