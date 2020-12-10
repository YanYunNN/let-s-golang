package main

import "fmt"

//Go 支持通过闭包来使用匿名函数
// 匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
func main() {
	// 斐波那契数列
	fmt.Println(fact(7))
}

// `face` 函数在到达 `face(0)` 前一直调用自身。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
