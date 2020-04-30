结构体  
    是一种数据类型

        type Person struct{ --类型定义（）
            name string
            sex byte
            age int
        }
普通变量定义和初始化
        1顺序初始化：
        var man Person=Person{"andy",'m',20}
        man:=Person{name:"rose,age:18}  未初始化成员变量，取默认值