package main

import "fmt"

// 相同类型的参数，可以简写成：
func add(x, y int) int {
	return x + y
}

// 可同时有两个返回值
func swap(x, y string) (string, string) {
	return y, x
}

// "Naked" return: 如果 return 右边不填，可直接返回参数当前的值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}


func main () {
	fmt.Println(add(55, 23))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(11))
}