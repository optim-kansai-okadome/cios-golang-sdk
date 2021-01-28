/*
 * Collection utility of MultipleDeviceModel Struct
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

type MultipleDeviceModelStream []MultipleDeviceModel

func MultipleDeviceModelStreamOf(arg ...MultipleDeviceModel) MultipleDeviceModelStream {
	return arg
}
func MultipleDeviceModelStreamFrom(arg []MultipleDeviceModel) MultipleDeviceModelStream {
	return arg
}
func CreateMultipleDeviceModelStream(arg ...MultipleDeviceModel) *MultipleDeviceModelStream {
	tmp := MultipleDeviceModelStreamOf(arg...)
	return &tmp
}
func GenerateMultipleDeviceModelStream(arg []MultipleDeviceModel) *MultipleDeviceModelStream {
	tmp := MultipleDeviceModelStreamFrom(arg)
	return &tmp
}

func (self *MultipleDeviceModelStream) Add(arg MultipleDeviceModel) *MultipleDeviceModelStream {
	return self.AddAll(arg)
}
func (self *MultipleDeviceModelStream) AddAll(arg ...MultipleDeviceModel) *MultipleDeviceModelStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleDeviceModelStream) AddSafe(arg *MultipleDeviceModel) *MultipleDeviceModelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleDeviceModelStream) Aggregate(fn func(MultipleDeviceModel, MultipleDeviceModel) MultipleDeviceModel) *MultipleDeviceModelStream {
	result := MultipleDeviceModelStreamOf()
	self.ForEach(func(v MultipleDeviceModel, i int) {
		if i == 0 {
			result.Add(fn(MultipleDeviceModel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleDeviceModelStream) AllMatch(fn func(MultipleDeviceModel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleDeviceModelStream) AnyMatch(fn func(MultipleDeviceModel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleDeviceModelStream) Clone() *MultipleDeviceModelStream {
	temp := make([]MultipleDeviceModel, self.Len())
	copy(temp, *self)
	return (*MultipleDeviceModelStream)(&temp)
}
func (self *MultipleDeviceModelStream) Copy() *MultipleDeviceModelStream {
	return self.Clone()
}
func (self *MultipleDeviceModelStream) Concat(arg []MultipleDeviceModel) *MultipleDeviceModelStream {
	return self.AddAll(arg...)
}
func (self *MultipleDeviceModelStream) Contains(arg MultipleDeviceModel) bool {
	return self.FindIndex(func(_arg MultipleDeviceModel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleDeviceModelStream) Clean() *MultipleDeviceModelStream {
	*self = MultipleDeviceModelStreamOf()
	return self
}
func (self *MultipleDeviceModelStream) Delete(index int) *MultipleDeviceModelStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleDeviceModelStream) DeleteRange(startIndex, endIndex int) *MultipleDeviceModelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleDeviceModelStream) Distinct() *MultipleDeviceModelStream {
	caches := map[string]bool{}
	result := MultipleDeviceModelStreamOf()
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
func (self *MultipleDeviceModelStream) Each(fn func(MultipleDeviceModel)) *MultipleDeviceModelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleDeviceModelStream) EachRight(fn func(MultipleDeviceModel)) *MultipleDeviceModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleDeviceModelStream) Equals(arr []MultipleDeviceModel) bool {
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
func (self *MultipleDeviceModelStream) Filter(fn func(MultipleDeviceModel, int) bool) *MultipleDeviceModelStream {
	result := MultipleDeviceModelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDeviceModelStream) FilterSlim(fn func(MultipleDeviceModel, int) bool) *MultipleDeviceModelStream {
	result := MultipleDeviceModelStreamOf()
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
func (self *MultipleDeviceModelStream) Find(fn func(MultipleDeviceModel, int) bool) *MultipleDeviceModel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleDeviceModelStream) FindOr(fn func(MultipleDeviceModel, int) bool, or MultipleDeviceModel) MultipleDeviceModel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleDeviceModelStream) FindIndex(fn func(MultipleDeviceModel, int) bool) int {
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
func (self *MultipleDeviceModelStream) First() *MultipleDeviceModel {
	return self.Get(0)
}
func (self *MultipleDeviceModelStream) FirstOr(arg MultipleDeviceModel) MultipleDeviceModel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDeviceModelStream) ForEach(fn func(MultipleDeviceModel, int)) *MultipleDeviceModelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleDeviceModelStream) ForEachRight(fn func(MultipleDeviceModel, int)) *MultipleDeviceModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleDeviceModelStream) GroupBy(fn func(MultipleDeviceModel, int) string) map[string][]MultipleDeviceModel {
	m := map[string][]MultipleDeviceModel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleDeviceModelStream) GroupByValues(fn func(MultipleDeviceModel, int) string) [][]MultipleDeviceModel {
	var tmp [][]MultipleDeviceModel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleDeviceModelStream) IndexOf(arg MultipleDeviceModel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleDeviceModelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleDeviceModelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleDeviceModelStream) Last() *MultipleDeviceModel {
	return self.Get(self.Len() - 1)
}
func (self *MultipleDeviceModelStream) LastOr(arg MultipleDeviceModel) MultipleDeviceModel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDeviceModelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleDeviceModelStream) Limit(limit int) *MultipleDeviceModelStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleDeviceModelStream) Map(fn func(MultipleDeviceModel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Map2Int(fn func(MultipleDeviceModel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Map2Int32(fn func(MultipleDeviceModel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Map2Int64(fn func(MultipleDeviceModel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Map2Float32(fn func(MultipleDeviceModel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Map2Float64(fn func(MultipleDeviceModel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Map2Bool(fn func(MultipleDeviceModel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Map2Bytes(fn func(MultipleDeviceModel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Map2String(fn func(MultipleDeviceModel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Max(fn func(MultipleDeviceModel, int) float64) *MultipleDeviceModel {
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
func (self *MultipleDeviceModelStream) Min(fn func(MultipleDeviceModel, int) float64) *MultipleDeviceModel {
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
func (self *MultipleDeviceModelStream) NoneMatch(fn func(MultipleDeviceModel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleDeviceModelStream) Get(index int) *MultipleDeviceModel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleDeviceModelStream) GetOr(index int, arg MultipleDeviceModel) MultipleDeviceModel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDeviceModelStream) Peek(fn func(*MultipleDeviceModel, int)) *MultipleDeviceModelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MultipleDeviceModelStream) Reduce(fn func(MultipleDeviceModel, MultipleDeviceModel, int) MultipleDeviceModel) *MultipleDeviceModelStream {
	return self.ReduceInit(fn, MultipleDeviceModel{})
}
func (self *MultipleDeviceModelStream) ReduceInit(fn func(MultipleDeviceModel, MultipleDeviceModel, int) MultipleDeviceModel, initialValue MultipleDeviceModel) *MultipleDeviceModelStream {
	result := MultipleDeviceModelStreamOf()
	self.ForEach(func(v MultipleDeviceModel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleDeviceModelStream) ReduceInterface(fn func(interface{}, MultipleDeviceModel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleDeviceModel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleDeviceModelStream) ReduceString(fn func(string, MultipleDeviceModel, int) string) []string {
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
func (self *MultipleDeviceModelStream) ReduceInt(fn func(int, MultipleDeviceModel, int) int) []int {
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
func (self *MultipleDeviceModelStream) ReduceInt32(fn func(int32, MultipleDeviceModel, int) int32) []int32 {
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
func (self *MultipleDeviceModelStream) ReduceInt64(fn func(int64, MultipleDeviceModel, int) int64) []int64 {
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
func (self *MultipleDeviceModelStream) ReduceFloat32(fn func(float32, MultipleDeviceModel, int) float32) []float32 {
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
func (self *MultipleDeviceModelStream) ReduceFloat64(fn func(float64, MultipleDeviceModel, int) float64) []float64 {
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
func (self *MultipleDeviceModelStream) ReduceBool(fn func(bool, MultipleDeviceModel, int) bool) []bool {
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
func (self *MultipleDeviceModelStream) Reverse() *MultipleDeviceModelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleDeviceModelStream) Replace(fn func(MultipleDeviceModel, int) MultipleDeviceModel) *MultipleDeviceModelStream {
	return self.ForEach(func(v MultipleDeviceModel, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleDeviceModelStream) Select(fn func(MultipleDeviceModel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleDeviceModelStream) Set(index int, val MultipleDeviceModel) *MultipleDeviceModelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleDeviceModelStream) Skip(skip int) *MultipleDeviceModelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleDeviceModelStream) SkippingEach(fn func(MultipleDeviceModel, int) int) *MultipleDeviceModelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleDeviceModelStream) Slice(startIndex, n int) *MultipleDeviceModelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleDeviceModel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleDeviceModelStream) Sort(fn func(i, j int) bool) *MultipleDeviceModelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleDeviceModelStream) Tail() *MultipleDeviceModel {
	return self.Last()
}
func (self *MultipleDeviceModelStream) TailOr(arg MultipleDeviceModel) MultipleDeviceModel {
	return self.LastOr(arg)
}
func (self *MultipleDeviceModelStream) ToList() []MultipleDeviceModel {
	return self.Val()
}
func (self *MultipleDeviceModelStream) Unique() *MultipleDeviceModelStream {
	return self.Distinct()
}
func (self *MultipleDeviceModelStream) Val() []MultipleDeviceModel {
	if self == nil {
		return []MultipleDeviceModel{}
	}
	return *self.Copy()
}
func (self *MultipleDeviceModelStream) While(fn func(MultipleDeviceModel, int) bool) *MultipleDeviceModelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleDeviceModelStream) Where(fn func(MultipleDeviceModel) bool) *MultipleDeviceModelStream {
	result := MultipleDeviceModelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDeviceModelStream) WhereSlim(fn func(MultipleDeviceModel) bool) *MultipleDeviceModelStream {
	result := MultipleDeviceModelStreamOf()
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
