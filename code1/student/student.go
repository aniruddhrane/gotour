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

func CreateStudent(s *Student) {
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
	fmt.Print("Number of subjects: ")
	fmt.Scan(&n)

	s.Subjects = make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Subject %d: ", i+1)
		fmt.Scan(&s.Subjects[i])
	}

	fmt.Print("Enter Guide Name: ")
	fmt.Scan(&s.Guide)
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




// package student
// //this will contain the struct with fields and some methods on it 
// import "fmt"

// type Student struct{
// 	PRN string
// 	Name string
// 	Branch string
// 	Year   int
// 	Division string
// 	RollNo   int
// 	CGPA     float64
// 	Subjects  []string
// 	Guide     string
// }

// func CreateStudnent(s *Student) {
//  fmt.Print("Enter PRN: ")
// fmt.Scan(&s.PRN)

// fmt.Print("Enter Name: ")
// fmt.Scan(&s.Name)

// fmt.Print("Enter Branch: ")
// fmt.Scan(&s.Branch)

// fmt.Print("Enter Year: ")
// fmt.Scan(&s.Year)

// fmt.Print("Enter Division: ")
// fmt.Scan(&s.Division)

// fmt.Print("Enter Roll Number: ")
// fmt.Scan(&s.RollNo)

// fmt.Print("Enter CGPA: ")
// fmt.Scan(&s.CGPA)
// var n int

// fmt.Print("Number of subjects: ")
// fmt.Scan(&n)

// s.Subjects = make([]string, n)

// for i := 0; i < n; i++ {
//     fmt.Printf("Subject %d: ", i+1)
//     fmt.Scan(&s.Subjects[i])
// }
// fmt.Print("Enter Guide Name: ")
// fmt.Scan(&s.Guide)
// }

// func (s *Student)UpdateYear(year int){
//    s.Year=year
// }
// func (s *Student)Print(){
//   fmt.Println(s.PRN); 
//   fmt.Println(s.Name);    
//   fmt.Println(s.Branch);
//   fmt.Println(s.Year);
//   fmt.Println(s.Division);
//   fmt.Println(s.RollNo);
//   fmt.Println(s.CGPA);
//   fmt.Println(s.Subjects);
//   fmt.Println(s.Guide);
// }
// func(s *Student)UpdateSubjects(sub string){
//     s.Subjects=append(s.Subjects,sub)
// }