package main

import "fmt"
func main(){
	//var name string ="golang"
	//fmt.Println(name)
	//infer
	//var name="golang"
	var isAdult bool = true
	fmt.Println(isAdult)
	var age int=30
	fmt.Println(age)


	//shorthand syntax 
	name := "golang" //:= this is use for when same time it is decleared
	fmt.Println(name)
	
	var name1 string // here is need to write type as string because not given input
	name1="golang"
	fmt.Println(name1)	
    
	//we assign 3 types var values
	//var price float32 = 50.5 
	//var price =50.5
	price := 50.5
	fmt.Println(price)
}
/* in terminal: go run 3_variables/variables.go
output:
true
30
golang
golang
50.5
*/




