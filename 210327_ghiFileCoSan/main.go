package main
 
// thêm vào 2 module chúng ta cần
import (
    "fmt"
    "io/ioutil"
    "os"
)

func main(){

	//ghi File moi

	DataFile := []byte("The data")

	error := ioutil.WriteFile("File1.data", DataFile, 0777)

	if error != nil {
		fmt.Println(error)
	}

	//doc FIle vua moi tao

	data, error1 := ioutil.ReadFile("File1.data")

	if error1 != nil {
		fmt.Println(error1)
	}

	fmt.Println(string(data))

	//ghi thong tin vao File co san 

	f, err := os
}