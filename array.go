package main
import (
	"fmt"
)



func main(){
	//ARRAYS
	var peliculas[4] string
	peliculas[0]="Pelicula 1"
	peliculas[1]="Pelicula 2"
	peliculas[2]="Pelicula 3"
	/*for i := 0; i < len(peliculas); i++ {
		fmt.Println(peliculas[i])
		
	}*/

	nombres:=[] string{
		"Andres",
		"Estuardo"}
	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])		
	}

	//MATRICES
	var cursos[2][2] string
	cursos[0][0]="Algoritmos"
	cursos[0][1]="Matematicas"
	cursos[1][0]="Ingles"
	cursos[1][1]="Estructura de datos"

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(cursos[i][j])		}
	}
	fmt.Println("--------------------------")	
	//USO DE APPEND
	nombres=append(nombres,"Antonio")
	nombres=append(nombres,"Maria")
	nombres=append(nombres,"Antonio")
	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])		
	}

}

