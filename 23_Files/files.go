/*
package main
import (
	"fmt"
	"os"
)
func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		//log the error
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		//log the error
		panic(err)
	}
	fmt.Println("file name:", fileInfo.Name())
}
/* in terminal first cd 23_Files
go run files.go
output: file name: example.txt
*/

/*
//Check is file or folder
package main
import (
	"fmt"
	"os"
)
func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		//log the error
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		//log the error
		panic(err)
	}
	fmt.Println("file name:", fileInfo.Name())
	fmt.Println("file or folder:", fileInfo.IsDir())
	fmt.Println("file size:", fileInfo.Size())
	fmt.Println("file permission:", fileInfo.Mode())
	fmt.Println("file modified at:", fileInfo.ModTime())

}
/*ouput:
file name: example.txt
file or folder: false
file size: 13
file permission: -rw-r--r--
*/

/*
//How to read file
package main

import "os"
func main (){
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 12)
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(buf); i++ {
		println("data", d, string(buf[i]))
	}

}
/*output:

data 12 h
data 12 e
data 12 l
data 12 l
data 12 o
data 12
data 12 G
data 12 o
data 12 l
data 12 a
data 12 n
data 12 g
*/

/*
// Simple method to read data
package main

import (
	"fmt"
	"os"
)
func main (){

	f, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println((string(f)))
}//output: hello Golang
//we can read small file if we read large then show error to loading
*/

/*
// how to read large files
package main

import (
	"fmt"
	"os"
)
func main(){
	//read folder
	dir, err := os.Open(".")
	if err != nil{
		panic(err)
	}

	defer dir.Close()
	fileInfo, err := dir.ReadDir(2)

	for _, fi := range fileInfo{
		fmt.Println(fi.Name())
	}

}/*output:
files.go
example.txt
*/

/*
// how to read large files with ../ back one folder
package main

import (
	"fmt"
	"os"
)
func main(){
	//read folder
	dir, err := os.Open("../") //back one folder
	if err != nil{
		panic(err)
	}

	defer dir.Close()
	fileInfo, err := dir.ReadDir(2)

	for _, fi := range fileInfo{
		fmt.Println(fi.Name(), fi.IsDir())
	}

}/*output: it give all folder
.DS_Store
6_if_else
*/

/*
//how to write files
package main

import (

	"os"
)
func main(){

	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//f.WriteString("hi I am Ruchi")

	// Or using byte we can also write
	bytes := []byte("Hello Golang")
	f.Write(bytes)

}
*/


/*
// read and write to another file(streaming fashsion)
package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)

	}
	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)

	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	Writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}			
			break
		}
		e := Writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}
	Writer.Flush()
	fmt.Println("written to new file successfully")
}
//output:written to new file successfully
*/


//how to delete files
package main

import (

	"fmt"
	"os"
)
func main(){
	
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file deleted successfully")
}
//ouput: file deleted successfully



