package main

import "fmt"

/**
if else switch
 */

func conditionUse(score int) string {
	res := "D"
	switch  {
		case score < 60:
			res = "F"
		case score < 70:
			res = "E"
		case score < 80:
		res = "D"
	}
	return res
}

func main() {

	// if else
	//filename := "a.txt"
	//content, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(string(content))
	//}
	// switch
	//fmt.Println(conditionUse(5))
	//var b int = 1000 >> 2
	//fmt.Println(b)
	// 定义数组，然后遍历
	//a := [4]int{1, 2, 3, 4,}
	//for k, v := range a {
	//	fmt.Printf("%v, %v", k, v)
	//	fmt.Println()
	//}

	// - 数组切片 slice
	// 创建一个数组，然后通过切片访问
	arr := [6]int{1, 2, 3, 4, 5, 6,}
	a := arr[0:4]
	b := arr[3:]
	c := arr[:]
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	// make 创建切片
	d := make([]int, 10)

	e := make([]int, 10, 15)
	fmt.Printf("%v, %d, %d\n", d, len(d), cap(d))
	fmt.Printf("%v, %d, %d\n", e, len(e), cap(e))

	// ... 方式创建切片
	f := [...]int{1, 2, 3, 4, 5, 6,}
	fmt.Printf("%v, %d, %d\n", f, len(f), cap(f))
	g := f[:]
	g = append(g, 7)
	fmt.Printf("%v, %d, %d\n", g, len(g), cap(g))
	//  output [1 2 3 4 5 6 7], 7, 12  长度增加1 容量扩充了一倍


	// map
	ma := map[int]string{1: "java", 2:"php"}
	fmt.Println(ma[1])

	mb := make(map[int]string)
	mb[1] = "scala"
	mb[2] = "php"
	mb[3] = "java"

	for k, v := range mb {
		fmt.Println("key=", k, "value=", v)
	}


	//
	mc := make(map[string]int)
	mc["php"] = 3
	mc["java"] = 1
	mc["scala"] = 2

	for k, v := range mc {
		fmt.Println("key=", k, "value=", v)
	}
}
