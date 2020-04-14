//结构体
type movie struct{
    Name string
    Rating float32
}

获得指针 需要&

a:=b 与 a:=&b 区别
1把b的值赋值给a 改变a的属性 b不变
2把b的地址值赋值给a a和b指向同一块地址 改变a的属性 b也改变