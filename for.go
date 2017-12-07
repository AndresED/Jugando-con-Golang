package main
import (
	"fmt"
)



func main(){

	peliculas:=[] string{
		"Pelicula 1",
		"Pelicula 2",
		"Pelicula 3"}
	for i := 0; i < len(peliculas); i++ {
		fmt.Println(peliculas[i])
	}
}

