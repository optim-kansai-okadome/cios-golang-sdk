/*
 * Collection utility of ApiCreateChannelRequest Struct
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

type ApiCreateChannelRequestStream []ApiCreateChannelRequest

func ApiCreateChannelRequestStreamOf(arg ...ApiCreateChannelRequest) ApiCreateChannelRequestStream {
	return arg
}
func ApiCreateChannelRequestStreamFrom(arg []ApiCreateChannelRequest) ApiCreateChannelRequestStream {
	return arg
}
func CreateApiCreateChannelRequestStream(arg ...ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	tmp := ApiCreateChannelRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiCreateChannelRequestStream(arg []ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	tmp := ApiCreateChannelRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiCreateChannelRequestStream) Add(arg ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	return self.AddAll(arg)
}
func (self *ApiCreateChannelRequestStream) AddAll(arg ...ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiCreateChannelRequestStream) AddSafe(arg *ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiCreateChannelRequestStream) Aggregate(fn func(ApiCreateChannelRequest, ApiCreateChannelRequest) ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	result := ApiCreateChannelRequestStreamOf()
	self.ForEach(func(v ApiCreateChannelRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiCreateChannelRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateChannelRequestStream) AllMatch(fn func(ApiCreateChannelRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiCreateChannelRequestStream) AnyMatch(fn func(ApiCreateChannelRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiCreateChannelRequestStream) Clone() *ApiCreateChannelRequestStream {
	temp := make([]ApiCreateChannelRequest, self.Len())
	copy(temp, *self)
	return (*ApiCreateChannelRequestStream)(&temp)
}
func (self *ApiCreateChannelRequestStream) Copy() *ApiCreateChannelRequestStream {
	return self.Clone()
}
func (self *ApiCreateChannelRequestStream) Concat(arg []ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiCreateChannelRequestStream) Contains(arg ApiCreateChannelRequest) bool {
	return self.FindIndex(func(_arg ApiCreateChannelRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiCreateChannelRequestStream) Clean() *ApiCreateChannelRequestStream {
	*self = ApiCreateChannelRequestStreamOf()
	return self
}
func (self *ApiCreateChannelRequestStream) Delete(index int) *ApiCreateChannelRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiCreateChannelRequestStream) DeleteRange(startIndex, endIndex int) *ApiCreateChannelRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiCreateChannelRequestStream) Distinct() *ApiCreateChannelRequestStream {
	caches := map[string]bool{}
	result := ApiCreateChannelRequestStreamOf()
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
func (self *ApiCreateChannelRequestStream) Each(fn func(ApiCreateChannelRequest)) *ApiCreateChannelRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiCreateChannelRequestStream) EachRight(fn func(ApiCreateChannelRequest)) *ApiCreateChannelRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiCreateChannelRequestStream) Equals(arr []ApiCreateChannelRequest) bool {
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
func (self *ApiCreateChannelRequestStream) Filter(fn func(ApiCreateChannelRequest, int) bool) *ApiCreateChannelRequestStream {
	result := ApiCreateChannelRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateChannelRequestStream) FilterSlim(fn func(ApiCreateChannelRequest, int) bool) *ApiCreateChannelRequestStream {
	result := ApiCreateChannelRequestStreamOf()
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
func (self *ApiCreateChannelRequestStream) Find(fn func(ApiCreateChannelRequest, int) bool) *ApiCreateChannelRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiCreateChannelRequestStream) FindOr(fn func(ApiCreateChannelRequest, int) bool, or ApiCreateChannelRequest) ApiCreateChannelRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiCreateChannelRequestStream) FindIndex(fn func(ApiCreateChannelRequest, int) bool) int {
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
func (self *ApiCreateChannelRequestStream) First() *ApiCreateChannelRequest {
	return self.Get(0)
}
func (self *ApiCreateChannelRequestStream) FirstOr(arg ApiCreateChannelRequest) ApiCreateChannelRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateChannelRequestStream) ForEach(fn func(ApiCreateChannelRequest, int)) *ApiCreateChannelRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiCreateChannelRequestStream) ForEachRight(fn func(ApiCreateChannelRequest, int)) *ApiCreateChannelRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiCreateChannelRequestStream) GroupBy(fn func(ApiCreateChannelRequest, int) string) map[string][]ApiCreateChannelRequest {
	m := map[string][]ApiCreateChannelRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiCreateChannelRequestStream) GroupByValues(fn func(ApiCreateChannelRequest, int) string) [][]ApiCreateChannelRequest {
	var tmp [][]ApiCreateChannelRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiCreateChannelRequestStream) IndexOf(arg ApiCreateChannelRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiCreateChannelRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiCreateChannelRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiCreateChannelRequestStream) Last() *ApiCreateChannelRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiCreateChannelRequestStream) LastOr(arg ApiCreateChannelRequest) ApiCreateChannelRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateChannelRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiCreateChannelRequestStream) Limit(limit int) *ApiCreateChannelRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiCreateChannelRequestStream) Map(fn func(ApiCreateChannelRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Map2Int(fn func(ApiCreateChannelRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Map2Int32(fn func(ApiCreateChannelRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Map2Int64(fn func(ApiCreateChannelRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Map2Float32(fn func(ApiCreateChannelRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Map2Float64(fn func(ApiCreateChannelRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Map2Bool(fn func(ApiCreateChannelRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Map2Bytes(fn func(ApiCreateChannelRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Map2String(fn func(ApiCreateChannelRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Max(fn func(ApiCreateChannelRequest, int) float64) *ApiCreateChannelRequest {
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
func (self *ApiCreateChannelRequestStream) Min(fn func(ApiCreateChannelRequest, int) float64) *ApiCreateChannelRequest {
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
func (self *ApiCreateChannelRequestStream) NoneMatch(fn func(ApiCreateChannelRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiCreateChannelRequestStream) Get(index int) *ApiCreateChannelRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiCreateChannelRequestStream) GetOr(index int, arg ApiCreateChannelRequest) ApiCreateChannelRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateChannelRequestStream) Peek(fn func(*ApiCreateChannelRequest, int)) *ApiCreateChannelRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiCreateChannelRequestStream) Reduce(fn func(ApiCreateChannelRequest, ApiCreateChannelRequest, int) ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	return self.ReduceInit(fn, ApiCreateChannelRequest{})
}
func (self *ApiCreateChannelRequestStream) ReduceInit(fn func(ApiCreateChannelRequest, ApiCreateChannelRequest, int) ApiCreateChannelRequest, initialValue ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	result := ApiCreateChannelRequestStreamOf()
	self.ForEach(func(v ApiCreateChannelRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateChannelRequestStream) ReduceInterface(fn func(interface{}, ApiCreateChannelRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiCreateChannelRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiCreateChannelRequestStream) ReduceString(fn func(string, ApiCreateChannelRequest, int) string) []string {
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
func (self *ApiCreateChannelRequestStream) ReduceInt(fn func(int, ApiCreateChannelRequest, int) int) []int {
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
func (self *ApiCreateChannelRequestStream) ReduceInt32(fn func(int32, ApiCreateChannelRequest, int) int32) []int32 {
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
func (self *ApiCreateChannelRequestStream) ReduceInt64(fn func(int64, ApiCreateChannelRequest, int) int64) []int64 {
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
func (self *ApiCreateChannelRequestStream) ReduceFloat32(fn func(float32, ApiCreateChannelRequest, int) float32) []float32 {
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
func (self *ApiCreateChannelRequestStream) ReduceFloat64(fn func(float64, ApiCreateChannelRequest, int) float64) []float64 {
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
func (self *ApiCreateChannelRequestStream) ReduceBool(fn func(bool, ApiCreateChannelRequest, int) bool) []bool {
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
func (self *ApiCreateChannelRequestStream) Reverse() *ApiCreateChannelRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiCreateChannelRequestStream) Replace(fn func(ApiCreateChannelRequest, int) ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	return self.ForEach(func(v ApiCreateChannelRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiCreateChannelRequestStream) Select(fn func(ApiCreateChannelRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiCreateChannelRequestStream) Set(index int, val ApiCreateChannelRequest) *ApiCreateChannelRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiCreateChannelRequestStream) Skip(skip int) *ApiCreateChannelRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiCreateChannelRequestStream) SkippingEach(fn func(ApiCreateChannelRequest, int) int) *ApiCreateChannelRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiCreateChannelRequestStream) Slice(startIndex, n int) *ApiCreateChannelRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiCreateChannelRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiCreateChannelRequestStream) Sort(fn func(i, j int) bool) *ApiCreateChannelRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiCreateChannelRequestStream) Tail() *ApiCreateChannelRequest {
	return self.Last()
}
func (self *ApiCreateChannelRequestStream) TailOr(arg ApiCreateChannelRequest) ApiCreateChannelRequest {
	return self.LastOr(arg)
}
func (self *ApiCreateChannelRequestStream) ToList() []ApiCreateChannelRequest {
	return self.Val()
}
func (self *ApiCreateChannelRequestStream) Unique() *ApiCreateChannelRequestStream {
	return self.Distinct()
}
func (self *ApiCreateChannelRequestStream) Val() []ApiCreateChannelRequest {
	if self == nil {
		return []ApiCreateChannelRequest{}
	}
	return *self.Copy()
}
func (self *ApiCreateChannelRequestStream) While(fn func(ApiCreateChannelRequest, int) bool) *ApiCreateChannelRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiCreateChannelRequestStream) Where(fn func(ApiCreateChannelRequest) bool) *ApiCreateChannelRequestStream {
	result := ApiCreateChannelRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateChannelRequestStream) WhereSlim(fn func(ApiCreateChannelRequest) bool) *ApiCreateChannelRequestStream {
	result := ApiCreateChannelRequestStreamOf()
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
