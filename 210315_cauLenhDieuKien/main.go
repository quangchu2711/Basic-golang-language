package main

import "fmt"
/*=======================if else=======================*/
//Note: Golang khong cho phep else xuong dong

// func main() {
// 	//if else 
// 	number := 10
// 	//if condition {//code here}
// 	if number == 10 {
// 		fmt.Println("number = 10")
// 	}
// 	if number < 10 {
// 		fmt.Println("number < 10")
// 	}else {
// 		fmt.Println("number = 10")
// 	}
// 	//if statement; condition{// code}
// 	if a := 100; a > 100 {
// 		fmt.Println("a > 100")
// 	}else{
// 		fmt.Println("a = 100")
// 	}
// }
/*=======================switch case=======================*/
/*Note: Khong can su dung break, golang tu truy cap toi case
		trong 1 case co the bat duoc nhieu gia tri*/
// func main(){
// 	number := 10
// 	switch number {
// 	case 1, 10, 300:
// 		fmt.Println("number = 1")
// 	case 2:
// 		fmt.Println("number = 2")
// 	case 3:
// 		fmt.Println("number = 3")
// 	case 4:
// 		fmt.Println("number = 4")
// 	case 5:
// 		fmt.Println("number = 5")
// 	default:
// 		fmt.Println("unknow")
// 	}
// }

// func main(){
// 	number := 10
// 	switch {
// 	case number > 10:
// 		fmt.Println("number > 10")
// 	case number == 10:
// 		fmt.Println("number = 10")
// 	}
//}
/*==================FallThrough, break, goto=======================*/
/*Note(FallThrough): Khi dung fallthrough thi cac case tiep theo van se
duoc thuc hien*/
// func main(){
// 	number := 10
// 	switch number {
// 	case 10: 
// 		fmt.Println("number = 1")
// 		break
// 	case 5: 
// 		fmt.Println("number = 10")
// 		fallthrough
// 	case 3: 
// 		fmt.Println("number = 3")
// 	}
// }

// func main(){
// 	number := 10
// 	switch number {
// 	case 10: 
// 		if number == 10 {
// 			goto handle
// 		}
// 		fmt.Println("number = 10")
// 		handle:
// 			fmt.Println("handle for case")
// 	case 2: 
// 		fmt.Println("number = 2")
// 		fallthrough
// 	case 3: 
// 		fmt.Println("number = 3")
// 	}
// }

/*==================For=======================*/

// func main(){
// 	loops
// 	for init; condition; post

// 	break; continue
// 	for i:= 0; i < 10; i++ {
// 		if i == 4 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// 	for i:= 0; i < 10; i++ {
// 		if i == 4 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}	
// }
/*==================while=======================*/
/*Note: Trong golang khong ho tro While*/

func main(){
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	for i, j:= 1, 2; i < 10 && j < 10; i, j = i + 1, j + 1 {
		fmt.Println(i)
		fmt.Println(j)		
	} 	
}


