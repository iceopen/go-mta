package api

import (
	"testing"

	"fmt"
)

// 应用每天的pv\uv\vv\iv数据。
func TestCtrCoreData(t *testing.T) {
	str := CtrCoreData("20180107", "20180109", "pv,uv,vv,iv")
	fmt.Println(str)
}
