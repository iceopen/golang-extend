package check

import (
	"fmt"
	"io"
	"strconv"
	"testing"

	. "github.com/pingcap/check"
)

func Test(t *testing.T) { TestingT(t) } // 继承testing的方法，可以直接使用go test命令运行

type MySuite struct{} // 创建测试套件结构体

var _ = Suite(&MySuite{})

var a int = 1

func (s *MySuite) SetUpSuite(c *C) {
	str3 := "第1次套件开始执行"
	fmt.Println(str3)

}

func (s *MySuite) TearDownSuite(c *C) {
	str4 := "第1次套件执行完成"
	fmt.Println(str4)
}

func (s *MySuite) SetUpTest(c *C) {
	str1 := "第" + strconv.Itoa(a) + "条用例开始执行"
	fmt.Println(str1)

}

func (s *MySuite) TearDownTest(c *C) {
	str2 := "第" + strconv.Itoa(a) + "条用例执行完成"
	fmt.Println(str2)
	a = a + 1
}

func (s *MySuite) TestHelloWorld(c *C) { // 声明TestHelloWorld方法为MySuite套件的测试用例
	c.Assert(42, Equals, 42)
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}
