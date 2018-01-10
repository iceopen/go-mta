package utils

import (
	"fmt"
	"testing"
)

// sig生成算法
func TestSigMake(t *testing.T) {
	m := make(map[string]string)
	m["hello"] = "echo hello"
	m["world"] = "echo world"
	m["go"] = "echo go"
	m["is"] = "echo is"
	m["cool"] = "echo cool"
	fmt.Println(SigMake(m))
}
