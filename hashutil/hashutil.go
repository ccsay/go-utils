package hashutil

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"liuchonglin.com/go-utils/intutil"
)

// 获取hash
func GetHash(prefix string, timestamp int64) string {
	data := bytes.Join([][]byte{
		[]byte(prefix),
		intutil.IntToBytes(timestamp),
	}, []byte{})
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
