package main

import "testing"

func Add(x int, y int) (z int) {
	z = x + y
	return
}

type ForTest struct {
	num int
}

func (this *ForTest) Loops() {
	for i := 0; i < 10000; i++ {
		this.num++
	}
}

type AddArray struct {
	result int
	add_1  int
	add_2  int
}

func BenchmarkLoops(b *testing.B) {
	var test ForTest
	ptr := &test
	// 必须循环 b.N 次 。 这个数字 b.N 会在运行中调整，以便最终达到合适的时间消耗。方便计算出合理的数据。 （ 免得数据全部是 0 ）
	for i := 0; i < b.N; i++ {
		ptr.Loops()
	}
}

func TestAdd(t *testing.T) {
	var test_data = [3]AddArray{
		{2, 1, 1},
		{5, 2, 3},
		{4, 2, 2},
	}
	for _, v := range test_data {
		if v.result != Add(v.add_1, v.add_2) {
			t.Errorf("Add( %d , %d ) != %d \n", v.add_1, v.add_2, v.result)
		}
	}
}
