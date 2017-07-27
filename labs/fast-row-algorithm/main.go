package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// get the random numbers
	origin := rand.Perm(10)
	fmt.Println("origin:", origin)
	quickSort(origin, 0, len(origin))
	randomQuickSort(origin, 0, len(origin))
	fmt.Println("quick sort:", origin)

}

func quickSort(list []int, start, end int) {
	if end-start > 1 {
		// get the pivot
		mid := partition(list, start, end)
		// left sort
		quickSort(list, start, mid)
		// right sort
		quickSort(list, mid+1, end)
	}
}

func partition(list []int, begin, end int) (i int) {
	cValue := list[begin]
	i = begin
	for j := i + 1; j < end; j++ {
		if list[j] < cValue {
			i++ // 这里一定要先加1后在交换值，因为支点现在不必交换，现在i 和 j（小于支点和大于支点）正在划分边界
			list[j], list[i] = list[i], list[j]
			fmt.Println("list:", list)
		}
	}
	/* 到此，i和j的边界已经划分完成，把i对应的值和支点对应的值交换后，就符合了快分的要求：i左边对应的值都小于等于且右边的都大于支点对应的值
	此时i的值就是新的支点, 对应的值就是新的主元
	*/
	list[i], list[begin] = list[begin], list[i]
	return i
}

func randomQuickSort(list []int, start, end int) {
	if end-start > 1 {
		// get the pivot
		mid := randomPartition(list, start, end)
		randomQuickSort(list, start, mid)
		randomQuickSort(list, mid+1, end)
	}
}

func randomPartition(list []int, begin, end int) int {
	// 生成真随机数
	i := randInt(begin, end)
	fmt.Println("random number:", i)
	// 下面这行是核心部分，随机选择主元， 如果没有此次交换，就是普通快排
	list[i], list[begin] = list[begin], list[i]
	return partition(list, begin, end)
}

// 真随机数
func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
