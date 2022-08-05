package main

import "fmt"

import "time"

func main(){

	for {
		fmt.Println("Quang")
		time.Sleep(1 * time.Second)
	}
}

func Sum(number int) (int) {
	sum := 0
	for i:= 0; i < number; i++ {
		sum += i
	}

	return sum
}

