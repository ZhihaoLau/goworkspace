package main

import "fmt"

var i, j int = 1, 2

func main() {
	// 如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
	var c, Python, Java = true, false, "no!"
	fmt.Println(i, j, c, Python, Java)
}