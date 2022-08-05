package main

import "fmt"

func main() {
	// fmt.Println("Quang")

	// Array :=[...]int {1, 2, 3}
	// var sum int = 0
	// for i := 0; i < len(Array); i++ {
	// 	sum += Array[i]
	// }
	// fmt.Println(sum)

	//arrayString :=[3]string {"1", "2", "3"}
	
	var arrayString string = "1234"
	var arrayNumber[4] int
	//arrayNumber := [...]int {0, 1, 2}
	var sum int = 0
	for i:= 0; i < len(arrayString); i++ {
		arrayNumber[i] = (int)(arrayString[i] - '0')
	}

	for i:= 0; i < len(arrayString); i++ {
		sum += arrayNumber[i]
	}
	fmt.Println(sum)

}