/*
package main
import "fmt"
func add(a int, b int) int{
	return  a+b
}

func main(){
	result := add(3,5)
	fmt.Println(result)
}//output: 8
*/

/*
package main
import "fmt"
func add(a, b int) int{ // (a, b int) b is before all element is integer
	return  a+b
}

func main(){
	result := add(3,5)
	fmt.Println(result)
}//output: 8
*/

/*
package main
import "fmt"
func getLanguages() (string, string, string){
	return "golang", "javascript", "c"
}

func main(){
	fmt.Println(getLanguages())
}
//output: golang javascript c
*/

/*
package main
import "fmt"
func getLanguages() (string, string, string){
	return "golang", "javascript", "c"
}

func main(){
	lang1, lang2, _ := getLanguages() //here _ is assign value but not used.
	fmt.Println(lang1, lang2)
}//output: golang javascript c
*/


package main
//func processIt(fn func(a int) int) {
	//fn(1)
//}

func processIt() func(a int) int {
	return func (a int) int{ 
		return 4
	}
}
func main(){
	//fn := func(a int) int{
		//return 2
	//}
	fn := processIt()
	fn(6)
}