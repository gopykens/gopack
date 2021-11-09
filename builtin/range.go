/*
 Copyright 2021 The GoPlus Authors (goplus.org)
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

package builtin

type RangeIter struct {
	data Range
	idx  int
}

func (p *RangeIter) Next() (val float64, ok bool) {
	if p.idx < len(p.data) {
		val, ok = p.data[p.idx], true
		p.idx = p.idx + 1
	}
	return
}

type Range []float64

func NewRange(start, end, step float64) Range {
	var data Range
	for i := start; i < end; i = i + step {
		data = append(data, i)
	}
	return data
}

func (p Range) Gop_Enum() *RangeIter {
	return &RangeIter{data: p}
}
