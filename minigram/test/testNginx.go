package test

import (
	"net/http"
	"strconv"
	"testing"
)

func testFunc() int {
	for i := 0; i < 10; i++ {
		num := strconv.Itoa(i)
		if i%2 == 0 {
			i++
			println("apple:", num, "i++ =", i)
		}
		println("fruit:", i, 3+i)
	}
	return -1
}

func testNginx() {
	http.Handle("/", http.FileServer(http.Dir("D:/")))
	http.ListenAndServe(":8088", nil)
}

func testHello() {
	println("hello world!")
	fruit := int16(32)
	apples := fruit + 1
	oranges := fruit + 1
	fruit = apples + oranges
	println(fruit)
}

func Hello() string { return "Hello, world." }

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
