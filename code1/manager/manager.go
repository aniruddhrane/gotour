package manager

import "CODE1/student"

type StudentManager struct {
	students []student.Student
}

func (m *StudentManager) AddStudent() {
	var s student.Student

	student.CreateStudent(&s)

	m.students = append(m.students, s)
}

func (m *StudentManager) RemoveStudent(prn string) bool {
	for i, s := range m.students {
		if s.PRN == prn {
			m.students = append(m.students[:i], m.students[i+1:]...)
			return true
		}
	}

	return false
}

func (m *StudentManager) FindAndUpdateSubject(prn string, sub string) bool {
	for i := range m.students {
		if m.students[i].PRN == prn {
			m.students[i].UpdateSubjects(sub)
			return true
		}
	}

	return false
}

func (m *StudentManager) FindAndUpdateYear(prn string, year int) bool {
	for i := range m.students {
		if m.students[i].PRN == prn {
			m.students[i].UpdateYear(year)
			return true
		}
	}

	return false
}

func (m *StudentManager) FindAndPrint(prn string) bool {
	for i := range m.students {
		if m.students[i].PRN == prn {
			m.students[i].Print()
			return true
		}
	}

	return false
}
func(m*StudentManager) List(){
	for _,s:=range m.students{
		s.Print()
	}
}
func(m*StudentManager) Length(){
	println(len(m.students))
}

// package manager
// import "CODE1/student"
// type StudentManager struct{
// 	students []student.Student
// }
// func (m* StudentManager) AddStudent(){
//    var s student.Student
//    student.createStudnent(&s)
//    m.students=append(m.students,s)
// }

// func(m* StudentManager) RemoveStudent(prn string)bool{
// 	for i,s:=range m.students{
// 		if s.PRN==prn{
// 			m.students=append(m.students[:i],m.students[i+1:]...)
// 			return true
// 		}
// 	}
// 	return false
// }
// func(m* StudentManager) FindandupdateSubject(prn stirng,sub string){
// 	for i,s =range m.students{
// 		if s.PRN=prn{
// 		 s.UpdateSubjects(sub)  
// 		}
// 	}
// }

// func(m* StudentManager) FindandupdateYear(prn stirng,year int){
// 	for i,s =range m.students{
// 		if s.PRN=prn{
// 		 s.UpdateYear(year)  
// 		}
// 	}
// }

// func(m* StudentManager) FindandPrint(prn stirng){
// 	for i,s =range m.students{
// 		if s.PRN=prn{
// 		  s.Print()
// 		}
// 	}
// }