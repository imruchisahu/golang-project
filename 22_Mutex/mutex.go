/* // mutex is used for deadlock 
package main

import "fmt"
type post struct{
	views int
}

func (p *post) inc() {
	p.views += 1
}
func main(){
	myPost := post{views: 0}
	myPost.inc()
	myPost.inc()
	fmt.Println(myPost.views)
}//output: 2
*/

/*
// use mutex
package main
import (
	"fmt"
	"sync"
)
type post struct{
	views int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()

	p.mu.Lock()
	p.views += 1
	p.mu.Unlock()
}
func main(){
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg) //for concurerently runnig we use go 
	}
	wg.Wait()
	fmt.Println(myPost.views)
}
//output: randomely
//100
//99
*/

/*
// If print multiple times
package main
import (
	"fmt"
	"sync"
)
type post struct{
	views int
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.views += 1
}
func main(){
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg) //for concurerently runnig we use go 
	}
	wg.Wait()
	fmt.Println(myPost.views)
}
//output: 100
//If you run multiple time it give 100 output only not in randam format
*/

// If print multiple times
package main
import (
	"fmt"
	"sync"
)
type post struct{
	views int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}() //put lock and unlock inside defer

	p.mu.Lock()
	p.views += 1
	
}
func main(){
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg) //for concurerently runnig we use go 
	}
	wg.Wait()
	fmt.Println(myPost.views)
}//output: 100
