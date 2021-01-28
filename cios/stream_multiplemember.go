/*
 * Collection utility of MultipleMember Struct
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

type MultipleMemberStream []MultipleMember

func MultipleMemberStreamOf(arg ...MultipleMember) MultipleMemberStream {
	return arg
}
func MultipleMemberStreamFrom(arg []MultipleMember) MultipleMemberStream {
	return arg
}
func CreateMultipleMemberStream(arg ...MultipleMember) *MultipleMemberStream {
	tmp := MultipleMemberStreamOf(arg...)
	return &tmp
}
func GenerateMultipleMemberStream(arg []MultipleMember) *MultipleMemberStream {
	tmp := MultipleMemberStreamFrom(arg)
	return &tmp
}

func (self *MultipleMemberStream) Add(arg MultipleMember) *MultipleMemberStream {
	return self.AddAll(arg)
}
func (self *MultipleMemberStream) AddAll(arg ...MultipleMember) *MultipleMemberStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleMemberStream) AddSafe(arg *MultipleMember) *MultipleMemberStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleMemberStream) Aggregate(fn func(MultipleMember, MultipleMember) MultipleMember) *MultipleMemberStream {
	result := MultipleMemberStreamOf()
	self.ForEach(func(v MultipleMember, i int) {
		if i == 0 {
			result.Add(fn(MultipleMember{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleMemberStream) AllMatch(fn func(MultipleMember, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleMemberStream) AnyMatch(fn func(MultipleMember, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleMemberStream) Clone() *MultipleMemberStream {
	temp := make([]MultipleMember, self.Len())
	copy(temp, *self)
	return (*MultipleMemberStream)(&temp)
}
func (self *MultipleMemberStream) Copy() *MultipleMemberStream {
	return self.Clone()
}
func (self *MultipleMemberStream) Concat(arg []MultipleMember) *MultipleMemberStream {
	return self.AddAll(arg...)
}
func (self *MultipleMemberStream) Contains(arg MultipleMember) bool {
	return self.FindIndex(func(_arg MultipleMember, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleMemberStream) Clean() *MultipleMemberStream {
	*self = MultipleMemberStreamOf()
	return self
}
func (self *MultipleMemberStream) Delete(index int) *MultipleMemberStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleMemberStream) DeleteRange(startIndex, endIndex int) *MultipleMemberStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleMemberStream) Distinct() *MultipleMemberStream {
	caches := map[string]bool{}
	result := MultipleMemberStreamOf()
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
func (self *MultipleMemberStream) Each(fn func(MultipleMember)) *MultipleMemberStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleMemberStream) EachRight(fn func(MultipleMember)) *MultipleMemberStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleMemberStream) Equals(arr []MultipleMember) bool {
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
func (self *MultipleMemberStream) Filter(fn func(MultipleMember, int) bool) *MultipleMemberStream {
	result := MultipleMemberStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleMemberStream) FilterSlim(fn func(MultipleMember, int) bool) *MultipleMemberStream {
	result := MultipleMemberStreamOf()
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
func (self *MultipleMemberStream) Find(fn func(MultipleMember, int) bool) *MultipleMember {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleMemberStream) FindOr(fn func(MultipleMember, int) bool, or MultipleMember) MultipleMember {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleMemberStream) FindIndex(fn func(MultipleMember, int) bool) int {
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
func (self *MultipleMemberStream) First() *MultipleMember {
	return self.Get(0)
}
func (self *MultipleMemberStream) FirstOr(arg MultipleMember) MultipleMember {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleMemberStream) ForEach(fn func(MultipleMember, int)) *MultipleMemberStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleMemberStream) ForEachRight(fn func(MultipleMember, int)) *MultipleMemberStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleMemberStream) GroupBy(fn func(MultipleMember, int) string) map[string][]MultipleMember {
	m := map[string][]MultipleMember{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleMemberStream) GroupByValues(fn func(MultipleMember, int) string) [][]MultipleMember {
	var tmp [][]MultipleMember
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleMemberStream) IndexOf(arg MultipleMember) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleMemberStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleMemberStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleMemberStream) Last() *MultipleMember {
	return self.Get(self.Len() - 1)
}
func (self *MultipleMemberStream) LastOr(arg MultipleMember) MultipleMember {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleMemberStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleMemberStream) Limit(limit int) *MultipleMemberStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleMemberStream) Map(fn func(MultipleMember, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberStream) Map2Int(fn func(MultipleMember, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberStream) Map2Int32(fn func(MultipleMember, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberStream) Map2Int64(fn func(MultipleMember, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberStream) Map2Float32(fn func(MultipleMember, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberStream) Map2Float64(fn func(MultipleMember, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberStream) Map2Bool(fn func(MultipleMember, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberStream) Map2Bytes(fn func(MultipleMember, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberStream) Map2String(fn func(MultipleMember, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberStream) Max(fn func(MultipleMember, int) float64) *MultipleMember {
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
func (self *MultipleMemberStream) Min(fn func(MultipleMember, int) float64) *MultipleMember {
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
func (self *MultipleMemberStream) NoneMatch(fn func(MultipleMember, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleMemberStream) Get(index int) *MultipleMember {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleMemberStream) GetOr(index int, arg MultipleMember) MultipleMember {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleMemberStream) Peek(fn func(*MultipleMember, int)) *MultipleMemberStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MultipleMemberStream) Reduce(fn func(MultipleMember, MultipleMember, int) MultipleMember) *MultipleMemberStream {
	return self.ReduceInit(fn, MultipleMember{})
}
func (self *MultipleMemberStream) ReduceInit(fn func(MultipleMember, MultipleMember, int) MultipleMember, initialValue MultipleMember) *MultipleMemberStream {
	result := MultipleMemberStreamOf()
	self.ForEach(func(v MultipleMember, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleMemberStream) ReduceInterface(fn func(interface{}, MultipleMember, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleMember{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleMemberStream) ReduceString(fn func(string, MultipleMember, int) string) []string {
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
func (self *MultipleMemberStream) ReduceInt(fn func(int, MultipleMember, int) int) []int {
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
func (self *MultipleMemberStream) ReduceInt32(fn func(int32, MultipleMember, int) int32) []int32 {
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
func (self *MultipleMemberStream) ReduceInt64(fn func(int64, MultipleMember, int) int64) []int64 {
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
func (self *MultipleMemberStream) ReduceFloat32(fn func(float32, MultipleMember, int) float32) []float32 {
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
func (self *MultipleMemberStream) ReduceFloat64(fn func(float64, MultipleMember, int) float64) []float64 {
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
func (self *MultipleMemberStream) ReduceBool(fn func(bool, MultipleMember, int) bool) []bool {
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
func (self *MultipleMemberStream) Reverse() *MultipleMemberStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleMemberStream) Replace(fn func(MultipleMember, int) MultipleMember) *MultipleMemberStream {
	return self.ForEach(func(v MultipleMember, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleMemberStream) Select(fn func(MultipleMember) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleMemberStream) Set(index int, val MultipleMember) *MultipleMemberStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleMemberStream) Skip(skip int) *MultipleMemberStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleMemberStream) SkippingEach(fn func(MultipleMember, int) int) *MultipleMemberStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleMemberStream) Slice(startIndex, n int) *MultipleMemberStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleMember{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleMemberStream) Sort(fn func(i, j int) bool) *MultipleMemberStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleMemberStream) Tail() *MultipleMember {
	return self.Last()
}
func (self *MultipleMemberStream) TailOr(arg MultipleMember) MultipleMember {
	return self.LastOr(arg)
}
func (self *MultipleMemberStream) ToList() []MultipleMember {
	return self.Val()
}
func (self *MultipleMemberStream) Unique() *MultipleMemberStream {
	return self.Distinct()
}
func (self *MultipleMemberStream) Val() []MultipleMember {
	if self == nil {
		return []MultipleMember{}
	}
	return *self.Copy()
}
func (self *MultipleMemberStream) While(fn func(MultipleMember, int) bool) *MultipleMemberStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleMemberStream) Where(fn func(MultipleMember) bool) *MultipleMemberStream {
	result := MultipleMemberStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleMemberStream) WhereSlim(fn func(MultipleMember) bool) *MultipleMemberStream {
	result := MultipleMemberStreamOf()
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
