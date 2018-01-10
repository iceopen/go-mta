package api

import (
	"testing"

	"fmt"

	"github.com/iceopen/go-mta/test"
)

// 应用每天的pv\uv\vv\iv数据。
func TestCtrCoreData(t *testing.T) {
	test.SetConfig()
	str := CtrCoreData("20180107", "20180109", "pv,uv,vv,iv")
	fmt.Println(str)
}
