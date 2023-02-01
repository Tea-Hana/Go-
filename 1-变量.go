package main

import "fmt"

func main() {
	//变量定义：var
	//常量定义： const

	//01-先定义变量，在赋值 var 变量 数据类型

	var name string

	name = "duke"

	var age int

	age = 20

	fmt.Println("name", name)

	fmt.Printf("name is:%s, %d\n", name, age)

	//02-定义时直接赋值

	var gender = "man"

	fmt.Println("gender", gender)

	//03-定义直接赋值，使用直接推导（最常用的）
	address := "上海"

	fmt.Println("address:", address)

	//04-平行赋值
	i, j := 10, 20 //同时定义两个变量
	fmt.Println("变换前=> i:", i, "j:", j)

	i, j = j, i

	fmt.Println("变换后=> i:", i, "j:", j)
}
