package main
import (
	"fmt"
)

type Persona struct{
	nombreyapellidos string
	edad int
	profesion string

}

func main(){
	var andres Persona = Persona{
		nombreyapellidos:"Estuardo Andrés Esquivel Díaz",
		edad:25,
		profesion:"Programador"}
	fmt.Println(andres)
	fmt.Println(andres.nombreyapellidos)
	fmt.Println(andres.edad)
	fmt.Println(andres.profesion)
}