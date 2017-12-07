/* Leer dos valores a,b e intercambiarlos
por ejemplo a=5 y b=2 , debe de imprimir a=2 b=5
*/
package main
import (
	"fmt"
)



func main(){
	a:=2
	b:=3
	auxiliar:= a
	a=b
	b=auxiliar
	fmt.Println(a)
	fmt.Println(b)
}