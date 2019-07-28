package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(parts)
	newParts := strings.Join(parts, "-")
	t.Log(newParts)
}
func TestConv(t *testing.T) {
	s := strconv.Itoa(10) //转换成字符串
	t.Log("str", s)
	// t.Log(strconv.Atoi("10"))
	if i, error := strconv.Atoi("10"); error == nil {
		t.Log(10 + i)
	}
}
