package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// hàm rand.Seed để thay đổi giá trị của rand mỗi khi thực hiện run chương trình
	rand.Seed(time.Now().UTC().UnixNano())
	// hàm trên gây ảnh hưởng xuống hàm rand.Intn. Nếu không có hàm trên thì rand.Intn chỉ trả về một giá trị duy nhất cho mọi lần thực thi
	fmt.Println("My favorite number is", rand.Intn(10))
}
