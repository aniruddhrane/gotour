package manager

import "CODE1/student"

type StudentManager struct {
	students []student.Student
}

func (m *StudentManager) AddStudent(s student.Student) {
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

func (m *StudentManager) ListStudents() {
	if len(m.students) == 0 {
		println("No students found.")
		return
	}

	for _, s := range m.students {
		s.Print()
	}
}

func (m *StudentManager) CountStudents() int {
	return len(m.students)
}