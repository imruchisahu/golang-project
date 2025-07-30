/*
package main
import "fmt"

func task(id int){
	fmt.Println("doing task", id)
}

func main(){
for i := 0; i <= 5; i++ {
	task(i)
}
}

/*output:
doing task 0
doing task 1
doing task 2
doing task 3
doing task 4
doing task 5
*/

/*
// how to run task paralelly using go
package main
import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("doing task", id)
}

func main(){
	for i := 0; i <= 5; i++ {
		go task(i) // go for parallel task
	}
	time.Sleep(time.Second * 2)
}
/*output: // this is not in order. it is run in concurrently.
doing task 2
doing task 3
doing task 1
doing task 5
doing task 4
doing task 0
*/

/*
package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("doing task", id)
}

func main(){
	for i := 0; i <= 5; i++ {
		//go task(i) // go for parallel task
		go func() {   // run anonymously function
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 2)
}
/*output:
5
4
1
2
0
3
*/

/*
package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("doing task", id)
}

func main(){
	for i := 0; i <= 5; i++ {
		//go task(i) // go for parallel task
		go func(i int) {   // run anonymously function
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 2)
}
/*output:
5
4
1
2
0
3
*/

//use waitgroup :

package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup){
	defer w.Done() // defer means it execute after run the function defer
	fmt.Println("doing task", id)
}

func main(){
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go task(i, &wg) // go for parallel task
		
	}
	wg.Wait()
}
/*output:
doing task 0
doing task 5
doing task 3
doing task 1
doing task 4
doing task 2
*/