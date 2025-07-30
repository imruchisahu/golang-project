/*
package main
import "fmt"
//numered sequence of specific length
func main() {
	//zeroed values
	//int -> 0 , string= "", bool=false

	var nums [4]int

	nums[0] = 1
	fmt.Println(nums[0]) // output: 1
	fmt.Println(nums) //output: [1 0 0 0]
	//array length
	//fmt.Println(len(nums)) // o/t: 4

	//var vals [4]bool
	//fmt.Println(vals) // output: [false false false false]


	var vals [4]bool
	vals[2] = true
	fmt.Println(vals) // output: [false false true false]

	//in string case
	var name [3]string
	name[0] = "golang"
	fmt.Println(name) // output: [golang  ]
}
*/


//how to declare arrays element when add
package main
import "fmt"
func main(){
	//to declare it in single line
	//nums := [3]int{1,2,3}
	//fmt.Println(nums) // output: [1 2 3]

	//2d arrays
	nums := [2][2]int{{3, 4},{5, 6}}
	fmt.Println(nums) //output: [[3 4] [5 6]]
}


// -fixed size, that is predictable
// - Memory optimazation
// - Contant time access
