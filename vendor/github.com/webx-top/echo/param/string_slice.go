/*

   Copyright 2016 Wenhui Shen <www.webx.top>

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

*/

package param

import (
	"strconv"
	"strings"
)

type StringSlice []string

func (p StringSlice) String() []string {
	return []string(p)
}

func (p StringSlice) Join(sep string) string {
	return strings.Join([]string(p), sep)
}

func (p StringSlice) Interface(filters ...func(int, string) bool) []interface{} {
	var filter func(int, string) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var ids []interface{}
	for idx, id := range p {
		if filter == nil || filter(idx, id) {
			ids = append(ids, interface{}(id))
		}
	}
	return ids
}

func (p StringSlice) Int(filters ...func(int, int) bool) []int {
	var filter func(int, int) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var ids []int
	for idx, id := range p {
		i, _ := strconv.Atoi(strings.TrimSpace(id))
		if filter == nil || filter(idx, i) {
			ids = append(ids, i)
		}
	}
	return ids
}

func (p StringSlice) Int64(filters ...func(int, int64) bool) []int64 {
	var filter func(int, int64) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var ids []int64
	for idx, id := range p {
		i, _ := strconv.ParseInt(strings.TrimSpace(id), 10, 64)
		if filter == nil || filter(idx, i) {
			ids = append(ids, i)
		}
	}
	return ids
}

func (p StringSlice) Int32(filters ...func(int, int32) bool) []int32 {
	var filter func(int, int32) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var ids []int32
	for idx, id := range p {
		i, _ := strconv.ParseInt(strings.TrimSpace(id), 10, 32)
		iv := int32(i)
		if filter == nil || filter(idx, iv) {
			ids = append(ids, iv)
		}
	}
	return ids
}

func (p StringSlice) Uint(filters ...func(int, uint) bool) []uint {
	var filter func(int, uint) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var ids []uint
	for idx, id := range p {
		i, _ := strconv.ParseUint(strings.TrimSpace(id), 10, 64)
		iv := uint(i)
		if filter == nil || filter(idx, iv) {
			ids = append(ids, iv)
		}
	}
	return ids
}

func (p StringSlice) Uint64(filters ...func(int, uint64) bool) []uint64 {
	var filter func(int, uint64) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var ids []uint64
	for idx, id := range p {
		i, _ := strconv.ParseUint(strings.TrimSpace(id), 10, 64)
		if filter == nil || filter(idx, i) {
			ids = append(ids, i)
		}
	}
	return ids
}

func (p StringSlice) Uint32(filters ...func(int, uint32) bool) []uint32 {
	var filter func(int, uint32) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var ids []uint32
	for idx, id := range p {
		i, _ := strconv.ParseUint(strings.TrimSpace(id), 10, 32)
		iv := uint32(i)
		if filter == nil || filter(idx, iv) {
			ids = append(ids, iv)
		}
	}
	return ids
}

func (p StringSlice) Float32(filters ...func(int, float32) bool) []float32 {
	var filter func(int, float32) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var values []float32
	for idx, v := range p {
		i, _ := strconv.ParseFloat(strings.TrimSpace(v), 32)
		iv := float32(i)
		if filter == nil || filter(idx, iv) {
			values = append(values, iv)
		}
	}
	return values
}

func (p StringSlice) Float64(filters ...func(int, float64) bool) []float64 {
	var filter func(int, float64) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var values []float64
	for idx, v := range p {
		i, _ := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if filter == nil || filter(idx, i) {
			values = append(values, i)
		}
	}
	return values
}

func (p StringSlice) Bool(filters ...func(int, bool) bool) []bool {
	var filter func(int, bool) bool
	if len(filters) > 0 {
		filter = filters[0]
	}
	var values []bool
	for idx, v := range p {
		i, e := strconv.ParseBool(strings.TrimSpace(v))
		if e != nil {
			continue
		}
		if filter == nil || filter(idx, i) {
			values = append(values, i)
		}
	}
	return values
}
