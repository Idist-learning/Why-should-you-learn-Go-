package main

import "fmt"

// vị trí các hàm không quan trọng. quan trọng là khi hàm thực thi nó chỉ thực hiện hàm main. các hàm khác đều được load trước main.

func main() {
	fmt.Println(add(42, 13))
}

func add(x, y int) int {
	return x + y
}
