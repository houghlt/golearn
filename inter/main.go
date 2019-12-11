package main

import (
	"fmt"
	data "main/inter/mysql"
)

// 定义一个数据库接口
type Database interface {
	Connect(string) bool
	Query(string) map[string]string
	Close()
}

// 定义一个操作该接口的方法
func Connect(d Database) bool {
	return d.Connect("mysql conning...")
}

func main() {

	//var d Database
	amap := map[string]string{"name":"hough", "age":"27"}
	//fmt.Println(amap)
	//d := data.MyMysql{"localhost:3306", "select * from fa_order", amap}
	d := data.MyMysql{
		Dsn: "localhost:3306",
		QueryStr: "select * from fa_order",
		Arr: amap,
	}
	fmt.Println(Connect(d))
}
