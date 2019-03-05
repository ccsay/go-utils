package intutil

import "strconv"

// int转byte数组
func IntToBytes(num int64) []byte {
	return []byte(strconv.FormatInt(num, 10))
}
