/* //for integer values
package main
import "fmt"

//enumerated types
type OrderStatus int
const(
	Received OrderStatus = iota
	confirmed
	Prepared
	Deleivered
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("changing order status to", status)
}
func main(){
	changeOrderStatus(Received)
}
*/

// for string values
package main

import "fmt"

// enumerated types
type OrderStatus string

const (
	Received   OrderStatus = "received"
	confirmed  OrderStatus = "confirmed"
	Prepared   OrderStatus = "prepared"
	Deleivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}
func main() {
	changeOrderStatus(Prepared)
}

//output: changing order status to prepared
