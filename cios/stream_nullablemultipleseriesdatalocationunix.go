/*
 * Collection utility of NullableMultipleSeriesDataLocationUnix Struct
 *
 * Generated by: Go Streamer
 */

package cios

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type NullableMultipleSeriesDataLocationUnixStream []NullableMultipleSeriesDataLocationUnix

func NullableMultipleSeriesDataLocationUnixStreamOf(arg ...NullableMultipleSeriesDataLocationUnix) NullableMultipleSeriesDataLocationUnixStream {
	return arg
}
func NullableMultipleSeriesDataLocationUnixStreamFrom(arg []NullableMultipleSeriesDataLocationUnix) NullableMultipleSeriesDataLocationUnixStream {
	return arg
}
func CreateNullableMultipleSeriesDataLocationUnixStream(arg ...NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	tmp := NullableMultipleSeriesDataLocationUnixStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleSeriesDataLocationUnixStream(arg []NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	tmp := NullableMultipleSeriesDataLocationUnixStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleSeriesDataLocationUnixStream) Add(arg NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) AddAll(arg ...NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) AddSafe(arg *NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Aggregate(fn func(NullableMultipleSeriesDataLocationUnix, NullableMultipleSeriesDataLocationUnix) NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	result := NullableMultipleSeriesDataLocationUnixStreamOf()
	self.ForEach(func(v NullableMultipleSeriesDataLocationUnix, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleSeriesDataLocationUnix{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) AllMatch(fn func(NullableMultipleSeriesDataLocationUnix, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleSeriesDataLocationUnixStream) AnyMatch(fn func(NullableMultipleSeriesDataLocationUnix, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Clone() *NullableMultipleSeriesDataLocationUnixStream {
	temp := make([]NullableMultipleSeriesDataLocationUnix, self.Len())
	copy(temp, *self)
	return (*NullableMultipleSeriesDataLocationUnixStream)(&temp)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Copy() *NullableMultipleSeriesDataLocationUnixStream {
	return self.Clone()
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Concat(arg []NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Contains(arg NullableMultipleSeriesDataLocationUnix) bool {
	return self.FindIndex(func(_arg NullableMultipleSeriesDataLocationUnix, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Clean() *NullableMultipleSeriesDataLocationUnixStream {
	*self = NullableMultipleSeriesDataLocationUnixStreamOf()
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Delete(index int) *NullableMultipleSeriesDataLocationUnixStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) DeleteRange(startIndex, endIndex int) *NullableMultipleSeriesDataLocationUnixStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Distinct() *NullableMultipleSeriesDataLocationUnixStream {
	caches := map[string]bool{}
	result := NullableMultipleSeriesDataLocationUnixStreamOf()
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if !f {
				result = append(result, v)
			}
		} else if caches[key] = true; !f {
			result = append(result, v)
		}

	}
	*self = result
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Each(fn func(NullableMultipleSeriesDataLocationUnix)) *NullableMultipleSeriesDataLocationUnixStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) EachRight(fn func(NullableMultipleSeriesDataLocationUnix)) *NullableMultipleSeriesDataLocationUnixStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Equals(arr []NullableMultipleSeriesDataLocationUnix) bool {
	if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
		return false
	}
	for i := range *self {
		if !reflect.DeepEqual((*self)[i], arr[i]) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Filter(fn func(NullableMultipleSeriesDataLocationUnix, int) bool) *NullableMultipleSeriesDataLocationUnixStream {
	result := NullableMultipleSeriesDataLocationUnixStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) FilterSlim(fn func(NullableMultipleSeriesDataLocationUnix, int) bool) *NullableMultipleSeriesDataLocationUnixStream {
	result := NullableMultipleSeriesDataLocationUnixStreamOf()
	caches := map[string]bool{}
	for i, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v, i); caches[key] {
			result.Add(v)

		}
	}
	*self = result
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Find(fn func(NullableMultipleSeriesDataLocationUnix, int) bool) *NullableMultipleSeriesDataLocationUnix {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleSeriesDataLocationUnixStream) FindOr(fn func(NullableMultipleSeriesDataLocationUnix, int) bool, or NullableMultipleSeriesDataLocationUnix) NullableMultipleSeriesDataLocationUnix {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleSeriesDataLocationUnixStream) FindIndex(fn func(NullableMultipleSeriesDataLocationUnix, int) bool) int {
	if self == nil {
		return -1
	}
	for i, v := range *self {
		if fn(v, i) {
			return i
		}
	}
	return -1
}
func (self *NullableMultipleSeriesDataLocationUnixStream) First() *NullableMultipleSeriesDataLocationUnix {
	return self.Get(0)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) FirstOr(arg NullableMultipleSeriesDataLocationUnix) NullableMultipleSeriesDataLocationUnix {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ForEach(fn func(NullableMultipleSeriesDataLocationUnix, int)) *NullableMultipleSeriesDataLocationUnixStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ForEachRight(fn func(NullableMultipleSeriesDataLocationUnix, int)) *NullableMultipleSeriesDataLocationUnixStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) GroupBy(fn func(NullableMultipleSeriesDataLocationUnix, int) string) map[string][]NullableMultipleSeriesDataLocationUnix {
	m := map[string][]NullableMultipleSeriesDataLocationUnix{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleSeriesDataLocationUnixStream) GroupByValues(fn func(NullableMultipleSeriesDataLocationUnix, int) string) [][]NullableMultipleSeriesDataLocationUnix {
	var tmp [][]NullableMultipleSeriesDataLocationUnix
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleSeriesDataLocationUnixStream) IndexOf(arg NullableMultipleSeriesDataLocationUnix) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleSeriesDataLocationUnixStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleSeriesDataLocationUnixStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Last() *NullableMultipleSeriesDataLocationUnix {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) LastOr(arg NullableMultipleSeriesDataLocationUnix) NullableMultipleSeriesDataLocationUnix {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Limit(limit int) *NullableMultipleSeriesDataLocationUnixStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleSeriesDataLocationUnixStream) Map(fn func(NullableMultipleSeriesDataLocationUnix, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Map2Int(fn func(NullableMultipleSeriesDataLocationUnix, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Map2Int32(fn func(NullableMultipleSeriesDataLocationUnix, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Map2Int64(fn func(NullableMultipleSeriesDataLocationUnix, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Map2Float32(fn func(NullableMultipleSeriesDataLocationUnix, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Map2Float64(fn func(NullableMultipleSeriesDataLocationUnix, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Map2Bool(fn func(NullableMultipleSeriesDataLocationUnix, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Map2Bytes(fn func(NullableMultipleSeriesDataLocationUnix, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Map2String(fn func(NullableMultipleSeriesDataLocationUnix, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Max(fn func(NullableMultipleSeriesDataLocationUnix, int) float64) *NullableMultipleSeriesDataLocationUnix {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Max(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Min(fn func(NullableMultipleSeriesDataLocationUnix, int) float64) *NullableMultipleSeriesDataLocationUnix {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Min(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) NoneMatch(fn func(NullableMultipleSeriesDataLocationUnix, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Get(index int) *NullableMultipleSeriesDataLocationUnix {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleSeriesDataLocationUnixStream) GetOr(index int, arg NullableMultipleSeriesDataLocationUnix) NullableMultipleSeriesDataLocationUnix {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Peek(fn func(*NullableMultipleSeriesDataLocationUnix, int)) *NullableMultipleSeriesDataLocationUnixStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMultipleSeriesDataLocationUnixStream) Reduce(fn func(NullableMultipleSeriesDataLocationUnix, NullableMultipleSeriesDataLocationUnix, int) NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	return self.ReduceInit(fn, NullableMultipleSeriesDataLocationUnix{})
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ReduceInit(fn func(NullableMultipleSeriesDataLocationUnix, NullableMultipleSeriesDataLocationUnix, int) NullableMultipleSeriesDataLocationUnix, initialValue NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	result := NullableMultipleSeriesDataLocationUnixStreamOf()
	self.ForEach(func(v NullableMultipleSeriesDataLocationUnix, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ReduceInterface(fn func(interface{}, NullableMultipleSeriesDataLocationUnix, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleSeriesDataLocationUnix{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ReduceString(fn func(string, NullableMultipleSeriesDataLocationUnix, int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ReduceInt(fn func(int, NullableMultipleSeriesDataLocationUnix, int) int) []int {
	result := []int{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ReduceInt32(fn func(int32, NullableMultipleSeriesDataLocationUnix, int) int32) []int32 {
	result := []int32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ReduceInt64(fn func(int64, NullableMultipleSeriesDataLocationUnix, int) int64) []int64 {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ReduceFloat32(fn func(float32, NullableMultipleSeriesDataLocationUnix, int) float32) []float32 {
	result := []float32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ReduceFloat64(fn func(float64, NullableMultipleSeriesDataLocationUnix, int) float64) []float64 {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ReduceBool(fn func(bool, NullableMultipleSeriesDataLocationUnix, int) bool) []bool {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Reverse() *NullableMultipleSeriesDataLocationUnixStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Replace(fn func(NullableMultipleSeriesDataLocationUnix, int) NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	return self.ForEach(func(v NullableMultipleSeriesDataLocationUnix, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Select(fn func(NullableMultipleSeriesDataLocationUnix) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Set(index int, val NullableMultipleSeriesDataLocationUnix) *NullableMultipleSeriesDataLocationUnixStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Skip(skip int) *NullableMultipleSeriesDataLocationUnixStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) SkippingEach(fn func(NullableMultipleSeriesDataLocationUnix, int) int) *NullableMultipleSeriesDataLocationUnixStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Slice(startIndex, n int) *NullableMultipleSeriesDataLocationUnixStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleSeriesDataLocationUnix{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Sort(fn func(i, j int) bool) *NullableMultipleSeriesDataLocationUnixStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleSeriesDataLocationUnixStream) Tail() *NullableMultipleSeriesDataLocationUnix {
	return self.Last()
}
func (self *NullableMultipleSeriesDataLocationUnixStream) TailOr(arg NullableMultipleSeriesDataLocationUnix) NullableMultipleSeriesDataLocationUnix {
	return self.LastOr(arg)
}
func (self *NullableMultipleSeriesDataLocationUnixStream) ToList() []NullableMultipleSeriesDataLocationUnix {
	return self.Val()
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Unique() *NullableMultipleSeriesDataLocationUnixStream {
	return self.Distinct()
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Val() []NullableMultipleSeriesDataLocationUnix {
	if self == nil {
		return []NullableMultipleSeriesDataLocationUnix{}
	}
	return *self.Copy()
}
func (self *NullableMultipleSeriesDataLocationUnixStream) While(fn func(NullableMultipleSeriesDataLocationUnix, int) bool) *NullableMultipleSeriesDataLocationUnixStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) Where(fn func(NullableMultipleSeriesDataLocationUnix) bool) *NullableMultipleSeriesDataLocationUnixStream {
	result := NullableMultipleSeriesDataLocationUnixStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleSeriesDataLocationUnixStream) WhereSlim(fn func(NullableMultipleSeriesDataLocationUnix) bool) *NullableMultipleSeriesDataLocationUnixStream {
	result := NullableMultipleSeriesDataLocationUnixStreamOf()
	caches := map[string]bool{}
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v); caches[key] {
			result.Add(v)
		}
	}
	*self = result
	return self
}
