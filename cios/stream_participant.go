/*
 * Collection utility of Participant Struct
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

type ParticipantStream []Participant

func ParticipantStreamOf(arg ...Participant) ParticipantStream {
	return arg
}
func ParticipantStreamFrom(arg []Participant) ParticipantStream {
	return arg
}
func CreateParticipantStream(arg ...Participant) *ParticipantStream {
	tmp := ParticipantStreamOf(arg...)
	return &tmp
}
func GenerateParticipantStream(arg []Participant) *ParticipantStream {
	tmp := ParticipantStreamFrom(arg)
	return &tmp
}

func (self *ParticipantStream) Add(arg Participant) *ParticipantStream {
	return self.AddAll(arg)
}
func (self *ParticipantStream) AddAll(arg ...Participant) *ParticipantStream {
	*self = append(*self, arg...)
	return self
}
func (self *ParticipantStream) AddSafe(arg *Participant) *ParticipantStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ParticipantStream) Aggregate(fn func(Participant, Participant) Participant) *ParticipantStream {
	result := ParticipantStreamOf()
	self.ForEach(func(v Participant, i int) {
		if i == 0 {
			result.Add(fn(Participant{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ParticipantStream) AllMatch(fn func(Participant, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ParticipantStream) AnyMatch(fn func(Participant, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ParticipantStream) Clone() *ParticipantStream {
	temp := make([]Participant, self.Len())
	copy(temp, *self)
	return (*ParticipantStream)(&temp)
}
func (self *ParticipantStream) Copy() *ParticipantStream {
	return self.Clone()
}
func (self *ParticipantStream) Concat(arg []Participant) *ParticipantStream {
	return self.AddAll(arg...)
}
func (self *ParticipantStream) Contains(arg Participant) bool {
	return self.FindIndex(func(_arg Participant, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ParticipantStream) Clean() *ParticipantStream {
	*self = ParticipantStreamOf()
	return self
}
func (self *ParticipantStream) Delete(index int) *ParticipantStream {
	return self.DeleteRange(index, index)
}
func (self *ParticipantStream) DeleteRange(startIndex, endIndex int) *ParticipantStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ParticipantStream) Distinct() *ParticipantStream {
	caches := map[string]bool{}
	result := ParticipantStreamOf()
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
func (self *ParticipantStream) Each(fn func(Participant)) *ParticipantStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ParticipantStream) EachRight(fn func(Participant)) *ParticipantStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ParticipantStream) Equals(arr []Participant) bool {
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
func (self *ParticipantStream) Filter(fn func(Participant, int) bool) *ParticipantStream {
	result := ParticipantStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ParticipantStream) FilterSlim(fn func(Participant, int) bool) *ParticipantStream {
	result := ParticipantStreamOf()
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
func (self *ParticipantStream) Find(fn func(Participant, int) bool) *Participant {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ParticipantStream) FindOr(fn func(Participant, int) bool, or Participant) Participant {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ParticipantStream) FindIndex(fn func(Participant, int) bool) int {
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
func (self *ParticipantStream) First() *Participant {
	return self.Get(0)
}
func (self *ParticipantStream) FirstOr(arg Participant) Participant {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ParticipantStream) ForEach(fn func(Participant, int)) *ParticipantStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ParticipantStream) ForEachRight(fn func(Participant, int)) *ParticipantStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ParticipantStream) GroupBy(fn func(Participant, int) string) map[string][]Participant {
	m := map[string][]Participant{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ParticipantStream) GroupByValues(fn func(Participant, int) string) [][]Participant {
	var tmp [][]Participant
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ParticipantStream) IndexOf(arg Participant) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ParticipantStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ParticipantStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ParticipantStream) Last() *Participant {
	return self.Get(self.Len() - 1)
}
func (self *ParticipantStream) LastOr(arg Participant) Participant {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ParticipantStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ParticipantStream) Limit(limit int) *ParticipantStream {
	self.Slice(0, limit)
	return self
}

func (self *ParticipantStream) Map(fn func(Participant, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ParticipantStream) Map2Int(fn func(Participant, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ParticipantStream) Map2Int32(fn func(Participant, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ParticipantStream) Map2Int64(fn func(Participant, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ParticipantStream) Map2Float32(fn func(Participant, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ParticipantStream) Map2Float64(fn func(Participant, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ParticipantStream) Map2Bool(fn func(Participant, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ParticipantStream) Map2Bytes(fn func(Participant, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ParticipantStream) Map2String(fn func(Participant, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ParticipantStream) Max(fn func(Participant, int) float64) *Participant {
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
func (self *ParticipantStream) Min(fn func(Participant, int) float64) *Participant {
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
func (self *ParticipantStream) NoneMatch(fn func(Participant, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ParticipantStream) Get(index int) *Participant {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ParticipantStream) GetOr(index int, arg Participant) Participant {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ParticipantStream) Peek(fn func(*Participant, int)) *ParticipantStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ParticipantStream) Reduce(fn func(Participant, Participant, int) Participant) *ParticipantStream {
	return self.ReduceInit(fn, Participant{})
}
func (self *ParticipantStream) ReduceInit(fn func(Participant, Participant, int) Participant, initialValue Participant) *ParticipantStream {
	result := ParticipantStreamOf()
	self.ForEach(func(v Participant, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ParticipantStream) ReduceInterface(fn func(interface{}, Participant, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Participant{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ParticipantStream) ReduceString(fn func(string, Participant, int) string) []string {
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
func (self *ParticipantStream) ReduceInt(fn func(int, Participant, int) int) []int {
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
func (self *ParticipantStream) ReduceInt32(fn func(int32, Participant, int) int32) []int32 {
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
func (self *ParticipantStream) ReduceInt64(fn func(int64, Participant, int) int64) []int64 {
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
func (self *ParticipantStream) ReduceFloat32(fn func(float32, Participant, int) float32) []float32 {
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
func (self *ParticipantStream) ReduceFloat64(fn func(float64, Participant, int) float64) []float64 {
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
func (self *ParticipantStream) ReduceBool(fn func(bool, Participant, int) bool) []bool {
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
func (self *ParticipantStream) Reverse() *ParticipantStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ParticipantStream) Replace(fn func(Participant, int) Participant) *ParticipantStream {
	return self.ForEach(func(v Participant, i int) { self.Set(i, fn(v, i)) })
}
func (self *ParticipantStream) Select(fn func(Participant) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ParticipantStream) Set(index int, val Participant) *ParticipantStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ParticipantStream) Skip(skip int) *ParticipantStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ParticipantStream) SkippingEach(fn func(Participant, int) int) *ParticipantStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ParticipantStream) Slice(startIndex, n int) *ParticipantStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Participant{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ParticipantStream) Sort(fn func(i, j int) bool) *ParticipantStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ParticipantStream) Tail() *Participant {
	return self.Last()
}
func (self *ParticipantStream) TailOr(arg Participant) Participant {
	return self.LastOr(arg)
}
func (self *ParticipantStream) ToList() []Participant {
	return self.Val()
}
func (self *ParticipantStream) Unique() *ParticipantStream {
	return self.Distinct()
}
func (self *ParticipantStream) Val() []Participant {
	if self == nil {
		return []Participant{}
	}
	return *self.Copy()
}
func (self *ParticipantStream) While(fn func(Participant, int) bool) *ParticipantStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ParticipantStream) Where(fn func(Participant) bool) *ParticipantStream {
	result := ParticipantStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ParticipantStream) WhereSlim(fn func(Participant) bool) *ParticipantStream {
	result := ParticipantStreamOf()
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
