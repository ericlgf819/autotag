package main

import (
	"fmt"

	"erictest.com/test/fibonacci"
)

func main() {
	fmt.Print("请输入数字：")
	var i int
	fmt.Scanln(&i)

	fmt.Printf("hello 斐波那契数 %d = %d", i, fibonacci.Fibonacci(i))
}
