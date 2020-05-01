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
