package gocsv

import (
	"encoding"
	"reflect"
)

// --------------------------------------------------------------------------
// Conversion interfaces

var (
	marshalerType              = reflect.TypeOf(new(TypeMarshaller)).Elem()
	textMarshalerType          = reflect.TypeOf(new(encoding.TextMarshaler)).Elem()
	unmarshalerType            = reflect.TypeOf(new(TypeUnmarshaller)).Elem()
	unmarshalCSVWithFieldsType = reflect.TypeOf(new(TypeUnmarshalCSVWithFields)).Elem()
)

// TypeMarshaller is implemented by any value that has a MarshalCSV method
// This converter is used to convert the value to it string representation
type TypeMarshaller interface {
	MarshalCSV() (string, error)
}

// TypeUnmarshaller is implemented by any value that has an UnmarshalCSV method
// This converter is used to convert a string to your value representation of that string
type TypeUnmarshaller interface {
	UnmarshalCSV(string) error
}

// TypeUnmarshalCSVWithFields can be implemented on whole structs to allow for whole structures to customized internal vs one off fields
type TypeUnmarshalCSVWithFields interface {
	UnmarshalCSVWithFields(key, value string) error
}

// NoUnmarshalFuncError is the custom error type to be raised in case there is no unmarshal function defined on type
type NoUnmarshalFuncError struct {
	t reflect.Type
}

func (e NoUnmarshalFuncError) Error() string { _ = "STUB: not implemented"; return "" }

// NoMarshalFuncError is the custom error type to be raised in case there is no marshal function defined on type
type NoMarshalFuncError struct {
	ty reflect.Type
}

func (e NoMarshalFuncError) Error() string { _ = "STUB: not implemented"; return "" }

// --------------------------------------------------------------------------
// Conversion helpers

func toString(in interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func toBool(in interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func toInt(in interface{}) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func toUint(in interface{}) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// support the float input

func toFloat(in interface{}) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func setField(field reflect.Value, value string, omitEmpty bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Not a native type, check for unmarshal method

// Could not unmarshal, check for kind, e.g. renamed type from basic type

func getFieldAsString(field reflect.Value) (str string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check if field is go native type

// Not a native type, check for marshal method

// If not marshal method, is field compatible with/renamed from native type

// --------------------------------------------------------------------------
// Un/serializations helpers

func canMarshal(t reflect.Type) bool {
	_ = "STUB: not implemented"
	// Struct that implements any of the text or CSV marshaling interfaces
	return false
}

// Pointer to a struct that implements any of the text or CSV marshaling interfaces

func unmarshall(field reflect.Value, value string) error { _ = "STUB: not implemented"; return nil }

// Otherwise try to use TextUnmarshaler

func marshall(field reflect.Value) (value string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Use TypeMarshaller when possible

// Otherwise try to use TextMarshaller

// Otherwise try to use Stringer
