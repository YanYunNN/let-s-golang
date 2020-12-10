package main

// _Slice_ 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口

import (
	"fmt"
)

func main() {
	// 不同于数组，slice 的类型仅有它所包含的元素决定（不像数组中还需要元素的个数）。
	//要创建一个长度非零的空slice，需要使用内建的方法 `make`。
	//这里我们创建了一 个长度为3的 `string` 类型 slice（初始化为零值）。
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 我们可以和数组一起设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// 作为基本操作的补充，slice 支持比数组更多的操作。
	// 其中一个是内建的 `append`，它返回一个包含了一个或者多个新值的 slice。
	// 注意我们接受返回由 append 返回的新的 slice 值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice 也可以被 `copy`。
	//这里我们创建一个空的和 `s` 有相同长度的 slice `c`，并且将 `s` 复制给 `c`。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	// Slice 支持通过 `slice[low:high]` 语法进行“切片”操作。
	// 左闭右开：序列2-4个
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这个 slice 从 `s[0]` 到（但是不包含）`s[5]`。
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个 slice 从（包含）`s[2]` 到 slice 的后一个值。
	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们可以在一行代码中申明并初始化一个 slice 变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不同，这和多位数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
