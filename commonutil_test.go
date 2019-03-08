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
package utils

import (
	"fmt"
	"reflect"
	"testing"
	"github.com/liuchonglin/go-utils/fieldutil"
	"time"
)

func TestCheckPointer(t *testing.T) {
	type args struct {
		p interface{}
	}
	var a *args
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				p: &args{p: nil},
			},
			wantErr: false,
		}, {
			name: "nil",
			args: args{
				p: nil,
			},
			wantErr: true,
		}, {
			name: "notPointer",
			args: args{
				p: args{},
			},
			wantErr: true,
		}, {
			name: "pointer is nil",
			args: args{
				p: a,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckPointer(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("CheckPointer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type TestStruct struct {
	Name    string
	Age     int
	Address *string
	List    []string
}

func TestIsEmpty(t *testing.T) {
	type args struct {
		o interface{}
	}
	var a *args
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				o: "s",
			},
			want: false,
		}, {
			name: "nil",
			args: args{
				o: nil,
			},
			want: true,
		}, {
			name: "string is empty",
			args: args{
				o: " ",
			},
			want: true,
		}, {
			name: "string not empty",
			args: args{
				o: "s",
			},
			want: false,
		}, {
			name: "bool is empty",
			args: args{
				o: false,
			},
			want: true,
		}, {
			name: "bool not empty",
			args: args{
				o: true,
			},
			want: false,
		}, {
			name: "int is empty",
			args: args{
				o: int(0),
			},
			want: true,
		}, {
			name: "int not empty",
			args: args{
				o: int(1),
			},
			want: false,
		}, {
			name: "uint empty",
			args: args{
				o: uint(0),
			},
			want: true,
		}, {
			name: "uint not empty",
			args: args{
				o: uint(1),
			},
			want: false,
		}, {
			name: "float64 empty",
			args: args{
				o: float64(0),
			},
			want: true,
		}, {
			name: "float64 not empty",
			args: args{
				o: float64(1),
			},
			want: false,
		}, {
			name: "pointer empty",
			args: args{
				o: a,
			},
			want: true,
		}, {
			name: "pointer not empty",
			args: args{
				o: &args{},
			},
			want: false,
		}, {
			name: "slice empty",
			args: args{
				o: []string{},
			},
			want: true,
		}, {
			name: "slice not empty",
			args: args{
				o: []string{"s"},
			},
			want: false,
		}, {
			name: "array empty",
			args: args{
				o: [1]string{},
			},
			want: false,
		}, {
			name: "slice not empty",
			args: args{
				o: [1]string{"s"},
			},
			want: false,
		}, {
			name: "map empty",
			args: args{
				o: map[string]interface{}{},
			},
			want: true,
		}, {
			name: "map not empty",
			args: args{
				o: map[string]interface{}{"name": "a"},
			},
			want: false,
		}, {
			name: "struct empty1",
			args: args{
				o: TestStruct{},
			},
			want: true,
		}, {
			name: "struct empty2",
			args: args{
				o: TestStruct{Name: "", Age: 0, Address: nil, List: nil},
			},
			want: true,
		},{
			name: "struct not empty",
			args: args{
				o: TestStruct{Name: "xxx", Age: 0, Address: nil, List: nil},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.o); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray(t *testing.T) {
	a := [3]string{" "}
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a[0] == "")
}

func TestIsValid(t *testing.T) {
	var a int = 0
	v := reflect.ValueOf(a)
	fmt.Println(v.IsValid())

	var b *int = fieldutil.IntToPtr(0)
	v1 := reflect.ValueOf(b)
	fmt.Println(v1.IsValid())
}

type Student struct {
	Name string
}

func TestPrintlnJson(t *testing.T) {
	type student struct {
		Name      string
		Age       int
		Pay       float64
		Time      time.Time
		Address   *string
		List      []string
		FriendMap map[string]interface{}
	}

	s := student{
		Name:      "小明",
		Age:       10,
		Pay:       1000.22,
		Time:      time.Now(),
		Address:   fieldutil.StringToPtr("笑笑啊"),
		List:      []string{"a", "b", "c", "d"},
		FriendMap: map[string]interface{}{"小傻": student{Name: "小傻", Age: 29}, "小赵": student{Name: "小赵", Age: 20}},
	}
	type args struct {
		a interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				a: &s,
			},
		}, {
			name: "int",
			args: args{
				a: 1,
			},
		}, {
			name: "nil",
			args: args{
				a: nil,
			},
		}, {
			name: "array",
			args: args{
				a: []string{"a", "b", "c", "d"},
			},
		}, {
			name: "map",
			args: args{
				a: map[string]interface{}{"a": 1, "b": 2, "c": 3, "d": 4, "e": []string{"aa", "bb", "cc"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintlnJson(tt.args.a)
		})
	}
}
