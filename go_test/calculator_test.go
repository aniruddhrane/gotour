package calculator

import(
	"testing"
	
)

func TestAdd(t *testing.T){
    sum:=add(2,3)
	if sum!=5{
		t.Errorf("expected 5,got %d",sum)
	}
}

func TestSubtract(t *testing.T){
    sub:=subtract(3,4)
	if sub!=-1{
		t.Errorf("expected -1 got %d",sub)
	}
}
func TestMultiply(t *testing.T){
	mul:=multiply(1,2)
    if mul!=2{
		t.Errorf("expected 2 got %d",mul)
	}
}
func TestDivision(t *testing.T){
	result:=division(3,2)

    if result!=1.5{
		t.Errorf("expected 1.5,got %f",result)
	}
}