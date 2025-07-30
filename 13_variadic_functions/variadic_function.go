/* 
A variadic function in Go is a function that accepts zero or more arguments of the same type. 
It allows you to pass a variable number of values to a single parameter. 
*/
/*
package main
import "fmt"
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main(){
	result := sum(3, 4, 5, 6)
	fmt.Println(result)
}
//output: 18
*/

package main
import "fmt"
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main(){
	nums := []int{3, 4, 5, 6}
	result := sum(nums...) // after nums put... (3 dots)
	fmt.Println(result)
}//output: 18

