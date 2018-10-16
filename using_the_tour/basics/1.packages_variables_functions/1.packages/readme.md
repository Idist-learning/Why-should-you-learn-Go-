Mọi chương trình của Go đều được tạo từ các packpage.

Các chương trình đều chạy từ package main

Chương trình này sử dụng package được import từ đường dẫn "fmt" và "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
Theo quy ước, một tên của package giống với 

Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number. (To see a different number, seed the number generator; see rand.Seed. Time is constant in the playground, so you will need to use something else as the seed.)

