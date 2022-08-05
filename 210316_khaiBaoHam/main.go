package main

import "fmt"

func Chao() {
	fmt.Println("Hello")
}

func sayHello(name string) {
	fmt.Println("hello", name)
}

func greeting(name string) string {
	//smt.Sprintf tra ve kieu string
	result := fmt.Sprintf("hello %s", name)
	return result
}

func returnInt(value int) int {
	x := 10
	x = x + value
	return x
}

func returnFloat(value float64) float64 {
	x := 10.1
	x = x + value
	return x
}

//Multiple return values
func multipleReturnValue(x, y int) (int, int, int) {
	z := x * y
	return x, y , z
}

//Named return values
func nameReturnValue(x, y int) (width int, height int, isSquare bool) {
	isSquare = x == y  //gan ket qua so sanh x voi y cho isSquare
	return x, y , isSquare
}


func main() {
	//Chao()
	//sayHello("Quang")
	// x := greeting("Quang")
	// fmt.Println(x)
	// y := returnInt(2)
	// fmt.Println(y)
	//z := returnFloat(2.2)
	//fmt.Println(z)
	X, Y, Z := nameReturnValue(100, 200)
	if Z {
		fmt.Println("TRUE")
	} else {
		fmt.Println("Width =", X)
		fmt.Println("Height =", Y)		
	}
}