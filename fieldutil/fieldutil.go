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

// 基本数据类型转与指针类型转换
package fieldutil

func IntToPtr(i int) *int {
	return &i
}

func Int8ToPtr(i int8) *int8 {
	return &i
}

func Int16ToPtr(i int16) *int16 {
	return &i
}

func Int32ToPtr(i int32) *int32 {
	return &i
}

func Int64ToPtr(i int64) *int64 {
	return &i
}

func UintToPtr(i uint) *uint {
	return &i
}

func Uint8ToPtr(i uint8) *uint8 {
	return &i
}

func Uint16ToPtr(i uint16) *uint16 {
	return &i
}

func Uint32ToPtr(i uint32) *uint32 {
	return &i
}

func Uint64ToPtr(i uint64) *uint64 {
	return &i
}

func StringToPtr(s string) *string {
	return &s
}

func BoolToPtr(b bool) *bool {
	return &b
}

func ToInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
func ToInt8(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}
func ToInt16(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}
func ToInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}
func ToInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func ToUint(i *uint) uint {
	if i == nil {
		return 0
	}
	return *i
}
func ToUint8(i *uint8) uint8 {
	if i == nil {
		return 0
	}
	return *i
}
func ToUint16(i *uint16) uint16 {
	if i == nil {
		return 0
	}
	return *i
}
func ToUint32(i *uint32) uint32 {
	if i == nil {
		return 0
	}
	return *i
}
func ToUint64(i *uint64) uint64 {
	if i == nil {
		return 0
	}
	return *i
}

func ToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ToBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
