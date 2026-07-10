package student

import "fmt"

type Student struct {
	PRN       string
	Name      string
	Branch    string
	Year      int
	Division  string
	RollNo    int
	CGPA      float64
	Subjects  []string
	Guide     string
}

func (s *Student) UpdateYear(year int) {
	s.Year = year
}

func (s *Student) UpdateSubjects(sub string) {
	s.Subjects = append(s.Subjects, sub)
}

func (s *Student) Print() {
	fmt.Println("PRN      :", s.PRN)
	fmt.Println("Name     :", s.Name)
	fmt.Println("Branch   :", s.Branch)
	fmt.Println("Year     :", s.Year)
	fmt.Println("Division :", s.Division)
	fmt.Println("Roll No  :", s.RollNo)
	fmt.Println("CGPA     :", s.CGPA)
	fmt.Println("Subjects :", s.Subjects)
	fmt.Println("Guide    :", s.Guide)
}