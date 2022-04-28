// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.

// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/switchupcb/copygen/examples/_tests/multi/complex"
	"github.com/switchupcb/copygen/examples/_tests/multi/external"
)

// Placeholder represents a basic type.
type Placeholder bool

// Collection represents a type that holds collection field types.

// empty represents a struct that contains an empty struct.

// freefloat serves the purpose of checking for free-floating comments.
type freefloat struct {
	A string
}

type Collection struct {
	Arr [16]byte
	S   []string
	M   map[string]bool
	C   chan int
	I   error
	F   func() int
}

type empty struct {
	e struct{}
}

// Array copies a [16]byte to a [16]byte.
func Array(tb [16]byte, fb [16]byte) {
	// [16]byte fields
	tb = fb
}

// ArrayComplex copies a [16]map[byte]string to a *complex.Collection.
func ArrayComplex(tC *complex.Collection, fm [16]map[byte]string) {
	// *complex.Collection fields
	tC.Arr = fm
}

// ArrayExternal copies a [16]Placeholder to a [16]Placeholder.
func ArrayExternal(tP [16]Placeholder, fP [16]Placeholder) {
	// [16]Placeholder fields
	tP = fP
}

// ArrayExternalComplex copies a [16]map[Collection]string to a *complex.ComplexCollection.
func ArrayExternalComplex(tC *complex.ComplexCollection, fm [16]map[Collection]string) {
	// *complex.ComplexCollection fields
	tC.Arr = fm
}

// ArraySimple copies a [16]byte to a *Collection.
func ArraySimple(tC *Collection, fb [16]byte) {
	// *Collection fields
	tC.Arr = fb
}

// Basic copies a bool to a bool.
func Basic(tb bool, fb bool) {
	// bool fields
	tb = fb
}

// BasicDoublePointer copies a *bool to a **bool.
func BasicDoublePointer(tb **bool, fb *bool) {
	// **bool fields
	tb = &fb
}

// BasicExternal copies a *external.Placeholder to a external.Placeholder.
func BasicExternal(tP external.Placeholder, fP *external.Placeholder) {
	// external.Placeholder fields
	tP = *fP
}

// BasicExternalMulti copies a *external.Placeholder to a external.Placeholder, *external.Placeholder.
func BasicExternalMulti(tP external.Placeholder, tP1 *external.Placeholder, fP *external.Placeholder) {
	// external.Placeholder fields
	tP = *fP

	// *external.Placeholder fields
	tP1 = fP
}

// BasicPointer copies a bool to a *bool.
func BasicPointer(tb *bool, fb bool) {
	// *bool fields
	tb = &fb
}

// BasicPointerMulti copies a *Placeholder to a *Placeholder, *Placeholder, string.
func BasicPointerMulti(tP *Placeholder, tP1 *Placeholder, ts string, fP *Placeholder) {
	// *Placeholder fields
	tP = fP

	// *Placeholder fields

	// string fields
}

// BasicSimple copies a Placeholder to a Placeholder.
func BasicSimple(tP Placeholder, fP Placeholder) {
	// Placeholder fields
	tP = fP
}

// BasicSimplePointer copies a Placeholder to a *Placeholder.
func BasicSimplePointer(tP *Placeholder, fP Placeholder) {
	// *Placeholder fields
	tP = &fP
}

// Chan copies a chan int to a chan int.
func Chan(tc chan int, fc chan int) {
	// chan int fields
	tc = fc
}

// ChanComplex copies a chan []int to a *complex.Collection.
func ChanComplex(tC *complex.Collection, fc chan []int) {
	// *complex.Collection fields
	tC.C = fc
}

// ChanExternal copies a chan Placeholder to a chan Placeholder.
func ChanExternal(tc chan Placeholder, fc chan Placeholder) {
	// chan Placeholder fields
	tc = fc
}

// ChanExternalComplex copies a chan []Collection to a complex.ComplexCollection.
func ChanExternalComplex(tC complex.ComplexCollection, fc chan []Collection) {
	// complex.ComplexCollection fields
	tC.C = fc
}

// ChanSimple copies a chan int to a *Collection.
func ChanSimple(tC *Collection, fc chan int) {
	// *Collection fields
	tC.C = fc
}

// EmptyStruct copies a struct{} to a empty.
func EmptyStruct(te empty, fs struct{}) {
	// empty fields
}

// Func copies a func() int to a func() int.
func Func(tf func() int, ff func() int) {
	// func() int fields
	tf = ff
}

// FuncComplex copies a func([]string, uint64) byte to a *complex.Collection.
func FuncComplex(tC *complex.Collection, ff func([]string, uint64) byte) {
	// *complex.Collection fields
	tC.F = ff
}

// FuncExternal copies a func(Placeholder) int to a func(Placeholder) int.
func FuncExternal(tf func(Placeholder) int, ff func(Placeholder) int) {
	// func(Placeholder) int fields
	tf = ff
}

// FuncExternalComplex copies a func(Collection) []string to a *complex.ComplexCollection.
func FuncExternalComplex(tC *complex.ComplexCollection, ff func(Collection) []string) {
	// *complex.ComplexCollection fields
	tC.F = ff
}

// FuncSimple copies a func() int to a *Collection.
func FuncSimple(tC *Collection, ff func() int) {
	// *Collection fields
	tC.F = ff
}

// Interface copies a interface{} to a interface{}.
func Interface(ti interface{}, fi interface{}) {
	// interface{} fields
	ti = fi
}

// InterfaceComplex copies a interface{func(rune) int; } to a *complex.Collection.
func InterfaceComplex(tC *complex.Collection, fi interface{ func(rune) int }) {
	// *complex.Collection fields
	tC.I = fi
}

// InterfaceExternal copies a error to a *external.Collection.
func InterfaceExternal(tC *external.Collection, fe error) {
	// *external.Collection fields
	tC.I = fe
}

// InterfaceExternalComplex copies a interface{func(string) map[Collection]bool; func() (int, byte); } to a complex.ComplexCollection.
func InterfaceExternalComplex(tC complex.ComplexCollection, fi interface {
	func(string) map[Collection]bool
	func() (int, byte)
}) {
	// complex.ComplexCollection fields
	tC.I = fi
}

// InterfaceSimple copies a error to a *Collection.
func InterfaceSimple(tC *Collection, fe error) {
	// *Collection fields
	tC.I = fe
}

// Map copies a map[string]bool to a map[string]bool.
func Map(tm map[string]bool, fm map[string]bool) {
	// map[string]bool fields
	tm = fm
}

// MapComplex copies a map[string]interface{func() string; } to a *complex.Collection.
func MapComplex(tC *complex.Collection, fm map[string]interface{ func() string }) {
	// *complex.Collection fields
}

// MapExternal copies a map[string]Placeholder to a map[string]Placeholder.
func MapExternal(tm map[string]Placeholder, fm map[string]Placeholder) {
	// map[string]Placeholder fields
	tm = fm
}

// MapExternalComplex copies a map[Collection]Placeholder to a *complex.ComplexCollection.
func MapExternalComplex(tC *complex.ComplexCollection, fm map[Collection]Placeholder) {
	// *complex.ComplexCollection fields
	tC.M = fm
}

// MapSimple copies a map[string]bool to a *Collection.
func MapSimple(tC *Collection, fm map[string]bool) {
	// *Collection fields
	tC.M = fm
}

// NoMatchArraySimple copies a [16]byte to a Collection.
func NoMatchArraySimple(tC Collection, fb [16]byte) {
	// Collection fields
}

// NoMatchBasic copies a Placeholder to a Placeholder.
func NoMatchBasic(tP Placeholder, fP Placeholder) {
	// Placeholder fields
}

// NoMatchBasicExternal copies a *Placeholder to a external.Placeholder, *external.Placeholder, bool.
func NoMatchBasicExternal(tP external.Placeholder, tP1 *external.Placeholder, tb bool, fP *Placeholder) {
	// external.Placeholder fields

	// *external.Placeholder fields

	// bool fields
}

// NoMatchChan copies a chan int to a Collection.
func NoMatchChan(tC Collection, fc chan int) {
	// Collection fields
}

// NoMatchComplex copies a []Collection to a []Collection.
func NoMatchComplex(tC []Collection, fC []Collection) {
	// []Collection fields
}

// NoMatchFunc copies a func() int to a Collection.
func NoMatchFunc(tC Collection, ff func() int) {
	// Collection fields
}

// NoMatchInterface copies a error to a Collection.
func NoMatchInterface(tC Collection, fe error) {
	// Collection fields
}

// NoMatchMap copies a map[string]bool to a Collection.
func NoMatchMap(tC Collection, fm map[string]bool) {
	// Collection fields
}

// NoMatchSliceSimple copies a []string to a Collection.
func NoMatchSliceSimple(tC Collection, fs []string) {
	// Collection fields
}

// Slice copies a []string to a []string.
func Slice(ts []string, fs []string) {
	// []string fields
	ts = fs
}

// SliceComplex copies a []map[string][16]int to a *complex.Collection.
func SliceComplex(tC *complex.Collection, fm []map[string][16]int) {
	// *complex.Collection fields
	tC.S = fm
}

// SliceExternal copies a []Placeholder to a []Placeholder.
func SliceExternal(tP []Placeholder, fP []Placeholder) {
	// []Placeholder fields
	tP = fP
}

// SliceExternalComplex copies a []map[string]func(Collection) string to a *complex.ComplexCollection.
func SliceExternalComplex(tC *complex.ComplexCollection, fm []map[string]func(Collection) string) {
	// *complex.ComplexCollection fields
	tC.S = fm
}

// SliceSimple copies a []string to a *Collection.
func SliceSimple(tC *Collection, fs []string) {
	// *Collection fields
	tC.S = fs
}

// Struct copies a Collection to a Collection.
func Struct(tC Collection, fC Collection) {
	// Collection fields
	tC = fC
}

// StructExternal copies a external.Collection to a *Collection.
func StructExternal(tC *Collection, fC external.Collection) {
	// *Collection fields
	tC.Arr = fC.Arr
	tC.S = fC.S
	tC.M = fC.M
	tC.C = fC.C
	tC.I = fC.I
	tC.F = fC.F
}
