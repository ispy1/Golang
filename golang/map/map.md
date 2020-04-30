map 
    字典 映射  key--value key ：唯一 无序  不能是引用类型数据（切片，函数）
    map 无需
    map 不能用cap
创建map   
var m1 map[int]string 声明一个map 没有初始化 为空（nil）map


m2:=map[int]string{} 

m3 :=make(map[int]string)  leng==0

m4 :=make(map[int]string,5) leng=5

初始化mapr

var m5 map[int]string=map[int]string{1:"2",250:"20"}

m7 :=map[int]string{1:"2",250:"20"}

map赋值
m8 :=make(map[int]string)
m8[100]="aa"
m8[30]="bb"
m8[30]="ccc"  //覆盖原有元素


map 使用
    遍历 map
    for k,v :=range m8{
		fmt.Println(k,v)
	}

    判断key 是否存在
    if v,has:=m8[1]; has{
		fmt.Println(ture)
	} else{
		fmt.Println("false")
	}

    删除map
    delete(m,key) 删除m中键为key的参数
    map 做函数参数和返回值 传引用 （地址8）