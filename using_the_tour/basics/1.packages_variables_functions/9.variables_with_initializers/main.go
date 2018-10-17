package main

import "fmt"

// khai báo có sử dụng thêm kiểu dữ liệu, nhưng không cần thiết
var i, j int = 1, 2

func main() {
	//khai báo đã được loại bỏ kiểu dữ liệu
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
