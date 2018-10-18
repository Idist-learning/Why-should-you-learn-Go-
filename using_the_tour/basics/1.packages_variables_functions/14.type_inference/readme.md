Khi khai báo một biến mà không chỉ định kiểu dữ liệu rõ ràng (hoặc bằng cách sử dụng cú pháp := hoặc var =), kiểu của biến sẽ được lấy ra từ giá trị ở phía bên phải.

Khi giá trị của phần bên phải của biến được đưa vào có kiểu là gì, biến mới sẽ có cùng kiểu với nó:

    var i int
    j := i // j is an int

Nhưng khi phía bên phải của nó chức một hằng số không có kiểu dữ liệu, biến mới có thể là int, float64 hoặc complex128 phụ thuộc vào độ chính xác của hằng số này:

    i := 42           // int
    f := 3.142        // float64
    g := 0.867 + 0.5i // complex128

thử thay đổi giá trị ban đầu của v trong ví dụ này và theo dõi xem kiểu dữ liệu của nó thay đổi như thế nào.