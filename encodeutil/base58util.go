// Copyright 2019 go-utils Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package encodeutil

import (
	"bytes"
	"math/big"
	"crypto/sha256"
	"github.com/liuchonglin/go-utils/arrayutil"
	"github.com/liuchonglin/go-utils/intutil"
)

const version = byte(0x00)

/**
Base58是用于Bitcoin中使用的一种独特的编码方式，主要用于产生Bitcoin的钱包地址。

相比Base64，Base58不使用数字"0"，字母大写"O"，字母大写"I"，和字母小写"l"，以及"+"和"/"符号

设计Base58主要的目的是：
1、避免混淆。在某些字体下，数字0和字母大写O，以及字母大写I和字母小写l会非常相似。
2、不使用"+"和"/"的原因是非字母或数字的字符串作为帐号较难被接受。
3、没有标点符号，通常不会被从中间分行。
4、大部分的软件支持双击选择整个字符串。
*/
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// 获取base58
func GetAddress(prefix string, timestamp int64) string {
	data := bytes.Join([][]byte{
		[]byte(prefix),
		intutil.IntToBytes(timestamp),
	}, []byte{})
	hash := sha256.Sum256(data)
	fullPayload := append([]byte{version}, hash[:24]...)
	address := string(Base58Encode(fullPayload))
	return address
}

// 将一个字节数组编码为Base58
func Base58Encode(input []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}
	// https://en.bitcoin.it/wiki/Base58Check_encoding#Version_bytes
	if input[0] == 0x00 {
		result = append(result, b58Alphabet[0])
	}
	arrayutil.ReverseBytes(result)
	return result
}

// 解码Base58编码的数据
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == b58Alphabet[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded
}
