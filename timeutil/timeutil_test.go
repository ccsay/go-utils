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

// 日期与时间工具
package timeutil

import (
	"reflect"
	"testing"
	"time"
	"fmt"
)

func TestTimeToString(t *testing.T) {
	type args struct {
		t      time.Time
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"ok",
			args:args{
				t:time.Now(),
				format:NoDivFormatTime,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeToString(tt.args.t, tt.args.format)
			fmt.Println(got)
		})
	}
}

func TestStringToTime(t *testing.T) {
	type args struct {
		timeStr string
		format  string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToTime(tt.args.timeStr, tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrentTime(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentTime(tt.args.format); got != tt.want {
				t.Errorf("GetCurrentTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddSubTime(t *testing.T) {
	type args struct {
		t  time.Time
		ts string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddSubTime(tt.args.t, tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSubTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
