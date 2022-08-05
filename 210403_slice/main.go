package main

import "fmt"

/*Note: Slice khong he chua bat cu gia tri nao cua chung ca ma chi la
        tham chieu den chinh nhung mang da ton tai thoi

        Cau truc: khi khoi tao mot silce, thi cau truc du lieu nay cung 
        mot mang duoc tao ra. Gia tri cua slice la gia tri cua cau truc 
        du lieu.

        Thanh phan: 
        - Con tro (pointer): Con tro nay tham chieu den mot mang
        - Kich thuoc (length): So luong phan tu trong slice
        - dung tich (Capacity): So luong phan tu trong mang tinh tu
          vi tri bat dau cua slice*/

/*================================================Khai bao slice================================================*/
//C1: Tao slice tu mang cho truoc _ sliceName := arrayName[start:end]
func example1() {
	languages := [6]string{"Golang", "C/C++", "Java", "Python", "PHP", "Vue"}
	sliceLangs := languages[0:4]
	fmt.Println(len(sliceLangs))
	fmt.Println(cap(sliceLangs))
	fmt.Println(sliceLangs)
}

func example2() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	s0 := a[:]
	s1 := s0[:]
	s2 := s1[1:3]
	s3 := s2[2:6]

	fmt.Println(a)
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(len(s2), cap(s2), s2)
	fmt.Println(len(s3), cap(s3), s3)	 
}

//C2: Khai bao tat _ name := []type {valueOfType1, valueOfType1...}
func example3() {
	a := []int {1, 2, 3, 4, 5, 6}
	a1 := a[1:4]
	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(a1), cap(a1), a1)	 
}

//C3: Su dung tu khoa Make _ name := make([]Type, lenghth, capacity)
//Note: Co the luoc bo di capacity neu length = capacity
func example4() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println("Slice ban dau")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s1 := make([]int, 5, 10)
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	//fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s1), cap(s2))		
}

/*================================================Cac phep toan voi slice ================================================*/
//Duyet su dung for...range
func example5() {
	langs := []int {1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("langs = %v, len: %d, cap: %d",langs, len(langs), cap(langs))
	fmt.Println()
	for i, lang := range langs {
		fmt.Printf("index = %d, value = %d\n", i, lang)
	}
}

//Cap nhat cac phan tu trong silde
/*Note: Khac voi mang, slice la kieu tham chieu den mot mang co ban. 
		Viec sua doi cac phan tu trong slice se sua doi cac phan tu 
		trong mang ma slice tham chieu den*/
func example6() {
	name := [6]string {"A", "B", "C", "D", "E", "F"}
	name1 := name[1:]
	name2 := name[3:]

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println()
	name[1] = "X"
	fmt.Println(name1)
	fmt.Println(name2)
}

//Them mot phan tu vao slice chung ta khoa append
func example7() {
	number :=[]int {1, 2, 3, 4, 5, 6}
	addNumber := append(number, 7, 8, 9, 10)
	fmt.Println(addNumber)
}

func main() {
	//fmt.Println("Quang")
	//example1()
	//example2()
	//example3()
	//example4()
	//example5()
	//example6()
	example7()
}