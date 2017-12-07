package main
import (
	"fmt"
)
func main(){
	//GO ES UN LENGUAJDE DE PROGRAMACION TIPADO
	var suma int=3+5; //TIPO DE DATO INT
	var resta int=8-6;//TIPO DE DATO  INT
	var nombre string="Andres Esquivel" //TIPO DE DATO STRING
	saludo:="Hola me llamo " //TIPO DE DATO STRING
	pais:="Perú" //TIPO DE DATO STRING
	edad:=25 //TIPO DE DATO INT
	edad2:="25" //TIPO DE DATO STRING
	var programador bool=true //TIPO DE DATO BOOLEAN
	var igv float32=0.09 //TIPO DE DATO FLOAT

	//DECLARANDO UNA CONSTANTE
	const year int =2017
	//year = 2018 => no se puede modificar porque es una constante
	fmt.Println (suma)
	fmt.Println (resta)
	fmt.Println(suma%2)
	fmt.Println(year)
	fmt.Println(edad)
	fmt.Println(programador)
	fmt.Println(igv)
	fmt.Println (saludo+nombre+" soy de "+pais+" y tengo "+edad2+" años")
}