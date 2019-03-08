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

// 通用工具
package utils

import (
	"reflect"
	"errors"
	"github.com/liuchonglin/go-utils/stringutil"
	"encoding/json"
	"fmt"
)

var (
	NilPointerError           = errors.New("is a nil pointer")
	NotPointerError           = errors.New("is not a pointer")
	InvalidMemoryAddressError = errors.New("invalid memory address")
)

// 将所有数据类型转换为json进行输出
// 使用时请注意字段需要是公共的
func PrintlnJson(a interface{}) {
	jsonData, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		fmt.Printf("%+v\n", a)
		return
	}
	fmt.Println(string(jsonData))
}

// 检查传入的接口类型不是指针类型，CheckPointer返回相应错误
func CheckPointer(p interface{}) error {
	if p == nil {
		return InvalidMemoryAddressError
	}
	if reflect.TypeOf(p).Kind() != reflect.Ptr {
		return NotPointerError
	}
	if isEmpty(reflect.ValueOf(p)) {
		return NilPointerError
	}
	return nil
}

// 如果传入的接口类型为默认值或nil，IsEmpty返回true
func IsEmpty(o interface{}) bool {
	if o == nil {
		return true
	}
	return isEmpty(reflect.ValueOf(o))
}

func isEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return len(stringutil.Trim(value.String())) == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return value.Complex() == 0
	case reflect.Interface, reflect.Ptr, reflect.Func:
		return value.IsNil()
	case reflect.Array, reflect.Slice, reflect.Chan:
		return value.Len() == 0
	case reflect.Map:
		return len(value.MapKeys()) == 0
	default:
		return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
	}
}
