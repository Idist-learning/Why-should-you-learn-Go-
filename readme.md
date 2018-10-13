[Source](https://medium.com/exploring-code/why-should-you-learn-go-f607681fad65 "Permalink to Why should you learn Go? – Exploring Code – Medium")

# Why should you learn Go? – Exploring Code – Medium
# Tại sao bạn nên học Go? – Khám phá về code – Medium


![](https://cdn-images-1.medium.com/max/800/1*vHUiXvBE0p0fLRwFHZuAYw.gif)

`“Go will be the server language of the future.” — Tobias Lütke, Shopify`
`"Go sẽ là ngôn ngữ của server trong tương lai" - Tobias Lutke, Shopify`

In past couple of years, there is a rise of new programming language: [Go or GoLang](https://golang.org/). Nothing makes a developer crazy than a new programming language, right? So, I started learning Go before 4 to 5 months and here I am going to tell you about why you should also learn this new language.
Trong viaf năm qua, đây là một ngôn ngữ mới nổi: [Go hoặc Golang](https://golang.org). không có gì khiến cho các developer phát điên hơn là một ngôn ngữ mới đúng không? Vì thế tôi bắt đầu học Go khoảng 4 đến 5 tháng trước và bây giờ tôi sẽ nói cho biết vì sao bạn nên học ngôn ngữ mới này.

I am not going to teach you, how you can write “Hello World!!” in this article. There are lots of other articles online for that. **I am going the explain current stage of computer hardware-software and why we need new language like Go?** Because if there isn’t any problem, then we don’t need solution, right?
Trong bài viết này tôi sẽ không dạy bạn viết "Hello World!!" như thế nào. Đã có rất nhiều bài viết khác thực hiện việc này. **Tôi sẽ giải thích dựa trên tình trạng của phần mềm - phần cứng của máy tính và tại sao chúng ta cần ngôn ngữ mới như Go** Vì nếu không có các vấn đề thì chúng ta cũng không cần các giải pháp đúng không? 
* * *

### Hardware limitations:
### Giới hạn của phần cứng:

[Luật của Moore](http://www.investopedia.com/terms/m/mooreslaw.asp) đang lỗi thời dần.

First Pentium 4 processor with 3.0GHz clock speed was [introduced back in 2004 by Intel](http://www.informit.com/articles/article.aspx?p=339073). Today, my [Mackbook Pro 2016](http://www.apple.com/macbook-pro/specs/) has clock speed of 2.9GHz. So, nearly in one decade, there is not too much gain in the raw processing power. You can see the comparison of increasing the processing power with the time in below chart.
Chip xử lý Pentium 4 đầu tiên với xung nhịp 3.0GHz được [giới thiệu từ trước năm 2004 bởi Intel](http://www.informit.com/articles/article.aspx?p=339073). Ngày nay, chiếc [Macbook Pro 2016](http://www.apple.com/macbook-pro/specs/) lại có xung nhịp là 2.9GHz. Vì thế, trong gần một thập kỉ, không có quá nhiều lợi ích từ sức mạnh trong quá trình xử lý nguyên thuỷ. Bạn có thể so sánh việc phát triển sức mạnh của việc xử lý theo thời gian trong biểu đồ sau.

![](https://cdn-images-1.medium.com/max/800/1*Azz7YwzYYR6lDKFj8iIGZg.png)

From the above chart you can see that the single-thread performance and the frequency of the processor remained steady for almost a decade. If you are thinking that adding more transistor is the solution, then you are wrong. This is because at smaller scale some quantum properties starts to emerge (like tunneling) and because it actually costs more to put more transistors ([why?](https://www.quora.com/What-is-Quantum-Tunneling-Limit-How-does-it-limit-the-size-of-a-transistor)) and the number of transistors you can add per dollar starts to fall.
Từ bảng trên bạn có thể thấy hiệu năng của một luồng và tần số của các vi xử lý vẫn ổn định trong gần thập kỉ qua. Nếu bạn nghĩ việc thêm nhiều transistor hơn nữa là một giải pháp thì bạn đã sai rồi. Việc này là do ở mức rất nhỏ, một số tính chất lượng tử bắt đầu xuất hiện (như hiện tượng đường hầm) và bởi vì nó thực sự tốn nhiều tiền hơn để đưa thêm số lượng transistor vào.

So, for the solution of above problem,
Vì thế, để giải quyết các vấn đề trên,

* Manufacturers started adding more and more cores to the processor. Nowadays we have quad-core and octa-core CPUs available.
* Các nhà sản xuất bắt đầu đưa thêm các lõi vào vi xử lý. Ngày nay chúng ta đã có CPU 4 lõi và 8 lõi
* We also introduced hyper-threading.
* Chúng ta cũng giới thiệu về đa luồng.
* Added more cache to the processor to increase the performance.
* Thêm nhiều cache vào vi xử lý hơn để tăng hiệu năng.

But above solutions have its own limitations too. We cannot add more and more cache to the processor to increase performance as cache have physical limits: the bigger the cache, the slower it gets. Adding more core to the processor has its cost too. Also, that cannot scale to indefinitely. These multi-core processors can run multiple threads simultaneously and that brings concurrency to the picture. We’ll discuss it later.
Nhưng những giải pháp trên cũng có những giới hạn của chúng. Chúng ta không thể thêm quá nhiều cache vào vi xử lý để tăng hiệu năng vì cache cũng có giới 
hạn vật lý: cache càng lớn thì tốc độ càng chậm. Thêm nhiều core vào vi xử lý thì nó cũng tốn nhiều chi phí. Ngoài ra nó cũng không thể mở rộng vô hạn. Các vi xử lý đa nhân có thể chạy nhiều luồng cùng một lúc và gây ra ảnh hưởng lẫn nhau vào tình huống này. Chúng ta sẽ thảo luận lại vấn đề này sau

So, if we cannot rely on the hardware improvements, the only way to go is more efficient software to increase the performance. But sadly, modern programming language are not much efficient.
Vì thế nếu chúng ta không thể áp dụng các cải thiện về mặt phần cứng, cách duy nhất để tăng hiệu năng là thực hiện cải thiện phần mềm. Nhưng  đáng buồn là nhiều ngôn ngữ hiện lại lại không được hiệu quả lắm.

``
“Modern processors are a like nitro fueled funny cars, they excel at the quarter mile. Unfortunately modern programming languages are like Monte Carlo, they are full of twists and turns.” — [David Ungar]
"Các vi xử lý hiện đại giống như những chiếc xe nitro vui nhộn, chúng khá nổi ở những cuộc đua ngắn. Tuy nhiên những ngôn ngữ lập trình hiện đại thì giống như Monte Carlo, chúng có mặt ở mọi ngóc ngách" [David Ungar]
``
* * *

### **Go has goroutines !!**
### **Go có các Goroutine**

As we discussed above, hardware manufacturers are adding more and more cores to the processors to increase the performance. All the data centers are running on those processors and we should expect increase in the number of cores in upcoming years. More to that, today’s applications using multiple micro-services for maintaining database connections, message queues and maintain caches. So, the software we develop and the programming languages should support concurrency easily and they should be scalable with increased number of cores.
Như chúng ta đã thảo luận ở trên, các nhà sản xuất phần cứng đang cố gắng bổ sung nhiều lõi hơn vào vi xử lý để tăng hiệu năng. Tất cả các dữ liệu trung tâm đều chạy trên những vi xử lý đó và chúng ta dựu kiến sẽ tăng số lượng lõi trong những năm sắp tới. Thêm vào đó, những ứng dụng ngày nay sử dụng đồng thời nhiều dịch vụ nhỏ khác để duy trì các kết nối tới cơ sở dữ liệu, các tin nhắn trên hàng đợi và duy trì bộ nhớ cache. Vì thế, những phần mềm chúng ta phát triển và những ngôn ngữ lập trình phải hỗ tính đồng thời một cách dễ dàng và nó phải có khả năng mở rộng khi tăng số lượng lõi lên.

But, most of the modern programming languages(like Java, Python etc.) are from the ’90s single threaded environment. Most of those programming languages supports multi-threading. But the real problem comes with concurrent execution, threading-locking, race conditions and deadlocks. Those things make it hard to create a multi-threading application on those languages.
Nhưng hầu hết các ngôn ngữ lập trình hiện đại (như Java, Python,...) đều xuất phát từ môi trường đơn luồng ở thập niên 90. Hầu hết các ngôn ngữ lập trình đều hỗ trợ đa luồng. Nhưng vấn đề thực sự là xử lý đồng thời, khoá luồng, loại điều kiện và khoá chết. Những điều này mới tạo ra một ứng dụng đa luồng trên các ngôn ngữ này.

For an example, creating new thread in Java is not memory efficient. As every thread consumes approx 1MB of the memory heap size and eventually if you start spinning thousands of threads, they will put tremendous pressure on the heap and will cause shut down due to out of memory. Also, if you want to communicate between two or more threads, it’s very difficult.
Để ví dụ, tạo một luồng mới trên Java không hiệu quả về mặt bộ nhớ. Vì mỗi luồng tiêu thụ khoảng 1M trong bộ nhớ heap và cuối cùng nếu bạn bắt đầu quay khoảng một nghìn luồng, chúng sẽ khiến heap chịu một áp lực rất lớn và bị shut down khi tràn bộ nhớ. thêm nữa là nếu bạn muốn gao tiếp giữa 3 hay nhiều luồng thì nó thực sự rất khó.

On the other hand, Go was released in 2009 when multi-core processors were already available. That’s why Go is built with keeping concurrency in mind. Go has goroutines instead of threads. They consume almost 2KB memory from the heap. So, you can spin millions of goroutines at any time.
Mặt khác, Go được công bố năm 2009 khi các vi xử lý đa nhân đã rất phổ biến. Đây là lý do vì sao Go được xây dựng để giữ đồng thời các luồng trong bộ nhớ. Go có goroutine thay cho các luồng. Chúng chỉ chiếm khoảng 2KB bộ nhớ heap. vì thế bạn có thể quay cả triệu goroutine cùng lúc.

![](https://cdn-images-1.medium.com/max/800/1*NFojvbkdRkxz0ZDbu4ysNA.jpeg)

**Other benefits are**:
**Các lợi ích khác**:

* Goroutines have growable segmented stacks. That means they will use more memory only when needed.
* Goroutine có thể phát triển các đoạn trong ngăn xếp. Nó có nghĩa là chúng sẽ sử dụng bộ nhớ nhiều hơn khi cần.
* Goroutines have a faster startup time than threads.
* Goroutines có thời gian khởi động nhanh hơn các luồng.
* Goroutines come with built-in primitives to communicate safely between themselves (channels).
* Goroutines đi kèm cả các tích hợp nguyên thuỷ để giao tiếp với nhau (channels).
* Goroutines allow you to avoid having to resort to mutex locking when sharing data structures.
* Goroutines cho phép bạn tránh việc sử dụng các khoá mutex khi chia sẻ cấu trúc dữ liệu.
* Also, goroutines and OS threads do not have 1:1 mapping. A single goroutine can run on multiple threads. Goroutines are multiplexed into small number of OS threads.
* Tương tự goroutines và luồng của hệ điều hành không map với nhau kiểu 1:1. Một luồng goroutine có thể chạy trên nhiều luồng. Goroutine được ghép lại từ một số nhỏ các luồng hệ điều hành. 

``You can see Rob Pike’s excellent talk concurrency is not parallelism to get more deep understanding on this.``
``Bạn có thể thấy sự tương tranh tuyệt vời của Rob Pike không phải tương đồng với việc hiểu sâu hơn về điều này``

All the above points, make Go very powerful to handle concurrency like Java, C and C++ while keeping concurrency execution code strait and beautiful like Erlang.
Những quan điểm trên khiến Go trở lên mạnh mẽ để xử lý đồng thời như Java, C và C++ trong khi giữ các tiến trình xử lý đồng thời và đẹp đẽ như Erlang.

![](https://cdn-images-1.medium.com/max/800/1*xbsHBQJReC5l_VO4XgNSIQ.png)

* * *

**Go runs directly on underlying hardware.**
**Go chạy trực tiếp trên phần cứng cơ bản**

One most considerable benefit of using C, C++ over other modern higher level languages like Java/Python is their performance. Because C/C++ are compiled and not interpreted.
Lợi ích đáng kể nhất khi sử dụng C, C++ so với các ngôn ngữ bậc cao khác như Java/Python là hiệu năng của chúng. Vì C/C++ được biên dịch mà không cần thông dịch.

Processors understand binaries. Generally, when you build an application using Java or other JVM-based languages when you compile your project, it compiles the human readable code to byte-code which can be understood by JVM or other virtual machines that run on top of underlying OS. While execution, VM interprets those bytecodes and convert them to the binaries that processors can understand.
Vi xử lý có thể hiểu được mã nhị phân. Nói chung là khi bạn xây dựng một ứng dụng bằng Java hoặc một ngôn ngữ nền tảng JVM khác khi bạn compile project của bạn, nó compile code mà người có thể đọc được sang byte code mà JVM hoặc các máy ảo khác có thể hiểu và chạy trên các hệ điều hành cơ bản. Khi thực thi, máy ảo thông dịch những bytecode này và chuyển chúng về dạng nhị phân cho phép vi xử lý có thể hiểu được.

![](https://cdn-images-1.medium.com/max/800/1*TVR-VLVg68KwCOLjqQmQAw.png)

While on the other side, C/C++ does not execute on VMs and that removes one step from the execution cycle and increases the performance. It directly compiles the human readable code to binaries.
Trong khi đó, C/C++ không cần thực thi trên các máy ảo và nó xoá bước này khỏi quá trình thực hiện và tăng hiệu năng. nó compile trực tiếp từ code người có thể đọc được sang dạng nhị phân.

![](https://cdn-images-1.medium.com/max/800/1*ii6xUkU_PchybiG8_GnOjA.png)

But, freeing and allocating variable in those languages is a huge pain. While most of the programming languages handle object allocation and removing using Garbage Collector or Reference Counting algorithms.
Nhưng việc giải phóng hay phân bố các biến trong những ngôn ngữ này là vấn đề lớn. trong khi hầu hết các ngôn ngữ lập trình xử lý phân bố các đối tượng và xoá bỏ bằng việc sử dụng Garbage Collector hoặc các thuật toán Reference Counting. 

Go brings best of both the worlds. Like lower level languages like C/C++, Go is compiled language. That means performance is almost nearer to lower level languages. It also uses garbage collection to allocation and removal of the object. So, no more malloc() and free() statements!!! Cool!!!
Go mang lại giải pháp tối ưu cho cả 2. Giống các ngôn ngữ bậc thấp như C/C++ , Go là một ngôn ngữ biên dịch. Nó có nghĩa là hiệu năng gần bằng với các ngôn ngữ bậc thấp. nó cũng sử dụng Garbage Collection để phân bố và xoá các đối tượng. vì vậy không cần các trạng thái như malloc() và free()!! Thật tuyệt vời !!


* * *

**Code written in Go is easy to maintain.**
**Code được viết bằng Go rất dễ để bảo trì**
Let me tell you one thing. Go does not have crazy programming syntax like other languages have. It has very neat and clean syntax.
Để tôi nói cho bạn 1 thứ. Go không có các cú pháp điên rồ như những ngôn ngữ lập trình khác. Cú pháp của nó rất gọn gàng và sạch sẽ.

The designers of the Go at google had this thing in mind when they were creating the language. As google has the very large code-base and thousands of developers were working on that same code-base, code should be simple to understand for other developers and one segment of code should has minimum side effect on another segment of the code. That will make code easily maintainable and easy to modify.
Những người thiết kế ra Go ở google đã nghĩ những thứ này trong tâm trí khi họ tạo ra ngôn ngữ này. Vì google có một nền code rất lớn và cả nghìn developer đang làm việc trên nền code này. do đó code phải cực kì đơn giản để developer khác có thể hiểu và mỗi đoạn code phải cố gắng tối thiểu việc ảnh hưởng tới những đoạn code khác. Việc này sẽ khiến code dễ bảo trì và sửa đổi hơn.

Go intentionally leaves out many features of modern OOP languages.
Go cố tình bỏ qua các tính năng của một ngôn ngữ OOP hiện đại.

* No classes. Every thing is divided into packages only. Go has only structs instead of classes.
* Không có các class. Tất cả mọi thứ đều được phân chia và những packages. Go chỉ có struct thay cho các class.
* Does not support inheritance. That will make code easy to modify. In other languages like Java/Python, if the class ABC inherits class XYZ and you make some changes in class XYZ, then that may produce some side effects in other classes that inherit XYZ. By removing inheritance, Go makes it easy to understand the code also (as there is no super class to look at while looking at a piece of code).
* Không hỗ trợ việc kế thừa. Điều này khiến code dễ sửa đổi hơn. Ở các ngôn ngữ khác như Java/Python, nếu một class ABC kế thừa từ class XYZ và bạn thực hiện việc thay đổi trong class XYZ, sau đó thì điều này có thể tạo ra một vài ảnh hưởng ngoại lệ lên class khác mà kế thừa từ XYZ. Việc xoá bỏ tính kế thừa, Go khiến code dễ hiểu hơn (vì không có những super class để xem xét khi phải xem xét một đoạn code)
* No constructors.
* Không có khởi tạo
* No annotations.
* không có chú thích
* No generics.
* không có đặc tính chung
* No exceptions.
* Không có các ngoại lệ

Above changes make Go very different from other languages and it makes programming in Go different from others. You may not like some points from above. But, it is not like you can not code your application without above features. All you have to do is write 2–3 more lines. But on the positive side, it will make your code cleaner and add more clarity to your code.
Những thay đổi trên khiến cho Go rất khác so với những ngôn ngữ khác và nó khiến việc lập trình bằng Go cũng khác so với chúng. Bạn có thể không thích những điểm trên. Nhưng nó không giống như việc bạn không thể code ứng dụng của bạn mà không có các tính năng trên. Tất cả những gì bạn phải làm là viết thêm 2-3 dòng code. Nhưng ngược lại nó khiến bạn code sạch và rõ ràng hơn.  

![](https://cdn-images-1.medium.com/max/800/1*nlpYI256BR71xMBWd1nlfg.png)

Above graph displays that Go is almost as efficient as C/C++, while keeping the code syntax simple as Ruby, Python and other languages. That is a win-win situation for both humans and processors!!!
Biểu đồ trên thể hiện rằng Go hầu như hiệu quả hơn C/C++, trong khi vẫn giữ được cú pháp code đơn giản như Ruby, Python và những ngôn ngữ khác. Đây là một thứ có lợi cho cả con người và các vi xử lý !!!

[Unlike other new languages like Swift](https://www.quora.com/Is-Swifts-syntax-still-changing), it’s syntax of Go is very stable. It remained same since the initial public release 1.0, back in year 2012. That makes it backward compatible.
[Không giống các ngôn ngữ khác như Swift](https://www.quora.com/Is-Swifts-syntax-still-changing), cú pháp của Go đã rất ổn định. Nó vẫn giữ nguyên kể từ bản phát hành 1.0 đầu tiên năm 2012. Việc này giúp nó có tính tương thích ngược.

**Go is backed by Google.**
**Go được hỗ trợ bởi Google**

* I know this is not a direct technical advantage. But, Go is designed and supported by Google. Google has one of the largest cloud infrastructures in the world and it is scaled massively. Go is designed by Google to solve their problems of supporting scalability and effectiveness. Those are the same issues you will face while creating your own servers.
* Tôi biết điều này không phải là lợi ích trực tiếp về mặt kĩ thuật. Nhưng Go được thiết kế và hỗ trợ bởi Google. Google có một cơ sở hạ tầng đám mây cực lớn trên thế giới và nó đang mở rộng một cách nhanh chóng. Go được thiết kế bởi Google để giải quyết các vấn đề của họ về việc hỗ trợ mở rộng và ít ảnh hưởng nhất. Đó cũng là những vấn đề tương tự khi bạn phải đối mặt khi bạn tạo những server của riêng mình.
* More to that Go is also used by some big companies like Adobe, BBC, IBM, Intel and even [Medium](https://medium.engineering/how-medium-goes-social-b7dbefa6d413#.r8nqjxjpk).(Source: [https://github.com/golang/go/wiki/GoUsers](https://github.com/golang/go/wiki/GoUsers))
* Thêm vào đó Go cũng được sử dụng bởi vài công ty như Adobe, BBC, IBM, Intel và cả [Medium](https://medium.engineering/how-medium-goes-social-b7dbefa6d413#.r8nqjxjpk).(Source: [https://github.com/golang/go/wiki/GoUsers](https://github.com/golang/go/wiki/GoUsers))

**Conclusion:**
**Kết luận:**

* Even though Go is very different from other object-oriented languages, it is still the same beast. Go provides you high performance like C/C++, super efficient concurrency handling like Java and fun to code like Python/Perl.
* Mặc dù Go rất khác so với những ngôn ngữ hướng đối tượng khác, nó vẫn còn rất thô. Go cung cấp cho bạn hiệu năng như C/C++, xử lý đồng thời siêu hiệu quả như Java và code thú vị như Python/Perl.

* If you don’t have any plans to learn Go, I will still say hardware limit puts pressure to us, software developers to write super efficient code. Developer needs to understand the hardware and make their program optimize accordingly. The optimized software can run on cheaper and slower hardware (like IOT devices) and overall better impact on end user experience.
* Nếu bạn không có dự định gì để học Go, tôi sẽ tiếp tục nói về giới hạn phần cứng gây áp lực cho chúng tôi, các developer phần mềm cần phải viết code siêu hiệu quả hơn. Developer cần hiểu là phần ứng và  tối ưu hoá chương trình của bạn phù hợp hơn. việc tối ưu phần mềm để có thể chạy trên phần cứng rẻ hơn, chậm hơn (như các thiết bị IOT) và vẫn có các tác động về tổng thể tốt hơn đối với trải nghiệm của người dùng.
