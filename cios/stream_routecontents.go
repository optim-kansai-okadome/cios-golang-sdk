/*
 * Collection utility of RouteContents Struct
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

type RouteContentsStream []RouteContents

func RouteContentsStreamOf(arg ...RouteContents) RouteContentsStream {
	return arg
}
func RouteContentsStreamFrom(arg []RouteContents) RouteContentsStream {
	return arg
}
func CreateRouteContentsStream(arg ...RouteContents) *RouteContentsStream {
	tmp := RouteContentsStreamOf(arg...)
	return &tmp
}
func GenerateRouteContentsStream(arg []RouteContents) *RouteContentsStream {
	tmp := RouteContentsStreamFrom(arg)
	return &tmp
}

func (self *RouteContentsStream) Add(arg RouteContents) *RouteContentsStream {
	return self.AddAll(arg)
}
func (self *RouteContentsStream) AddAll(arg ...RouteContents) *RouteContentsStream {
	*self = append(*self, arg...)
	return self
}
func (self *RouteContentsStream) AddSafe(arg *RouteContents) *RouteContentsStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *RouteContentsStream) Aggregate(fn func(RouteContents, RouteContents) RouteContents) *RouteContentsStream {
	result := RouteContentsStreamOf()
	self.ForEach(func(v RouteContents, i int) {
		if i == 0 {
			result.Add(fn(RouteContents{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *RouteContentsStream) AllMatch(fn func(RouteContents, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *RouteContentsStream) AnyMatch(fn func(RouteContents, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *RouteContentsStream) Clone() *RouteContentsStream {
	temp := make([]RouteContents, self.Len())
	copy(temp, *self)
	return (*RouteContentsStream)(&temp)
}
func (self *RouteContentsStream) Copy() *RouteContentsStream {
	return self.Clone()
}
func (self *RouteContentsStream) Concat(arg []RouteContents) *RouteContentsStream {
	return self.AddAll(arg...)
}
func (self *RouteContentsStream) Contains(arg RouteContents) bool {
	return self.FindIndex(func(_arg RouteContents, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *RouteContentsStream) Clean() *RouteContentsStream {
	*self = RouteContentsStreamOf()
	return self
}
func (self *RouteContentsStream) Delete(index int) *RouteContentsStream {
	return self.DeleteRange(index, index)
}
func (self *RouteContentsStream) DeleteRange(startIndex, endIndex int) *RouteContentsStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *RouteContentsStream) Distinct() *RouteContentsStream {
	caches := map[string]bool{}
	result := RouteContentsStreamOf()
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
func (self *RouteContentsStream) Each(fn func(RouteContents)) *RouteContentsStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *RouteContentsStream) EachRight(fn func(RouteContents)) *RouteContentsStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *RouteContentsStream) Equals(arr []RouteContents) bool {
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
func (self *RouteContentsStream) Filter(fn func(RouteContents, int) bool) *RouteContentsStream {
	result := RouteContentsStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *RouteContentsStream) FilterSlim(fn func(RouteContents, int) bool) *RouteContentsStream {
	result := RouteContentsStreamOf()
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
func (self *RouteContentsStream) Find(fn func(RouteContents, int) bool) *RouteContents {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *RouteContentsStream) FindOr(fn func(RouteContents, int) bool, or RouteContents) RouteContents {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *RouteContentsStream) FindIndex(fn func(RouteContents, int) bool) int {
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
func (self *RouteContentsStream) First() *RouteContents {
	return self.Get(0)
}
func (self *RouteContentsStream) FirstOr(arg RouteContents) RouteContents {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *RouteContentsStream) ForEach(fn func(RouteContents, int)) *RouteContentsStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *RouteContentsStream) ForEachRight(fn func(RouteContents, int)) *RouteContentsStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *RouteContentsStream) GroupBy(fn func(RouteContents, int) string) map[string][]RouteContents {
	m := map[string][]RouteContents{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *RouteContentsStream) GroupByValues(fn func(RouteContents, int) string) [][]RouteContents {
	var tmp [][]RouteContents
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *RouteContentsStream) IndexOf(arg RouteContents) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *RouteContentsStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *RouteContentsStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *RouteContentsStream) Last() *RouteContents {
	return self.Get(self.Len() - 1)
}
func (self *RouteContentsStream) LastOr(arg RouteContents) RouteContents {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *RouteContentsStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *RouteContentsStream) Limit(limit int) *RouteContentsStream {
	self.Slice(0, limit)
	return self
}

func (self *RouteContentsStream) Map(fn func(RouteContents, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteContentsStream) Map2Int(fn func(RouteContents, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteContentsStream) Map2Int32(fn func(RouteContents, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteContentsStream) Map2Int64(fn func(RouteContents, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteContentsStream) Map2Float32(fn func(RouteContents, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteContentsStream) Map2Float64(fn func(RouteContents, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteContentsStream) Map2Bool(fn func(RouteContents, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteContentsStream) Map2Bytes(fn func(RouteContents, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteContentsStream) Map2String(fn func(RouteContents, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteContentsStream) Max(fn func(RouteContents, int) float64) *RouteContents {
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
func (self *RouteContentsStream) Min(fn func(RouteContents, int) float64) *RouteContents {
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
func (self *RouteContentsStream) NoneMatch(fn func(RouteContents, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *RouteContentsStream) Get(index int) *RouteContents {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *RouteContentsStream) GetOr(index int, arg RouteContents) RouteContents {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *RouteContentsStream) Peek(fn func(*RouteContents, int)) *RouteContentsStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *RouteContentsStream) Reduce(fn func(RouteContents, RouteContents, int) RouteContents) *RouteContentsStream {
	return self.ReduceInit(fn, RouteContents{})
}
func (self *RouteContentsStream) ReduceInit(fn func(RouteContents, RouteContents, int) RouteContents, initialValue RouteContents) *RouteContentsStream {
	result := RouteContentsStreamOf()
	self.ForEach(func(v RouteContents, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *RouteContentsStream) ReduceInterface(fn func(interface{}, RouteContents, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(RouteContents{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *RouteContentsStream) ReduceString(fn func(string, RouteContents, int) string) []string {
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
func (self *RouteContentsStream) ReduceInt(fn func(int, RouteContents, int) int) []int {
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
func (self *RouteContentsStream) ReduceInt32(fn func(int32, RouteContents, int) int32) []int32 {
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
func (self *RouteContentsStream) ReduceInt64(fn func(int64, RouteContents, int) int64) []int64 {
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
func (self *RouteContentsStream) ReduceFloat32(fn func(float32, RouteContents, int) float32) []float32 {
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
func (self *RouteContentsStream) ReduceFloat64(fn func(float64, RouteContents, int) float64) []float64 {
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
func (self *RouteContentsStream) ReduceBool(fn func(bool, RouteContents, int) bool) []bool {
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
func (self *RouteContentsStream) Reverse() *RouteContentsStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *RouteContentsStream) Replace(fn func(RouteContents, int) RouteContents) *RouteContentsStream {
	return self.ForEach(func(v RouteContents, i int) { self.Set(i, fn(v, i)) })
}
func (self *RouteContentsStream) Select(fn func(RouteContents) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *RouteContentsStream) Set(index int, val RouteContents) *RouteContentsStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *RouteContentsStream) Skip(skip int) *RouteContentsStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *RouteContentsStream) SkippingEach(fn func(RouteContents, int) int) *RouteContentsStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *RouteContentsStream) Slice(startIndex, n int) *RouteContentsStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []RouteContents{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *RouteContentsStream) Sort(fn func(i, j int) bool) *RouteContentsStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *RouteContentsStream) Tail() *RouteContents {
	return self.Last()
}
func (self *RouteContentsStream) TailOr(arg RouteContents) RouteContents {
	return self.LastOr(arg)
}
func (self *RouteContentsStream) ToList() []RouteContents {
	return self.Val()
}
func (self *RouteContentsStream) Unique() *RouteContentsStream {
	return self.Distinct()
}
func (self *RouteContentsStream) Val() []RouteContents {
	if self == nil {
		return []RouteContents{}
	}
	return *self.Copy()
}
func (self *RouteContentsStream) While(fn func(RouteContents, int) bool) *RouteContentsStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *RouteContentsStream) Where(fn func(RouteContents) bool) *RouteContentsStream {
	result := RouteContentsStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *RouteContentsStream) WhereSlim(fn func(RouteContents) bool) *RouteContentsStream {
	result := RouteContentsStreamOf()
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
