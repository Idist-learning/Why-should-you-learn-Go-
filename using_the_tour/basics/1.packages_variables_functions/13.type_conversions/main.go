package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	//var z uint = uint(f)

	// ví dụ:

	var f float64 = math.Sqrt((x*x + y*y))
	var z uint = (f)

	//
	//báo lỗi
	//# command-line-arguments
	//./main.go:15:33: cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt
	//./main.go:16:6: cannot use f (type float64) as type uint in assignment

	fmt.Println(x, y, z)
}
