/*
package main
import "fmt"
//iterating over data structure
func main(){
	nums := []int {6, 7, 8}

/*
	for i:=0; i<len(nums); i++{
		fmt.Println(nums[i])
	}
	output: 6
	7
	8 
*/

/*
//or 
for _, num := range nums {
	fmt.Println(num)
}
/*output:
6
7
8
}*/



/*
//sum of element
package main
import "fmt"
//iterating over data structure
func main(){
	nums := []int {6, 7, 8}
	sum := 0
	for _, num := range nums{ //_ is a index
		sum = sum + num
	}
	fmt.Println(sum) //output:21
}
*/

/*
//without using _ in this place use i index
package main
import "fmt"

func main(){
	nums := []int {6, 7, 8}

	for i, num := range nums{ 
		fmt.Println(num, i)
	}
}
/*output:
6 0
7 1
8 2
*/

/*
//using map
package main
import "fmt"
//iterating over data structure
func main(){
	m := map[string]string{"fname": "john", "lname": "doe"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
}
}
/*output:
name john
lname doe
fname
lname
*/

//Unicode code point rune
//starting byte of rune
//255 -> 1 byte , 2 byte
package main
import "fmt"
//iterating over data structure
func main(){
	//i is starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
/*output:

0 103
0 g
1 111
1 o
2 108
2 l
3 97
3 a
4 110
4 n
5 103
5 g
*/