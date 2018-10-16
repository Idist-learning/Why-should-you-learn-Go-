Trên Go, một tên biến được đọc ra nếu nó bắt đầu bằng chữ viết hoa. Ví dụ, Pizza là một tên có thể đọc cũng như Pi, nó được đọc ra từ package math.

pizza và pi không bắt đầu bằng chữ viết hoa nên không thể lấy giá trị của nó ra được

Khi thêm một package, bạn chỉ có thể xem các tên được xuất ra của nó. mọi tên biến " không được xuất ra" không thể truy cập từ bên ngoài package.

Chạy đoạn code, sẽ thấy thông báo lỗi.

Để sửa lỗi này, đổi math.pi thành math.Pi và thử lại.