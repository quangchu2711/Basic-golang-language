package main
 
// thêm vào 2 module chúng ta cần
import (
    "fmt"
    "io/ioutil"
)
 
func main() {
    // Đọc nội dung trong Quang.data
    data, error := ioutil.ReadFile("Quang.data")
    // Nếu chương trình không thể đọc file
    // in ra nguyên nhân tại sao
    if error != nil {
        fmt.Println(error)
    }
 
    // Nếu đọc file thành công thì
    // in ra nội dung như một string
    fmt.Print(string(data))
 
}