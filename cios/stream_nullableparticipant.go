/*
 * Collection utility of NullableParticipant Struct
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

type NullableParticipantStream []NullableParticipant

func NullableParticipantStreamOf(arg ...NullableParticipant) NullableParticipantStream {
	return arg
}
func NullableParticipantStreamFrom(arg []NullableParticipant) NullableParticipantStream {
	return arg
}
func CreateNullableParticipantStream(arg ...NullableParticipant) *NullableParticipantStream {
	tmp := NullableParticipantStreamOf(arg...)
	return &tmp
}
func GenerateNullableParticipantStream(arg []NullableParticipant) *NullableParticipantStream {
	tmp := NullableParticipantStreamFrom(arg)
	return &tmp
}

func (self *NullableParticipantStream) Add(arg NullableParticipant) *NullableParticipantStream {
	return self.AddAll(arg)
}
func (self *NullableParticipantStream) AddAll(arg ...NullableParticipant) *NullableParticipantStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableParticipantStream) AddSafe(arg *NullableParticipant) *NullableParticipantStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableParticipantStream) Aggregate(fn func(NullableParticipant, NullableParticipant) NullableParticipant) *NullableParticipantStream {
	result := NullableParticipantStreamOf()
	self.ForEach(func(v NullableParticipant, i int) {
		if i == 0 {
			result.Add(fn(NullableParticipant{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableParticipantStream) AllMatch(fn func(NullableParticipant, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableParticipantStream) AnyMatch(fn func(NullableParticipant, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableParticipantStream) Clone() *NullableParticipantStream {
	temp := make([]NullableParticipant, self.Len())
	copy(temp, *self)
	return (*NullableParticipantStream)(&temp)
}
func (self *NullableParticipantStream) Copy() *NullableParticipantStream {
	return self.Clone()
}
func (self *NullableParticipantStream) Concat(arg []NullableParticipant) *NullableParticipantStream {
	return self.AddAll(arg...)
}
func (self *NullableParticipantStream) Contains(arg NullableParticipant) bool {
	return self.FindIndex(func(_arg NullableParticipant, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableParticipantStream) Clean() *NullableParticipantStream {
	*self = NullableParticipantStreamOf()
	return self
}
func (self *NullableParticipantStream) Delete(index int) *NullableParticipantStream {
	return self.DeleteRange(index, index)
}
func (self *NullableParticipantStream) DeleteRange(startIndex, endIndex int) *NullableParticipantStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableParticipantStream) Distinct() *NullableParticipantStream {
	caches := map[string]bool{}
	result := NullableParticipantStreamOf()
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
func (self *NullableParticipantStream) Each(fn func(NullableParticipant)) *NullableParticipantStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableParticipantStream) EachRight(fn func(NullableParticipant)) *NullableParticipantStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableParticipantStream) Equals(arr []NullableParticipant) bool {
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
func (self *NullableParticipantStream) Filter(fn func(NullableParticipant, int) bool) *NullableParticipantStream {
	result := NullableParticipantStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableParticipantStream) FilterSlim(fn func(NullableParticipant, int) bool) *NullableParticipantStream {
	result := NullableParticipantStreamOf()
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
func (self *NullableParticipantStream) Find(fn func(NullableParticipant, int) bool) *NullableParticipant {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableParticipantStream) FindOr(fn func(NullableParticipant, int) bool, or NullableParticipant) NullableParticipant {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableParticipantStream) FindIndex(fn func(NullableParticipant, int) bool) int {
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
func (self *NullableParticipantStream) First() *NullableParticipant {
	return self.Get(0)
}
func (self *NullableParticipantStream) FirstOr(arg NullableParticipant) NullableParticipant {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableParticipantStream) ForEach(fn func(NullableParticipant, int)) *NullableParticipantStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableParticipantStream) ForEachRight(fn func(NullableParticipant, int)) *NullableParticipantStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableParticipantStream) GroupBy(fn func(NullableParticipant, int) string) map[string][]NullableParticipant {
	m := map[string][]NullableParticipant{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableParticipantStream) GroupByValues(fn func(NullableParticipant, int) string) [][]NullableParticipant {
	var tmp [][]NullableParticipant
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableParticipantStream) IndexOf(arg NullableParticipant) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableParticipantStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableParticipantStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableParticipantStream) Last() *NullableParticipant {
	return self.Get(self.Len() - 1)
}
func (self *NullableParticipantStream) LastOr(arg NullableParticipant) NullableParticipant {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableParticipantStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableParticipantStream) Limit(limit int) *NullableParticipantStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableParticipantStream) Map(fn func(NullableParticipant, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableParticipantStream) Map2Int(fn func(NullableParticipant, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableParticipantStream) Map2Int32(fn func(NullableParticipant, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableParticipantStream) Map2Int64(fn func(NullableParticipant, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableParticipantStream) Map2Float32(fn func(NullableParticipant, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableParticipantStream) Map2Float64(fn func(NullableParticipant, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableParticipantStream) Map2Bool(fn func(NullableParticipant, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableParticipantStream) Map2Bytes(fn func(NullableParticipant, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableParticipantStream) Map2String(fn func(NullableParticipant, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableParticipantStream) Max(fn func(NullableParticipant, int) float64) *NullableParticipant {
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
func (self *NullableParticipantStream) Min(fn func(NullableParticipant, int) float64) *NullableParticipant {
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
func (self *NullableParticipantStream) NoneMatch(fn func(NullableParticipant, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableParticipantStream) Get(index int) *NullableParticipant {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableParticipantStream) GetOr(index int, arg NullableParticipant) NullableParticipant {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableParticipantStream) Peek(fn func(*NullableParticipant, int)) *NullableParticipantStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableParticipantStream) Reduce(fn func(NullableParticipant, NullableParticipant, int) NullableParticipant) *NullableParticipantStream {
	return self.ReduceInit(fn, NullableParticipant{})
}
func (self *NullableParticipantStream) ReduceInit(fn func(NullableParticipant, NullableParticipant, int) NullableParticipant, initialValue NullableParticipant) *NullableParticipantStream {
	result := NullableParticipantStreamOf()
	self.ForEach(func(v NullableParticipant, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableParticipantStream) ReduceInterface(fn func(interface{}, NullableParticipant, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableParticipant{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableParticipantStream) ReduceString(fn func(string, NullableParticipant, int) string) []string {
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
func (self *NullableParticipantStream) ReduceInt(fn func(int, NullableParticipant, int) int) []int {
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
func (self *NullableParticipantStream) ReduceInt32(fn func(int32, NullableParticipant, int) int32) []int32 {
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
func (self *NullableParticipantStream) ReduceInt64(fn func(int64, NullableParticipant, int) int64) []int64 {
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
func (self *NullableParticipantStream) ReduceFloat32(fn func(float32, NullableParticipant, int) float32) []float32 {
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
func (self *NullableParticipantStream) ReduceFloat64(fn func(float64, NullableParticipant, int) float64) []float64 {
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
func (self *NullableParticipantStream) ReduceBool(fn func(bool, NullableParticipant, int) bool) []bool {
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
func (self *NullableParticipantStream) Reverse() *NullableParticipantStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableParticipantStream) Replace(fn func(NullableParticipant, int) NullableParticipant) *NullableParticipantStream {
	return self.ForEach(func(v NullableParticipant, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableParticipantStream) Select(fn func(NullableParticipant) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableParticipantStream) Set(index int, val NullableParticipant) *NullableParticipantStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableParticipantStream) Skip(skip int) *NullableParticipantStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableParticipantStream) SkippingEach(fn func(NullableParticipant, int) int) *NullableParticipantStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableParticipantStream) Slice(startIndex, n int) *NullableParticipantStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableParticipant{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableParticipantStream) Sort(fn func(i, j int) bool) *NullableParticipantStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableParticipantStream) Tail() *NullableParticipant {
	return self.Last()
}
func (self *NullableParticipantStream) TailOr(arg NullableParticipant) NullableParticipant {
	return self.LastOr(arg)
}
func (self *NullableParticipantStream) ToList() []NullableParticipant {
	return self.Val()
}
func (self *NullableParticipantStream) Unique() *NullableParticipantStream {
	return self.Distinct()
}
func (self *NullableParticipantStream) Val() []NullableParticipant {
	if self == nil {
		return []NullableParticipant{}
	}
	return *self.Copy()
}
func (self *NullableParticipantStream) While(fn func(NullableParticipant, int) bool) *NullableParticipantStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableParticipantStream) Where(fn func(NullableParticipant) bool) *NullableParticipantStream {
	result := NullableParticipantStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableParticipantStream) WhereSlim(fn func(NullableParticipant) bool) *NullableParticipantStream {
	result := NullableParticipantStreamOf()
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
