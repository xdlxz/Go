package main

import "fmt"

/*
	go没有默认参数值
	实参通过值传递
	如果实参包括引用类型，如指针、slice(切片)、map、function、channel 等类型，
	实参可能会由于函数的间接引用被修改
*/

//测试值传递
type Data struct {
	complex  []int
	instance InnerData
	ptr      *InnerData
}
type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	fmt.Printf("inFunc value:%+v\n", inFunc)
	fmt.Printf("inFunc ptr:%p\n", &inFunc)
	return inFunc
}
func main() {
	in := Data{
		complex: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{6},
	}

	fmt.Printf("main value:%+v\n", in)
	fmt.Printf("main ptr:%p\n", &in)

	out := passByValue(in)
	fmt.Printf("out value:%+v\n", out)
	fmt.Printf("out ptr:%p\n", &out)

}
