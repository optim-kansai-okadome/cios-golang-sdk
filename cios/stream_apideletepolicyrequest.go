/*
 * Collection utility of ApiDeletePolicyRequest Struct
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

type ApiDeletePolicyRequestStream []ApiDeletePolicyRequest

func ApiDeletePolicyRequestStreamOf(arg ...ApiDeletePolicyRequest) ApiDeletePolicyRequestStream {
	return arg
}
func ApiDeletePolicyRequestStreamFrom(arg []ApiDeletePolicyRequest) ApiDeletePolicyRequestStream {
	return arg
}
func CreateApiDeletePolicyRequestStream(arg ...ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	tmp := ApiDeletePolicyRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiDeletePolicyRequestStream(arg []ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	tmp := ApiDeletePolicyRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiDeletePolicyRequestStream) Add(arg ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	return self.AddAll(arg)
}
func (self *ApiDeletePolicyRequestStream) AddAll(arg ...ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiDeletePolicyRequestStream) AddSafe(arg *ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) Aggregate(fn func(ApiDeletePolicyRequest, ApiDeletePolicyRequest) ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	result := ApiDeletePolicyRequestStreamOf()
	self.ForEach(func(v ApiDeletePolicyRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiDeletePolicyRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiDeletePolicyRequestStream) AllMatch(fn func(ApiDeletePolicyRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiDeletePolicyRequestStream) AnyMatch(fn func(ApiDeletePolicyRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiDeletePolicyRequestStream) Clone() *ApiDeletePolicyRequestStream {
	temp := make([]ApiDeletePolicyRequest, self.Len())
	copy(temp, *self)
	return (*ApiDeletePolicyRequestStream)(&temp)
}
func (self *ApiDeletePolicyRequestStream) Copy() *ApiDeletePolicyRequestStream {
	return self.Clone()
}
func (self *ApiDeletePolicyRequestStream) Concat(arg []ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiDeletePolicyRequestStream) Contains(arg ApiDeletePolicyRequest) bool {
	return self.FindIndex(func(_arg ApiDeletePolicyRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiDeletePolicyRequestStream) Clean() *ApiDeletePolicyRequestStream {
	*self = ApiDeletePolicyRequestStreamOf()
	return self
}
func (self *ApiDeletePolicyRequestStream) Delete(index int) *ApiDeletePolicyRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiDeletePolicyRequestStream) DeleteRange(startIndex, endIndex int) *ApiDeletePolicyRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiDeletePolicyRequestStream) Distinct() *ApiDeletePolicyRequestStream {
	caches := map[string]bool{}
	result := ApiDeletePolicyRequestStreamOf()
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
func (self *ApiDeletePolicyRequestStream) Each(fn func(ApiDeletePolicyRequest)) *ApiDeletePolicyRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) EachRight(fn func(ApiDeletePolicyRequest)) *ApiDeletePolicyRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) Equals(arr []ApiDeletePolicyRequest) bool {
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
func (self *ApiDeletePolicyRequestStream) Filter(fn func(ApiDeletePolicyRequest, int) bool) *ApiDeletePolicyRequestStream {
	result := ApiDeletePolicyRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeletePolicyRequestStream) FilterSlim(fn func(ApiDeletePolicyRequest, int) bool) *ApiDeletePolicyRequestStream {
	result := ApiDeletePolicyRequestStreamOf()
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
func (self *ApiDeletePolicyRequestStream) Find(fn func(ApiDeletePolicyRequest, int) bool) *ApiDeletePolicyRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiDeletePolicyRequestStream) FindOr(fn func(ApiDeletePolicyRequest, int) bool, or ApiDeletePolicyRequest) ApiDeletePolicyRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiDeletePolicyRequestStream) FindIndex(fn func(ApiDeletePolicyRequest, int) bool) int {
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
func (self *ApiDeletePolicyRequestStream) First() *ApiDeletePolicyRequest {
	return self.Get(0)
}
func (self *ApiDeletePolicyRequestStream) FirstOr(arg ApiDeletePolicyRequest) ApiDeletePolicyRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeletePolicyRequestStream) ForEach(fn func(ApiDeletePolicyRequest, int)) *ApiDeletePolicyRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) ForEachRight(fn func(ApiDeletePolicyRequest, int)) *ApiDeletePolicyRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) GroupBy(fn func(ApiDeletePolicyRequest, int) string) map[string][]ApiDeletePolicyRequest {
	m := map[string][]ApiDeletePolicyRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiDeletePolicyRequestStream) GroupByValues(fn func(ApiDeletePolicyRequest, int) string) [][]ApiDeletePolicyRequest {
	var tmp [][]ApiDeletePolicyRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiDeletePolicyRequestStream) IndexOf(arg ApiDeletePolicyRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiDeletePolicyRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiDeletePolicyRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiDeletePolicyRequestStream) Last() *ApiDeletePolicyRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiDeletePolicyRequestStream) LastOr(arg ApiDeletePolicyRequest) ApiDeletePolicyRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeletePolicyRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiDeletePolicyRequestStream) Limit(limit int) *ApiDeletePolicyRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiDeletePolicyRequestStream) Map(fn func(ApiDeletePolicyRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Map2Int(fn func(ApiDeletePolicyRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Map2Int32(fn func(ApiDeletePolicyRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Map2Int64(fn func(ApiDeletePolicyRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Map2Float32(fn func(ApiDeletePolicyRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Map2Float64(fn func(ApiDeletePolicyRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Map2Bool(fn func(ApiDeletePolicyRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Map2Bytes(fn func(ApiDeletePolicyRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Map2String(fn func(ApiDeletePolicyRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Max(fn func(ApiDeletePolicyRequest, int) float64) *ApiDeletePolicyRequest {
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
func (self *ApiDeletePolicyRequestStream) Min(fn func(ApiDeletePolicyRequest, int) float64) *ApiDeletePolicyRequest {
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
func (self *ApiDeletePolicyRequestStream) NoneMatch(fn func(ApiDeletePolicyRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiDeletePolicyRequestStream) Get(index int) *ApiDeletePolicyRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiDeletePolicyRequestStream) GetOr(index int, arg ApiDeletePolicyRequest) ApiDeletePolicyRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeletePolicyRequestStream) Peek(fn func(*ApiDeletePolicyRequest, int)) *ApiDeletePolicyRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiDeletePolicyRequestStream) Reduce(fn func(ApiDeletePolicyRequest, ApiDeletePolicyRequest, int) ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	return self.ReduceInit(fn, ApiDeletePolicyRequest{})
}
func (self *ApiDeletePolicyRequestStream) ReduceInit(fn func(ApiDeletePolicyRequest, ApiDeletePolicyRequest, int) ApiDeletePolicyRequest, initialValue ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	result := ApiDeletePolicyRequestStreamOf()
	self.ForEach(func(v ApiDeletePolicyRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiDeletePolicyRequestStream) ReduceInterface(fn func(interface{}, ApiDeletePolicyRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiDeletePolicyRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiDeletePolicyRequestStream) ReduceString(fn func(string, ApiDeletePolicyRequest, int) string) []string {
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
func (self *ApiDeletePolicyRequestStream) ReduceInt(fn func(int, ApiDeletePolicyRequest, int) int) []int {
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
func (self *ApiDeletePolicyRequestStream) ReduceInt32(fn func(int32, ApiDeletePolicyRequest, int) int32) []int32 {
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
func (self *ApiDeletePolicyRequestStream) ReduceInt64(fn func(int64, ApiDeletePolicyRequest, int) int64) []int64 {
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
func (self *ApiDeletePolicyRequestStream) ReduceFloat32(fn func(float32, ApiDeletePolicyRequest, int) float32) []float32 {
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
func (self *ApiDeletePolicyRequestStream) ReduceFloat64(fn func(float64, ApiDeletePolicyRequest, int) float64) []float64 {
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
func (self *ApiDeletePolicyRequestStream) ReduceBool(fn func(bool, ApiDeletePolicyRequest, int) bool) []bool {
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
func (self *ApiDeletePolicyRequestStream) Reverse() *ApiDeletePolicyRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) Replace(fn func(ApiDeletePolicyRequest, int) ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	return self.ForEach(func(v ApiDeletePolicyRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiDeletePolicyRequestStream) Select(fn func(ApiDeletePolicyRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiDeletePolicyRequestStream) Set(index int, val ApiDeletePolicyRequest) *ApiDeletePolicyRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) Skip(skip int) *ApiDeletePolicyRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiDeletePolicyRequestStream) SkippingEach(fn func(ApiDeletePolicyRequest, int) int) *ApiDeletePolicyRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) Slice(startIndex, n int) *ApiDeletePolicyRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiDeletePolicyRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) Sort(fn func(i, j int) bool) *ApiDeletePolicyRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiDeletePolicyRequestStream) Tail() *ApiDeletePolicyRequest {
	return self.Last()
}
func (self *ApiDeletePolicyRequestStream) TailOr(arg ApiDeletePolicyRequest) ApiDeletePolicyRequest {
	return self.LastOr(arg)
}
func (self *ApiDeletePolicyRequestStream) ToList() []ApiDeletePolicyRequest {
	return self.Val()
}
func (self *ApiDeletePolicyRequestStream) Unique() *ApiDeletePolicyRequestStream {
	return self.Distinct()
}
func (self *ApiDeletePolicyRequestStream) Val() []ApiDeletePolicyRequest {
	if self == nil {
		return []ApiDeletePolicyRequest{}
	}
	return *self.Copy()
}
func (self *ApiDeletePolicyRequestStream) While(fn func(ApiDeletePolicyRequest, int) bool) *ApiDeletePolicyRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiDeletePolicyRequestStream) Where(fn func(ApiDeletePolicyRequest) bool) *ApiDeletePolicyRequestStream {
	result := ApiDeletePolicyRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeletePolicyRequestStream) WhereSlim(fn func(ApiDeletePolicyRequest) bool) *ApiDeletePolicyRequestStream {
	result := ApiDeletePolicyRequestStreamOf()
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
