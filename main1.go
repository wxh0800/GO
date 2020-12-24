// package main

// import "fmt"

// example 1
// func main()  {
// 	fmt.Println("Hello World!")
// }
// func main() {
// 	// var age int      // 变量声明
// 	// var age int = 29 // 声明变量并初始化
// 	var age = 29 // 可以推断类型
// 	fmt.Println("my age is", age)
// 	age = 29 // 赋值
// 	fmt.Println("my age is", age)
// 	age = 54 // 赋值
// 	fmt.Println("my new age is", age)

// 	var width, height int = 100, 50 // 声明多个变量

// 	fmt.Println("width is", width, "height is", height)

// 	var (
// 		name    = "naveen"
// 		age1    = 29
// 		height1 int
// 	)
// 	fmt.Println("my name is", name, ", age is", age1, "and height is", height1)

// 	name, age2 := "naveen", 29 // 简短声明

// 	fmt.Println("my name is", name, "age is", age2)
// }

// func rectProps(length, width float64) (float64, float64) {
// 	var area = length * width
// 	var perimeter = (length + width) * 2
// 	return area, perimeter
// }

// func main() {
// 	area, perimeter := rectProps(10.8, 5.6)
// 	fmt.Printf("Area %f Perimeter %f", area, perimeter)
// }
package main

import (
	"fmt"
)

// func main() {
// 	pls := [][]string{
// 		{"C", "C++"},
// 		{"JavaScript"},
// 		{"Go", "Rust"},
// 	}
// 	for _, v1 := range pls {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%s ", v2)
// 		}
// 		fmt.Printf("\n")
// 	}
// }

// func find(num int, nums ...int) {
// 	fmt.Printf("type of nums is %T\n", nums)
// 	found := false
// 	for i, v := range nums {
// 		if v == num {
// 			fmt.Println(num, "found at index", i, "in", nums)
// 			found = true
// 		}
// 	}
// 	if !found {
// 		fmt.Println(num, "not found in ", nums)
// 	}
// 	fmt.Printf("\n")
// }
// func main() {
// 	find(89, 89, 90, 95)
// 	find(45, 56, 67, 45, 90, 109)
// 	find(78, 38, 56, 98)
// 	find(87)
// }

// func find(num int, nums ...int) {
// 	fmt.Printf("type of nums is %T\n", nums)
// 	found := false
// 	for i, v := range nums {
// 		if v == num {
// 			fmt.Println(num, "found at index", i, "in", nums)
// 			found = true
// 		}
// 	}
// 	if !found {
// 		fmt.Println(num, "not found in ", nums)
// 	}
// 	fmt.Printf("\n")
// }
// func main() {
// 	nums := []int{89, 90, 95}
// 	find(89, nums...)
// }

func main() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)

}
