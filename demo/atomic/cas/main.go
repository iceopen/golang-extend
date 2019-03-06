package main

import (
	"fmt"
	"sync/atomic"
)

var value int32

func main() {
	fmt.Println("======old value=======")
	fmt.Println(value)
	fmt.Println("======CAS value=======")
	addValue(3)
	fmt.Println(value)

}

//不断地尝试原子地更新value的值,直到操作成功为止
func addValue(delta int32) {
	//在被操作值被频繁变更的情况下,CAS操作并不那么容易成功
	//so 不得不利用for循环以进行多次尝试
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			//在函数的结果值为true时,退出循环
			break
		}
		//操作失败的缘由总会是value的旧值已不与v的值相等了.
		//CAS操作虽然不会让某个Goroutine阻塞在某条语句上,但是仍可能会使流产的执行暂时停一下,不过时间大都极其短暂.
	}
}
