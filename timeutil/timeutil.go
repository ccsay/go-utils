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
	"fmt"
	"time"
	"github.com/liuchonglin/go-utils/stringutil"
)

const (
	FormatDay  = "2006-01-02"
	FormatTime = "2006-01-02 15:04:05"
)

// 时间转字符串
func TimeToString(t time.Time, format string) string {
	return t.Format(format)
}

// 字符串转时间
func StringToTime(timeStr string, format string) (time.Time, error) {
	timeStr = stringutil.Trim(timeStr)
	var t time.Time
	local, err := time.LoadLocation("Local")
	if err != nil {
		return t, fmt.Errorf("加载本地时区失败：%v", err)
	}

	t, err = time.ParseInLocation(format, timeStr, local)
	if err != nil {
		return t, fmt.Errorf("[string To Time] 失败：%v", err)
	}
	return t, nil
}

// 获取格式化后的当前时间
func GetCurrentTime(format string) string {
	return time.Now().Format(format)
}

// 加减时间
func AddSubTime(t time.Time, ts string) time.Time {
	//ParseDuration解析一个时间段字符串，如"300ms"、"-1.5h"、"2h45m"。
	//合法的单位有"ns"纳秒,"us"微秒,"µs"微秒、"ms"毫秒、"s"秒、"m"分钟、"h"小时。
	timePart, err := time.ParseDuration(ts)
	if err != nil {
		return t
	}
	return t.Add(timePart)
}
