package main

import "fmt"

func main() {
	foo01()
	foo02()
	foo03()
	foo04()

	foo := []int{1, 2, 3, 4, 5}
	calc(1, 2, 3)
	calc(foo...)
	calc(1, 2, 3, 4)
}

func foo01() {
	fmt.Println("foo01")
	var foo [3]int
	var bar []int
	test := []int{}

	if bar == nil {
		fmt.Printf("%v\n", foo)
		fmt.Printf("%p\n", &foo)
		fmt.Printf("%p\n", bar)
		fmt.Printf("%p\n", test)
	}
	fmt.Println("---")
}

func foo02() {
	fmt.Println("foo02")
	x := []int{1, 2, 3} // x是slice
	y := x[1:]          // y = [2, 3]
	x = append(x, 4)
	y[0] = 100 // y = [100, 3]
	fmt.Println(x)
	fmt.Println(y)
	y = append(y, 1000) // y = [100, 3, 1000]
	y[0] = 99           // y = [99, 3, 1000]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("---")
}

func foo03() {
	fmt.Println("foo03")
	x := [3]int{1, 2, 3}
	//x = append(x, 4) // x是array，所以這裡會有error
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println(x)
	fmt.Println("---")
}

func foo04() {
	fmt.Println("foo04")
	x := []int{1, 2, 3} // x是slice

	func(arr []int) {
		arr[0] = 7
		x = append(x, 4)
		fmt.Println(x)
	}(x)
	fmt.Println(x)
	fmt.Println("---")
}

func calc(vals ...int) {
	fmt.Println("calc")
	foo := 1
	if len(vals) > 0 {
		foo = vals[0]
	}
	fmt.Println(foo)
	fmt.Println("array count:", len(vals))
	for index, val := range vals {
		fmt.Println(index, ":", val)
	}
	fmt.Println("---")
}
