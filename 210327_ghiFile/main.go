package main
 
import (
    "fmt"
    "io/ioutil"
)
 
func main() {

    var NameFile string = "Quang.txt"
    fmt.Println(NameFile)
 
    //Ghi du lieu vao NewFlie
 
    mydata := []byte("du lieu da duoc ghi vao")
 
    // Phương thức WriteFile trả về lỗi nêú không thành công
    error := ioutil.WriteFile(NameFile, mydata, 0)
    // xử lý lỗi này
    if error != nil {
        // in ra thông tin lỗi
        fmt.Println(error)
    }

    //Doc du lieu tu File vua luu
 
    data, err := ioutil.ReadFile("NameFile.txt")
    if err != nil {
        fmt.Println(err)
    }
 
    fmt.Print(string(data))
 
}