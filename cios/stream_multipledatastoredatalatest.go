/*
 * Collection utility of MultipleDataStoreDataLatest Struct
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

type MultipleDataStoreDataLatestStream []MultipleDataStoreDataLatest

func MultipleDataStoreDataLatestStreamOf(arg ...MultipleDataStoreDataLatest) MultipleDataStoreDataLatestStream {
	return arg
}
func MultipleDataStoreDataLatestStreamFrom(arg []MultipleDataStoreDataLatest) MultipleDataStoreDataLatestStream {
	return arg
}
func CreateMultipleDataStoreDataLatestStream(arg ...MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	tmp := MultipleDataStoreDataLatestStreamOf(arg...)
	return &tmp
}
func GenerateMultipleDataStoreDataLatestStream(arg []MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	tmp := MultipleDataStoreDataLatestStreamFrom(arg)
	return &tmp
}

func (self *MultipleDataStoreDataLatestStream) Add(arg MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	return self.AddAll(arg)
}
func (self *MultipleDataStoreDataLatestStream) AddAll(arg ...MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleDataStoreDataLatestStream) AddSafe(arg *MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) Aggregate(fn func(MultipleDataStoreDataLatest, MultipleDataStoreDataLatest) MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	result := MultipleDataStoreDataLatestStreamOf()
	self.ForEach(func(v MultipleDataStoreDataLatest, i int) {
		if i == 0 {
			result.Add(fn(MultipleDataStoreDataLatest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleDataStoreDataLatestStream) AllMatch(fn func(MultipleDataStoreDataLatest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleDataStoreDataLatestStream) AnyMatch(fn func(MultipleDataStoreDataLatest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleDataStoreDataLatestStream) Clone() *MultipleDataStoreDataLatestStream {
	temp := make([]MultipleDataStoreDataLatest, self.Len())
	copy(temp, *self)
	return (*MultipleDataStoreDataLatestStream)(&temp)
}
func (self *MultipleDataStoreDataLatestStream) Copy() *MultipleDataStoreDataLatestStream {
	return self.Clone()
}
func (self *MultipleDataStoreDataLatestStream) Concat(arg []MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	return self.AddAll(arg...)
}
func (self *MultipleDataStoreDataLatestStream) Contains(arg MultipleDataStoreDataLatest) bool {
	return self.FindIndex(func(_arg MultipleDataStoreDataLatest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleDataStoreDataLatestStream) Clean() *MultipleDataStoreDataLatestStream {
	*self = MultipleDataStoreDataLatestStreamOf()
	return self
}
func (self *MultipleDataStoreDataLatestStream) Delete(index int) *MultipleDataStoreDataLatestStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleDataStoreDataLatestStream) DeleteRange(startIndex, endIndex int) *MultipleDataStoreDataLatestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleDataStoreDataLatestStream) Distinct() *MultipleDataStoreDataLatestStream {
	caches := map[string]bool{}
	result := MultipleDataStoreDataLatestStreamOf()
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
func (self *MultipleDataStoreDataLatestStream) Each(fn func(MultipleDataStoreDataLatest)) *MultipleDataStoreDataLatestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) EachRight(fn func(MultipleDataStoreDataLatest)) *MultipleDataStoreDataLatestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) Equals(arr []MultipleDataStoreDataLatest) bool {
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
func (self *MultipleDataStoreDataLatestStream) Filter(fn func(MultipleDataStoreDataLatest, int) bool) *MultipleDataStoreDataLatestStream {
	result := MultipleDataStoreDataLatestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDataStoreDataLatestStream) FilterSlim(fn func(MultipleDataStoreDataLatest, int) bool) *MultipleDataStoreDataLatestStream {
	result := MultipleDataStoreDataLatestStreamOf()
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
func (self *MultipleDataStoreDataLatestStream) Find(fn func(MultipleDataStoreDataLatest, int) bool) *MultipleDataStoreDataLatest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleDataStoreDataLatestStream) FindOr(fn func(MultipleDataStoreDataLatest, int) bool, or MultipleDataStoreDataLatest) MultipleDataStoreDataLatest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleDataStoreDataLatestStream) FindIndex(fn func(MultipleDataStoreDataLatest, int) bool) int {
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
func (self *MultipleDataStoreDataLatestStream) First() *MultipleDataStoreDataLatest {
	return self.Get(0)
}
func (self *MultipleDataStoreDataLatestStream) FirstOr(arg MultipleDataStoreDataLatest) MultipleDataStoreDataLatest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDataStoreDataLatestStream) ForEach(fn func(MultipleDataStoreDataLatest, int)) *MultipleDataStoreDataLatestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) ForEachRight(fn func(MultipleDataStoreDataLatest, int)) *MultipleDataStoreDataLatestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) GroupBy(fn func(MultipleDataStoreDataLatest, int) string) map[string][]MultipleDataStoreDataLatest {
	m := map[string][]MultipleDataStoreDataLatest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleDataStoreDataLatestStream) GroupByValues(fn func(MultipleDataStoreDataLatest, int) string) [][]MultipleDataStoreDataLatest {
	var tmp [][]MultipleDataStoreDataLatest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleDataStoreDataLatestStream) IndexOf(arg MultipleDataStoreDataLatest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleDataStoreDataLatestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleDataStoreDataLatestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleDataStoreDataLatestStream) Last() *MultipleDataStoreDataLatest {
	return self.Get(self.Len() - 1)
}
func (self *MultipleDataStoreDataLatestStream) LastOr(arg MultipleDataStoreDataLatest) MultipleDataStoreDataLatest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDataStoreDataLatestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleDataStoreDataLatestStream) Limit(limit int) *MultipleDataStoreDataLatestStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleDataStoreDataLatestStream) Map(fn func(MultipleDataStoreDataLatest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Map2Int(fn func(MultipleDataStoreDataLatest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Map2Int32(fn func(MultipleDataStoreDataLatest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Map2Int64(fn func(MultipleDataStoreDataLatest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Map2Float32(fn func(MultipleDataStoreDataLatest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Map2Float64(fn func(MultipleDataStoreDataLatest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Map2Bool(fn func(MultipleDataStoreDataLatest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Map2Bytes(fn func(MultipleDataStoreDataLatest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Map2String(fn func(MultipleDataStoreDataLatest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Max(fn func(MultipleDataStoreDataLatest, int) float64) *MultipleDataStoreDataLatest {
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
func (self *MultipleDataStoreDataLatestStream) Min(fn func(MultipleDataStoreDataLatest, int) float64) *MultipleDataStoreDataLatest {
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
func (self *MultipleDataStoreDataLatestStream) NoneMatch(fn func(MultipleDataStoreDataLatest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleDataStoreDataLatestStream) Get(index int) *MultipleDataStoreDataLatest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleDataStoreDataLatestStream) GetOr(index int, arg MultipleDataStoreDataLatest) MultipleDataStoreDataLatest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDataStoreDataLatestStream) Peek(fn func(*MultipleDataStoreDataLatest, int)) *MultipleDataStoreDataLatestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MultipleDataStoreDataLatestStream) Reduce(fn func(MultipleDataStoreDataLatest, MultipleDataStoreDataLatest, int) MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	return self.ReduceInit(fn, MultipleDataStoreDataLatest{})
}
func (self *MultipleDataStoreDataLatestStream) ReduceInit(fn func(MultipleDataStoreDataLatest, MultipleDataStoreDataLatest, int) MultipleDataStoreDataLatest, initialValue MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	result := MultipleDataStoreDataLatestStreamOf()
	self.ForEach(func(v MultipleDataStoreDataLatest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleDataStoreDataLatestStream) ReduceInterface(fn func(interface{}, MultipleDataStoreDataLatest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleDataStoreDataLatest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleDataStoreDataLatestStream) ReduceString(fn func(string, MultipleDataStoreDataLatest, int) string) []string {
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
func (self *MultipleDataStoreDataLatestStream) ReduceInt(fn func(int, MultipleDataStoreDataLatest, int) int) []int {
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
func (self *MultipleDataStoreDataLatestStream) ReduceInt32(fn func(int32, MultipleDataStoreDataLatest, int) int32) []int32 {
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
func (self *MultipleDataStoreDataLatestStream) ReduceInt64(fn func(int64, MultipleDataStoreDataLatest, int) int64) []int64 {
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
func (self *MultipleDataStoreDataLatestStream) ReduceFloat32(fn func(float32, MultipleDataStoreDataLatest, int) float32) []float32 {
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
func (self *MultipleDataStoreDataLatestStream) ReduceFloat64(fn func(float64, MultipleDataStoreDataLatest, int) float64) []float64 {
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
func (self *MultipleDataStoreDataLatestStream) ReduceBool(fn func(bool, MultipleDataStoreDataLatest, int) bool) []bool {
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
func (self *MultipleDataStoreDataLatestStream) Reverse() *MultipleDataStoreDataLatestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) Replace(fn func(MultipleDataStoreDataLatest, int) MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	return self.ForEach(func(v MultipleDataStoreDataLatest, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleDataStoreDataLatestStream) Select(fn func(MultipleDataStoreDataLatest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleDataStoreDataLatestStream) Set(index int, val MultipleDataStoreDataLatest) *MultipleDataStoreDataLatestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) Skip(skip int) *MultipleDataStoreDataLatestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleDataStoreDataLatestStream) SkippingEach(fn func(MultipleDataStoreDataLatest, int) int) *MultipleDataStoreDataLatestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) Slice(startIndex, n int) *MultipleDataStoreDataLatestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleDataStoreDataLatest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) Sort(fn func(i, j int) bool) *MultipleDataStoreDataLatestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleDataStoreDataLatestStream) Tail() *MultipleDataStoreDataLatest {
	return self.Last()
}
func (self *MultipleDataStoreDataLatestStream) TailOr(arg MultipleDataStoreDataLatest) MultipleDataStoreDataLatest {
	return self.LastOr(arg)
}
func (self *MultipleDataStoreDataLatestStream) ToList() []MultipleDataStoreDataLatest {
	return self.Val()
}
func (self *MultipleDataStoreDataLatestStream) Unique() *MultipleDataStoreDataLatestStream {
	return self.Distinct()
}
func (self *MultipleDataStoreDataLatestStream) Val() []MultipleDataStoreDataLatest {
	if self == nil {
		return []MultipleDataStoreDataLatest{}
	}
	return *self.Copy()
}
func (self *MultipleDataStoreDataLatestStream) While(fn func(MultipleDataStoreDataLatest, int) bool) *MultipleDataStoreDataLatestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleDataStoreDataLatestStream) Where(fn func(MultipleDataStoreDataLatest) bool) *MultipleDataStoreDataLatestStream {
	result := MultipleDataStoreDataLatestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDataStoreDataLatestStream) WhereSlim(fn func(MultipleDataStoreDataLatest) bool) *MultipleDataStoreDataLatestStream {
	result := MultipleDataStoreDataLatestStreamOf()
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
