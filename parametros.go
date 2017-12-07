package main
import (
	"fmt"
)



func main(){
	personas("Estuardo","Andres","Antonio","Alcides","Maria","Elena")
}

func personas(nombres ...string){
	for _,nombre :=range nombres{
		fmt.Println(nombre);
	}
}