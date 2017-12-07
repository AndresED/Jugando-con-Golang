package main
import (
	"fmt"
)



func main(){
	fmt.Println(getDatos(2,3))
}

//UN CLOSURE ES UNA FUNCION QUE  EJECUTA INTERNAMENTE UNA FUNCION ANONIMA
func getDatos(num1 int,num2 int) int{
	//FUNCION ANONIMA
	dato:=func() int{
		return num1+num2
	}
	return dato()
}
