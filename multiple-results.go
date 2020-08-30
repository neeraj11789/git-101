package main 

import "fmt"

func swap(a, b string) (string, string){
	return b, a 
}

func main(){
	x, y := swap("a string", "b string")
	fmt.Println("String a " + x)
	fmt.Println("String b " + y)
}