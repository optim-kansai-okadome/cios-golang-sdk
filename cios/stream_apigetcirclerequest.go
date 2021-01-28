/*
 * Collection utility of ApiGetCircleRequest Struct
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

type ApiGetCircleRequestStream []ApiGetCircleRequest

func ApiGetCircleRequestStreamOf(arg ...ApiGetCircleRequest) ApiGetCircleRequestStream {
	return arg
}
func ApiGetCircleRequestStreamFrom(arg []ApiGetCircleRequest) ApiGetCircleRequestStream {
	return arg
}
func CreateApiGetCircleRequestStream(arg ...ApiGetCircleRequest) *ApiGetCircleRequestStream {
	tmp := ApiGetCircleRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetCircleRequestStream(arg []ApiGetCircleRequest) *ApiGetCircleRequestStream {
	tmp := ApiGetCircleRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetCircleRequestStream) Add(arg ApiGetCircleRequest) *ApiGetCircleRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetCircleRequestStream) AddAll(arg ...ApiGetCircleRequest) *ApiGetCircleRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetCircleRequestStream) AddSafe(arg *ApiGetCircleRequest) *ApiGetCircleRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetCircleRequestStream) Aggregate(fn func(ApiGetCircleRequest, ApiGetCircleRequest) ApiGetCircleRequest) *ApiGetCircleRequestStream {
	result := ApiGetCircleRequestStreamOf()
	self.ForEach(func(v ApiGetCircleRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetCircleRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetCircleRequestStream) AllMatch(fn func(ApiGetCircleRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetCircleRequestStream) AnyMatch(fn func(ApiGetCircleRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetCircleRequestStream) Clone() *ApiGetCircleRequestStream {
	temp := make([]ApiGetCircleRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetCircleRequestStream)(&temp)
}
func (self *ApiGetCircleRequestStream) Copy() *ApiGetCircleRequestStream {
	return self.Clone()
}
func (self *ApiGetCircleRequestStream) Concat(arg []ApiGetCircleRequest) *ApiGetCircleRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetCircleRequestStream) Contains(arg ApiGetCircleRequest) bool {
	return self.FindIndex(func(_arg ApiGetCircleRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetCircleRequestStream) Clean() *ApiGetCircleRequestStream {
	*self = ApiGetCircleRequestStreamOf()
	return self
}
func (self *ApiGetCircleRequestStream) Delete(index int) *ApiGetCircleRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetCircleRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetCircleRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetCircleRequestStream) Distinct() *ApiGetCircleRequestStream {
	caches := map[string]bool{}
	result := ApiGetCircleRequestStreamOf()
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
func (self *ApiGetCircleRequestStream) Each(fn func(ApiGetCircleRequest)) *ApiGetCircleRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetCircleRequestStream) EachRight(fn func(ApiGetCircleRequest)) *ApiGetCircleRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetCircleRequestStream) Equals(arr []ApiGetCircleRequest) bool {
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
func (self *ApiGetCircleRequestStream) Filter(fn func(ApiGetCircleRequest, int) bool) *ApiGetCircleRequestStream {
	result := ApiGetCircleRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetCircleRequestStream) FilterSlim(fn func(ApiGetCircleRequest, int) bool) *ApiGetCircleRequestStream {
	result := ApiGetCircleRequestStreamOf()
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
func (self *ApiGetCircleRequestStream) Find(fn func(ApiGetCircleRequest, int) bool) *ApiGetCircleRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetCircleRequestStream) FindOr(fn func(ApiGetCircleRequest, int) bool, or ApiGetCircleRequest) ApiGetCircleRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetCircleRequestStream) FindIndex(fn func(ApiGetCircleRequest, int) bool) int {
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
func (self *ApiGetCircleRequestStream) First() *ApiGetCircleRequest {
	return self.Get(0)
}
func (self *ApiGetCircleRequestStream) FirstOr(arg ApiGetCircleRequest) ApiGetCircleRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetCircleRequestStream) ForEach(fn func(ApiGetCircleRequest, int)) *ApiGetCircleRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetCircleRequestStream) ForEachRight(fn func(ApiGetCircleRequest, int)) *ApiGetCircleRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetCircleRequestStream) GroupBy(fn func(ApiGetCircleRequest, int) string) map[string][]ApiGetCircleRequest {
	m := map[string][]ApiGetCircleRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetCircleRequestStream) GroupByValues(fn func(ApiGetCircleRequest, int) string) [][]ApiGetCircleRequest {
	var tmp [][]ApiGetCircleRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetCircleRequestStream) IndexOf(arg ApiGetCircleRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetCircleRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetCircleRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetCircleRequestStream) Last() *ApiGetCircleRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetCircleRequestStream) LastOr(arg ApiGetCircleRequest) ApiGetCircleRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetCircleRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetCircleRequestStream) Limit(limit int) *ApiGetCircleRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetCircleRequestStream) Map(fn func(ApiGetCircleRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Map2Int(fn func(ApiGetCircleRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Map2Int32(fn func(ApiGetCircleRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Map2Int64(fn func(ApiGetCircleRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Map2Float32(fn func(ApiGetCircleRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Map2Float64(fn func(ApiGetCircleRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Map2Bool(fn func(ApiGetCircleRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Map2Bytes(fn func(ApiGetCircleRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Map2String(fn func(ApiGetCircleRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Max(fn func(ApiGetCircleRequest, int) float64) *ApiGetCircleRequest {
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
func (self *ApiGetCircleRequestStream) Min(fn func(ApiGetCircleRequest, int) float64) *ApiGetCircleRequest {
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
func (self *ApiGetCircleRequestStream) NoneMatch(fn func(ApiGetCircleRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetCircleRequestStream) Get(index int) *ApiGetCircleRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetCircleRequestStream) GetOr(index int, arg ApiGetCircleRequest) ApiGetCircleRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetCircleRequestStream) Peek(fn func(*ApiGetCircleRequest, int)) *ApiGetCircleRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetCircleRequestStream) Reduce(fn func(ApiGetCircleRequest, ApiGetCircleRequest, int) ApiGetCircleRequest) *ApiGetCircleRequestStream {
	return self.ReduceInit(fn, ApiGetCircleRequest{})
}
func (self *ApiGetCircleRequestStream) ReduceInit(fn func(ApiGetCircleRequest, ApiGetCircleRequest, int) ApiGetCircleRequest, initialValue ApiGetCircleRequest) *ApiGetCircleRequestStream {
	result := ApiGetCircleRequestStreamOf()
	self.ForEach(func(v ApiGetCircleRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetCircleRequestStream) ReduceInterface(fn func(interface{}, ApiGetCircleRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetCircleRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetCircleRequestStream) ReduceString(fn func(string, ApiGetCircleRequest, int) string) []string {
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
func (self *ApiGetCircleRequestStream) ReduceInt(fn func(int, ApiGetCircleRequest, int) int) []int {
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
func (self *ApiGetCircleRequestStream) ReduceInt32(fn func(int32, ApiGetCircleRequest, int) int32) []int32 {
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
func (self *ApiGetCircleRequestStream) ReduceInt64(fn func(int64, ApiGetCircleRequest, int) int64) []int64 {
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
func (self *ApiGetCircleRequestStream) ReduceFloat32(fn func(float32, ApiGetCircleRequest, int) float32) []float32 {
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
func (self *ApiGetCircleRequestStream) ReduceFloat64(fn func(float64, ApiGetCircleRequest, int) float64) []float64 {
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
func (self *ApiGetCircleRequestStream) ReduceBool(fn func(bool, ApiGetCircleRequest, int) bool) []bool {
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
func (self *ApiGetCircleRequestStream) Reverse() *ApiGetCircleRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetCircleRequestStream) Replace(fn func(ApiGetCircleRequest, int) ApiGetCircleRequest) *ApiGetCircleRequestStream {
	return self.ForEach(func(v ApiGetCircleRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetCircleRequestStream) Select(fn func(ApiGetCircleRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetCircleRequestStream) Set(index int, val ApiGetCircleRequest) *ApiGetCircleRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetCircleRequestStream) Skip(skip int) *ApiGetCircleRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetCircleRequestStream) SkippingEach(fn func(ApiGetCircleRequest, int) int) *ApiGetCircleRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetCircleRequestStream) Slice(startIndex, n int) *ApiGetCircleRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetCircleRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetCircleRequestStream) Sort(fn func(i, j int) bool) *ApiGetCircleRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetCircleRequestStream) Tail() *ApiGetCircleRequest {
	return self.Last()
}
func (self *ApiGetCircleRequestStream) TailOr(arg ApiGetCircleRequest) ApiGetCircleRequest {
	return self.LastOr(arg)
}
func (self *ApiGetCircleRequestStream) ToList() []ApiGetCircleRequest {
	return self.Val()
}
func (self *ApiGetCircleRequestStream) Unique() *ApiGetCircleRequestStream {
	return self.Distinct()
}
func (self *ApiGetCircleRequestStream) Val() []ApiGetCircleRequest {
	if self == nil {
		return []ApiGetCircleRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetCircleRequestStream) While(fn func(ApiGetCircleRequest, int) bool) *ApiGetCircleRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetCircleRequestStream) Where(fn func(ApiGetCircleRequest) bool) *ApiGetCircleRequestStream {
	result := ApiGetCircleRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetCircleRequestStream) WhereSlim(fn func(ApiGetCircleRequest) bool) *ApiGetCircleRequestStream {
	result := ApiGetCircleRequestStreamOf()
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
