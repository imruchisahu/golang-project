/*
// in terminal :
// go mod init mypp
// create package
package main

import (
	"fmt"

	"github.com/codersgyan/podcast/auth"
	"github.com/codersgyan/podcast/user"
)
func main(){
	auth.LoginWithCredentials("codersgyan", "secret")

	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		Name: "John Doe",
	}
	fmt.Println(user.Email, user.Name)


}
//in terminal first cd 24_packages 
// go run packages.go
/*output:
login user using codersgyan secret
session loggedin
user@email.com John Doe
*/


// in terminal :
// go mod init mypp
// create package
package main

import (
	"fmt"

	"github.com/codersgyan/podcast/auth"
	"github.com/codersgyan/podcast/user"
	"github.com/fatih/color"
)
func main(){
	auth.LoginWithCredentials("codersgyan", "secret")

	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		//Name: "John Doe",
	}
	//fmt.Println(user.Email, user.Name)
	color.Red(user.Email)

}
/*ouput:
ogin user using codersgyan secret
session loggedin
user@email.com //this change color only red color
*/
//fixed all the mod 
//in terminal run: go mod tidy
// run code 
//go run 