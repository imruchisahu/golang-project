
/*package main
import "fmt"

//by value
func changeNum(num int){
	num = 5
	fmt.Println("In changeNum", num)
}

func main(){
	num := 1
	changeNum(num)
	fmt.Println("After changeNum in main", num)
}
//output: In changeNum 5
*/

package main
import "fmt"

//by reference
func changeNum(num *int){
	*num = 5
	fmt.Println("in changeNum", num)
}
func main(){
	num := 1
	changeNum(&num)
	fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main", num)
}
/*output:
in changeNum 0xc0000100c0
Memory address 0xc0000100c0
After changeNum in main 5
*/
