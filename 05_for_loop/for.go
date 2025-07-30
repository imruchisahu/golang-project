package main

import "fmt"

// for -> only construct in go for looping
func main() {
	/*//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}*/
	/*output: go ru
	n 5_For_loop/for.go
	1
	2
	3*/

	//infinite loop
	//for{
	//println("1")
	//} // infinite output print for this cancel ctrl +c

	/*//Classic for loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}*/

	/*//break use:
	for i := 0; i <= 3; i++ {
		if i == 3 {
			break
		} // stop the iteration
		fmt.Println(i)
	}*/
	/*output:
	0
	1
	2
	*/

	/*//continue- Basically skip the current iteration
	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(n)
	}*/
	/*output:
	0
	1
	3
	*/
	//

	// use range keyword1.22 range
	for i := range 4{
		fmt.Println(i)
	}
}

