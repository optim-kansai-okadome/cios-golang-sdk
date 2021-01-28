/*
 * Collection utility of DataStoreObject Struct
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

type DataStoreObjectStream []DataStoreObject

func DataStoreObjectStreamOf(arg ...DataStoreObject) DataStoreObjectStream {
	return arg
}
func DataStoreObjectStreamFrom(arg []DataStoreObject) DataStoreObjectStream {
	return arg
}
func CreateDataStoreObjectStream(arg ...DataStoreObject) *DataStoreObjectStream {
	tmp := DataStoreObjectStreamOf(arg...)
	return &tmp
}
func GenerateDataStoreObjectStream(arg []DataStoreObject) *DataStoreObjectStream {
	tmp := DataStoreObjectStreamFrom(arg)
	return &tmp
}

func (self *DataStoreObjectStream) Add(arg DataStoreObject) *DataStoreObjectStream {
	return self.AddAll(arg)
}
func (self *DataStoreObjectStream) AddAll(arg ...DataStoreObject) *DataStoreObjectStream {
	*self = append(*self, arg...)
	return self
}
func (self *DataStoreObjectStream) AddSafe(arg *DataStoreObject) *DataStoreObjectStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *DataStoreObjectStream) Aggregate(fn func(DataStoreObject, DataStoreObject) DataStoreObject) *DataStoreObjectStream {
	result := DataStoreObjectStreamOf()
	self.ForEach(func(v DataStoreObject, i int) {
		if i == 0 {
			result.Add(fn(DataStoreObject{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *DataStoreObjectStream) AllMatch(fn func(DataStoreObject, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *DataStoreObjectStream) AnyMatch(fn func(DataStoreObject, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *DataStoreObjectStream) Clone() *DataStoreObjectStream {
	temp := make([]DataStoreObject, self.Len())
	copy(temp, *self)
	return (*DataStoreObjectStream)(&temp)
}
func (self *DataStoreObjectStream) Copy() *DataStoreObjectStream {
	return self.Clone()
}
func (self *DataStoreObjectStream) Concat(arg []DataStoreObject) *DataStoreObjectStream {
	return self.AddAll(arg...)
}
func (self *DataStoreObjectStream) Contains(arg DataStoreObject) bool {
	return self.FindIndex(func(_arg DataStoreObject, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *DataStoreObjectStream) Clean() *DataStoreObjectStream {
	*self = DataStoreObjectStreamOf()
	return self
}
func (self *DataStoreObjectStream) Delete(index int) *DataStoreObjectStream {
	return self.DeleteRange(index, index)
}
func (self *DataStoreObjectStream) DeleteRange(startIndex, endIndex int) *DataStoreObjectStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *DataStoreObjectStream) Distinct() *DataStoreObjectStream {
	caches := map[string]bool{}
	result := DataStoreObjectStreamOf()
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
func (self *DataStoreObjectStream) Each(fn func(DataStoreObject)) *DataStoreObjectStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *DataStoreObjectStream) EachRight(fn func(DataStoreObject)) *DataStoreObjectStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *DataStoreObjectStream) Equals(arr []DataStoreObject) bool {
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
func (self *DataStoreObjectStream) Filter(fn func(DataStoreObject, int) bool) *DataStoreObjectStream {
	result := DataStoreObjectStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DataStoreObjectStream) FilterSlim(fn func(DataStoreObject, int) bool) *DataStoreObjectStream {
	result := DataStoreObjectStreamOf()
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
func (self *DataStoreObjectStream) Find(fn func(DataStoreObject, int) bool) *DataStoreObject {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *DataStoreObjectStream) FindOr(fn func(DataStoreObject, int) bool, or DataStoreObject) DataStoreObject {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *DataStoreObjectStream) FindIndex(fn func(DataStoreObject, int) bool) int {
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
func (self *DataStoreObjectStream) First() *DataStoreObject {
	return self.Get(0)
}
func (self *DataStoreObjectStream) FirstOr(arg DataStoreObject) DataStoreObject {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *DataStoreObjectStream) ForEach(fn func(DataStoreObject, int)) *DataStoreObjectStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *DataStoreObjectStream) ForEachRight(fn func(DataStoreObject, int)) *DataStoreObjectStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *DataStoreObjectStream) GroupBy(fn func(DataStoreObject, int) string) map[string][]DataStoreObject {
	m := map[string][]DataStoreObject{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *DataStoreObjectStream) GroupByValues(fn func(DataStoreObject, int) string) [][]DataStoreObject {
	var tmp [][]DataStoreObject
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *DataStoreObjectStream) IndexOf(arg DataStoreObject) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *DataStoreObjectStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *DataStoreObjectStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *DataStoreObjectStream) Last() *DataStoreObject {
	return self.Get(self.Len() - 1)
}
func (self *DataStoreObjectStream) LastOr(arg DataStoreObject) DataStoreObject {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *DataStoreObjectStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *DataStoreObjectStream) Limit(limit int) *DataStoreObjectStream {
	self.Slice(0, limit)
	return self
}

func (self *DataStoreObjectStream) Map(fn func(DataStoreObject, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreObjectStream) Map2Int(fn func(DataStoreObject, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreObjectStream) Map2Int32(fn func(DataStoreObject, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreObjectStream) Map2Int64(fn func(DataStoreObject, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreObjectStream) Map2Float32(fn func(DataStoreObject, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreObjectStream) Map2Float64(fn func(DataStoreObject, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreObjectStream) Map2Bool(fn func(DataStoreObject, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreObjectStream) Map2Bytes(fn func(DataStoreObject, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreObjectStream) Map2String(fn func(DataStoreObject, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreObjectStream) Max(fn func(DataStoreObject, int) float64) *DataStoreObject {
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
func (self *DataStoreObjectStream) Min(fn func(DataStoreObject, int) float64) *DataStoreObject {
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
func (self *DataStoreObjectStream) NoneMatch(fn func(DataStoreObject, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *DataStoreObjectStream) Get(index int) *DataStoreObject {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *DataStoreObjectStream) GetOr(index int, arg DataStoreObject) DataStoreObject {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *DataStoreObjectStream) Peek(fn func(*DataStoreObject, int)) *DataStoreObjectStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *DataStoreObjectStream) Reduce(fn func(DataStoreObject, DataStoreObject, int) DataStoreObject) *DataStoreObjectStream {
	return self.ReduceInit(fn, DataStoreObject{})
}
func (self *DataStoreObjectStream) ReduceInit(fn func(DataStoreObject, DataStoreObject, int) DataStoreObject, initialValue DataStoreObject) *DataStoreObjectStream {
	result := DataStoreObjectStreamOf()
	self.ForEach(func(v DataStoreObject, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *DataStoreObjectStream) ReduceInterface(fn func(interface{}, DataStoreObject, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(DataStoreObject{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *DataStoreObjectStream) ReduceString(fn func(string, DataStoreObject, int) string) []string {
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
func (self *DataStoreObjectStream) ReduceInt(fn func(int, DataStoreObject, int) int) []int {
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
func (self *DataStoreObjectStream) ReduceInt32(fn func(int32, DataStoreObject, int) int32) []int32 {
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
func (self *DataStoreObjectStream) ReduceInt64(fn func(int64, DataStoreObject, int) int64) []int64 {
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
func (self *DataStoreObjectStream) ReduceFloat32(fn func(float32, DataStoreObject, int) float32) []float32 {
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
func (self *DataStoreObjectStream) ReduceFloat64(fn func(float64, DataStoreObject, int) float64) []float64 {
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
func (self *DataStoreObjectStream) ReduceBool(fn func(bool, DataStoreObject, int) bool) []bool {
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
func (self *DataStoreObjectStream) Reverse() *DataStoreObjectStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *DataStoreObjectStream) Replace(fn func(DataStoreObject, int) DataStoreObject) *DataStoreObjectStream {
	return self.ForEach(func(v DataStoreObject, i int) { self.Set(i, fn(v, i)) })
}
func (self *DataStoreObjectStream) Select(fn func(DataStoreObject) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *DataStoreObjectStream) Set(index int, val DataStoreObject) *DataStoreObjectStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *DataStoreObjectStream) Skip(skip int) *DataStoreObjectStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *DataStoreObjectStream) SkippingEach(fn func(DataStoreObject, int) int) *DataStoreObjectStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *DataStoreObjectStream) Slice(startIndex, n int) *DataStoreObjectStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []DataStoreObject{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *DataStoreObjectStream) Sort(fn func(i, j int) bool) *DataStoreObjectStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *DataStoreObjectStream) Tail() *DataStoreObject {
	return self.Last()
}
func (self *DataStoreObjectStream) TailOr(arg DataStoreObject) DataStoreObject {
	return self.LastOr(arg)
}
func (self *DataStoreObjectStream) ToList() []DataStoreObject {
	return self.Val()
}
func (self *DataStoreObjectStream) Unique() *DataStoreObjectStream {
	return self.Distinct()
}
func (self *DataStoreObjectStream) Val() []DataStoreObject {
	if self == nil {
		return []DataStoreObject{}
	}
	return *self.Copy()
}
func (self *DataStoreObjectStream) While(fn func(DataStoreObject, int) bool) *DataStoreObjectStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *DataStoreObjectStream) Where(fn func(DataStoreObject) bool) *DataStoreObjectStream {
	result := DataStoreObjectStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DataStoreObjectStream) WhereSlim(fn func(DataStoreObject) bool) *DataStoreObjectStream {
	result := DataStoreObjectStreamOf()
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
