// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef = __mainPackageInFile_types_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __mainPackageInFile_types_Ref() ref.Ref {
	p := types.PackageDef{
		NamedTypes: types.MapOfStringToTypeRefDef{

			"Pitch": types.MakeStructTypeRef("Pitch",
				[]types.Field{
					types.Field{"X", types.MakePrimitiveTypeRef(types.Float64Kind), false},
					types.Field{"Z", types.MakePrimitiveTypeRef(types.Float64Kind), false},
				},
				types.Choices{},
			),
		},
	}.New()
	return types.RegisterPackage(&p)
}

// ListOfMapOfStringToValue

type ListOfMapOfStringToValue struct {
	l types.List
}

func NewListOfMapOfStringToValue() ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{types.NewList()}
}

type ListOfMapOfStringToValueDef []MapOfStringToValueDef

func (def ListOfMapOfStringToValueDef) New() ListOfMapOfStringToValue {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New().NomsValue()
	}
	return ListOfMapOfStringToValue{types.NewList(l...)}
}

func (l ListOfMapOfStringToValue) Def() ListOfMapOfStringToValueDef {
	d := make([]MapOfStringToValueDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = MapOfStringToValueFromVal(l.l.Get(i)).Def()
	}
	return d
}

func ListOfMapOfStringToValueFromVal(val types.Value) ListOfMapOfStringToValue {
	// TODO: Validate here
	return ListOfMapOfStringToValue{val.(types.List)}
}

func (l ListOfMapOfStringToValue) NomsValue() types.Value {
	return l.l
}

func (l ListOfMapOfStringToValue) Equals(other types.Value) bool {
	if other, ok := other.(ListOfMapOfStringToValue); ok {
		return l.l.Equals(other.l)
	}
	return false
}

func (l ListOfMapOfStringToValue) Ref() ref.Ref {
	return l.l.Ref()
}

func (l ListOfMapOfStringToValue) Chunks() (futures []types.Future) {
	futures = append(futures, l.TypeRef().Chunks()...)
	futures = append(futures, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfMapOfStringToValue.
var __typeRefForListOfMapOfStringToValue types.TypeRef

func (m ListOfMapOfStringToValue) TypeRef() types.TypeRef {
	return __typeRefForListOfMapOfStringToValue
}

func init() {
	__typeRefForListOfMapOfStringToValue = types.MakeCompoundTypeRef("", types.ListKind, types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.ValueKind)))
	types.RegisterFromValFunction(__typeRefForListOfMapOfStringToValue, func(v types.Value) types.NomsValue {
		return ListOfMapOfStringToValueFromVal(v)
	})
}

func (l ListOfMapOfStringToValue) Len() uint64 {
	return l.l.Len()
}

func (l ListOfMapOfStringToValue) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfMapOfStringToValue) Get(i uint64) MapOfStringToValue {
	return MapOfStringToValueFromVal(l.l.Get(i))
}

func (l ListOfMapOfStringToValue) Slice(idx uint64, end uint64) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Slice(idx, end)}
}

func (l ListOfMapOfStringToValue) Set(i uint64, val MapOfStringToValue) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Set(i, val.NomsValue())}
}

func (l ListOfMapOfStringToValue) Append(v ...MapOfStringToValue) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfMapOfStringToValue) Insert(idx uint64, v ...MapOfStringToValue) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfMapOfStringToValue) Remove(idx uint64, end uint64) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Remove(idx, end)}
}

func (l ListOfMapOfStringToValue) RemoveAt(idx uint64) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{(l.l.RemoveAt(idx))}
}

func (l ListOfMapOfStringToValue) fromElemSlice(p []MapOfStringToValue) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

type ListOfMapOfStringToValueIterCallback func(v MapOfStringToValue, i uint64) (stop bool)

func (l ListOfMapOfStringToValue) Iter(cb ListOfMapOfStringToValueIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(MapOfStringToValueFromVal(v), i)
	})
}

type ListOfMapOfStringToValueIterAllCallback func(v MapOfStringToValue, i uint64)

func (l ListOfMapOfStringToValue) IterAll(cb ListOfMapOfStringToValueIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(MapOfStringToValueFromVal(v), i)
	})
}

type ListOfMapOfStringToValueFilterCallback func(v MapOfStringToValue, i uint64) (keep bool)

func (l ListOfMapOfStringToValue) Filter(cb ListOfMapOfStringToValueFilterCallback) ListOfMapOfStringToValue {
	nl := NewListOfMapOfStringToValue()
	l.IterAll(func(v MapOfStringToValue, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// MapOfStringToValue

type MapOfStringToValue struct {
	m types.Map
}

func NewMapOfStringToValue() MapOfStringToValue {
	return MapOfStringToValue{types.NewMap()}
}

type MapOfStringToValueDef map[string]types.Value

func (def MapOfStringToValueDef) New() MapOfStringToValue {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v)
	}
	return MapOfStringToValue{types.NewMap(kv...)}
}

func (m MapOfStringToValue) Def() MapOfStringToValueDef {
	def := make(map[string]types.Value)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v
		return false
	})
	return def
}

func MapOfStringToValueFromVal(p types.Value) MapOfStringToValue {
	// TODO: Validate here
	return MapOfStringToValue{p.(types.Map)}
}

func (m MapOfStringToValue) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToValue) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToValue); ok {
		return m.m.Equals(other.m)
	}
	return false
}

func (m MapOfStringToValue) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToValue) Chunks() (futures []types.Future) {
	futures = append(futures, m.TypeRef().Chunks()...)
	futures = append(futures, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToValue.
var __typeRefForMapOfStringToValue types.TypeRef

func (m MapOfStringToValue) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToValue
}

func init() {
	__typeRefForMapOfStringToValue = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.ValueKind))
	types.RegisterFromValFunction(__typeRefForMapOfStringToValue, func(v types.Value) types.NomsValue {
		return MapOfStringToValueFromVal(v)
	})
}

func (m MapOfStringToValue) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToValue) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToValue) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToValue) Get(p string) types.Value {
	return m.m.Get(types.NewString(p))
}

func (m MapOfStringToValue) MaybeGet(p string) (types.Value, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return types.Bool(false), false
	}
	return v, ok
}

func (m MapOfStringToValue) Set(k string, v types.Value) MapOfStringToValue {
	return MapOfStringToValue{m.m.Set(types.NewString(k), v)}
}

// TODO: Implement SetM?

func (m MapOfStringToValue) Remove(p string) MapOfStringToValue {
	return MapOfStringToValue{m.m.Remove(types.NewString(p))}
}

type MapOfStringToValueIterCallback func(k string, v types.Value) (stop bool)

func (m MapOfStringToValue) Iter(cb MapOfStringToValueIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueIterAllCallback func(k string, v types.Value)

func (m MapOfStringToValue) IterAll(cb MapOfStringToValueIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueFilterCallback func(k string, v types.Value) (keep bool)

func (m MapOfStringToValue) Filter(cb MapOfStringToValueFilterCallback) MapOfStringToValue {
	nm := NewMapOfStringToValue()
	m.IterAll(func(k string, v types.Value) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// MapOfStringToListOfPitch

type MapOfStringToListOfPitch struct {
	m types.Map
}

func NewMapOfStringToListOfPitch() MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{types.NewMap()}
}

type MapOfStringToListOfPitchDef map[string]ListOfPitchDef

func (def MapOfStringToListOfPitchDef) New() MapOfStringToListOfPitch {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New().NomsValue())
	}
	return MapOfStringToListOfPitch{types.NewMap(kv...)}
}

func (m MapOfStringToListOfPitch) Def() MapOfStringToListOfPitchDef {
	def := make(map[string]ListOfPitchDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = ListOfPitchFromVal(v).Def()
		return false
	})
	return def
}

func MapOfStringToListOfPitchFromVal(p types.Value) MapOfStringToListOfPitch {
	// TODO: Validate here
	return MapOfStringToListOfPitch{p.(types.Map)}
}

func (m MapOfStringToListOfPitch) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToListOfPitch) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToListOfPitch); ok {
		return m.m.Equals(other.m)
	}
	return false
}

func (m MapOfStringToListOfPitch) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToListOfPitch) Chunks() (futures []types.Future) {
	futures = append(futures, m.TypeRef().Chunks()...)
	futures = append(futures, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToListOfPitch.
var __typeRefForMapOfStringToListOfPitch types.TypeRef

func (m MapOfStringToListOfPitch) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToListOfPitch
}

func init() {
	__typeRefForMapOfStringToListOfPitch = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeCompoundTypeRef("", types.ListKind, types.MakeTypeRef("Pitch", __mainPackageInFile_types_CachedRef)))
	types.RegisterFromValFunction(__typeRefForMapOfStringToListOfPitch, func(v types.Value) types.NomsValue {
		return MapOfStringToListOfPitchFromVal(v)
	})
}

func (m MapOfStringToListOfPitch) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToListOfPitch) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToListOfPitch) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToListOfPitch) Get(p string) ListOfPitch {
	return ListOfPitchFromVal(m.m.Get(types.NewString(p)))
}

func (m MapOfStringToListOfPitch) MaybeGet(p string) (ListOfPitch, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewListOfPitch(), false
	}
	return ListOfPitchFromVal(v), ok
}

func (m MapOfStringToListOfPitch) Set(k string, v ListOfPitch) MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{m.m.Set(types.NewString(k), v.NomsValue())}
}

// TODO: Implement SetM?

func (m MapOfStringToListOfPitch) Remove(p string) MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{m.m.Remove(types.NewString(p))}
}

type MapOfStringToListOfPitchIterCallback func(k string, v ListOfPitch) (stop bool)

func (m MapOfStringToListOfPitch) Iter(cb MapOfStringToListOfPitchIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), ListOfPitchFromVal(v))
	})
}

type MapOfStringToListOfPitchIterAllCallback func(k string, v ListOfPitch)

func (m MapOfStringToListOfPitch) IterAll(cb MapOfStringToListOfPitchIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), ListOfPitchFromVal(v))
	})
}

type MapOfStringToListOfPitchFilterCallback func(k string, v ListOfPitch) (keep bool)

func (m MapOfStringToListOfPitch) Filter(cb MapOfStringToListOfPitchFilterCallback) MapOfStringToListOfPitch {
	nm := NewMapOfStringToListOfPitch()
	m.IterAll(func(k string, v ListOfPitch) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// ListOfPitch

type ListOfPitch struct {
	l types.List
}

func NewListOfPitch() ListOfPitch {
	return ListOfPitch{types.NewList()}
}

type ListOfPitchDef []PitchDef

func (def ListOfPitchDef) New() ListOfPitch {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New().NomsValue()
	}
	return ListOfPitch{types.NewList(l...)}
}

func (l ListOfPitch) Def() ListOfPitchDef {
	d := make([]PitchDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = PitchFromVal(l.l.Get(i)).Def()
	}
	return d
}

func ListOfPitchFromVal(val types.Value) ListOfPitch {
	// TODO: Validate here
	return ListOfPitch{val.(types.List)}
}

func (l ListOfPitch) NomsValue() types.Value {
	return l.l
}

func (l ListOfPitch) Equals(other types.Value) bool {
	if other, ok := other.(ListOfPitch); ok {
		return l.l.Equals(other.l)
	}
	return false
}

func (l ListOfPitch) Ref() ref.Ref {
	return l.l.Ref()
}

func (l ListOfPitch) Chunks() (futures []types.Future) {
	futures = append(futures, l.TypeRef().Chunks()...)
	futures = append(futures, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfPitch.
var __typeRefForListOfPitch types.TypeRef

func (m ListOfPitch) TypeRef() types.TypeRef {
	return __typeRefForListOfPitch
}

func init() {
	__typeRefForListOfPitch = types.MakeCompoundTypeRef("", types.ListKind, types.MakeTypeRef("Pitch", __mainPackageInFile_types_CachedRef))
	types.RegisterFromValFunction(__typeRefForListOfPitch, func(v types.Value) types.NomsValue {
		return ListOfPitchFromVal(v)
	})
}

func (l ListOfPitch) Len() uint64 {
	return l.l.Len()
}

func (l ListOfPitch) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfPitch) Get(i uint64) Pitch {
	return PitchFromVal(l.l.Get(i))
}

func (l ListOfPitch) Slice(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Slice(idx, end)}
}

func (l ListOfPitch) Set(i uint64, val Pitch) ListOfPitch {
	return ListOfPitch{l.l.Set(i, val.NomsValue())}
}

func (l ListOfPitch) Append(v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfPitch) Insert(idx uint64, v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfPitch) Remove(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Remove(idx, end)}
}

func (l ListOfPitch) RemoveAt(idx uint64) ListOfPitch {
	return ListOfPitch{(l.l.RemoveAt(idx))}
}

func (l ListOfPitch) fromElemSlice(p []Pitch) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

type ListOfPitchIterCallback func(v Pitch, i uint64) (stop bool)

func (l ListOfPitch) Iter(cb ListOfPitchIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(PitchFromVal(v), i)
	})
}

type ListOfPitchIterAllCallback func(v Pitch, i uint64)

func (l ListOfPitch) IterAll(cb ListOfPitchIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(PitchFromVal(v), i)
	})
}

type ListOfPitchFilterCallback func(v Pitch, i uint64) (keep bool)

func (l ListOfPitch) Filter(cb ListOfPitchFilterCallback) ListOfPitch {
	nl := NewListOfPitch()
	l.IterAll(func(v Pitch, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// Pitch

type Pitch struct {
	m types.Map
}

func NewPitch() Pitch {
	return Pitch{types.NewMap(
		types.NewString("$name"), types.NewString("Pitch"),
		types.NewString("$type"), types.MakeTypeRef("Pitch", __mainPackageInFile_types_CachedRef),
		types.NewString("X"), types.Float64(0),
		types.NewString("Z"), types.Float64(0),
	)}
}

type PitchDef struct {
	X float64
	Z float64
}

func (def PitchDef) New() Pitch {
	return Pitch{
		types.NewMap(
			types.NewString("$name"), types.NewString("Pitch"),
			types.NewString("$type"), types.MakeTypeRef("Pitch", __mainPackageInFile_types_CachedRef),
			types.NewString("X"), types.Float64(def.X),
			types.NewString("Z"), types.Float64(def.Z),
		)}
}

func (s Pitch) Def() (d PitchDef) {
	d.X = float64(s.m.Get(types.NewString("X")).(types.Float64))
	d.Z = float64(s.m.Get(types.NewString("Z")).(types.Float64))
	return
}

var __typeRefForPitch = types.MakeTypeRef("Pitch", __mainPackageInFile_types_CachedRef)

func (m Pitch) TypeRef() types.TypeRef {
	return __typeRefForPitch
}

func init() {
	types.RegisterFromValFunction(__typeRefForPitch, func(v types.Value) types.NomsValue {
		return PitchFromVal(v)
	})
}

func PitchFromVal(val types.Value) Pitch {
	// TODO: Validate here
	return Pitch{val.(types.Map)}
}

func (s Pitch) NomsValue() types.Value {
	return s.m
}

func (s Pitch) Equals(other types.Value) bool {
	if other, ok := other.(Pitch); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s Pitch) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Pitch) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s Pitch) X() float64 {
	return float64(s.m.Get(types.NewString("X")).(types.Float64))
}

func (s Pitch) SetX(val float64) Pitch {
	return Pitch{s.m.Set(types.NewString("X"), types.Float64(val))}
}

func (s Pitch) Z() float64 {
	return float64(s.m.Get(types.NewString("Z")).(types.Float64))
}

func (s Pitch) SetZ(val float64) Pitch {
	return Pitch{s.m.Set(types.NewString("Z"), types.Float64(val))}
}

// MapOfStringToString

type MapOfStringToString struct {
	m types.Map
}

func NewMapOfStringToString() MapOfStringToString {
	return MapOfStringToString{types.NewMap()}
}

type MapOfStringToStringDef map[string]string

func (def MapOfStringToStringDef) New() MapOfStringToString {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), types.NewString(v))
	}
	return MapOfStringToString{types.NewMap(kv...)}
}

func (m MapOfStringToString) Def() MapOfStringToStringDef {
	def := make(map[string]string)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(types.String).String()
		return false
	})
	return def
}

func MapOfStringToStringFromVal(p types.Value) MapOfStringToString {
	// TODO: Validate here
	return MapOfStringToString{p.(types.Map)}
}

func (m MapOfStringToString) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToString) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToString); ok {
		return m.m.Equals(other.m)
	}
	return false
}

func (m MapOfStringToString) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToString) Chunks() (futures []types.Future) {
	futures = append(futures, m.TypeRef().Chunks()...)
	futures = append(futures, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToString.
var __typeRefForMapOfStringToString types.TypeRef

func (m MapOfStringToString) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToString
}

func init() {
	__typeRefForMapOfStringToString = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.StringKind))
	types.RegisterFromValFunction(__typeRefForMapOfStringToString, func(v types.Value) types.NomsValue {
		return MapOfStringToStringFromVal(v)
	})
}

func (m MapOfStringToString) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToString) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToString) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToString) Get(p string) string {
	return m.m.Get(types.NewString(p)).(types.String).String()
}

func (m MapOfStringToString) MaybeGet(p string) (string, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return "", false
	}
	return v.(types.String).String(), ok
}

func (m MapOfStringToString) Set(k string, v string) MapOfStringToString {
	return MapOfStringToString{m.m.Set(types.NewString(k), types.NewString(v))}
}

// TODO: Implement SetM?

func (m MapOfStringToString) Remove(p string) MapOfStringToString {
	return MapOfStringToString{m.m.Remove(types.NewString(p))}
}

type MapOfStringToStringIterCallback func(k string, v string) (stop bool)

func (m MapOfStringToString) Iter(cb MapOfStringToStringIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(types.String).String())
	})
}

type MapOfStringToStringIterAllCallback func(k string, v string)

func (m MapOfStringToString) IterAll(cb MapOfStringToStringIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(types.String).String())
	})
}

type MapOfStringToStringFilterCallback func(k string, v string) (keep bool)

func (m MapOfStringToString) Filter(cb MapOfStringToStringFilterCallback) MapOfStringToString {
	nm := NewMapOfStringToString()
	m.IterAll(func(k string, v string) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}
