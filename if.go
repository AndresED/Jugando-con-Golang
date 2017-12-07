package main
import (
	"fmt"
	"os"
	"strconv"
)



func main(){
	edad,err:= strconv.Atoi(os.Args[1])
	fmt.Println(err)
	if edad>=18 {
		fmt.Println("Soy mayor de edad")
	}else{
		fmt.Println("Soy menor de edad")
	}

}

