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
package hashutil

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/liuchonglin/go-utils/intutil"
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
