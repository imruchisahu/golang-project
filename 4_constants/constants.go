// A constant in Golang is a fixed value that cannot be changed during program execution. 
//It is declared using the const keyword and must be assigned a value at the time of declaration. 

/*
// This code is for single constant values
package main
import "fmt"
const age =30

func main(){
	// :=
	//const name="golang"
	//const age=30
	fmt.Println(age)
}
//output : 30
*/

// this is for multiple constant values
package main
import "fmt"
func main(){
	const(
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}
//output: 5000 localhost