package main

import "fmt"

func main() {
	i := 20

	i++

	//++i 	//这种也是错误的    使语音更加明确
	//fmt.Println("i:", i++) //这是错误的，不允许和其他代码放在一起，必须单独一行
	fmt.Println("i:", i)

}
