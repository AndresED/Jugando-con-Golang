/* 
SUMA DE LOS N PRIMEROS CUADRADOS NATURALES
a = (n*(n+1)*(2*n+1))/6
*/
package main
import (
	"fmt"
)



func main(){
	var n int=20
	var a int= (n*(n+1)*(2*n+1))/6
	fmt.Println(a)
}