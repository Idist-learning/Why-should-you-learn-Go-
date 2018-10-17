Biểu thức T(v) chuyển giá trị v sang kiểu T.

Một vài kiểu ép kiểu:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)

Không giống như C, đối với Go thì việc ép kiểu giữa các kiểu phải rõ ràng. Hãy thử loại bỏ `float64` hay `uint` trên ví dụ về ép kiểu và xem điều gì sẽ xảy ra.