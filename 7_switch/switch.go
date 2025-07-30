/*
//simple switch
package main

import "fmt"
func main(){
	
	
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3: 
		fmt.Println("three")
	default :
		fmt.Println("other")
	
	}
*/

/*
// multiple condtion switch
package main
import "fmt"
import "time"
func main(){
	
	
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("it's workday")
	}

	//output: it is weekend
}
*/


//type switch
package main
import "fmt"
func main(){

	//type switch
	whoAmI := func(i interface{}) { // interface accepts multiple inputs
		switch t := i.(type){
		case int:
			fmt.Println("its an Interger")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI("golang")
}
//output: its a string


