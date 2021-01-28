/*
 * Collection utility of ApiCreateDeviceEntitiesLifecycleRequest Struct
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

type ApiCreateDeviceEntitiesLifecycleRequestStream []ApiCreateDeviceEntitiesLifecycleRequest

func ApiCreateDeviceEntitiesLifecycleRequestStreamOf(arg ...ApiCreateDeviceEntitiesLifecycleRequest) ApiCreateDeviceEntitiesLifecycleRequestStream {
	return arg
}
func ApiCreateDeviceEntitiesLifecycleRequestStreamFrom(arg []ApiCreateDeviceEntitiesLifecycleRequest) ApiCreateDeviceEntitiesLifecycleRequestStream {
	return arg
}
func CreateApiCreateDeviceEntitiesLifecycleRequestStream(arg ...ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	tmp := ApiCreateDeviceEntitiesLifecycleRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiCreateDeviceEntitiesLifecycleRequestStream(arg []ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	tmp := ApiCreateDeviceEntitiesLifecycleRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Add(arg ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	return self.AddAll(arg)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) AddAll(arg ...ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) AddSafe(arg *ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Aggregate(fn func(ApiCreateDeviceEntitiesLifecycleRequest, ApiCreateDeviceEntitiesLifecycleRequest) ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	result := ApiCreateDeviceEntitiesLifecycleRequestStreamOf()
	self.ForEach(func(v ApiCreateDeviceEntitiesLifecycleRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiCreateDeviceEntitiesLifecycleRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) AllMatch(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) AnyMatch(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Clone() *ApiCreateDeviceEntitiesLifecycleRequestStream {
	temp := make([]ApiCreateDeviceEntitiesLifecycleRequest, self.Len())
	copy(temp, *self)
	return (*ApiCreateDeviceEntitiesLifecycleRequestStream)(&temp)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Copy() *ApiCreateDeviceEntitiesLifecycleRequestStream {
	return self.Clone()
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Concat(arg []ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Contains(arg ApiCreateDeviceEntitiesLifecycleRequest) bool {
	return self.FindIndex(func(_arg ApiCreateDeviceEntitiesLifecycleRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Clean() *ApiCreateDeviceEntitiesLifecycleRequestStream {
	*self = ApiCreateDeviceEntitiesLifecycleRequestStreamOf()
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Delete(index int) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) DeleteRange(startIndex, endIndex int) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Distinct() *ApiCreateDeviceEntitiesLifecycleRequestStream {
	caches := map[string]bool{}
	result := ApiCreateDeviceEntitiesLifecycleRequestStreamOf()
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Each(fn func(ApiCreateDeviceEntitiesLifecycleRequest)) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) EachRight(fn func(ApiCreateDeviceEntitiesLifecycleRequest)) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Equals(arr []ApiCreateDeviceEntitiesLifecycleRequest) bool {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Filter(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	result := ApiCreateDeviceEntitiesLifecycleRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) FilterSlim(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	result := ApiCreateDeviceEntitiesLifecycleRequestStreamOf()
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Find(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool) *ApiCreateDeviceEntitiesLifecycleRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) FindOr(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool, or ApiCreateDeviceEntitiesLifecycleRequest) ApiCreateDeviceEntitiesLifecycleRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) FindIndex(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool) int {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) First() *ApiCreateDeviceEntitiesLifecycleRequest {
	return self.Get(0)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) FirstOr(arg ApiCreateDeviceEntitiesLifecycleRequest) ApiCreateDeviceEntitiesLifecycleRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ForEach(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int)) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ForEachRight(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int)) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) GroupBy(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) string) map[string][]ApiCreateDeviceEntitiesLifecycleRequest {
	m := map[string][]ApiCreateDeviceEntitiesLifecycleRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) GroupByValues(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) string) [][]ApiCreateDeviceEntitiesLifecycleRequest {
	var tmp [][]ApiCreateDeviceEntitiesLifecycleRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) IndexOf(arg ApiCreateDeviceEntitiesLifecycleRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Last() *ApiCreateDeviceEntitiesLifecycleRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) LastOr(arg ApiCreateDeviceEntitiesLifecycleRequest) ApiCreateDeviceEntitiesLifecycleRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Limit(limit int) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Map(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Map2Int(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Map2Int32(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Map2Int64(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Map2Float32(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Map2Float64(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Map2Bool(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Map2Bytes(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Map2String(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Max(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) float64) *ApiCreateDeviceEntitiesLifecycleRequest {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Min(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) float64) *ApiCreateDeviceEntitiesLifecycleRequest {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) NoneMatch(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Get(index int) *ApiCreateDeviceEntitiesLifecycleRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) GetOr(index int, arg ApiCreateDeviceEntitiesLifecycleRequest) ApiCreateDeviceEntitiesLifecycleRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Peek(fn func(*ApiCreateDeviceEntitiesLifecycleRequest, int)) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Reduce(fn func(ApiCreateDeviceEntitiesLifecycleRequest, ApiCreateDeviceEntitiesLifecycleRequest, int) ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	return self.ReduceInit(fn, ApiCreateDeviceEntitiesLifecycleRequest{})
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ReduceInit(fn func(ApiCreateDeviceEntitiesLifecycleRequest, ApiCreateDeviceEntitiesLifecycleRequest, int) ApiCreateDeviceEntitiesLifecycleRequest, initialValue ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	result := ApiCreateDeviceEntitiesLifecycleRequestStreamOf()
	self.ForEach(func(v ApiCreateDeviceEntitiesLifecycleRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ReduceInterface(fn func(interface{}, ApiCreateDeviceEntitiesLifecycleRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiCreateDeviceEntitiesLifecycleRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ReduceString(fn func(string, ApiCreateDeviceEntitiesLifecycleRequest, int) string) []string {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ReduceInt(fn func(int, ApiCreateDeviceEntitiesLifecycleRequest, int) int) []int {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ReduceInt32(fn func(int32, ApiCreateDeviceEntitiesLifecycleRequest, int) int32) []int32 {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ReduceInt64(fn func(int64, ApiCreateDeviceEntitiesLifecycleRequest, int) int64) []int64 {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ReduceFloat32(fn func(float32, ApiCreateDeviceEntitiesLifecycleRequest, int) float32) []float32 {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ReduceFloat64(fn func(float64, ApiCreateDeviceEntitiesLifecycleRequest, int) float64) []float64 {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ReduceBool(fn func(bool, ApiCreateDeviceEntitiesLifecycleRequest, int) bool) []bool {
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
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Reverse() *ApiCreateDeviceEntitiesLifecycleRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Replace(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	return self.ForEach(func(v ApiCreateDeviceEntitiesLifecycleRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Select(fn func(ApiCreateDeviceEntitiesLifecycleRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Set(index int, val ApiCreateDeviceEntitiesLifecycleRequest) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Skip(skip int) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) SkippingEach(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) int) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Slice(startIndex, n int) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiCreateDeviceEntitiesLifecycleRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Sort(fn func(i, j int) bool) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Tail() *ApiCreateDeviceEntitiesLifecycleRequest {
	return self.Last()
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) TailOr(arg ApiCreateDeviceEntitiesLifecycleRequest) ApiCreateDeviceEntitiesLifecycleRequest {
	return self.LastOr(arg)
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) ToList() []ApiCreateDeviceEntitiesLifecycleRequest {
	return self.Val()
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Unique() *ApiCreateDeviceEntitiesLifecycleRequestStream {
	return self.Distinct()
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Val() []ApiCreateDeviceEntitiesLifecycleRequest {
	if self == nil {
		return []ApiCreateDeviceEntitiesLifecycleRequest{}
	}
	return *self.Copy()
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) While(fn func(ApiCreateDeviceEntitiesLifecycleRequest, int) bool) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) Where(fn func(ApiCreateDeviceEntitiesLifecycleRequest) bool) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	result := ApiCreateDeviceEntitiesLifecycleRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateDeviceEntitiesLifecycleRequestStream) WhereSlim(fn func(ApiCreateDeviceEntitiesLifecycleRequest) bool) *ApiCreateDeviceEntitiesLifecycleRequestStream {
	result := ApiCreateDeviceEntitiesLifecycleRequestStreamOf()
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
