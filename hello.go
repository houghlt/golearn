package main

import "fmt"

// 变量
var name  = "hough"
var age  = 27
var sex = "man"

func variable() {
	var a, b = 2, 4
	fmt.Println(a, b)
}

// 常量
const filename = "a.txt"
// 枚举类型 go 没有枚举类型，通过 const 实现

func mains() {
	//fmt.Println("hello")
	//variable()
	//fmt.Println(name, sex, age)
	//const age int = 27
	//fmt.Println(filename, age)
	const (
		a = 1
		b = 2
		c = 3
	)
	// iota 自增赋值 _ 下划线跳过值赋值
	const(
		java = iota
		_ // 跳过 2 的赋值
		python
		php
		javascript
	)
	fmt.Println(a, b, c)
	fmt.Println(java, python, php, javascript)
}
