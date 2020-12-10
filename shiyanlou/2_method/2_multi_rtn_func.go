package main

import "fmt"

// Go 内建_多返回值_ 支持
// 例如用来同时返回一个函数的结果和错误信息。
func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 如果你仅仅想返回值的一部分的话，你可以使用空白定义符 `_`
	_, c := vals()
	fmt.Println(c)

}

// `(int, int)` 在这个函数中标志着这个函数返回 2 个 `int`。无参多返回
func vals() (int, int) {
	return 3, 7
}
