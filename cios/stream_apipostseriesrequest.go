/*
 * Collection utility of ApiPostSeriesRequest Struct
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

type ApiPostSeriesRequestStream []ApiPostSeriesRequest

func ApiPostSeriesRequestStreamOf(arg ...ApiPostSeriesRequest) ApiPostSeriesRequestStream {
	return arg
}
func ApiPostSeriesRequestStreamFrom(arg []ApiPostSeriesRequest) ApiPostSeriesRequestStream {
	return arg
}
func CreateApiPostSeriesRequestStream(arg ...ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	tmp := ApiPostSeriesRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiPostSeriesRequestStream(arg []ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	tmp := ApiPostSeriesRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiPostSeriesRequestStream) Add(arg ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	return self.AddAll(arg)
}
func (self *ApiPostSeriesRequestStream) AddAll(arg ...ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiPostSeriesRequestStream) AddSafe(arg *ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiPostSeriesRequestStream) Aggregate(fn func(ApiPostSeriesRequest, ApiPostSeriesRequest) ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	result := ApiPostSeriesRequestStreamOf()
	self.ForEach(func(v ApiPostSeriesRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiPostSeriesRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiPostSeriesRequestStream) AllMatch(fn func(ApiPostSeriesRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiPostSeriesRequestStream) AnyMatch(fn func(ApiPostSeriesRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiPostSeriesRequestStream) Clone() *ApiPostSeriesRequestStream {
	temp := make([]ApiPostSeriesRequest, self.Len())
	copy(temp, *self)
	return (*ApiPostSeriesRequestStream)(&temp)
}
func (self *ApiPostSeriesRequestStream) Copy() *ApiPostSeriesRequestStream {
	return self.Clone()
}
func (self *ApiPostSeriesRequestStream) Concat(arg []ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiPostSeriesRequestStream) Contains(arg ApiPostSeriesRequest) bool {
	return self.FindIndex(func(_arg ApiPostSeriesRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiPostSeriesRequestStream) Clean() *ApiPostSeriesRequestStream {
	*self = ApiPostSeriesRequestStreamOf()
	return self
}
func (self *ApiPostSeriesRequestStream) Delete(index int) *ApiPostSeriesRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiPostSeriesRequestStream) DeleteRange(startIndex, endIndex int) *ApiPostSeriesRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiPostSeriesRequestStream) Distinct() *ApiPostSeriesRequestStream {
	caches := map[string]bool{}
	result := ApiPostSeriesRequestStreamOf()
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
func (self *ApiPostSeriesRequestStream) Each(fn func(ApiPostSeriesRequest)) *ApiPostSeriesRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiPostSeriesRequestStream) EachRight(fn func(ApiPostSeriesRequest)) *ApiPostSeriesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiPostSeriesRequestStream) Equals(arr []ApiPostSeriesRequest) bool {
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
func (self *ApiPostSeriesRequestStream) Filter(fn func(ApiPostSeriesRequest, int) bool) *ApiPostSeriesRequestStream {
	result := ApiPostSeriesRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiPostSeriesRequestStream) FilterSlim(fn func(ApiPostSeriesRequest, int) bool) *ApiPostSeriesRequestStream {
	result := ApiPostSeriesRequestStreamOf()
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
func (self *ApiPostSeriesRequestStream) Find(fn func(ApiPostSeriesRequest, int) bool) *ApiPostSeriesRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiPostSeriesRequestStream) FindOr(fn func(ApiPostSeriesRequest, int) bool, or ApiPostSeriesRequest) ApiPostSeriesRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiPostSeriesRequestStream) FindIndex(fn func(ApiPostSeriesRequest, int) bool) int {
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
func (self *ApiPostSeriesRequestStream) First() *ApiPostSeriesRequest {
	return self.Get(0)
}
func (self *ApiPostSeriesRequestStream) FirstOr(arg ApiPostSeriesRequest) ApiPostSeriesRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiPostSeriesRequestStream) ForEach(fn func(ApiPostSeriesRequest, int)) *ApiPostSeriesRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiPostSeriesRequestStream) ForEachRight(fn func(ApiPostSeriesRequest, int)) *ApiPostSeriesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiPostSeriesRequestStream) GroupBy(fn func(ApiPostSeriesRequest, int) string) map[string][]ApiPostSeriesRequest {
	m := map[string][]ApiPostSeriesRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiPostSeriesRequestStream) GroupByValues(fn func(ApiPostSeriesRequest, int) string) [][]ApiPostSeriesRequest {
	var tmp [][]ApiPostSeriesRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiPostSeriesRequestStream) IndexOf(arg ApiPostSeriesRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiPostSeriesRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiPostSeriesRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiPostSeriesRequestStream) Last() *ApiPostSeriesRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiPostSeriesRequestStream) LastOr(arg ApiPostSeriesRequest) ApiPostSeriesRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiPostSeriesRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiPostSeriesRequestStream) Limit(limit int) *ApiPostSeriesRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiPostSeriesRequestStream) Map(fn func(ApiPostSeriesRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Map2Int(fn func(ApiPostSeriesRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Map2Int32(fn func(ApiPostSeriesRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Map2Int64(fn func(ApiPostSeriesRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Map2Float32(fn func(ApiPostSeriesRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Map2Float64(fn func(ApiPostSeriesRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Map2Bool(fn func(ApiPostSeriesRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Map2Bytes(fn func(ApiPostSeriesRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Map2String(fn func(ApiPostSeriesRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Max(fn func(ApiPostSeriesRequest, int) float64) *ApiPostSeriesRequest {
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
func (self *ApiPostSeriesRequestStream) Min(fn func(ApiPostSeriesRequest, int) float64) *ApiPostSeriesRequest {
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
func (self *ApiPostSeriesRequestStream) NoneMatch(fn func(ApiPostSeriesRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiPostSeriesRequestStream) Get(index int) *ApiPostSeriesRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiPostSeriesRequestStream) GetOr(index int, arg ApiPostSeriesRequest) ApiPostSeriesRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiPostSeriesRequestStream) Peek(fn func(*ApiPostSeriesRequest, int)) *ApiPostSeriesRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiPostSeriesRequestStream) Reduce(fn func(ApiPostSeriesRequest, ApiPostSeriesRequest, int) ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	return self.ReduceInit(fn, ApiPostSeriesRequest{})
}
func (self *ApiPostSeriesRequestStream) ReduceInit(fn func(ApiPostSeriesRequest, ApiPostSeriesRequest, int) ApiPostSeriesRequest, initialValue ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	result := ApiPostSeriesRequestStreamOf()
	self.ForEach(func(v ApiPostSeriesRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiPostSeriesRequestStream) ReduceInterface(fn func(interface{}, ApiPostSeriesRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiPostSeriesRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiPostSeriesRequestStream) ReduceString(fn func(string, ApiPostSeriesRequest, int) string) []string {
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
func (self *ApiPostSeriesRequestStream) ReduceInt(fn func(int, ApiPostSeriesRequest, int) int) []int {
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
func (self *ApiPostSeriesRequestStream) ReduceInt32(fn func(int32, ApiPostSeriesRequest, int) int32) []int32 {
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
func (self *ApiPostSeriesRequestStream) ReduceInt64(fn func(int64, ApiPostSeriesRequest, int) int64) []int64 {
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
func (self *ApiPostSeriesRequestStream) ReduceFloat32(fn func(float32, ApiPostSeriesRequest, int) float32) []float32 {
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
func (self *ApiPostSeriesRequestStream) ReduceFloat64(fn func(float64, ApiPostSeriesRequest, int) float64) []float64 {
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
func (self *ApiPostSeriesRequestStream) ReduceBool(fn func(bool, ApiPostSeriesRequest, int) bool) []bool {
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
func (self *ApiPostSeriesRequestStream) Reverse() *ApiPostSeriesRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiPostSeriesRequestStream) Replace(fn func(ApiPostSeriesRequest, int) ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	return self.ForEach(func(v ApiPostSeriesRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiPostSeriesRequestStream) Select(fn func(ApiPostSeriesRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiPostSeriesRequestStream) Set(index int, val ApiPostSeriesRequest) *ApiPostSeriesRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiPostSeriesRequestStream) Skip(skip int) *ApiPostSeriesRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiPostSeriesRequestStream) SkippingEach(fn func(ApiPostSeriesRequest, int) int) *ApiPostSeriesRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiPostSeriesRequestStream) Slice(startIndex, n int) *ApiPostSeriesRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiPostSeriesRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiPostSeriesRequestStream) Sort(fn func(i, j int) bool) *ApiPostSeriesRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiPostSeriesRequestStream) Tail() *ApiPostSeriesRequest {
	return self.Last()
}
func (self *ApiPostSeriesRequestStream) TailOr(arg ApiPostSeriesRequest) ApiPostSeriesRequest {
	return self.LastOr(arg)
}
func (self *ApiPostSeriesRequestStream) ToList() []ApiPostSeriesRequest {
	return self.Val()
}
func (self *ApiPostSeriesRequestStream) Unique() *ApiPostSeriesRequestStream {
	return self.Distinct()
}
func (self *ApiPostSeriesRequestStream) Val() []ApiPostSeriesRequest {
	if self == nil {
		return []ApiPostSeriesRequest{}
	}
	return *self.Copy()
}
func (self *ApiPostSeriesRequestStream) While(fn func(ApiPostSeriesRequest, int) bool) *ApiPostSeriesRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiPostSeriesRequestStream) Where(fn func(ApiPostSeriesRequest) bool) *ApiPostSeriesRequestStream {
	result := ApiPostSeriesRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiPostSeriesRequestStream) WhereSlim(fn func(ApiPostSeriesRequest) bool) *ApiPostSeriesRequestStream {
	result := ApiPostSeriesRequestStreamOf()
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
