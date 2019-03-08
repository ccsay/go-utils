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
	"testing"
)

func TestEqual(t *testing.T) {
	var arrayX []string
	var mapX map[string]interface{}
	var funcX func()
	type args struct {
		x interface{}
		y interface{}
	}

	var ptrX *args
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				x: 1,
				y: 1,
			},
			want: true,
		}, {
			name: "x nil",
			args: args{
				x: nil,
				y: 1,
			},
			want: false,
		}, {
			name: "y nil",
			args: args{
				x: 1,
				y: nil,
			},
			want: false,
		}, {
			name: "y,x nil",
			args: args{
				x: nil,
				y: nil,
			},
			want: true,
		}, {
			name: "type not same",
			args: args{
				x: 1,
				y: args{},
			},
			want: false,
		}, {
			name: "x,y nil",
			args: args{
				x: nil,
				y: nil,
			},
			want: true,
		}, {
			name: "x,y nil",
			args: args{
				x: nil,
				y: nil,
			},
			want: true,
		}, {
			name: "x,y bool",
			args: args{
				x: true,
				y: true,
			},
			want: true,
		}, {
			name: "x,y string",
			args: args{
				x: "s",
				y: "s",
			},
			want: true,
		}, {
			name: "x,y int",
			args: args{
				x: int(1),
				y: int(1),
			},
			want: true,
		}, {
			name: "x,y uint",
			args: args{
				x: uint(1),
				y: uint(1),
			},
			want: true,
		}, {
			name: "x,y float",
			args: args{
				x: float64(1),
				y: float64(1),
			},
			want: true,
		}, {
			name: "x,y float",
			args: args{
				x: 3.2 + 12i,
				y: 3.2 + 12i,
			},
			want: true,
		}, {
			name: "x,y chan",
			args: args{
				x: make(chan int),
				y: make(chan int),
			},
			want: false,
		}, {
			name: "x,y ptr",
			args: args{
				x: &args{},
				y: &args{},
			},
			want: true,
		}, {
			name: "x,y array ok",
			args: args{
				x: [1]string{"s"},
				y: [1]string{"s"},
			},
			want: true,
		}, {
			name: "x,y slice ok",
			args: args{
				x: []string{"s"},
				y: []string{"s"},
			},
			want: true,
		}, {
			name: "x,y array len not same",
			args: args{
				x: [2]string{"s", "11"},
				y: [2]string{"s"},
			},
			want: false,
		}, {
			name: "x,y slice not",
			args: args{
				x: []string{"s", "a"},
				y: []string{"s", "b"},
			},
			want: false,
		}, {
			name: "x slice nil,y slice empty",
			args: args{
				x: arrayX,
				y: []string{},
			},
			want: true,
		}, {
			name: "x,y struct not",
			args: args{
				x: args{
					x: "1",
					y: "2",
				},
				y: args{
					x: "2",
					y: "1",
				},
			},
			want: false,
		}, {
			name: "x,y map len not same",
			args: args{
				x: map[string]interface{}{"name": "s", "age": 1},
				y: map[string]interface{}{"name": "s"},
			},
			want: false,
		}, {
			name: "x,y map ptr same",
			args: args{
				x: mapX,
				y: mapX,
			},
			want: true,
		}, {
			name: "x,y map ok",
			args: args{
				x: map[string]interface{}{"name": "s"},
				y: map[string]interface{}{"name": "s"},
			},
			want: true,
		}, {
			name: "x,y map not",
			args: args{
				x: map[string]interface{}{"name": "s", "age": 1},
				y: map[string]interface{}{"name": "s", "age": 2},
			},
			want: false,
		}, {
			name: "x map nil,y map empty",
			args: args{
				x: mapX,
				y: map[string]interface{}{},
			},
			want: true,
		}, {
			name: "x,y func ok",
			args: args{
				x: funcX,
				y: funcX,
			},
			want: true,
		}, {
			name: "x,y func not",
			args: args{
				x: func() {},
				y: func() {},
			},
			want: false,
		}, {
			name: "x,y ptr same",
			args: args{
				x: ptrX,
				y: ptrX,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}