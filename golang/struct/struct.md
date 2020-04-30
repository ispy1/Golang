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
    
交媾提变量比较
        ==比较
        相同结构体（成员变量类型，个数，顺序一致） 变量之间可以直接赋值
结构体传参
        unSAfe.Sizeof(变量名)  类型站的内存空间
        将结构体变量的值拷贝一份，传递  --- 几乎不用，内存消耗大，效率低
    
结构体地址：
        结构体元素的地址

指针变量定义和初始化
        1顺序初始化，依次将结构体内部所有成员变量初始化
            var man *Person=&Person{"andy",'m','3'}
        2 new(Person)
    