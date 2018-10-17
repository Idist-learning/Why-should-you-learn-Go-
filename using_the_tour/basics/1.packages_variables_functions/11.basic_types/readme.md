Các kiểu dữ liệu cơ bản của Go:

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

Ví dụ trên biểu diễn các biến với nhiều kiểu dữ liệu, khi khai báo biến thì tên biến và kiểu biến được khai báo cùng nhau, kèm theo giá trị khởi tạo

các kiểu dữ liệu như `int`, `uint` và `uintptr` thường rộng khoảng 32bit trên các hệ thống 32bit và rộng 64bit trên hệ thống 64bit. Khi bạn cần một giá trị nguyên, ít nhất thì bạn nên sử dụng kiểu int, khi bạn có lý do đặc biệt để sử dụng như kiểu dữ liệu số nguyên với thuộc tính size hay unsigned