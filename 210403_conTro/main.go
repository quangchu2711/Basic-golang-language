package main

 import "fmt"

/*Note: nill: Gia tri 0 cho con tro*/

/*Khai bao nhung khong gan dia chi*/
 func example1() {
 	var p *int

 	fmt.Println("p =  ", p)
 }
 
  /*Khai bao va gan dia chi*/
 func example2() {
 	var a = 5.67
 	var p = &a
 
 	fmt.Println("Value stored in variable a = ", a)
 	fmt.Println("Address of variable a = ", &a)
 	fmt.Println("Value stored in variable p = ", p)
}

/*Thay doi gia tri duoc luu tru trong bien thong qua con tro*/
func example3() {
	var a = 100
	var p = &a

	fmt.Println("a = ", a)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	*p = 2000
	fmt.Println("a (after) = ", a)
}

/*a = 7.98
  address of a = 0x01
  address of p = 0x03
  p = 0x01
  pp = 0x03

  *pp = 0x01
  **pp = 0x03 */


/*Con tro tro toi con tro*/
func example4() {
	var a = 7.98
	var p = &a
	var pp = &p

	fmt.Println("a = ", a)
	fmt.Println("address of a = ", &a)

	fmt.Println("p = ", p)
	fmt.Println("address of p = ", &p)

	fmt.Println("pp = ", pp)

	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)
}
 func main() {
 	//example1()
 	//example2()
 	//example3()
 	example4()
 }