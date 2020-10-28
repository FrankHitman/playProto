package main

import (
	"fmt"
)

func main() {
	var c = [...]int{2: 3, 1: 2}
	fmt.Printf("c: %#v \n", c)

	var d = [...]int{1, 2, 4: 5, 6}
	fmt.Printf("d type: %T, info: %#v \n", d, d)

}

// ---output----
// c: [3]int{0, 2, 3}  index:value 形式的初始化，因此元素的初始化值出现顺序比较随意。
// d type: [6]int, info: [6]int{1, 2, 0, 0, 5, 6}  混合方式初始化，前面两个元素采用顺序初始化，第三第四个元素零值初始化，
// 第五个元素通过索引初始化，最后一个元素跟在前面的第五个元素之后采用顺序初始化。”
//
// 通过数组指针操作数组的方式和通过数组本身的操作类似，而且数组指针赋值时只会拷贝一个指针。
// 但是数组指针类型依然不够灵活，因为数组的长度是数组类型的组成部分，指向不同长度数组的数组指针类型也是完全不同的。”
//
// c1
//