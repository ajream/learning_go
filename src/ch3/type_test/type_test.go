// 1.不支持隐式转换
// 2.支持指针类型，但指针不可运算
// 3.字符串是值类型，默认初始化值是空字符串，而不是空（nil）
package type_test

import (
	"fmt"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)

	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	fmt.Print(a, aPtr)
	fmt.Printf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	fmt.Print("*" + s + "*")
	fmt.Print(len(s))
}
