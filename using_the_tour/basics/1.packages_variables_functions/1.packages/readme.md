Mọi chương trình của Go đều được tạo từ các packpage.

Các chương trình đều chạy từ package main

Chương trình này sử dụng package được import từ đường dẫn "fmt" và "math/rand".

Theo quy ước, một tên của package giống với thành phần cuối cùng của đường dẫn. Ví dụ, package "math/rand" sẽ chứa các file mà được bắt đầu với package chứa thành phần rand.

Chú ý: môi trường mà các chương trình này thực thi đã được xác định vì thế mỗi lần bạn chạy các ví dụ của hàm rand.Intn đều sẽ trả về cùng một giá trị. ( để nhận được số khác nhau, việc tạo ra số này tham khảo hàm rand.Seed.Time như một hằng số trên playground, vì thế bạn sẽ sử dụng một thứ khác để tạo ra nó.)
