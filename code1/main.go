package main

import (
	"CODE1/manager"
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
		fmt.Println("4. Add subject")
		fmt.Println("5. Update Year")
		fmt.Println("6. List all students")
		fmt.Println("7.Count all students")
		fmt.Println("7. Exit")

		var choice int

		fmt.Print("Enter choice: ")
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
			m.AddStudent()
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
			fmt.Print("Enter the prn of whose subject is to be added")
			fmt.Scan(&prn)
			fmt.Print("Enter the subject ")
			fmt.Scan(&sub)
			if m.FindAndUpdateSubject(prn, sub) {
				fmt.Println("Subject updated successfully")
			} else {
				fmt.Println("Student not found")
			}

		case 5:
			var prn string
			var year int
			fmt.Print("Enter the prn of whose here needed to be updated")
			fmt.Scan(&prn)
			fmt.Print("Enter the year")
			fmt.Scan(&year)
			if m.FindAndUpdateYear(prn, year) {
				fmt.Println("Year updated successfully")
			} else {
				fmt.Println("Student not found")
			}
		case 6:
			m.List()
		case 7:
			m.Length()
		case 8:
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}

// package main
// import "fmt"

// func main(){
// 	for{
// 	fmt.Println(`
// 	 This is a student system
// 	 Enter:
// 	 1.Print student detail by prn
// 	 2.Add student
// 	 3.Delete student
// 	`)
// 	var choice int
// 	fmt.Scan(&chocie)

// 	switch choice{
// 	case 1:
// 		fmt.Println("Enter the prn")
// 		prn string
// 		scan(&prn)
// 		manager.FindandPrint(prn)
// 	case 2:
// 		manager.AddStudent()
// 	case 3:
// 		fmt.Println("Enter the prn")
// 		prn string
// 		scan(&prn)
// 		manager.RemoveStudent(prn)
// 	case 4:
// 		return;
// 	}
//  }

// }
