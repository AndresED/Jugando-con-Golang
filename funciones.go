package main
import (
	"fmt"
)



func main(){
	holaMundo();
	fmt.Println(suma(4.5,2.1))
	fmt.Println(resta(4.5,2.1))
	fmt.Println(multiplicacion(4.5,2.1))
	fmt.Println(division(4.5,2.1))
}


func holaMundo(){
	fmt.Println("Hola mundo");
}

func suma(n1 float32,n2 float32) float32 {
	return (n1+n2);
}

func resta(n1 float32,n2 float32) float32 {
	return (n1-n2);
}
func multiplicacion(n1 float32,n2 float32) float32{
	return (n1*n2);
}
func division(n1 float32,n2 float32) float32 {
	return (n1/n2);
}