package gocsv

import (
	"reflect"
	"sync"
)

// --------------------------------------------------------------------------
// Reflection helpers

type structInfo struct {
	Fields []fieldInfo
}

// fieldInfo is a struct field that should be mapped to a CSV column, or vice-versa
// Each IndexChain element before the last is the index of an the embedded struct field
// that defines Key as a tag
type fieldInfo struct {
	keys         []string
	omitEmpty    bool
	IndexChain   []int
	defaultValue string
	partial      bool
}

func (f fieldInfo) getFirstKey() string { _ = "STUB: not implemented"; return "" }

func (f fieldInfo) matchesKey(key string) bool { _ = "STUB: not implemented"; return false }

// zwchs is Zero Width Characters map
var zwchs = map[rune]struct{}{
	'\u200B': {}, // zero width space (U+200B)
	'\uFEFF': {}, // zero width no-break space (U+FEFF)
	'\u200D': {}, // zero width joiner (U+200D)
	'\u200C': {}, // zero width non-joiner (U+200C)
}

func removeZeroWidthChars(s string) string { _ = "STUB: not implemented"; return "" }

var structInfoCache sync.Map
var structMap = make(map[reflect.Type]*structInfo)
var structMapMutex sync.RWMutex

func getStructInfo(rType reflect.Type) *structInfo { _ = "STUB: not implemented"; return nil }

func getFieldInfos(rType reflect.Type, parentIndexChain []int, parentKeys []string) []fieldInfo {
	_ = "STUB: not implemented"
	return nil
}

// ignore nested structs with - tag

// create cartesian product of keys
// eg: parent keys x field keys

// handle struct

// if the field is a pointer, follow the pointer

// if the field is a struct, create a fieldInfo for each of its fields

// Structs that implement any of the text or CSV marshaling methods
// should result in one value and not have their fields exposed

// if the field is an embedded struct, pass along parent keys

// if the field is an embedded struct, ignore the csv tag

// if the field is a slice or an array, see if it has a `csv[n]` tag

// slices or arrays of Struct get special handling

// if no special csv[] tag was supplied, just include the field directly

// When the field is a slice/array of structs, create a fieldInfo for each index and each field

// copy index chain and append array index

// copy array index chain and append array index

// create cartesian product of keys
// eg: array field keys x struct field keys

// When the field is a slice/array of primitives, create a fieldInfo for each index

// copy index chain and append array index

func filterTags(tagName string, indexChain []int, field reflect.StructField) (*fieldInfo, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handles cases like `csv:"foo, omitempty, default=test"`

// Must be an exact match: using HasPrefix here would mistake a column name that
// happens to start with "partial" (e.g. `csv:"partial_delivery_number"`) for the
// "partial" option, dropping the real column name. See issue #274.

func getConcreteContainerInnerType(in reflect.Type) (inInnerWasPointer bool, inInnerType reflect.Type) {
	_ = "STUB: not implemented"
	return false, *new(reflect.Type)
}

func getConcreteReflectValueAndType(in interface{}) (reflect.Value, reflect.Type) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Type)
}

var errorInterface = reflect.TypeOf((*error)(nil)).Elem()

func isErrorType(outType reflect.Type) bool { _ = "STUB: not implemented"; return false }
