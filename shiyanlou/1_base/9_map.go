package main

// _map_ 是 Go 内置[关联数据类型]，在一些其他的语言中称为_哈希_ 或者_字典_

import "fmt"

func main() {
	// 要创建一个空 map，需要使用内建的 `make`:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// 使用 `name[key]` 来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	// 找不到会返回零值
	v2 := m["k3"]
	fmt.Println("v2:", v2)

	// 当对一个 map 调用内建的 `len` 时，返回的是键值对数目
	fmt.Println("len:", len(m))

	// 内建的 `delete` 可以从一个 map 中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。
	//这可以用来消除键不存在和键有零值， 像 `0` 或者 `""` 而产生的歧义。
	//TODO
	_, get := m["k2"]
	fmt.Println("get", get)

	// 你也可以通过这个语法在同一行申明和初始化一个新的map。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
