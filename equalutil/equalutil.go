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

// 通过反射对两个值进行深度相等判断
package equalutil

import (
	"reflect"
	"unsafe"
	"fmt"
)

// 用于防止重复比较
type comparison struct {
	v1, v2 unsafe.Pointer
	t      reflect.Type
}

// 使用反射类型深度相等判断，equal()函数是一个递归函数
func equal(v1, v2 reflect.Value, seen map[comparison]bool) bool {
	// 判断v1,v2本身是否是零值
	// 1.未包含任何数据：
	// var v reflect.Value
	// v.IsValid() 结果false
	// 2.包含一个 nil 指针：
	// var v = reflect.ValueOf(nil)
	// v.IsValid() 结果false
	// 3.其他:
	// 结果true
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	// 判断不同类型
	if v1.Type() != v2.Type() {
		return false
	}

	// 避免遇到的任何引用循环
	hard := func(k reflect.Kind) bool {
		switch k {
		case reflect.Map, reflect.Slice, reflect.Ptr, reflect.Interface:
			return true
		}
		return false
	}

	// 为了确保算法对于有环的数据结构也能正常退出，
	// 我们必须记录每次已经比较的变量，从而避免进入第二次的比较

	// CanAddr() 判断v1,v2值是否可寻址
	// 1.指针的 Elem() 可寻址
	// 2.切片的元素可寻址
	// 3.可寻址数组的元素可寻址
	// 4.可寻址结构体的字段可寻址，方法不可寻址
	if v1.CanAddr() && v2.CanAddr() && hard(v1.Kind()) {
		v1Ptr := unsafe.Pointer(v1.UnsafeAddr())
		v2Ptr := unsafe.Pointer(v2.UnsafeAddr())
		if uintptr(v1Ptr) > uintptr(v2Ptr) {
			// 规范化顺序以减少访问中的条目数
			// 假设不动的垃圾收集器
			v1Ptr, v2Ptr = v2Ptr, v1Ptr
		}

		t := v1.Type()
		c := comparison{v1Ptr, v2Ptr, t}
		fmt.Println(c)
		if seen[c] {
			return true
		}

		seen[c] = true
	}

	switch v1.Kind() {
	case reflect.Bool:
		return v1.Bool() == v2.Bool()
	case reflect.String:
		return v1.String() == v2.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v1.Uint() == v2.Uint()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	case reflect.Complex64, reflect.Complex128:
		return v1.Complex() == v2.Complex()
	case reflect.Chan, reflect.UnsafePointer:
		return v1.Pointer() == v2.Pointer()
	case reflect.Interface:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		return equal(v1.Elem(), v2.Elem(), seen)
	case reflect.Ptr:
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		return equal(v1.Elem(), v2.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if v1.Len() != v2.Len() {
			return false
		}
		for i := 0; i < v1.Len(); i++ {
			if !equal(v1.Index(i), v2.Index(i), seen) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i, n := 0, v1.NumField(); i < n; i++ {
			if !equal(v1.Field(i), v2.Field(i), seen) {
				return false
			}
		}
		return true
	case reflect.Map:
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			if !equal(v1.MapIndex(k), v2.MapIndex(k), seen) {
				return false
			}
		}
		return true
	case reflect.Func:
		if v1.IsNil() && v2.IsNil() {
			return true
		}
		return false
	}
	panic("kind error")
}

// Equal()函数可以对两个值进行深度相等判断，可以支持任意的数据类型
// Equal()与reflect.DeepEqual()基本一致，下面将说明不同之处
// 例如1：DeepEqual()将一个nil值的map和非nil值但是空的map视作不相等
// 例如2：DeepEqual()将一个nil值的slice和非nil但是空的slice也视作不相等
// 如果您不需要上述特性，推荐使用reflect.DeepEqual()进行相等判断
func Equal(x, y interface{}) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if v1.Type() != v2.Type() {
		return false
	}
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), make(map[comparison]bool))
}
