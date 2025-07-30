//Channel is use between transfer the data between goroutine.
//channel is use between two goroutines for communication.
/*
package main

import "fmt"

func main(){
	messageChan := make(chan string)
	messageChan <- "ping" //blocking
	msg := <-messageChan
	fmt.Println(msg)
}
/*output:
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/Shield/Desktop - Rohit's Mac/Vs code/Golang/21_channels/channels.go:9 +0x36
exit status 2
*/

/*
package main

import "time"
import "fmt"
func processNum(numChan chan int){
	fmt.Println("processing number", <-numChan)
}


func main(){
	numChan := make(chan int)
	go processNum(numChan)
	numChan <- 5
	time.Sleep(time.Second * 2)

} //output: processing number 5
*/

/*
package main

import (
	"fmt"
	"math/rand"
	"time"
)
//send data
func processNum(numChan chan int){
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
}

func main(){
	numChan := make(chan int)
	go processNum(numChan)
	for {
		numChan <- rand.Intn(100)
	}
}
//output: infinite output.....
*/

/*
package main
import (
	"fmt"

)
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func main(){
	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result // send and receiveing , blocking
	fmt.Println(res)
	}
//output: 9
*/

/*
package main
import (
	"fmt"

)
//goroutine synchronyser
func task(done chan bool){
	//use defer
	defer func() { done <- true}()

	fmt.Println("processing...")
}

func main(){

	done := make(chan bool)
	go task(done)
	<- done //block
}//output: processing...
*/

/*
// How to implement queue goroutine
package main

import (
	"fmt"
	"time"
)

//implement queue
func emailSender(emailChan chan string, done chan bool){
	defer func() { done <- true}()
	for email := range emailChan{
		fmt.Println("Sending email to", email)
	time.Sleep(time.Second)
	}
}
func main(){
	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i:=0; i < 5; i++ {
		emailChan <-fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("done sending.")
	//this is important to close channel
	close(emailChan)
	<-done //block
	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

}
/* output:

done sending.
Sending email to 0@gmail.com
Sending email to 1@gmail.com
Sending email to 2@gmail.com
Sending email to 3@gmail.com
Sending email to 4@gmail.com
*/

/*
//Implement multiple channels
package main

import (
	"fmt"

)


func main(){
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func(){
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <- chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <- chan2:
			fmt.Println("received data from chan2", chan2Val)

		}
	}
}
/* output:
received data from chan2 pong
received data from chan1 10
*/

package main

import (
	"fmt"
	"time"
)

//implement queue
func emailSender(emailChan <-chan string, done chan bool){
	defer func() { done <- true}()

	<-done
	for email := range emailChan{
		fmt.Println("Sending email to", email)
	time.Sleep(time.Second)
	}
}

func main(){
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func(){
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <- chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <- chan2:
			fmt.Println("received data from chan2", chan2Val)

		}
	}
}
