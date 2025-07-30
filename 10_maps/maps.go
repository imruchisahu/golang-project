package main
import "fmt"
import "maps"
func main(){
/*
	//creating map
	m := make(map[string]string)
	//setting an element
	m["name"] = "golang"
	m["area"] = "backend"
	//get an element
	fmt.Println(m["name"], m["area"]) //output: golang backend
*/

/*
	//Imp: if key does not exits in the map then it returns zero value
	m := make(map[string]int)
	m["age"] = 30
	m["price"] = 40
	fmt.Println(m["phone"]) // output: 0
	fmt.Println(m["age"]) // output: 30
	fmt.Println(len(m)) //output: 2

	//delete element
	//delete(m, "price")
	//fmt.Println(m) //output: map[age:30]

	//clear function: for empty list
	clear(m)
	fmt.Println(m) //output: map[]
*/

/*
//without using make function make map
	m := map[string]int{"price": 40, "phones": 3}
	fmt.Println(m) // output: map[phones:3 price:40]
*/

/* 
//get an element if element exit in list
	m := map[string]int{"price": 40, "phones": 3}
	_, ok := m["price"]
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}
	//output: all ok
*/

/*
//if element doesnot exit
	m := map[string]int{"price": 40, "phones": 3}
	v, ok := m["phone"]
	fmt.Println(v) // output: 0 it returns the value 0 because phone is not declare
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}
	//output: not ok

*/


//check equality of the map
	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 4}
	fmt.Println(maps.Equal(m1, m2)) //output: false
}
