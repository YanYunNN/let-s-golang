package main

import (
	"fmt"
	"math"
)

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}
	measure(s)
	measure(c)
}

// 这里是一个几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// 让 `square` 和 `circle` 实现这个接口
type square struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。这里我们让 `square` 实现了 `geometry` 接口。
func (s square) area() float64 {
	return s.width * s.height
}

// 可以为值类型或者指针类型的接收器定义方法。这里是一个值类型接收器的例子。
func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

// `circle` 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量的是接口类型，那么我们可以调用这个被命名的接口中的方法。
// 这里有一个一通用的 `measure` 函数，利用这个特性，它可以用在任何 `geometry` 上。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
