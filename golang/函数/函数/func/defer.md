关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。
关键字 defer 允许我们进行一些函数执行完成后的收尾工作
关闭文件流 
// open a file  
defer file.Close()
解锁一个加锁的资源 
mu.Lock()  
defer mu.Unlock() 
打印最终报告
printHeader()  
defer printFooter()
关闭数据库链接
// open a database connection  
defer disconnectFromDB()
使用 defer 语句实现代码追踪

一个基础但十分实用的实现代码执行追踪的方案就是在进入和离开某个函数打印相关的消息，即可以提炼为下面两个函数：

func trace(s string) { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }