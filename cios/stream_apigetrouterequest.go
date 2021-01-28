/*
 * Collection utility of ApiGetRouteRequest Struct
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

type ApiGetRouteRequestStream []ApiGetRouteRequest

func ApiGetRouteRequestStreamOf(arg ...ApiGetRouteRequest) ApiGetRouteRequestStream {
	return arg
}
func ApiGetRouteRequestStreamFrom(arg []ApiGetRouteRequest) ApiGetRouteRequestStream {
	return arg
}
func CreateApiGetRouteRequestStream(arg ...ApiGetRouteRequest) *ApiGetRouteRequestStream {
	tmp := ApiGetRouteRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetRouteRequestStream(arg []ApiGetRouteRequest) *ApiGetRouteRequestStream {
	tmp := ApiGetRouteRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetRouteRequestStream) Add(arg ApiGetRouteRequest) *ApiGetRouteRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetRouteRequestStream) AddAll(arg ...ApiGetRouteRequest) *ApiGetRouteRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetRouteRequestStream) AddSafe(arg *ApiGetRouteRequest) *ApiGetRouteRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetRouteRequestStream) Aggregate(fn func(ApiGetRouteRequest, ApiGetRouteRequest) ApiGetRouteRequest) *ApiGetRouteRequestStream {
	result := ApiGetRouteRequestStreamOf()
	self.ForEach(func(v ApiGetRouteRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetRouteRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetRouteRequestStream) AllMatch(fn func(ApiGetRouteRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetRouteRequestStream) AnyMatch(fn func(ApiGetRouteRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetRouteRequestStream) Clone() *ApiGetRouteRequestStream {
	temp := make([]ApiGetRouteRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetRouteRequestStream)(&temp)
}
func (self *ApiGetRouteRequestStream) Copy() *ApiGetRouteRequestStream {
	return self.Clone()
}
func (self *ApiGetRouteRequestStream) Concat(arg []ApiGetRouteRequest) *ApiGetRouteRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetRouteRequestStream) Contains(arg ApiGetRouteRequest) bool {
	return self.FindIndex(func(_arg ApiGetRouteRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetRouteRequestStream) Clean() *ApiGetRouteRequestStream {
	*self = ApiGetRouteRequestStreamOf()
	return self
}
func (self *ApiGetRouteRequestStream) Delete(index int) *ApiGetRouteRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetRouteRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetRouteRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetRouteRequestStream) Distinct() *ApiGetRouteRequestStream {
	caches := map[string]bool{}
	result := ApiGetRouteRequestStreamOf()
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
func (self *ApiGetRouteRequestStream) Each(fn func(ApiGetRouteRequest)) *ApiGetRouteRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetRouteRequestStream) EachRight(fn func(ApiGetRouteRequest)) *ApiGetRouteRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetRouteRequestStream) Equals(arr []ApiGetRouteRequest) bool {
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
func (self *ApiGetRouteRequestStream) Filter(fn func(ApiGetRouteRequest, int) bool) *ApiGetRouteRequestStream {
	result := ApiGetRouteRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetRouteRequestStream) FilterSlim(fn func(ApiGetRouteRequest, int) bool) *ApiGetRouteRequestStream {
	result := ApiGetRouteRequestStreamOf()
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
func (self *ApiGetRouteRequestStream) Find(fn func(ApiGetRouteRequest, int) bool) *ApiGetRouteRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetRouteRequestStream) FindOr(fn func(ApiGetRouteRequest, int) bool, or ApiGetRouteRequest) ApiGetRouteRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetRouteRequestStream) FindIndex(fn func(ApiGetRouteRequest, int) bool) int {
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
func (self *ApiGetRouteRequestStream) First() *ApiGetRouteRequest {
	return self.Get(0)
}
func (self *ApiGetRouteRequestStream) FirstOr(arg ApiGetRouteRequest) ApiGetRouteRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetRouteRequestStream) ForEach(fn func(ApiGetRouteRequest, int)) *ApiGetRouteRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetRouteRequestStream) ForEachRight(fn func(ApiGetRouteRequest, int)) *ApiGetRouteRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetRouteRequestStream) GroupBy(fn func(ApiGetRouteRequest, int) string) map[string][]ApiGetRouteRequest {
	m := map[string][]ApiGetRouteRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetRouteRequestStream) GroupByValues(fn func(ApiGetRouteRequest, int) string) [][]ApiGetRouteRequest {
	var tmp [][]ApiGetRouteRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetRouteRequestStream) IndexOf(arg ApiGetRouteRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetRouteRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetRouteRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetRouteRequestStream) Last() *ApiGetRouteRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetRouteRequestStream) LastOr(arg ApiGetRouteRequest) ApiGetRouteRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetRouteRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetRouteRequestStream) Limit(limit int) *ApiGetRouteRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetRouteRequestStream) Map(fn func(ApiGetRouteRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Map2Int(fn func(ApiGetRouteRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Map2Int32(fn func(ApiGetRouteRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Map2Int64(fn func(ApiGetRouteRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Map2Float32(fn func(ApiGetRouteRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Map2Float64(fn func(ApiGetRouteRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Map2Bool(fn func(ApiGetRouteRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Map2Bytes(fn func(ApiGetRouteRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Map2String(fn func(ApiGetRouteRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Max(fn func(ApiGetRouteRequest, int) float64) *ApiGetRouteRequest {
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
func (self *ApiGetRouteRequestStream) Min(fn func(ApiGetRouteRequest, int) float64) *ApiGetRouteRequest {
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
func (self *ApiGetRouteRequestStream) NoneMatch(fn func(ApiGetRouteRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetRouteRequestStream) Get(index int) *ApiGetRouteRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetRouteRequestStream) GetOr(index int, arg ApiGetRouteRequest) ApiGetRouteRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetRouteRequestStream) Peek(fn func(*ApiGetRouteRequest, int)) *ApiGetRouteRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetRouteRequestStream) Reduce(fn func(ApiGetRouteRequest, ApiGetRouteRequest, int) ApiGetRouteRequest) *ApiGetRouteRequestStream {
	return self.ReduceInit(fn, ApiGetRouteRequest{})
}
func (self *ApiGetRouteRequestStream) ReduceInit(fn func(ApiGetRouteRequest, ApiGetRouteRequest, int) ApiGetRouteRequest, initialValue ApiGetRouteRequest) *ApiGetRouteRequestStream {
	result := ApiGetRouteRequestStreamOf()
	self.ForEach(func(v ApiGetRouteRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetRouteRequestStream) ReduceInterface(fn func(interface{}, ApiGetRouteRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetRouteRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetRouteRequestStream) ReduceString(fn func(string, ApiGetRouteRequest, int) string) []string {
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
func (self *ApiGetRouteRequestStream) ReduceInt(fn func(int, ApiGetRouteRequest, int) int) []int {
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
func (self *ApiGetRouteRequestStream) ReduceInt32(fn func(int32, ApiGetRouteRequest, int) int32) []int32 {
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
func (self *ApiGetRouteRequestStream) ReduceInt64(fn func(int64, ApiGetRouteRequest, int) int64) []int64 {
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
func (self *ApiGetRouteRequestStream) ReduceFloat32(fn func(float32, ApiGetRouteRequest, int) float32) []float32 {
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
func (self *ApiGetRouteRequestStream) ReduceFloat64(fn func(float64, ApiGetRouteRequest, int) float64) []float64 {
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
func (self *ApiGetRouteRequestStream) ReduceBool(fn func(bool, ApiGetRouteRequest, int) bool) []bool {
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
func (self *ApiGetRouteRequestStream) Reverse() *ApiGetRouteRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetRouteRequestStream) Replace(fn func(ApiGetRouteRequest, int) ApiGetRouteRequest) *ApiGetRouteRequestStream {
	return self.ForEach(func(v ApiGetRouteRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetRouteRequestStream) Select(fn func(ApiGetRouteRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetRouteRequestStream) Set(index int, val ApiGetRouteRequest) *ApiGetRouteRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetRouteRequestStream) Skip(skip int) *ApiGetRouteRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetRouteRequestStream) SkippingEach(fn func(ApiGetRouteRequest, int) int) *ApiGetRouteRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetRouteRequestStream) Slice(startIndex, n int) *ApiGetRouteRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetRouteRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetRouteRequestStream) Sort(fn func(i, j int) bool) *ApiGetRouteRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetRouteRequestStream) Tail() *ApiGetRouteRequest {
	return self.Last()
}
func (self *ApiGetRouteRequestStream) TailOr(arg ApiGetRouteRequest) ApiGetRouteRequest {
	return self.LastOr(arg)
}
func (self *ApiGetRouteRequestStream) ToList() []ApiGetRouteRequest {
	return self.Val()
}
func (self *ApiGetRouteRequestStream) Unique() *ApiGetRouteRequestStream {
	return self.Distinct()
}
func (self *ApiGetRouteRequestStream) Val() []ApiGetRouteRequest {
	if self == nil {
		return []ApiGetRouteRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetRouteRequestStream) While(fn func(ApiGetRouteRequest, int) bool) *ApiGetRouteRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetRouteRequestStream) Where(fn func(ApiGetRouteRequest) bool) *ApiGetRouteRequestStream {
	result := ApiGetRouteRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetRouteRequestStream) WhereSlim(fn func(ApiGetRouteRequest) bool) *ApiGetRouteRequestStream {
	result := ApiGetRouteRequestStreamOf()
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
