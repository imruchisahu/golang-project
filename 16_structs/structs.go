/*
A struct in Go is a composite data type that groups together variables (called fields) under a single name. It is used to create custom data types that represent real-world entities with multiple properties.
*/

/*
package main

import (
	"fmt"
	"time"
)

//order struct
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time //nanosecond precision
}
func main() {
	myOrder := order{
		id: "1",
		amount: 50.00,
		status: "received",
	}
	myOrder.createdAt = time.Now()
	fmt.Println(myOrder.status)
	fmt.Println("Order struct", myOrder)
}
/*Output:
Received
Order struct {1 50 received {13986461455606107608 110642 0x55af4c0}}
*/

/*
package main

import (
	"fmt"
	"time"
)

//order struct
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time //nanosecond precision
}
func main() {

	myOrder := order{
		id: "1",
		amount: 50.00,
		status: "received",
	}
	myOrder.createdAt = time.Now()
	fmt.Println(myOrder.status)


	myOrder2 := order{
		id: "2",
		amount: 100,
		status: "delivered",
		createdAt: time.Now(),
	}
	myOrder.status ="paid"
	fmt.Println("Order struct", myOrder2)
	fmt.Println("Order struct", myOrder)
}
/*ouput:
received
Order struct {2 100 delivered {13986462409579232144 157657 0xee594c0}}
Order struct {1 50 paid {13986462409579201144 127465 0xee594c0}}
*/

/*
package main

import (
	"fmt"
	"time"
)

//order struct
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time //nanosecond precision
}


//receiver type
func (o *order) changeStatus(status string){
	o.status = status
}

func (o order) getAmount() float32{
	return o.amount
}
func main() {

	// if you don't set any field, default value is zero value
	//int -> 0, float=>,string="", bool=false
	myOrder := order{
		id: "1",
		amount: 50.00,
		status: "received",
	}

	myOrder.changeStatus(("confirmed"))
	fmt.Println(myOrder.getAmount())
}
//output: 50
*/

/*
package main

import (
	"fmt"
	"time"
)

//order struct
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time //nanosecond precision
}

func newOrder(id string, amount float32, status string) *order {
	//initial setup goes here...
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
	}
	return &myOrder
}
//receiver type
func (o *order) changeStatus(status string){
	o.status = status
}

func (o order) getAmount() float32{
	return o.amount
}
func main() {
	myOrder := newOrder("1", 30.50, "received")
	fmt.Println(myOrder.amount)
	// if you don't set any field, default value is zero value
	//int -> 0, float=>,string="", bool=false
// 	myOrder := order{
// 		id: "1",
// 		amount: 50.00,
// 		status: "received",
// 	}

// 	myOrder.changeStatus(("confirmed"))
// 	fmt.Println(myOrder.getAmount())
// }
}

*/

/*
package main
import "fmt"
func main(){
	language := struct {
		name string
		isGood bool
	}{"golang", true}
	fmt.Println(language)
}
//ouput: {golang true}
*/

// embeded struct
package main

import (
	"fmt"
	"time"
)

//order struct

type customer struct {
	name string
	phone string
}
//coposition
type order struct {
	id string
	amount float32
	status string
	customer
	createdAt time.Time //nanosecond precision
}
func main() {
	//newCustomer := customer{
		//name: "john",
		//phone: "1234567890",
	//}
	newOrder := order{
		id: "1",
		amount: 30,
		status: "received",
		customer: customer{
			name: "john",
			phone: "1234567890",
		},
	}
	newOrder.customer.name = "robin"
	fmt.Println(newOrder.customer)
}
//ouput: {robin 1234567890}