打开 创建 文件
    创建文件 creat 文件不存在创建 文件存在 清空原文件
            参数 name 打开文件的路径
    
    打开文件 open 以只读方式打开文件  文件不存在打开失败

    
    打开文件 openfile 以 只读 只写 读写 方式打开文件

    f, err := os.Open("d:/a.txt")
	if err != nil {

    f, err := os.Create("d:/a.txt",os.O_RDONLY,6)
	if err != nil {

    f, err := os.Open("d:/a.txt")     
        1参数 name  路径   
        2 权限 O_RDONLY  O_WEONLY   O_RDWR 
        3 一般传6

写文件：
    按字符串写:  WriteString  
    f, err := os.OpenFile("d:/a.txt",os.O_RDWR,6)
	if err != nil {
        fmt.Println(err)
        return    
    }
    defer f.Close()
    n,err:=f.WriteString("hello")
    fi err !=nil{
        fmt.Println("err")
        return
    }
    fmt.Println(n)


    按位置写：
        seek()修改文件的读写指针 
            参数1 偏移量
            参数2 偏移起始位置
            io.SeekStart 文件起始位置
            io.SeekCurrent 当前
            io.SeekEnd  结尾位置
            返回值 表示从文件起始位置，到当前文件读写指针的偏移量

    按字节写
    writeAt() 在文件指定便宜位置 写入[]byte  通常搭配Seek()
        参数1 带写入数据
        参数2 偏移量
        返回值 实际写入数


读文件：
    按行读

    按自己读