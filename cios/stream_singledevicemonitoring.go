/*
 * Collection utility of SingleDeviceMonitoring Struct
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

type SingleDeviceMonitoringStream []SingleDeviceMonitoring

func SingleDeviceMonitoringStreamOf(arg ...SingleDeviceMonitoring) SingleDeviceMonitoringStream {
	return arg
}
func SingleDeviceMonitoringStreamFrom(arg []SingleDeviceMonitoring) SingleDeviceMonitoringStream {
	return arg
}
func CreateSingleDeviceMonitoringStream(arg ...SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	tmp := SingleDeviceMonitoringStreamOf(arg...)
	return &tmp
}
func GenerateSingleDeviceMonitoringStream(arg []SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	tmp := SingleDeviceMonitoringStreamFrom(arg)
	return &tmp
}

func (self *SingleDeviceMonitoringStream) Add(arg SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	return self.AddAll(arg)
}
func (self *SingleDeviceMonitoringStream) AddAll(arg ...SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	*self = append(*self, arg...)
	return self
}
func (self *SingleDeviceMonitoringStream) AddSafe(arg *SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SingleDeviceMonitoringStream) Aggregate(fn func(SingleDeviceMonitoring, SingleDeviceMonitoring) SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	result := SingleDeviceMonitoringStreamOf()
	self.ForEach(func(v SingleDeviceMonitoring, i int) {
		if i == 0 {
			result.Add(fn(SingleDeviceMonitoring{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *SingleDeviceMonitoringStream) AllMatch(fn func(SingleDeviceMonitoring, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SingleDeviceMonitoringStream) AnyMatch(fn func(SingleDeviceMonitoring, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SingleDeviceMonitoringStream) Clone() *SingleDeviceMonitoringStream {
	temp := make([]SingleDeviceMonitoring, self.Len())
	copy(temp, *self)
	return (*SingleDeviceMonitoringStream)(&temp)
}
func (self *SingleDeviceMonitoringStream) Copy() *SingleDeviceMonitoringStream {
	return self.Clone()
}
func (self *SingleDeviceMonitoringStream) Concat(arg []SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	return self.AddAll(arg...)
}
func (self *SingleDeviceMonitoringStream) Contains(arg SingleDeviceMonitoring) bool {
	return self.FindIndex(func(_arg SingleDeviceMonitoring, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SingleDeviceMonitoringStream) Clean() *SingleDeviceMonitoringStream {
	*self = SingleDeviceMonitoringStreamOf()
	return self
}
func (self *SingleDeviceMonitoringStream) Delete(index int) *SingleDeviceMonitoringStream {
	return self.DeleteRange(index, index)
}
func (self *SingleDeviceMonitoringStream) DeleteRange(startIndex, endIndex int) *SingleDeviceMonitoringStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SingleDeviceMonitoringStream) Distinct() *SingleDeviceMonitoringStream {
	caches := map[string]bool{}
	result := SingleDeviceMonitoringStreamOf()
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
func (self *SingleDeviceMonitoringStream) Each(fn func(SingleDeviceMonitoring)) *SingleDeviceMonitoringStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SingleDeviceMonitoringStream) EachRight(fn func(SingleDeviceMonitoring)) *SingleDeviceMonitoringStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SingleDeviceMonitoringStream) Equals(arr []SingleDeviceMonitoring) bool {
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
func (self *SingleDeviceMonitoringStream) Filter(fn func(SingleDeviceMonitoring, int) bool) *SingleDeviceMonitoringStream {
	result := SingleDeviceMonitoringStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleDeviceMonitoringStream) FilterSlim(fn func(SingleDeviceMonitoring, int) bool) *SingleDeviceMonitoringStream {
	result := SingleDeviceMonitoringStreamOf()
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
func (self *SingleDeviceMonitoringStream) Find(fn func(SingleDeviceMonitoring, int) bool) *SingleDeviceMonitoring {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SingleDeviceMonitoringStream) FindOr(fn func(SingleDeviceMonitoring, int) bool, or SingleDeviceMonitoring) SingleDeviceMonitoring {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SingleDeviceMonitoringStream) FindIndex(fn func(SingleDeviceMonitoring, int) bool) int {
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
func (self *SingleDeviceMonitoringStream) First() *SingleDeviceMonitoring {
	return self.Get(0)
}
func (self *SingleDeviceMonitoringStream) FirstOr(arg SingleDeviceMonitoring) SingleDeviceMonitoring {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SingleDeviceMonitoringStream) ForEach(fn func(SingleDeviceMonitoring, int)) *SingleDeviceMonitoringStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SingleDeviceMonitoringStream) ForEachRight(fn func(SingleDeviceMonitoring, int)) *SingleDeviceMonitoringStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SingleDeviceMonitoringStream) GroupBy(fn func(SingleDeviceMonitoring, int) string) map[string][]SingleDeviceMonitoring {
	m := map[string][]SingleDeviceMonitoring{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SingleDeviceMonitoringStream) GroupByValues(fn func(SingleDeviceMonitoring, int) string) [][]SingleDeviceMonitoring {
	var tmp [][]SingleDeviceMonitoring
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SingleDeviceMonitoringStream) IndexOf(arg SingleDeviceMonitoring) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SingleDeviceMonitoringStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SingleDeviceMonitoringStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SingleDeviceMonitoringStream) Last() *SingleDeviceMonitoring {
	return self.Get(self.Len() - 1)
}
func (self *SingleDeviceMonitoringStream) LastOr(arg SingleDeviceMonitoring) SingleDeviceMonitoring {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SingleDeviceMonitoringStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SingleDeviceMonitoringStream) Limit(limit int) *SingleDeviceMonitoringStream {
	self.Slice(0, limit)
	return self
}

func (self *SingleDeviceMonitoringStream) Map(fn func(SingleDeviceMonitoring, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Map2Int(fn func(SingleDeviceMonitoring, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Map2Int32(fn func(SingleDeviceMonitoring, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Map2Int64(fn func(SingleDeviceMonitoring, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Map2Float32(fn func(SingleDeviceMonitoring, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Map2Float64(fn func(SingleDeviceMonitoring, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Map2Bool(fn func(SingleDeviceMonitoring, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Map2Bytes(fn func(SingleDeviceMonitoring, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Map2String(fn func(SingleDeviceMonitoring, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Max(fn func(SingleDeviceMonitoring, int) float64) *SingleDeviceMonitoring {
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
func (self *SingleDeviceMonitoringStream) Min(fn func(SingleDeviceMonitoring, int) float64) *SingleDeviceMonitoring {
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
func (self *SingleDeviceMonitoringStream) NoneMatch(fn func(SingleDeviceMonitoring, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SingleDeviceMonitoringStream) Get(index int) *SingleDeviceMonitoring {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SingleDeviceMonitoringStream) GetOr(index int, arg SingleDeviceMonitoring) SingleDeviceMonitoring {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SingleDeviceMonitoringStream) Peek(fn func(*SingleDeviceMonitoring, int)) *SingleDeviceMonitoringStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *SingleDeviceMonitoringStream) Reduce(fn func(SingleDeviceMonitoring, SingleDeviceMonitoring, int) SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	return self.ReduceInit(fn, SingleDeviceMonitoring{})
}
func (self *SingleDeviceMonitoringStream) ReduceInit(fn func(SingleDeviceMonitoring, SingleDeviceMonitoring, int) SingleDeviceMonitoring, initialValue SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	result := SingleDeviceMonitoringStreamOf()
	self.ForEach(func(v SingleDeviceMonitoring, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SingleDeviceMonitoringStream) ReduceInterface(fn func(interface{}, SingleDeviceMonitoring, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(SingleDeviceMonitoring{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SingleDeviceMonitoringStream) ReduceString(fn func(string, SingleDeviceMonitoring, int) string) []string {
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
func (self *SingleDeviceMonitoringStream) ReduceInt(fn func(int, SingleDeviceMonitoring, int) int) []int {
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
func (self *SingleDeviceMonitoringStream) ReduceInt32(fn func(int32, SingleDeviceMonitoring, int) int32) []int32 {
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
func (self *SingleDeviceMonitoringStream) ReduceInt64(fn func(int64, SingleDeviceMonitoring, int) int64) []int64 {
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
func (self *SingleDeviceMonitoringStream) ReduceFloat32(fn func(float32, SingleDeviceMonitoring, int) float32) []float32 {
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
func (self *SingleDeviceMonitoringStream) ReduceFloat64(fn func(float64, SingleDeviceMonitoring, int) float64) []float64 {
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
func (self *SingleDeviceMonitoringStream) ReduceBool(fn func(bool, SingleDeviceMonitoring, int) bool) []bool {
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
func (self *SingleDeviceMonitoringStream) Reverse() *SingleDeviceMonitoringStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SingleDeviceMonitoringStream) Replace(fn func(SingleDeviceMonitoring, int) SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	return self.ForEach(func(v SingleDeviceMonitoring, i int) { self.Set(i, fn(v, i)) })
}
func (self *SingleDeviceMonitoringStream) Select(fn func(SingleDeviceMonitoring) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *SingleDeviceMonitoringStream) Set(index int, val SingleDeviceMonitoring) *SingleDeviceMonitoringStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SingleDeviceMonitoringStream) Skip(skip int) *SingleDeviceMonitoringStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SingleDeviceMonitoringStream) SkippingEach(fn func(SingleDeviceMonitoring, int) int) *SingleDeviceMonitoringStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SingleDeviceMonitoringStream) Slice(startIndex, n int) *SingleDeviceMonitoringStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []SingleDeviceMonitoring{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SingleDeviceMonitoringStream) Sort(fn func(i, j int) bool) *SingleDeviceMonitoringStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SingleDeviceMonitoringStream) Tail() *SingleDeviceMonitoring {
	return self.Last()
}
func (self *SingleDeviceMonitoringStream) TailOr(arg SingleDeviceMonitoring) SingleDeviceMonitoring {
	return self.LastOr(arg)
}
func (self *SingleDeviceMonitoringStream) ToList() []SingleDeviceMonitoring {
	return self.Val()
}
func (self *SingleDeviceMonitoringStream) Unique() *SingleDeviceMonitoringStream {
	return self.Distinct()
}
func (self *SingleDeviceMonitoringStream) Val() []SingleDeviceMonitoring {
	if self == nil {
		return []SingleDeviceMonitoring{}
	}
	return *self.Copy()
}
func (self *SingleDeviceMonitoringStream) While(fn func(SingleDeviceMonitoring, int) bool) *SingleDeviceMonitoringStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *SingleDeviceMonitoringStream) Where(fn func(SingleDeviceMonitoring) bool) *SingleDeviceMonitoringStream {
	result := SingleDeviceMonitoringStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleDeviceMonitoringStream) WhereSlim(fn func(SingleDeviceMonitoring) bool) *SingleDeviceMonitoringStream {
	result := SingleDeviceMonitoringStreamOf()
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
