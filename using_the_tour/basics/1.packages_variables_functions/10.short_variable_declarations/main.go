package main

import "fmt"

//// Khai báo không hợp lệ
//z:=1

func main() {
	var i, j int = 1, 2
	// Khai báo hợp lệ
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
