package main

import "fmt"
//Array - Cau truc du lieu mang
//1. Khai niem
//2. Demo
func main() {
	//Khai bao array
	var myArray [4]int 
	fmt.Println(myArray)

	//gan gia tri cho array
	myArray[0] = 123
	fmt.Println(myArray[0])

	//khai bao 1 array co khoi tao gia tri
	arrays1 := [3]int {1, 2, 3}
	fmt.Println(arrays1)

	arrays2 := [3]int {1}
	fmt.Println(arrays2)

	//size mang
	fmt.Println(len(arrays1))

	//khai bao mang ko can set size
	arrays3 := [...]string {"a", "b", "c"}
	fmt.Println(arrays3)

	//array ka value type khong phai la refernce type
	
	//c1
	for i:= 0; i < len(arrays1) ; i++ {
		fmt.Println(arrays1[i])
	}
	//c2
	for index, value := range arrays1 {
		fmt.Printf(" i = %d - value = %d", index, value)
		fmt.Println()
	}

	//mang hai chieu 
	matrix := [4][2]int {
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
	}
	fmt.Println(matrix)
	for i := 0;  i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}