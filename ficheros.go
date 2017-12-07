package main
import (
	"fmt"
	"io/ioutil"
	"os"
)



func main(){
	fichero,errorDeFichero:= ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)
	fmt.Println(fichero)
	fmt.Println("--------------")
	nuevo_texto:=[]byte(os.Args[1])
	fmt.Println(nuevo_texto)
	//escribirioutil
	
}

func showError(e error){
	if (e!=nil){
		panic(e)
	}
}
