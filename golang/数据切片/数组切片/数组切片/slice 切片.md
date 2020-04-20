切片的初始化格式是：var slice1 []type = arr1[start:end]。

这表示 slice1 是由数组 arr1 从 start 索引到 end-1 索引之间的元素构成的子集（切分数组，start:end 被称为 slice 表达式）。所以 slice1[0] 就等于 arr1[start]。这可以在 arr1 被填充前就定义好。

var slice1 []type = arr1[:] 那么 slice1 就等于完整的 arr1 数组（所以这种表示方式是 arr1[0:len(arr1)] 的一种缩写）。另外一种表述方式是：slice1 = &arr1。

arr1[2:] 和 arr1[2:len(arr1)] 相同，都包含了数组从第三个到最后的所有元素。

arr1[:3] 和 arr1[0:3] 相同，包含了从第一个到第三个元素（不包括第四个）
想去掉 slice1 的最后一个元素，只要 slice1 = slice1[:len(slice1)-1]

绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!!

当相关数组还没有定义时，我们可以使用 make() 函数来创建一个切片 同时创建好相关数组：var slice1 []type = make([]type, len)。

也可以简写为 slice1 := make([]type, len)，这里 len 是数组的长度并且也是 slice 的初始长度
所以定义 s2 := make([]int, 10)，那么 cap(s2) == len(s2) == 10
cap()函数返回的是数组切片分配的空间大小

如果你想创建一个 slice1，它不占用整个数组，而只是占用以 len 为个数个项，那么只要：slice1 := make([]type, len, cap)

这种构建方法可以应用于数组和切片:

for ix, value := range slice1 {
	...
}
第一个返回值 ix 是数组或者切片的索引，第二个是在该索引位置的值；他们都是仅在 for 循环内部可见的局部变量。value 只是 slice1 某个索引位置的值的一个拷贝，不能用来修改 slice1 该索引位置的值
ackage main

import "fmt"

func main() {
	var slice1 []int = make([]int, 4)

	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
}

slices_forrange2.go 给出了一个关于字符串的例子， _ 可以用于忽略索引。

如果你只需要索引，你可以忽略第二个变量，例如：

for ix := range seasons {
	fmt.Printf("%d", ix)
}
// Output: 0 1 2 3