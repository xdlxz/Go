package main

import "fmt"

/**
函数也是一种类型，可以保存为变量
func()  函数类型
*/

func func_type() {
	fmt.Println("func type")
}

func main() {
	var f func()
	f = func_type
	f()
}
