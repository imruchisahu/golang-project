/*
package main

import "fmt"

type payment struct{}

func(p payment) makePayment(amount float32){
	razorpayPaymentGw := razorpay{}
	razorpayPaymentGw.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float32){
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

func main(){
	newPayment := payment{}
	newPayment.makePayment(100)
}//ouput: making payment using razorpay 100
*/

/*
package main
import "fmt"
type payment struct{
	
}

//open close principle
func(p payment) makePayment(amount float32){
	//razorpayPaymentGw := razorpay{}
	//razorpayPaymentGw.pay(amount)
	stripePayemntGw := stripe{}
	stripePayemntGw.pay(amount)
	
}

type razorpay struct {}

func (r razorpay) pay(amount float32){
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct {}
func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe", amount)
}

func main(){
	newPayment := payment{}
	newPayment.makePayment(100)
}//output: making payment using stripe 100
*/

/*
package main
import "fmt"

type paymenter interface {
	pay(amount float32)
}
type payment struct{
	gateway paymenter
}
//open close principle
func(p payment) makePayment(amount float32){
	p.gateway.pay(amount)
	
}

type razorpay struct {}
func (r razorpay) pay(amount float32){
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct {}
// func (s stripe) pay(amount float32){
// 	fmt.Println("making payment using stripe", amount)
// }

type fakepayment struct{}
func (f fakepayment) pay(amount float32){
	fmt.Println("making payment using fake gateway for testing purpose")
}
func main(){
	//stripePayemntGw := stripe{}
	//razorpayPaymentGw := razorpay{}
	fakeGw := fakepayment{}
	newPayment := payment{
		//gateway: stripePayemntGw,
		gateway: fakeGw,
	}
	newPayment.makePayment(100)
}//output: making payment using fake gateway for testing purpose
*/

package main
import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}
type payment struct{
	gateway paymenter
}
//open close principle
func(p payment) makePayment(amount float32){
	p.gateway.pay(amount)
	
}

type razorpay struct {}
func (r razorpay) pay(amount float32){
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct {}
// func (s stripe) pay(amount float32){
// 	fmt.Println("making payment using stripe", amount)
// }

type fakepayment struct{}
func (f fakepayment) pay(amount float32){
	fmt.Println("making payment using fake gateway for testing purpose")
}

type paypal struct {}
func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal", amount)
}
func (p paypal) refund(amount float32, account string){
	
}
func main(){
	//stripePayemntGw := stripe{}
	//razorpayPaymentGw := razorpay{}
	//fakeGw := fakepayment{}
	paypalGw := paypal {}
	newPayment := payment{
		//gateway: stripePayemntGw,
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}//output: making payment using paypal 100