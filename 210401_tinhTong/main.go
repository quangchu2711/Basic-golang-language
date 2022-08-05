/*Code by Quang _ Date: 2021/04/01*/
/*Tính tổng: 1, Khai báo một chuỗi string chứa các số 
			 2, Xây dựng hàm con tính tổng chẵn và tính tổng lẻ trong chuỗi string
			 3, Ghi kết quả vào file*/
package main

import (
	"fmt"
	"io/ioutil"
)

/*Tinh tong chan*/
func sumEvenNUmber(array[9] int) int {
	var sum int = 0
	for i:= 0; i < len(array); i++ {
		if array[i] % 2 == 0 {
			sum += array[i]
		}
	}
	return sum
}

/*Tinh tong le*/
func sumOddNumber(array[9] int) int {
	var sum int = 0
	for i:= 0; i < len(array); i++ {
		if array[i] % 2 == 1 {
			sum += array[i]
		}
	}
	return sum
}

/*Dem so cac chu so*/
func numberOffDigits(number int) int {
	var num int = 0
	for number > 0 {
		number /= 10
		num += 1
	}
	return num
}

/*Chuyen doi so nguyen sang mot chuoi*/
func convertIntToString(number int) string{
   var s string = ""
   var sizeArray int = numberOffDigits(number)
   var arrayInt [10]int

   for i:= 0; i < sizeArray; i++ {
   		s += " "
   }

   i := sizeArray - 1
   for number > 0 {
   		arrayInt[i] = number % 10
   		number /= 10
   		i -= 1
   }

   runes := []rune(s)
   for i:= 0; i < sizeArray; i++ {
   		runes[i] = (rune)(arrayInt[i] + 48)
   		//fmt.Printf("%c ",runes[i])
   }
   s = string (runes)
   return s
}

func main(){
	/*========================Tinh tong========================================*/
	var arrayString string = "123456789"
	var arrayInt[9]int
	
	/*khong su dung ham con*/
	fmt.Println("== Don't Use function ==")
	var sum int = 0
	var sumEven int = 0
	var sumOdd int = 0

	for i:= 0; i < len(arrayString); i++ {
		arrayInt[i] = (int)(arrayString[i] - '0')
		sum += arrayInt[i]
		if arrayInt[i] % 2 == 0 {
			sumEven += arrayInt[i]
		}else {
			sumOdd += arrayInt[i]
		}
	}

	fmt.Printf("Sum = %d\n", sum)
	fmt.Printf("Sum even number = %d\n", sumEven)
	fmt.Printf("Sum odd number = %d\n",sumOdd)
	fmt.Println()

	/*Su dung ham con*/
	fmt.Println("== Use function ==")
	for i:= 0; i < len(arrayString); i++ {
		arrayInt[i] = (int)(arrayString[i] - '0')
	}

	var x int = sumEvenNUmber(arrayInt)
	var y int = sumOddNumber(arrayInt)
	fmt.Printf("Sum even number = %d\n", x)
	fmt.Printf("Sum odd number = %d\n", y)

	/*========================Ghi File========================================*/
	string1 := convertIntToString(sum)
	string2 := convertIntToString(sumEven)
	string3 := convertIntToString(sumOdd)
	var sumString string = "Sum: " + string1 + "\nSum even number: " + string2 + "\nSum odd number: " + string3
	
	var nameFile string = "fileOutPut.txt"
	outputFile := []byte(sumString)

 	error := ioutil.WriteFile(nameFile, outputFile, 0)

 	if error != nil {
 		fmt.Println(error)
 	}
}