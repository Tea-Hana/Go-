package main

import "fmt"

func main() {

	//1.只有false才能代码逻辑假，数字0和nil不能(只有false/true才能表示逻辑真假，数字不能代表逻辑值)

	//if 0 {
	//	fmt.Println("hello")
	//}

	//if nil {
	//	fmt.Println("hello")
	//}

	if true {
		fmt.Println("hello")
	}
	//a,b := 1,2
	//x := a > b ? 0:-1

}
