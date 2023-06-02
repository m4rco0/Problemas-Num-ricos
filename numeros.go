// fazendo exiba todos o numeros divizivel por 3 de 1 até 100
package main

import "fmt"

func brincar(i int){
	if i % 3 == 0 && i % 5 == 0{
		fmt.Println("PinPan")
	}else if (i % 3 == 0){
		fmt.Println("Pin")
	} else if (i % 5 == 0){
		fmt.Println("Pan")
	}
}
func main(){
	for i := 1; i <= 100; i++{
		brincar(i)
	}
}