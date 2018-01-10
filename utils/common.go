package utils

import (
	"crypto/md5"
	"fmt"

	"github.com/iceopen/go-mta/config"
)

// sig生成算法
func SigMake(params map[string]string) string {
	secret_key := config.SECRET_KEY
	for k, v := range params {
		secret_key = secret_key + k + "=" + v
	}
	data := []byte(secret_key)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
