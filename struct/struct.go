package main

import "fmt"

type User struct {
	name string
	age int
}

func main() {


	ma := make(map[int]User)
	sunhuafu := User{
		name: "hough",
		age:  27,
	}
	ma[1] = sunhuafu

	sunhuafu.name = "xiaoguo"
	// 修改值之后还要整体替换
	ma[1] = sunhuafu

	fmt.Printf("%v\n", ma)

}
