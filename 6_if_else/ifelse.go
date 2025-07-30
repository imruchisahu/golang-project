/*
package main
import "fmt"
func main(){
	age := 16

	if age >= 18{
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}
}*/
/*output:
Person is not an adult
*/

/*
//if else and else if
package main
import "fmt"
func main(){
	age := 16

	if age >= 18{
		fmt.Println("Person is an adult")
	} else if age >=12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is kid")
	}
}//output: Person is teeenager
*/


/*
//logical opeator && or ||
package main
import "fmt"
func main(){
	var role = "admin"
	var hasPermissions = true //false
	if role == "admin" || hasPermissions {
		fmt.Println("yes")
	}
}//output: yes
//output if false permissons than output: nothing 
*/


//we can declare a variable 
package main 
import "fmt"
func main(){
	if age := 15; age >= 18{
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager", age)
	}
}
//output: Person is teenager 15



// go does not have ternary, you will have to use normal if else