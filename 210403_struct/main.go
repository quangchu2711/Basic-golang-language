package main

import (
	"fmt"
)

//Defining a struct type
type Person struct {
	FirstName string
	LastName string
	Age 	 int
}

//Khai bao p co kieu du lieu la person
func exmaple1() {

	/*1*/
	var p1 Person
	p1.LastName = "Chu"
	p1.FirstName = "Quang"
	p1.Age = 20
	fmt.Println("Information of the first person: ", p1)

	/*2*/
	p2 := Person{"A", "Nguyen Van", 22}
	fmt.Println("Information of the second person: ", p2)

	/*3*/
	p3 := Person{
		FirstName: "X",
		LastName:  "Ngo",
		Age:       23,
	}
	fmt.Println("Information of the thỉrd person: ", p3)
}

//Truy xuat den cac truong trong struct
func example2() {
	p := Person{"Chu", "Quang", 20}
	fmt.Println("LastName: ", p.LastName)
	fmt.Println("FirstName: ", p.FirstName)
}

//Xay dung mang struct 
func example3() {
	var infor[3] Person
	infor[0] = Person{"A", "Nguyen", 19}
	infor[1] = Person{"B", "Chu", 20}
	infor[2] = Person{"C", "Ngo", 21}

	for i:= 0; i < len(infor); i++ {
		fmt.Println("LastName: ", infor[i].FirstName)
		fmt.Println("FirstName: ", infor[i].LastName)
		fmt.Println("Age: ", infor[i].Age)
		fmt.Println()
	}
}

//Struct trong struct
type employee struct {
	Person
	job string
}

func example4() {
	var infor[3] employee
	infor[0] = employee{Person{"A", "Nguyen", 19}, "X"}
	infor[1] = employee{Person{"B", "Chu", 20}, "Y"}
	infor[2] = employee{Person{"C", "Ngo", 21}, "Z"}

	for i:= 0; i < len(infor); i++ {
		fmt.Println("LastName: ", infor[i].FirstName)
		fmt.Println("FirstName: ", infor[i].LastName)
		fmt.Println("Age: ", infor[i].Age)
		fmt.Println("Job: ", infor[i].job)
		fmt.Println()
	}
}

func main() {
	//exmaple1()
	//example2()
	//example3()
	example4()
}
