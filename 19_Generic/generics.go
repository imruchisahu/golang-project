/*
package main

import "fmt"

func printSlice(items [] int){
	for _, item := range items{
		fmt.Println(item)
	}
}

func main(){
	printSlice([]int{1, 2, 3}) // you have interger slice
}
/*output:
1
2
3
*/

/*
//Now if you have string slice
package main
import "fmt"
func printStringSlice(items []string){
	for _, item := range items{
		fmt.Println(item)
	}
}

func main(){
	names := []string{"golang", "typescript"}
	printStringSlice(names)
}
/*output:
golang
typescript
*/

/*
//If we pass any type of function then use any
package main
import "fmt"
func printSlice[T any](items []T){
	for _, item := range items{
		fmt.Println(item)
	}
}

func main(){
	names := []string{"golang", "typescript"}
	printSlice(names)
}
/*output:
golang
typescript
*/

/*
//If we use in place of any to interface
package main
import "fmt"
func printSlice[T interface{}](items []T){
	for _, item := range items{
		fmt.Println(item)
	}
}

func main(){
	names := []string{"golang", "typescript"}
	printSlice(names)
}
/*output:
golang
typescript
*/

/*
//if we use | 
package main
import "fmt"
func printSlice[T int | string | bool](items []T){
	for _, item := range items{
		fmt.Println(item)
	}
}

func main(){
	//names := []string{"golang", "typescript"}
	vals := []bool{true, false, true}
	printSlice(vals)
}
/*output:
true
false
true
*/

/*
//How to apply this in struct 
package main
import "fmt"
//LIFO
type stack struct{
	elements []int // here in only accept int not string for this both use we use generics
}

func main(){
	myStack := stack{
		elements: []int{1, 2, 3},
	}
	fmt.Println(myStack)
}
//output: {[1 2 3]}
*/

/*
//for int and string both used we use generic
package main
import "fmt"
//LIFO
type stack[T any] struct{
	elements []T // here in only accept int not string for this both use we use generics
}

func main(){
	myStack := stack[string]{
		elements: []string{"golang"},
	}
	fmt.Println(myStack)
}//output: {[golang]}
*/

/*
//use comparable keywords
package main
import "fmt"
//LIFO
func printSlice[T comparable](items []T){ //compare is an interface i.e implemented by all comparables types(bool, numbers, string, pointer, channel, etc)
	for _, item := range items{
		fmt.Println(item)
	}
}

func main(){
	vals := []bool{true, false, true}
	printSlice(vals)
}
/*output:
true
false
true
*/

/* //use stack
package main 
import "fmt"
//LIFO
type stack struct{
	elements []int
}


func main(){
	myStack := stack{
		elements: []int{1, 2, 3},
	}
	fmt.Println(myStack)
}//output: {[1 2 3]}
*/

//use T generic for int and string accept in stack
package main
import "fmt"
//LIFO
type stack[T any] struct{
	elements []T
}


func main(){
	myStack := stack[string]{
		elements: []string {"goland"},
	}
	fmt.Println(myStack)
}//output: {[golang]}
