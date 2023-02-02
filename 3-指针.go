package main

import "fmt"

func main() {
	//go语义也有指针
	//结构体成员调用时：c语言  ptr -> name go语言：ptr.name

	//go语言在使用指针时，会使用内部的垃圾回收机制（gc：garbage collector）,开发人员不需要手动释放内存

	//c语言不允许返回栈上的指针，go语言可以返回栈上的指针，程序会在编译的时候就确定了变量的分配位置：
	//编译的时候，如果发现有必要的话，就将变量分配到堆上

	/**
	&(取址符)： 对变量取址
	& ：获取变量在计算机内存中的地址， &age，取出变量age所在内存地址，一般地址是十六进制。

	*(声明指针) ： 对指针取值
	申明指针*age， 打印指针内存单元的值 **age.如下面 x *int, *x就是指针所对应的值。
	*/
	name := "lily"
	ptr := &name

	fmt.Println("name", *ptr)
	fmt.Println("name ptr", ptr)

	//02使用new关键词定义
	name2Ptr := new(string)

	*name2Ptr = "Duke"

	fmt.Println("name2 ", *name2Ptr)
	fmt.Println("name2 ptr", name2Ptr)

	//可以返回栈上的指针，编译器在编译程序时,会自动判断这段代码，将city变量分配在堆上
	res := testPtr()
	fmt.Println("res city", *res, "address:", res)

	//if两端不用加()
	//即使有一行代码,也必须使用{}
	if res == nil {
		fmt.Println("res 是空 nil")
	} else {
		fmt.Println("res 是非空")
	}
}

// 定义一个函数，返回一个string类型的指针，go语言返回写在参数列表后面
func testPtr() *string {
	city := "深圳"
	ptr := &city
	return ptr
}
