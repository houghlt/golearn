package data

import "fmt"
// 定义一个结构实现 Database
type MyMysql struct {
	Dsn string
	QueryStr string
	Arr map[string]string
}


// 实现接口的方法  注意提供包外使用时保证首字母大写
func (m MyMysql) Connect(s string) bool {
	if s != ""{
		return true
	}
	return false
}


func (m MyMysql) Query(s string) map[string]string {
	m.Arr["name"] = s
	m.Arr["age"] = "27"
	m.Arr["sex"] = "male"
	return m.Arr
}

func (m MyMysql) Close() {
	fmt.Println("mysql have close the connect!")
}

func main() {
	
}
