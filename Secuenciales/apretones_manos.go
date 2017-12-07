/* 
A una reunion asistieron a n personas ,
¿Cuantos apretones de mano se dieron?

DEFINIENDO VARIABLES
n = numero de personas
a = numero de apretones
Si tenemos n personas entonces 1 persona debera de saludar a n-1 personas ,
entonces habran n*(n-1) y como en cada saludo intervienen 2 personas entonces
a = (n*(n-1))/2
*/
package main
import (
	"fmt"
)



func main(){
	var n int=20
	var a int= (n*(n-1))/2
	fmt.Println(a)
}