package main

import "fmt"

//_函数_ 是 Go 的中心
func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}

func plus(a int, b int) int {
	// Go 需要明确的返回值，例如，它不会自动返回最后一个表达式的值
	return a + b
}
