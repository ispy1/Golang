数组是具有相同 唯一类型 的一组已编号且长度固定的数据项序列这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。数组长度必须是一个常量表达式，并且必须是一个非负整数。数组长度也是数组类型的一部分，所以[5]int和[10]int是属于不同类型的。数组的编译时值初始化是按照数组顺序完成的


Go 语言中的数组是一种 值类型（不像 C/C++ 中是指向首元素的指针），所以可以通过 new() 来创建： var arr1 = new([5]int)。

那么这种方式和 var arr2 [5]int 的区别是什么呢？arr1 的类型是 *[5]int，而 arr2的类型是 [5]int。

所以在函数中数组作为参数传入时，如 func1(arr2)，会产生一次数组拷贝，func1 方法不会修改原始的数组 arr2。

如果你想修改原数组，那么 arr2 必须通过&操作符以引用方式传过来，例如 func1(&arr2）
package main
import "fmt"
func f(a [3]int) { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

func main() {
	var ar [3]int
	f(ar) 	// passes(传递) a copy of ar
	fp(&ar) // passes a pointer to ar
}

key: value 语法

var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
只有索引 3 和 4 被赋予实际的值，其他元素都被设置为空的字符串，所以输出结果为：

Person at 0 is
Person at 1 is
Person at 2 is
Person at 3 is Chris
Person at 4 is Ron