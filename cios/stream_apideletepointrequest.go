/*
 * Collection utility of ApiDeletePointRequest Struct
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

type ApiDeletePointRequestStream []ApiDeletePointRequest

func ApiDeletePointRequestStreamOf(arg ...ApiDeletePointRequest) ApiDeletePointRequestStream {
	return arg
}
func ApiDeletePointRequestStreamFrom(arg []ApiDeletePointRequest) ApiDeletePointRequestStream {
	return arg
}
func CreateApiDeletePointRequestStream(arg ...ApiDeletePointRequest) *ApiDeletePointRequestStream {
	tmp := ApiDeletePointRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiDeletePointRequestStream(arg []ApiDeletePointRequest) *ApiDeletePointRequestStream {
	tmp := ApiDeletePointRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiDeletePointRequestStream) Add(arg ApiDeletePointRequest) *ApiDeletePointRequestStream {
	return self.AddAll(arg)
}
func (self *ApiDeletePointRequestStream) AddAll(arg ...ApiDeletePointRequest) *ApiDeletePointRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiDeletePointRequestStream) AddSafe(arg *ApiDeletePointRequest) *ApiDeletePointRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiDeletePointRequestStream) Aggregate(fn func(ApiDeletePointRequest, ApiDeletePointRequest) ApiDeletePointRequest) *ApiDeletePointRequestStream {
	result := ApiDeletePointRequestStreamOf()
	self.ForEach(func(v ApiDeletePointRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiDeletePointRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiDeletePointRequestStream) AllMatch(fn func(ApiDeletePointRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiDeletePointRequestStream) AnyMatch(fn func(ApiDeletePointRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiDeletePointRequestStream) Clone() *ApiDeletePointRequestStream {
	temp := make([]ApiDeletePointRequest, self.Len())
	copy(temp, *self)
	return (*ApiDeletePointRequestStream)(&temp)
}
func (self *ApiDeletePointRequestStream) Copy() *ApiDeletePointRequestStream {
	return self.Clone()
}
func (self *ApiDeletePointRequestStream) Concat(arg []ApiDeletePointRequest) *ApiDeletePointRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiDeletePointRequestStream) Contains(arg ApiDeletePointRequest) bool {
	return self.FindIndex(func(_arg ApiDeletePointRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiDeletePointRequestStream) Clean() *ApiDeletePointRequestStream {
	*self = ApiDeletePointRequestStreamOf()
	return self
}
func (self *ApiDeletePointRequestStream) Delete(index int) *ApiDeletePointRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiDeletePointRequestStream) DeleteRange(startIndex, endIndex int) *ApiDeletePointRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiDeletePointRequestStream) Distinct() *ApiDeletePointRequestStream {
	caches := map[string]bool{}
	result := ApiDeletePointRequestStreamOf()
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
func (self *ApiDeletePointRequestStream) Each(fn func(ApiDeletePointRequest)) *ApiDeletePointRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiDeletePointRequestStream) EachRight(fn func(ApiDeletePointRequest)) *ApiDeletePointRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiDeletePointRequestStream) Equals(arr []ApiDeletePointRequest) bool {
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
func (self *ApiDeletePointRequestStream) Filter(fn func(ApiDeletePointRequest, int) bool) *ApiDeletePointRequestStream {
	result := ApiDeletePointRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeletePointRequestStream) FilterSlim(fn func(ApiDeletePointRequest, int) bool) *ApiDeletePointRequestStream {
	result := ApiDeletePointRequestStreamOf()
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
func (self *ApiDeletePointRequestStream) Find(fn func(ApiDeletePointRequest, int) bool) *ApiDeletePointRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiDeletePointRequestStream) FindOr(fn func(ApiDeletePointRequest, int) bool, or ApiDeletePointRequest) ApiDeletePointRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiDeletePointRequestStream) FindIndex(fn func(ApiDeletePointRequest, int) bool) int {
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
func (self *ApiDeletePointRequestStream) First() *ApiDeletePointRequest {
	return self.Get(0)
}
func (self *ApiDeletePointRequestStream) FirstOr(arg ApiDeletePointRequest) ApiDeletePointRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeletePointRequestStream) ForEach(fn func(ApiDeletePointRequest, int)) *ApiDeletePointRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiDeletePointRequestStream) ForEachRight(fn func(ApiDeletePointRequest, int)) *ApiDeletePointRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiDeletePointRequestStream) GroupBy(fn func(ApiDeletePointRequest, int) string) map[string][]ApiDeletePointRequest {
	m := map[string][]ApiDeletePointRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiDeletePointRequestStream) GroupByValues(fn func(ApiDeletePointRequest, int) string) [][]ApiDeletePointRequest {
	var tmp [][]ApiDeletePointRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiDeletePointRequestStream) IndexOf(arg ApiDeletePointRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiDeletePointRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiDeletePointRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiDeletePointRequestStream) Last() *ApiDeletePointRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiDeletePointRequestStream) LastOr(arg ApiDeletePointRequest) ApiDeletePointRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeletePointRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiDeletePointRequestStream) Limit(limit int) *ApiDeletePointRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiDeletePointRequestStream) Map(fn func(ApiDeletePointRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Map2Int(fn func(ApiDeletePointRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Map2Int32(fn func(ApiDeletePointRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Map2Int64(fn func(ApiDeletePointRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Map2Float32(fn func(ApiDeletePointRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Map2Float64(fn func(ApiDeletePointRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Map2Bool(fn func(ApiDeletePointRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Map2Bytes(fn func(ApiDeletePointRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Map2String(fn func(ApiDeletePointRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Max(fn func(ApiDeletePointRequest, int) float64) *ApiDeletePointRequest {
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
func (self *ApiDeletePointRequestStream) Min(fn func(ApiDeletePointRequest, int) float64) *ApiDeletePointRequest {
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
func (self *ApiDeletePointRequestStream) NoneMatch(fn func(ApiDeletePointRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiDeletePointRequestStream) Get(index int) *ApiDeletePointRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiDeletePointRequestStream) GetOr(index int, arg ApiDeletePointRequest) ApiDeletePointRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeletePointRequestStream) Peek(fn func(*ApiDeletePointRequest, int)) *ApiDeletePointRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiDeletePointRequestStream) Reduce(fn func(ApiDeletePointRequest, ApiDeletePointRequest, int) ApiDeletePointRequest) *ApiDeletePointRequestStream {
	return self.ReduceInit(fn, ApiDeletePointRequest{})
}
func (self *ApiDeletePointRequestStream) ReduceInit(fn func(ApiDeletePointRequest, ApiDeletePointRequest, int) ApiDeletePointRequest, initialValue ApiDeletePointRequest) *ApiDeletePointRequestStream {
	result := ApiDeletePointRequestStreamOf()
	self.ForEach(func(v ApiDeletePointRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiDeletePointRequestStream) ReduceInterface(fn func(interface{}, ApiDeletePointRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiDeletePointRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiDeletePointRequestStream) ReduceString(fn func(string, ApiDeletePointRequest, int) string) []string {
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
func (self *ApiDeletePointRequestStream) ReduceInt(fn func(int, ApiDeletePointRequest, int) int) []int {
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
func (self *ApiDeletePointRequestStream) ReduceInt32(fn func(int32, ApiDeletePointRequest, int) int32) []int32 {
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
func (self *ApiDeletePointRequestStream) ReduceInt64(fn func(int64, ApiDeletePointRequest, int) int64) []int64 {
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
func (self *ApiDeletePointRequestStream) ReduceFloat32(fn func(float32, ApiDeletePointRequest, int) float32) []float32 {
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
func (self *ApiDeletePointRequestStream) ReduceFloat64(fn func(float64, ApiDeletePointRequest, int) float64) []float64 {
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
func (self *ApiDeletePointRequestStream) ReduceBool(fn func(bool, ApiDeletePointRequest, int) bool) []bool {
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
func (self *ApiDeletePointRequestStream) Reverse() *ApiDeletePointRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiDeletePointRequestStream) Replace(fn func(ApiDeletePointRequest, int) ApiDeletePointRequest) *ApiDeletePointRequestStream {
	return self.ForEach(func(v ApiDeletePointRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiDeletePointRequestStream) Select(fn func(ApiDeletePointRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiDeletePointRequestStream) Set(index int, val ApiDeletePointRequest) *ApiDeletePointRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiDeletePointRequestStream) Skip(skip int) *ApiDeletePointRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiDeletePointRequestStream) SkippingEach(fn func(ApiDeletePointRequest, int) int) *ApiDeletePointRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiDeletePointRequestStream) Slice(startIndex, n int) *ApiDeletePointRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiDeletePointRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiDeletePointRequestStream) Sort(fn func(i, j int) bool) *ApiDeletePointRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiDeletePointRequestStream) Tail() *ApiDeletePointRequest {
	return self.Last()
}
func (self *ApiDeletePointRequestStream) TailOr(arg ApiDeletePointRequest) ApiDeletePointRequest {
	return self.LastOr(arg)
}
func (self *ApiDeletePointRequestStream) ToList() []ApiDeletePointRequest {
	return self.Val()
}
func (self *ApiDeletePointRequestStream) Unique() *ApiDeletePointRequestStream {
	return self.Distinct()
}
func (self *ApiDeletePointRequestStream) Val() []ApiDeletePointRequest {
	if self == nil {
		return []ApiDeletePointRequest{}
	}
	return *self.Copy()
}
func (self *ApiDeletePointRequestStream) While(fn func(ApiDeletePointRequest, int) bool) *ApiDeletePointRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiDeletePointRequestStream) Where(fn func(ApiDeletePointRequest) bool) *ApiDeletePointRequestStream {
	result := ApiDeletePointRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeletePointRequestStream) WhereSlim(fn func(ApiDeletePointRequest) bool) *ApiDeletePointRequestStream {
	result := ApiDeletePointRequestStreamOf()
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
