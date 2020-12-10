package main

//可变参数函数 可以用任意数量的参数调用
//例如，`fmt.Println` 是一个常见的变参函数。
import "fmt"

// Go 内建_多返回值_ 支持
// 例如用来同时返回一个函数的结果和错误信息。
func main() {
	// 变参函数使用常规的调用方式，除了参数比较特殊。
	sum(1, 2)
	sum(1, 2, 3)

	// 如果你的 slice 已经有了多个值，想把它们作为变参使用，调用 `func(slice...)`
	nums := []int{1, 2, 3, 4}
	sum(nums...)

}

// `(int, int)` 在这个函数中标志着这个函数返回 2 个 `int`。无参多返回
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
