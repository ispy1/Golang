<font size=5>fmt 输出注意</font>
fmt.Println不格式化像这样的东西%d。而是使用其参数的默认格式，并在参数之间添加空格。

fmt.Println("Hello, playground",i)  // Hello, playground 5
如果要使用printf样式格式，请使用fmt.Printf。

fmt.Printf("Hello, playground %d\n",i)
而且您不必特别讲究类型。%v通常会弄清楚。

fmt.Printf("Hello, playground %v\n",i)