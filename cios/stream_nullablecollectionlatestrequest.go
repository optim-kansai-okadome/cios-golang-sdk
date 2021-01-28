/*
 * Collection utility of NullableCollectionLatestRequest Struct
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

type NullableCollectionLatestRequestStream []NullableCollectionLatestRequest

func NullableCollectionLatestRequestStreamOf(arg ...NullableCollectionLatestRequest) NullableCollectionLatestRequestStream {
	return arg
}
func NullableCollectionLatestRequestStreamFrom(arg []NullableCollectionLatestRequest) NullableCollectionLatestRequestStream {
	return arg
}
func CreateNullableCollectionLatestRequestStream(arg ...NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	tmp := NullableCollectionLatestRequestStreamOf(arg...)
	return &tmp
}
func GenerateNullableCollectionLatestRequestStream(arg []NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	tmp := NullableCollectionLatestRequestStreamFrom(arg)
	return &tmp
}

func (self *NullableCollectionLatestRequestStream) Add(arg NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	return self.AddAll(arg)
}
func (self *NullableCollectionLatestRequestStream) AddAll(arg ...NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableCollectionLatestRequestStream) AddSafe(arg *NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) Aggregate(fn func(NullableCollectionLatestRequest, NullableCollectionLatestRequest) NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	result := NullableCollectionLatestRequestStreamOf()
	self.ForEach(func(v NullableCollectionLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(NullableCollectionLatestRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableCollectionLatestRequestStream) AllMatch(fn func(NullableCollectionLatestRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableCollectionLatestRequestStream) AnyMatch(fn func(NullableCollectionLatestRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableCollectionLatestRequestStream) Clone() *NullableCollectionLatestRequestStream {
	temp := make([]NullableCollectionLatestRequest, self.Len())
	copy(temp, *self)
	return (*NullableCollectionLatestRequestStream)(&temp)
}
func (self *NullableCollectionLatestRequestStream) Copy() *NullableCollectionLatestRequestStream {
	return self.Clone()
}
func (self *NullableCollectionLatestRequestStream) Concat(arg []NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	return self.AddAll(arg...)
}
func (self *NullableCollectionLatestRequestStream) Contains(arg NullableCollectionLatestRequest) bool {
	return self.FindIndex(func(_arg NullableCollectionLatestRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableCollectionLatestRequestStream) Clean() *NullableCollectionLatestRequestStream {
	*self = NullableCollectionLatestRequestStreamOf()
	return self
}
func (self *NullableCollectionLatestRequestStream) Delete(index int) *NullableCollectionLatestRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableCollectionLatestRequestStream) DeleteRange(startIndex, endIndex int) *NullableCollectionLatestRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableCollectionLatestRequestStream) Distinct() *NullableCollectionLatestRequestStream {
	caches := map[string]bool{}
	result := NullableCollectionLatestRequestStreamOf()
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
func (self *NullableCollectionLatestRequestStream) Each(fn func(NullableCollectionLatestRequest)) *NullableCollectionLatestRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) EachRight(fn func(NullableCollectionLatestRequest)) *NullableCollectionLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) Equals(arr []NullableCollectionLatestRequest) bool {
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
func (self *NullableCollectionLatestRequestStream) Filter(fn func(NullableCollectionLatestRequest, int) bool) *NullableCollectionLatestRequestStream {
	result := NullableCollectionLatestRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableCollectionLatestRequestStream) FilterSlim(fn func(NullableCollectionLatestRequest, int) bool) *NullableCollectionLatestRequestStream {
	result := NullableCollectionLatestRequestStreamOf()
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
func (self *NullableCollectionLatestRequestStream) Find(fn func(NullableCollectionLatestRequest, int) bool) *NullableCollectionLatestRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableCollectionLatestRequestStream) FindOr(fn func(NullableCollectionLatestRequest, int) bool, or NullableCollectionLatestRequest) NullableCollectionLatestRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableCollectionLatestRequestStream) FindIndex(fn func(NullableCollectionLatestRequest, int) bool) int {
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
func (self *NullableCollectionLatestRequestStream) First() *NullableCollectionLatestRequest {
	return self.Get(0)
}
func (self *NullableCollectionLatestRequestStream) FirstOr(arg NullableCollectionLatestRequest) NullableCollectionLatestRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableCollectionLatestRequestStream) ForEach(fn func(NullableCollectionLatestRequest, int)) *NullableCollectionLatestRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) ForEachRight(fn func(NullableCollectionLatestRequest, int)) *NullableCollectionLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) GroupBy(fn func(NullableCollectionLatestRequest, int) string) map[string][]NullableCollectionLatestRequest {
	m := map[string][]NullableCollectionLatestRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableCollectionLatestRequestStream) GroupByValues(fn func(NullableCollectionLatestRequest, int) string) [][]NullableCollectionLatestRequest {
	var tmp [][]NullableCollectionLatestRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableCollectionLatestRequestStream) IndexOf(arg NullableCollectionLatestRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableCollectionLatestRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableCollectionLatestRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableCollectionLatestRequestStream) Last() *NullableCollectionLatestRequest {
	return self.Get(self.Len() - 1)
}
func (self *NullableCollectionLatestRequestStream) LastOr(arg NullableCollectionLatestRequest) NullableCollectionLatestRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableCollectionLatestRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableCollectionLatestRequestStream) Limit(limit int) *NullableCollectionLatestRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableCollectionLatestRequestStream) Map(fn func(NullableCollectionLatestRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Map2Int(fn func(NullableCollectionLatestRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Map2Int32(fn func(NullableCollectionLatestRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Map2Int64(fn func(NullableCollectionLatestRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Map2Float32(fn func(NullableCollectionLatestRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Map2Float64(fn func(NullableCollectionLatestRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Map2Bool(fn func(NullableCollectionLatestRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Map2Bytes(fn func(NullableCollectionLatestRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Map2String(fn func(NullableCollectionLatestRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Max(fn func(NullableCollectionLatestRequest, int) float64) *NullableCollectionLatestRequest {
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
func (self *NullableCollectionLatestRequestStream) Min(fn func(NullableCollectionLatestRequest, int) float64) *NullableCollectionLatestRequest {
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
func (self *NullableCollectionLatestRequestStream) NoneMatch(fn func(NullableCollectionLatestRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableCollectionLatestRequestStream) Get(index int) *NullableCollectionLatestRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableCollectionLatestRequestStream) GetOr(index int, arg NullableCollectionLatestRequest) NullableCollectionLatestRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableCollectionLatestRequestStream) Peek(fn func(*NullableCollectionLatestRequest, int)) *NullableCollectionLatestRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableCollectionLatestRequestStream) Reduce(fn func(NullableCollectionLatestRequest, NullableCollectionLatestRequest, int) NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	return self.ReduceInit(fn, NullableCollectionLatestRequest{})
}
func (self *NullableCollectionLatestRequestStream) ReduceInit(fn func(NullableCollectionLatestRequest, NullableCollectionLatestRequest, int) NullableCollectionLatestRequest, initialValue NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	result := NullableCollectionLatestRequestStreamOf()
	self.ForEach(func(v NullableCollectionLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableCollectionLatestRequestStream) ReduceInterface(fn func(interface{}, NullableCollectionLatestRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableCollectionLatestRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableCollectionLatestRequestStream) ReduceString(fn func(string, NullableCollectionLatestRequest, int) string) []string {
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
func (self *NullableCollectionLatestRequestStream) ReduceInt(fn func(int, NullableCollectionLatestRequest, int) int) []int {
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
func (self *NullableCollectionLatestRequestStream) ReduceInt32(fn func(int32, NullableCollectionLatestRequest, int) int32) []int32 {
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
func (self *NullableCollectionLatestRequestStream) ReduceInt64(fn func(int64, NullableCollectionLatestRequest, int) int64) []int64 {
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
func (self *NullableCollectionLatestRequestStream) ReduceFloat32(fn func(float32, NullableCollectionLatestRequest, int) float32) []float32 {
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
func (self *NullableCollectionLatestRequestStream) ReduceFloat64(fn func(float64, NullableCollectionLatestRequest, int) float64) []float64 {
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
func (self *NullableCollectionLatestRequestStream) ReduceBool(fn func(bool, NullableCollectionLatestRequest, int) bool) []bool {
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
func (self *NullableCollectionLatestRequestStream) Reverse() *NullableCollectionLatestRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) Replace(fn func(NullableCollectionLatestRequest, int) NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	return self.ForEach(func(v NullableCollectionLatestRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableCollectionLatestRequestStream) Select(fn func(NullableCollectionLatestRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableCollectionLatestRequestStream) Set(index int, val NullableCollectionLatestRequest) *NullableCollectionLatestRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) Skip(skip int) *NullableCollectionLatestRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableCollectionLatestRequestStream) SkippingEach(fn func(NullableCollectionLatestRequest, int) int) *NullableCollectionLatestRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) Slice(startIndex, n int) *NullableCollectionLatestRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableCollectionLatestRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) Sort(fn func(i, j int) bool) *NullableCollectionLatestRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableCollectionLatestRequestStream) Tail() *NullableCollectionLatestRequest {
	return self.Last()
}
func (self *NullableCollectionLatestRequestStream) TailOr(arg NullableCollectionLatestRequest) NullableCollectionLatestRequest {
	return self.LastOr(arg)
}
func (self *NullableCollectionLatestRequestStream) ToList() []NullableCollectionLatestRequest {
	return self.Val()
}
func (self *NullableCollectionLatestRequestStream) Unique() *NullableCollectionLatestRequestStream {
	return self.Distinct()
}
func (self *NullableCollectionLatestRequestStream) Val() []NullableCollectionLatestRequest {
	if self == nil {
		return []NullableCollectionLatestRequest{}
	}
	return *self.Copy()
}
func (self *NullableCollectionLatestRequestStream) While(fn func(NullableCollectionLatestRequest, int) bool) *NullableCollectionLatestRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableCollectionLatestRequestStream) Where(fn func(NullableCollectionLatestRequest) bool) *NullableCollectionLatestRequestStream {
	result := NullableCollectionLatestRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableCollectionLatestRequestStream) WhereSlim(fn func(NullableCollectionLatestRequest) bool) *NullableCollectionLatestRequestStream {
	result := NullableCollectionLatestRequestStreamOf()
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
