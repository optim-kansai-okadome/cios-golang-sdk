/*
 * Collection utility of ApiCreateInventoryRequest Struct
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

type ApiCreateInventoryRequestStream []ApiCreateInventoryRequest

func ApiCreateInventoryRequestStreamOf(arg ...ApiCreateInventoryRequest) ApiCreateInventoryRequestStream {
	return arg
}
func ApiCreateInventoryRequestStreamFrom(arg []ApiCreateInventoryRequest) ApiCreateInventoryRequestStream {
	return arg
}
func CreateApiCreateInventoryRequestStream(arg ...ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	tmp := ApiCreateInventoryRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiCreateInventoryRequestStream(arg []ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	tmp := ApiCreateInventoryRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiCreateInventoryRequestStream) Add(arg ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	return self.AddAll(arg)
}
func (self *ApiCreateInventoryRequestStream) AddAll(arg ...ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiCreateInventoryRequestStream) AddSafe(arg *ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) Aggregate(fn func(ApiCreateInventoryRequest, ApiCreateInventoryRequest) ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	result := ApiCreateInventoryRequestStreamOf()
	self.ForEach(func(v ApiCreateInventoryRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiCreateInventoryRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateInventoryRequestStream) AllMatch(fn func(ApiCreateInventoryRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiCreateInventoryRequestStream) AnyMatch(fn func(ApiCreateInventoryRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiCreateInventoryRequestStream) Clone() *ApiCreateInventoryRequestStream {
	temp := make([]ApiCreateInventoryRequest, self.Len())
	copy(temp, *self)
	return (*ApiCreateInventoryRequestStream)(&temp)
}
func (self *ApiCreateInventoryRequestStream) Copy() *ApiCreateInventoryRequestStream {
	return self.Clone()
}
func (self *ApiCreateInventoryRequestStream) Concat(arg []ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiCreateInventoryRequestStream) Contains(arg ApiCreateInventoryRequest) bool {
	return self.FindIndex(func(_arg ApiCreateInventoryRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiCreateInventoryRequestStream) Clean() *ApiCreateInventoryRequestStream {
	*self = ApiCreateInventoryRequestStreamOf()
	return self
}
func (self *ApiCreateInventoryRequestStream) Delete(index int) *ApiCreateInventoryRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiCreateInventoryRequestStream) DeleteRange(startIndex, endIndex int) *ApiCreateInventoryRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiCreateInventoryRequestStream) Distinct() *ApiCreateInventoryRequestStream {
	caches := map[string]bool{}
	result := ApiCreateInventoryRequestStreamOf()
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
func (self *ApiCreateInventoryRequestStream) Each(fn func(ApiCreateInventoryRequest)) *ApiCreateInventoryRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) EachRight(fn func(ApiCreateInventoryRequest)) *ApiCreateInventoryRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) Equals(arr []ApiCreateInventoryRequest) bool {
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
func (self *ApiCreateInventoryRequestStream) Filter(fn func(ApiCreateInventoryRequest, int) bool) *ApiCreateInventoryRequestStream {
	result := ApiCreateInventoryRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateInventoryRequestStream) FilterSlim(fn func(ApiCreateInventoryRequest, int) bool) *ApiCreateInventoryRequestStream {
	result := ApiCreateInventoryRequestStreamOf()
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
func (self *ApiCreateInventoryRequestStream) Find(fn func(ApiCreateInventoryRequest, int) bool) *ApiCreateInventoryRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiCreateInventoryRequestStream) FindOr(fn func(ApiCreateInventoryRequest, int) bool, or ApiCreateInventoryRequest) ApiCreateInventoryRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiCreateInventoryRequestStream) FindIndex(fn func(ApiCreateInventoryRequest, int) bool) int {
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
func (self *ApiCreateInventoryRequestStream) First() *ApiCreateInventoryRequest {
	return self.Get(0)
}
func (self *ApiCreateInventoryRequestStream) FirstOr(arg ApiCreateInventoryRequest) ApiCreateInventoryRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateInventoryRequestStream) ForEach(fn func(ApiCreateInventoryRequest, int)) *ApiCreateInventoryRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) ForEachRight(fn func(ApiCreateInventoryRequest, int)) *ApiCreateInventoryRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) GroupBy(fn func(ApiCreateInventoryRequest, int) string) map[string][]ApiCreateInventoryRequest {
	m := map[string][]ApiCreateInventoryRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiCreateInventoryRequestStream) GroupByValues(fn func(ApiCreateInventoryRequest, int) string) [][]ApiCreateInventoryRequest {
	var tmp [][]ApiCreateInventoryRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiCreateInventoryRequestStream) IndexOf(arg ApiCreateInventoryRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiCreateInventoryRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiCreateInventoryRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiCreateInventoryRequestStream) Last() *ApiCreateInventoryRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiCreateInventoryRequestStream) LastOr(arg ApiCreateInventoryRequest) ApiCreateInventoryRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateInventoryRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiCreateInventoryRequestStream) Limit(limit int) *ApiCreateInventoryRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiCreateInventoryRequestStream) Map(fn func(ApiCreateInventoryRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Map2Int(fn func(ApiCreateInventoryRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Map2Int32(fn func(ApiCreateInventoryRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Map2Int64(fn func(ApiCreateInventoryRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Map2Float32(fn func(ApiCreateInventoryRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Map2Float64(fn func(ApiCreateInventoryRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Map2Bool(fn func(ApiCreateInventoryRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Map2Bytes(fn func(ApiCreateInventoryRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Map2String(fn func(ApiCreateInventoryRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Max(fn func(ApiCreateInventoryRequest, int) float64) *ApiCreateInventoryRequest {
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
func (self *ApiCreateInventoryRequestStream) Min(fn func(ApiCreateInventoryRequest, int) float64) *ApiCreateInventoryRequest {
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
func (self *ApiCreateInventoryRequestStream) NoneMatch(fn func(ApiCreateInventoryRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiCreateInventoryRequestStream) Get(index int) *ApiCreateInventoryRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiCreateInventoryRequestStream) GetOr(index int, arg ApiCreateInventoryRequest) ApiCreateInventoryRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateInventoryRequestStream) Peek(fn func(*ApiCreateInventoryRequest, int)) *ApiCreateInventoryRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiCreateInventoryRequestStream) Reduce(fn func(ApiCreateInventoryRequest, ApiCreateInventoryRequest, int) ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	return self.ReduceInit(fn, ApiCreateInventoryRequest{})
}
func (self *ApiCreateInventoryRequestStream) ReduceInit(fn func(ApiCreateInventoryRequest, ApiCreateInventoryRequest, int) ApiCreateInventoryRequest, initialValue ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	result := ApiCreateInventoryRequestStreamOf()
	self.ForEach(func(v ApiCreateInventoryRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateInventoryRequestStream) ReduceInterface(fn func(interface{}, ApiCreateInventoryRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiCreateInventoryRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiCreateInventoryRequestStream) ReduceString(fn func(string, ApiCreateInventoryRequest, int) string) []string {
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
func (self *ApiCreateInventoryRequestStream) ReduceInt(fn func(int, ApiCreateInventoryRequest, int) int) []int {
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
func (self *ApiCreateInventoryRequestStream) ReduceInt32(fn func(int32, ApiCreateInventoryRequest, int) int32) []int32 {
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
func (self *ApiCreateInventoryRequestStream) ReduceInt64(fn func(int64, ApiCreateInventoryRequest, int) int64) []int64 {
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
func (self *ApiCreateInventoryRequestStream) ReduceFloat32(fn func(float32, ApiCreateInventoryRequest, int) float32) []float32 {
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
func (self *ApiCreateInventoryRequestStream) ReduceFloat64(fn func(float64, ApiCreateInventoryRequest, int) float64) []float64 {
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
func (self *ApiCreateInventoryRequestStream) ReduceBool(fn func(bool, ApiCreateInventoryRequest, int) bool) []bool {
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
func (self *ApiCreateInventoryRequestStream) Reverse() *ApiCreateInventoryRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) Replace(fn func(ApiCreateInventoryRequest, int) ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	return self.ForEach(func(v ApiCreateInventoryRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiCreateInventoryRequestStream) Select(fn func(ApiCreateInventoryRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiCreateInventoryRequestStream) Set(index int, val ApiCreateInventoryRequest) *ApiCreateInventoryRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) Skip(skip int) *ApiCreateInventoryRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiCreateInventoryRequestStream) SkippingEach(fn func(ApiCreateInventoryRequest, int) int) *ApiCreateInventoryRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) Slice(startIndex, n int) *ApiCreateInventoryRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiCreateInventoryRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) Sort(fn func(i, j int) bool) *ApiCreateInventoryRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiCreateInventoryRequestStream) Tail() *ApiCreateInventoryRequest {
	return self.Last()
}
func (self *ApiCreateInventoryRequestStream) TailOr(arg ApiCreateInventoryRequest) ApiCreateInventoryRequest {
	return self.LastOr(arg)
}
func (self *ApiCreateInventoryRequestStream) ToList() []ApiCreateInventoryRequest {
	return self.Val()
}
func (self *ApiCreateInventoryRequestStream) Unique() *ApiCreateInventoryRequestStream {
	return self.Distinct()
}
func (self *ApiCreateInventoryRequestStream) Val() []ApiCreateInventoryRequest {
	if self == nil {
		return []ApiCreateInventoryRequest{}
	}
	return *self.Copy()
}
func (self *ApiCreateInventoryRequestStream) While(fn func(ApiCreateInventoryRequest, int) bool) *ApiCreateInventoryRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiCreateInventoryRequestStream) Where(fn func(ApiCreateInventoryRequest) bool) *ApiCreateInventoryRequestStream {
	result := ApiCreateInventoryRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateInventoryRequestStream) WhereSlim(fn func(ApiCreateInventoryRequest) bool) *ApiCreateInventoryRequestStream {
	result := ApiCreateInventoryRequestStreamOf()
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
