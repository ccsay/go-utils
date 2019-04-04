package encodeutil

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

// 对字符串进行gzip压缩，并进行base64编码
func CompressAndEncode(plain string) (string, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err := zw.Write([]byte(plain))
	if err != nil {
		return "", err
	}
	if err := zw.Close(); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// 对base64加密压缩的字符串，进行解压缩并解码
func UncompressAndDecode(cypher string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cypher)
	if err != nil {
		return "", err
	}
	rdata := bytes.NewReader(data)
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return "", err
	}
	plain, err := ioutil.ReadAll(r)
	return string(plain), err
}
