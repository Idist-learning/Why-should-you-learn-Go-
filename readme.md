[Source](https://medium.com/exploring-code/why-should-you-learn-go-f607681fad65 "Permalink to Why should you learn Go? – Exploring Code – Medium")

# Tại sao bạn nên học Go? – Khám phá về code – Medium


![](https://cdn-images-1.medium.com/max/800/1*vHUiXvBE0p0fLRwFHZuAYw.gif)

`"Go sẽ là ngôn ngữ của server trong tương lai" - Tobias Lutke, Shopify`

Trong vài năm qua, đây là một ngôn ngữ mới nổi: [Go hoặc Golang](https://golang.org). không có gì khiến cho các developer phát điên hơn là một ngôn ngữ mới đúng không? Vì thế tôi bắt đầu học Go khoảng 4 đến 5 tháng trước và bây giờ tôi sẽ nói cho biết vì sao bạn nên học ngôn ngữ mới này.

Trong bài viết này tôi sẽ không dạy bạn viết "Hello World!!" như thế nào. Đã có rất nhiều bài viết khác thực hiện việc này. **Tôi sẽ giải thích dựa trên tình trạng của phần mềm - phần cứng của máy tính và tại sao chúng ta cần ngôn ngữ mới như Go** Vì nếu không có các vấn đề thì chúng ta cũng không cần các giải pháp đúng không? 
* * *

### Giới hạn của phần cứng:

[Luật của Moore](http://www.investopedia.com/terms/m/mooreslaw.asp) đang lỗi thời dần.

Chip xử lý Pentium 4 đầu tiên với xung nhịp 3.0GHz được [giới thiệu từ trước năm 2004 bởi Intel](http://www.informit.com/articles/article.aspx?p=339073). Ngày nay, chiếc [Macbook Pro 2016](http://www.apple.com/macbook-pro/specs/) lại có xung nhịp là 2.9GHz. Vì thế, trong gần một thập kỉ, không có quá nhiều sự gia tăng về hiệu năng bộ xử lý. Bạn có thể so sánh việc phát triển hiệu năng xử lý theo thời gian trong biểu đồ sau.

![](https://cdn-images-1.medium.com/max/800/1*Azz7YwzYYR6lDKFj8iIGZg.png)

Từ bảng trên bạn có thể thấy hiệu năng của một luồng và tần số của các vi xử lý vẫn ổn định trong gần thập kỉ qua. Nếu bạn nghĩ việc thêm nhiều transistor hơn nữa là một giải pháp thì bạn đã sai rồi. Việc này là do ở mức rất nhỏ, một số tính chất lượng tử bắt đầu xuất hiện (như hiện tượng đường hầm) và bởi vì nó thực sự tốn nhiều tiền hơn để đưa thêm số lượng transistor vào.

Vì thế, để giải quyết các vấn đề trên,

* Các nhà sản xuất bắt đầu đưa thêm các lõi vào vi xử lý. Ngày nay chúng ta đã có CPU 4 lõi và 8 lõi
* Chúng ta cũng giới thiệu về đa luồng.
* Thêm nhiều cache vào vi xử lý hơn để tăng hiệu năng.

Nhưng những giải pháp trên cũng có những giới hạn của chúng. Chúng ta không thể thêm quá nhiều cache vào vi xử lý để tăng hiệu năng vì cache cũng có giới hạn vật lý: cache càng lớn thì tốc độ càng chậm. Thêm nhiều core vào vi xử lý thì nó cũng tốn nhiều chi phí. Ngoài ra nó cũng không thể mở rộng vô hạn. Các vi xử lý đa nhân có thể chạy nhiều luồng cùng một lúc và gây ra ảnh hưởng lẫn nhau vào tình huống này. Chúng ta sẽ thảo luận lại vấn đề này sau

Vì thế nếu chúng ta không thể áp dụng các cải thiện về mặt phần cứng, cách duy nhất để tăng hiệu năng là thực hiện cải thiện phần mềm. Nhưng  đáng buồn là nhiều ngôn ngữ hiện lại lại không được hiệu quả lắm.

``
"Các vi xử lý hiện đại giống như những chiếc xe nitro vui nhộn, chúng khá nổi ở những cuộc đua ngắn. Tuy nhiên những ngôn ngữ lập trình hiện đại thì giống như Monte Carlo, chúng có mặt ở mọi ngóc ngách" [David Ungar]
``
* * *

### **Go có các Goroutine**

Như chúng ta đã thảo luận ở trên, các nhà sản xuất phần cứng đang cố gắng bổ sung nhiều lõi hơn vào vi xử lý để tăng hiệu năng. Tất cả các dữ liệu trung tâm đều chạy trên những vi xử lý đó và chúng ta kỳ vọng sẽ tăng số lượng lõi trong những năm sắp tới. Thêm vào đó, những ứng dụng ngày nay sử dụng đồng thời nhiều dịch vụ nhỏ khác để duy trì các kết nối tới cơ sở dữ liệu, các tin nhắn trên hàng đợi và duy trì bộ nhớ cache. Vì thế, những phần mềm chúng ta phát triển và những ngôn ngữ lập trình phải hỗ tính đồng thời một cách dễ dàng và nó phải có khả năng mở rộng khi tăng số lượng lõi lên.

Nhưng hầu hết các ngôn ngữ lập trình hiện đại (như Java, Python,...) đều xuất phát từ môi trường đơn luồng ở thập niên 90. Hầu hết các ngôn ngữ lập trình đều hỗ trợ đa luồng. Nhưng vấn đề thực sự là xử lý đồng thời, khoá luồng, loại điều kiện và khoá chết. Những điều này mới tạo ra một ứng dụng đa luồng trên các ngôn ngữ này.

Để ví dụ, tạo một luồng mới trên Java không hiệu quả về mặt bộ nhớ. Vì mỗi luồng tiêu thụ khoảng 1M trong bộ nhớ heap và cuối cùng nếu bạn bắt đầu quay khoảng một nghìn luồng, chúng sẽ khiến heap chịu một áp lực rất lớn và bị shut down khi tràn bộ nhớ. thêm nữa là nếu bạn muốn gao tiếp giữa 3 hay nhiều luồng thì nó thực sự rất khó.

Mặt khác, Go được công bố năm 2009 khi các vi xử lý đa nhân đã rất phổ biến. Đây là lý do vì sao Go được xây dựng để giữ đồng thời các luồng trong bộ nhớ. Go có goroutine thay cho các luồng. Chúng chỉ chiếm khoảng 2KB bộ nhớ heap. vì thế bạn có thể quay cả triệu goroutine cùng lúc.

![](https://cdn-images-1.medium.com/max/800/1*NFojvbkdRkxz0ZDbu4ysNA.jpeg)

**Các lợi ích khác**:

* Goroutine có thể phát triển các đoạn trong ngăn xếp. Nó có nghĩa là chúng sẽ sử dụng bộ nhớ nhiều hơn khi cần.
* Goroutines có thời gian khởi động nhanh hơn các luồng.
* Goroutines đi kèm cả các tích hợp nguyên thuỷ để giao tiếp với nhau (channels).
* Goroutines cho phép bạn tránh việc sử dụng các khoá mutex khi chia sẻ cấu trúc dữ liệu.
* Tương tự goroutines và luồng của hệ điều hành không map với nhau kiểu 1:1. Một luồng goroutine có thể chạy trên nhiều luồng. Goroutine được ghép lại từ một số nhỏ các luồng hệ điều hành. 

``Bạn có thể thấy sự tương tranh tuyệt vời của Rob Pike không phải tương đồng với việc hiểu sâu hơn về điều này``

Những quan điểm trên khiến Go trở lên mạnh mẽ để xử lý đồng thời như Java, C và C++ trong khi giữ các tiến trình xử lý đồng thời và đẹp đẽ như Erlang.

![](https://cdn-images-1.medium.com/max/800/1*xbsHBQJReC5l_VO4XgNSIQ.png)

* * *

**Go chạy trực tiếp trên nền tảng phần cứng**

Lợi ích đáng kể nhất khi sử dụng C, C++ so với các ngôn ngữ bậc cao khác như Java/Python là hiệu năng của chúng. Vì C/C++ được biên dịch mà không cần thông dịch.

Vi xử lý có thể hiểu được mã nhị phân. Nói chung là khi bạn xây dựng một ứng dụng bằng Java hoặc một ngôn ngữ nền tảng JVM khác khi bạn compile project của bạn, nó compile code mà người có thể đọc được sang byte code mà JVM hoặc các máy ảo khác có thể hiểu và chạy trên các hệ điều hành cơ bản. Khi thực thi, máy ảo thông dịch những bytecode này và chuyển chúng về dạng nhị phân cho phép vi xử lý có thể hiểu được.

![](https://cdn-images-1.medium.com/max/800/1*TVR-VLVg68KwCOLjqQmQAw.png)

Trong khi đó, C/C++ không cần thực thi trên các máy ảo và nó xoá bước này khỏi quá trình thực hiện và tăng hiệu năng. nó compile trực tiếp từ code người có thể đọc được sang dạng nhị phân.

![](https://cdn-images-1.medium.com/max/800/1*ii6xUkU_PchybiG8_GnOjA.png)

Nhưng việc giải phóng hay phân bố các biến trong những ngôn ngữ này là vấn đề lớn. trong khi hầu hết các ngôn ngữ lập trình xử lý phân bố các đối tượng và xoá bỏ bằng việc sử dụng Garbage Collector hoặc các thuật toán Reference Counting. 

Go mang lại giải pháp tối ưu cho cả 2. Giống các ngôn ngữ bậc thấp như C/C++ , Go là một ngôn ngữ biên dịch. Nó có nghĩa là hiệu năng gần bằng với các ngôn ngữ bậc thấp. nó cũng sử dụng Garbage Collection để phân bố và xoá các đối tượng. vì vậy không cần các trạng thái như malloc() và free()!! Thật tuyệt vời !!


* * *

**Code được viết bằng Go rất dễ để bảo trì**

Để tôi nói cho bạn 1 thứ. Go không có các cú pháp điên rồ như những ngôn ngữ lập trình khác. Cú pháp của nó rất gọn gàng và sạch sẽ.

Những người thiết kế ra Go ở google đã nghĩ những thứ này trong tâm trí khi họ tạo ra ngôn ngữ này. Vì google có một nền code rất lớn và cả nghìn developer đang làm việc trên nền code này. do đó code phải cực kì đơn giản để developer khác có thể hiểu và mỗi đoạn code phải cố gắng tối thiểu việc ảnh hưởng tới những đoạn code khác. Việc này sẽ khiến code dễ bảo trì và sửa đổi hơn.

Go cố tình bỏ qua các tính năng của một ngôn ngữ OOP hiện đại.

* Không có các class. Tất cả mọi thứ đều được phân chia và những packages. Go chỉ có struct thay cho các class.
* Không hỗ trợ việc kế thừa. Điều này khiến code dễ sửa đổi hơn. Ở các ngôn ngữ khác như Java/Python, nếu một class ABC kế thừa từ class XYZ và bạn thực hiện việc thay đổi trong class XYZ, sau đó thì điều này có thể tạo ra một vài ảnh hưởng ngoại lệ lên class khác mà kế thừa từ XYZ. Việc xoá bỏ tính kế thừa, Go khiến code dễ hiểu hơn (vì không có những super class để xem xét khi phải xem xét một đoạn code)
* Không có khởi tạo
* không có chú thích
* không có đặc tính chung
* Không có các ngoại lệ

Những thay đổi trên khiến cho Go rất khác so với những ngôn ngữ khác và nó khiến việc lập trình bằng Go cũng khác so với chúng. Bạn có thể không thích những điểm trên. Nhưng nó không giống như việc bạn không thể code ứng dụng của bạn mà không có các tính năng trên. Tất cả những gì bạn phải làm là viết thêm 2-3 dòng code. Nhưng ngược lại nó khiến bạn code sạch và rõ ràng hơn.  

![](https://cdn-images-1.medium.com/max/800/1*nlpYI256BR71xMBWd1nlfg.png)

Biểu đồ trên thể hiện rằng Go hầu như hiệu quả hơn C/C++, trong khi vẫn giữ được cú pháp code đơn giản như Ruby, Python và những ngôn ngữ khác. Đây là một thứ có lợi cho cả con người và các vi xử lý !!!

[Không giống các ngôn ngữ khác như Swift](https://www.quora.com/Is-Swifts-syntax-still-changing), cú pháp của Go đã rất ổn định. Nó vẫn giữ nguyên kể từ bản phát hành 1.0 đầu tiên năm 2012. Việc này giúp nó có tính tương thích ngược.

**Go được hỗ trợ bởi Google**

* Tôi biết điều này không phải là lợi ích trực tiếp về mặt kĩ thuật. Nhưng Go được thiết kế và hỗ trợ bởi Google. Google có một cơ sở hạ tầng đám mây cực lớn trên thế giới và nó đang mở rộng một cách nhanh chóng. Go được thiết kế bởi Google để giải quyết các vấn đề của họ về việc hỗ trợ mở rộng và ít ảnh hưởng nhất. Đó cũng là những vấn đề tương tự khi bạn phải đối mặt khi bạn tạo những server của riêng mình.
* Thêm vào đó Go cũng được sử dụng bởi vài công ty như Adobe, BBC, IBM, Intel và cả [Medium](https://medium.engineering/how-medium-goes-social-b7dbefa6d413#.r8nqjxjpk).(Source: [https://github.com/golang/go/wiki/GoUsers](https://github.com/golang/go/wiki/GoUsers))

**Kết luận:**

* Mặc dù Go rất khác so với những ngôn ngữ hướng đối tượng khác, nó vẫn còn rất thô. Go cung cấp cho bạn hiệu năng như C/C++, xử lý đồng thời siêu hiệu quả như Java và code thú vị như Python/Perl.

* Nếu bạn không có dự định gì để học Go, tôi sẽ tiếp tục nói về giới hạn phần cứng gây áp lực cho chúng tôi, các developer phần mềm cần phải viết code siêu hiệu quả hơn. Developer cần hiểu là phần ứng và  tối ưu hoá chương trình của bạn phù hợp hơn. việc tối ưu phần mềm để có thể chạy trên phần cứng rẻ hơn, chậm hơn (như các thiết bị IOT) và vẫn có các tác động về tổng thể tốt hơn đối với trải nghiệm của người dùng.
