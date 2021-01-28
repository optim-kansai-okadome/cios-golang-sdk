/*
 * Collection utility of NullableInventory Struct
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

type NullableInventoryStream []NullableInventory

func NullableInventoryStreamOf(arg ...NullableInventory) NullableInventoryStream {
	return arg
}
func NullableInventoryStreamFrom(arg []NullableInventory) NullableInventoryStream {
	return arg
}
func CreateNullableInventoryStream(arg ...NullableInventory) *NullableInventoryStream {
	tmp := NullableInventoryStreamOf(arg...)
	return &tmp
}
func GenerateNullableInventoryStream(arg []NullableInventory) *NullableInventoryStream {
	tmp := NullableInventoryStreamFrom(arg)
	return &tmp
}

func (self *NullableInventoryStream) Add(arg NullableInventory) *NullableInventoryStream {
	return self.AddAll(arg)
}
func (self *NullableInventoryStream) AddAll(arg ...NullableInventory) *NullableInventoryStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableInventoryStream) AddSafe(arg *NullableInventory) *NullableInventoryStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableInventoryStream) Aggregate(fn func(NullableInventory, NullableInventory) NullableInventory) *NullableInventoryStream {
	result := NullableInventoryStreamOf()
	self.ForEach(func(v NullableInventory, i int) {
		if i == 0 {
			result.Add(fn(NullableInventory{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableInventoryStream) AllMatch(fn func(NullableInventory, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableInventoryStream) AnyMatch(fn func(NullableInventory, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableInventoryStream) Clone() *NullableInventoryStream {
	temp := make([]NullableInventory, self.Len())
	copy(temp, *self)
	return (*NullableInventoryStream)(&temp)
}
func (self *NullableInventoryStream) Copy() *NullableInventoryStream {
	return self.Clone()
}
func (self *NullableInventoryStream) Concat(arg []NullableInventory) *NullableInventoryStream {
	return self.AddAll(arg...)
}
func (self *NullableInventoryStream) Contains(arg NullableInventory) bool {
	return self.FindIndex(func(_arg NullableInventory, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableInventoryStream) Clean() *NullableInventoryStream {
	*self = NullableInventoryStreamOf()
	return self
}
func (self *NullableInventoryStream) Delete(index int) *NullableInventoryStream {
	return self.DeleteRange(index, index)
}
func (self *NullableInventoryStream) DeleteRange(startIndex, endIndex int) *NullableInventoryStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableInventoryStream) Distinct() *NullableInventoryStream {
	caches := map[string]bool{}
	result := NullableInventoryStreamOf()
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
func (self *NullableInventoryStream) Each(fn func(NullableInventory)) *NullableInventoryStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableInventoryStream) EachRight(fn func(NullableInventory)) *NullableInventoryStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableInventoryStream) Equals(arr []NullableInventory) bool {
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
func (self *NullableInventoryStream) Filter(fn func(NullableInventory, int) bool) *NullableInventoryStream {
	result := NullableInventoryStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableInventoryStream) FilterSlim(fn func(NullableInventory, int) bool) *NullableInventoryStream {
	result := NullableInventoryStreamOf()
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
func (self *NullableInventoryStream) Find(fn func(NullableInventory, int) bool) *NullableInventory {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableInventoryStream) FindOr(fn func(NullableInventory, int) bool, or NullableInventory) NullableInventory {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableInventoryStream) FindIndex(fn func(NullableInventory, int) bool) int {
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
func (self *NullableInventoryStream) First() *NullableInventory {
	return self.Get(0)
}
func (self *NullableInventoryStream) FirstOr(arg NullableInventory) NullableInventory {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableInventoryStream) ForEach(fn func(NullableInventory, int)) *NullableInventoryStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableInventoryStream) ForEachRight(fn func(NullableInventory, int)) *NullableInventoryStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableInventoryStream) GroupBy(fn func(NullableInventory, int) string) map[string][]NullableInventory {
	m := map[string][]NullableInventory{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableInventoryStream) GroupByValues(fn func(NullableInventory, int) string) [][]NullableInventory {
	var tmp [][]NullableInventory
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableInventoryStream) IndexOf(arg NullableInventory) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableInventoryStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableInventoryStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableInventoryStream) Last() *NullableInventory {
	return self.Get(self.Len() - 1)
}
func (self *NullableInventoryStream) LastOr(arg NullableInventory) NullableInventory {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableInventoryStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableInventoryStream) Limit(limit int) *NullableInventoryStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableInventoryStream) Map(fn func(NullableInventory, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInventoryStream) Map2Int(fn func(NullableInventory, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInventoryStream) Map2Int32(fn func(NullableInventory, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInventoryStream) Map2Int64(fn func(NullableInventory, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInventoryStream) Map2Float32(fn func(NullableInventory, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInventoryStream) Map2Float64(fn func(NullableInventory, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInventoryStream) Map2Bool(fn func(NullableInventory, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInventoryStream) Map2Bytes(fn func(NullableInventory, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInventoryStream) Map2String(fn func(NullableInventory, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInventoryStream) Max(fn func(NullableInventory, int) float64) *NullableInventory {
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
func (self *NullableInventoryStream) Min(fn func(NullableInventory, int) float64) *NullableInventory {
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
func (self *NullableInventoryStream) NoneMatch(fn func(NullableInventory, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableInventoryStream) Get(index int) *NullableInventory {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableInventoryStream) GetOr(index int, arg NullableInventory) NullableInventory {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableInventoryStream) Peek(fn func(*NullableInventory, int)) *NullableInventoryStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableInventoryStream) Reduce(fn func(NullableInventory, NullableInventory, int) NullableInventory) *NullableInventoryStream {
	return self.ReduceInit(fn, NullableInventory{})
}
func (self *NullableInventoryStream) ReduceInit(fn func(NullableInventory, NullableInventory, int) NullableInventory, initialValue NullableInventory) *NullableInventoryStream {
	result := NullableInventoryStreamOf()
	self.ForEach(func(v NullableInventory, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableInventoryStream) ReduceInterface(fn func(interface{}, NullableInventory, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableInventory{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableInventoryStream) ReduceString(fn func(string, NullableInventory, int) string) []string {
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
func (self *NullableInventoryStream) ReduceInt(fn func(int, NullableInventory, int) int) []int {
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
func (self *NullableInventoryStream) ReduceInt32(fn func(int32, NullableInventory, int) int32) []int32 {
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
func (self *NullableInventoryStream) ReduceInt64(fn func(int64, NullableInventory, int) int64) []int64 {
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
func (self *NullableInventoryStream) ReduceFloat32(fn func(float32, NullableInventory, int) float32) []float32 {
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
func (self *NullableInventoryStream) ReduceFloat64(fn func(float64, NullableInventory, int) float64) []float64 {
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
func (self *NullableInventoryStream) ReduceBool(fn func(bool, NullableInventory, int) bool) []bool {
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
func (self *NullableInventoryStream) Reverse() *NullableInventoryStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableInventoryStream) Replace(fn func(NullableInventory, int) NullableInventory) *NullableInventoryStream {
	return self.ForEach(func(v NullableInventory, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableInventoryStream) Select(fn func(NullableInventory) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableInventoryStream) Set(index int, val NullableInventory) *NullableInventoryStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableInventoryStream) Skip(skip int) *NullableInventoryStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableInventoryStream) SkippingEach(fn func(NullableInventory, int) int) *NullableInventoryStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableInventoryStream) Slice(startIndex, n int) *NullableInventoryStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableInventory{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableInventoryStream) Sort(fn func(i, j int) bool) *NullableInventoryStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableInventoryStream) Tail() *NullableInventory {
	return self.Last()
}
func (self *NullableInventoryStream) TailOr(arg NullableInventory) NullableInventory {
	return self.LastOr(arg)
}
func (self *NullableInventoryStream) ToList() []NullableInventory {
	return self.Val()
}
func (self *NullableInventoryStream) Unique() *NullableInventoryStream {
	return self.Distinct()
}
func (self *NullableInventoryStream) Val() []NullableInventory {
	if self == nil {
		return []NullableInventory{}
	}
	return *self.Copy()
}
func (self *NullableInventoryStream) While(fn func(NullableInventory, int) bool) *NullableInventoryStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableInventoryStream) Where(fn func(NullableInventory) bool) *NullableInventoryStream {
	result := NullableInventoryStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableInventoryStream) WhereSlim(fn func(NullableInventory) bool) *NullableInventoryStream {
	result := NullableInventoryStreamOf()
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
