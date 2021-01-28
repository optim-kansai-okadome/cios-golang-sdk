/*
 * Collection utility of NullableMultipleConstitutionComponent Struct
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

type NullableMultipleConstitutionComponentStream []NullableMultipleConstitutionComponent

func NullableMultipleConstitutionComponentStreamOf(arg ...NullableMultipleConstitutionComponent) NullableMultipleConstitutionComponentStream {
	return arg
}
func NullableMultipleConstitutionComponentStreamFrom(arg []NullableMultipleConstitutionComponent) NullableMultipleConstitutionComponentStream {
	return arg
}
func CreateNullableMultipleConstitutionComponentStream(arg ...NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	tmp := NullableMultipleConstitutionComponentStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleConstitutionComponentStream(arg []NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	tmp := NullableMultipleConstitutionComponentStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleConstitutionComponentStream) Add(arg NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleConstitutionComponentStream) AddAll(arg ...NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleConstitutionComponentStream) AddSafe(arg *NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) Aggregate(fn func(NullableMultipleConstitutionComponent, NullableMultipleConstitutionComponent) NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	result := NullableMultipleConstitutionComponentStreamOf()
	self.ForEach(func(v NullableMultipleConstitutionComponent, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleConstitutionComponent{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleConstitutionComponentStream) AllMatch(fn func(NullableMultipleConstitutionComponent, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleConstitutionComponentStream) AnyMatch(fn func(NullableMultipleConstitutionComponent, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleConstitutionComponentStream) Clone() *NullableMultipleConstitutionComponentStream {
	temp := make([]NullableMultipleConstitutionComponent, self.Len())
	copy(temp, *self)
	return (*NullableMultipleConstitutionComponentStream)(&temp)
}
func (self *NullableMultipleConstitutionComponentStream) Copy() *NullableMultipleConstitutionComponentStream {
	return self.Clone()
}
func (self *NullableMultipleConstitutionComponentStream) Concat(arg []NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleConstitutionComponentStream) Contains(arg NullableMultipleConstitutionComponent) bool {
	return self.FindIndex(func(_arg NullableMultipleConstitutionComponent, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleConstitutionComponentStream) Clean() *NullableMultipleConstitutionComponentStream {
	*self = NullableMultipleConstitutionComponentStreamOf()
	return self
}
func (self *NullableMultipleConstitutionComponentStream) Delete(index int) *NullableMultipleConstitutionComponentStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleConstitutionComponentStream) DeleteRange(startIndex, endIndex int) *NullableMultipleConstitutionComponentStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleConstitutionComponentStream) Distinct() *NullableMultipleConstitutionComponentStream {
	caches := map[string]bool{}
	result := NullableMultipleConstitutionComponentStreamOf()
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
func (self *NullableMultipleConstitutionComponentStream) Each(fn func(NullableMultipleConstitutionComponent)) *NullableMultipleConstitutionComponentStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) EachRight(fn func(NullableMultipleConstitutionComponent)) *NullableMultipleConstitutionComponentStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) Equals(arr []NullableMultipleConstitutionComponent) bool {
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
func (self *NullableMultipleConstitutionComponentStream) Filter(fn func(NullableMultipleConstitutionComponent, int) bool) *NullableMultipleConstitutionComponentStream {
	result := NullableMultipleConstitutionComponentStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleConstitutionComponentStream) FilterSlim(fn func(NullableMultipleConstitutionComponent, int) bool) *NullableMultipleConstitutionComponentStream {
	result := NullableMultipleConstitutionComponentStreamOf()
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
func (self *NullableMultipleConstitutionComponentStream) Find(fn func(NullableMultipleConstitutionComponent, int) bool) *NullableMultipleConstitutionComponent {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleConstitutionComponentStream) FindOr(fn func(NullableMultipleConstitutionComponent, int) bool, or NullableMultipleConstitutionComponent) NullableMultipleConstitutionComponent {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleConstitutionComponentStream) FindIndex(fn func(NullableMultipleConstitutionComponent, int) bool) int {
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
func (self *NullableMultipleConstitutionComponentStream) First() *NullableMultipleConstitutionComponent {
	return self.Get(0)
}
func (self *NullableMultipleConstitutionComponentStream) FirstOr(arg NullableMultipleConstitutionComponent) NullableMultipleConstitutionComponent {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleConstitutionComponentStream) ForEach(fn func(NullableMultipleConstitutionComponent, int)) *NullableMultipleConstitutionComponentStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) ForEachRight(fn func(NullableMultipleConstitutionComponent, int)) *NullableMultipleConstitutionComponentStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) GroupBy(fn func(NullableMultipleConstitutionComponent, int) string) map[string][]NullableMultipleConstitutionComponent {
	m := map[string][]NullableMultipleConstitutionComponent{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleConstitutionComponentStream) GroupByValues(fn func(NullableMultipleConstitutionComponent, int) string) [][]NullableMultipleConstitutionComponent {
	var tmp [][]NullableMultipleConstitutionComponent
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleConstitutionComponentStream) IndexOf(arg NullableMultipleConstitutionComponent) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleConstitutionComponentStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleConstitutionComponentStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleConstitutionComponentStream) Last() *NullableMultipleConstitutionComponent {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleConstitutionComponentStream) LastOr(arg NullableMultipleConstitutionComponent) NullableMultipleConstitutionComponent {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleConstitutionComponentStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleConstitutionComponentStream) Limit(limit int) *NullableMultipleConstitutionComponentStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleConstitutionComponentStream) Map(fn func(NullableMultipleConstitutionComponent, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Map2Int(fn func(NullableMultipleConstitutionComponent, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Map2Int32(fn func(NullableMultipleConstitutionComponent, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Map2Int64(fn func(NullableMultipleConstitutionComponent, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Map2Float32(fn func(NullableMultipleConstitutionComponent, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Map2Float64(fn func(NullableMultipleConstitutionComponent, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Map2Bool(fn func(NullableMultipleConstitutionComponent, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Map2Bytes(fn func(NullableMultipleConstitutionComponent, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Map2String(fn func(NullableMultipleConstitutionComponent, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Max(fn func(NullableMultipleConstitutionComponent, int) float64) *NullableMultipleConstitutionComponent {
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
func (self *NullableMultipleConstitutionComponentStream) Min(fn func(NullableMultipleConstitutionComponent, int) float64) *NullableMultipleConstitutionComponent {
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
func (self *NullableMultipleConstitutionComponentStream) NoneMatch(fn func(NullableMultipleConstitutionComponent, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleConstitutionComponentStream) Get(index int) *NullableMultipleConstitutionComponent {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleConstitutionComponentStream) GetOr(index int, arg NullableMultipleConstitutionComponent) NullableMultipleConstitutionComponent {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleConstitutionComponentStream) Peek(fn func(*NullableMultipleConstitutionComponent, int)) *NullableMultipleConstitutionComponentStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMultipleConstitutionComponentStream) Reduce(fn func(NullableMultipleConstitutionComponent, NullableMultipleConstitutionComponent, int) NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	return self.ReduceInit(fn, NullableMultipleConstitutionComponent{})
}
func (self *NullableMultipleConstitutionComponentStream) ReduceInit(fn func(NullableMultipleConstitutionComponent, NullableMultipleConstitutionComponent, int) NullableMultipleConstitutionComponent, initialValue NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	result := NullableMultipleConstitutionComponentStreamOf()
	self.ForEach(func(v NullableMultipleConstitutionComponent, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleConstitutionComponentStream) ReduceInterface(fn func(interface{}, NullableMultipleConstitutionComponent, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleConstitutionComponent{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleConstitutionComponentStream) ReduceString(fn func(string, NullableMultipleConstitutionComponent, int) string) []string {
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
func (self *NullableMultipleConstitutionComponentStream) ReduceInt(fn func(int, NullableMultipleConstitutionComponent, int) int) []int {
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
func (self *NullableMultipleConstitutionComponentStream) ReduceInt32(fn func(int32, NullableMultipleConstitutionComponent, int) int32) []int32 {
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
func (self *NullableMultipleConstitutionComponentStream) ReduceInt64(fn func(int64, NullableMultipleConstitutionComponent, int) int64) []int64 {
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
func (self *NullableMultipleConstitutionComponentStream) ReduceFloat32(fn func(float32, NullableMultipleConstitutionComponent, int) float32) []float32 {
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
func (self *NullableMultipleConstitutionComponentStream) ReduceFloat64(fn func(float64, NullableMultipleConstitutionComponent, int) float64) []float64 {
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
func (self *NullableMultipleConstitutionComponentStream) ReduceBool(fn func(bool, NullableMultipleConstitutionComponent, int) bool) []bool {
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
func (self *NullableMultipleConstitutionComponentStream) Reverse() *NullableMultipleConstitutionComponentStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) Replace(fn func(NullableMultipleConstitutionComponent, int) NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	return self.ForEach(func(v NullableMultipleConstitutionComponent, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleConstitutionComponentStream) Select(fn func(NullableMultipleConstitutionComponent) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleConstitutionComponentStream) Set(index int, val NullableMultipleConstitutionComponent) *NullableMultipleConstitutionComponentStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) Skip(skip int) *NullableMultipleConstitutionComponentStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleConstitutionComponentStream) SkippingEach(fn func(NullableMultipleConstitutionComponent, int) int) *NullableMultipleConstitutionComponentStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) Slice(startIndex, n int) *NullableMultipleConstitutionComponentStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleConstitutionComponent{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) Sort(fn func(i, j int) bool) *NullableMultipleConstitutionComponentStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleConstitutionComponentStream) Tail() *NullableMultipleConstitutionComponent {
	return self.Last()
}
func (self *NullableMultipleConstitutionComponentStream) TailOr(arg NullableMultipleConstitutionComponent) NullableMultipleConstitutionComponent {
	return self.LastOr(arg)
}
func (self *NullableMultipleConstitutionComponentStream) ToList() []NullableMultipleConstitutionComponent {
	return self.Val()
}
func (self *NullableMultipleConstitutionComponentStream) Unique() *NullableMultipleConstitutionComponentStream {
	return self.Distinct()
}
func (self *NullableMultipleConstitutionComponentStream) Val() []NullableMultipleConstitutionComponent {
	if self == nil {
		return []NullableMultipleConstitutionComponent{}
	}
	return *self.Copy()
}
func (self *NullableMultipleConstitutionComponentStream) While(fn func(NullableMultipleConstitutionComponent, int) bool) *NullableMultipleConstitutionComponentStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleConstitutionComponentStream) Where(fn func(NullableMultipleConstitutionComponent) bool) *NullableMultipleConstitutionComponentStream {
	result := NullableMultipleConstitutionComponentStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleConstitutionComponentStream) WhereSlim(fn func(NullableMultipleConstitutionComponent) bool) *NullableMultipleConstitutionComponentStream {
	result := NullableMultipleConstitutionComponentStreamOf()
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
