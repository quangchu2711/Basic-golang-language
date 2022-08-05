package main

import (
    "fmt"
)


/*Note: Range cho phep lap lai toan bo cac phan tu o trong mang hoac slice.
		Dung de duyet qua mot slice hoac map
		gia tri dau tien la chi so cua phan tu
		gia tri thu hai la ban sao cua phan tu do*/ 

//Dung ca hai gia tri
func example1() {
	var pow = []int {1, 2, 4, 8}
	for i, v := range pow {
		fmt.Printf("2 ^ %d = %d\n", i, v)
	}
}

//Su dung mot trong hai gia tri tra ve dung dau gach duoi thay the
func example2() {
	var pow = []int {1, 2, 4, 6, 8}
	for _, v := range pow {
		fmt.Printf("v = %d\n", v)
	}
}

/*Note: Map la tap hop cac phan tu duoc luu tru duoi dang key-value. key trong
		map co kieu du lieu so sanh duoc va khong bi trung lap. De tao mao ta
		dung ham make() voi cong thuc sau: make(map[type of key] type of value)*/
func example3() {
	languages := make(map[string]float32)

	//Them phan tu
	languages["go"] = 0.63
	languages["java"] = 1.03

	fmt.Printf("languages go: %f\n", languages["go"])

	//Cap nhat phan tu
	languages["go"] = 0.23
	fmt.Println(languages)

	//xoa phan tu
	delete(languages, "go")
	fmt.Println(languages)
}
 
func main() {
	fmt.Println("Quang")
	//example1()
	example3()
}