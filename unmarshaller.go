package gocsv

import (
	"encoding/csv"
	"reflect"
)

// Unmarshaller is a CSV to struct unmarshaller.
type Unmarshaller struct {
	reader                 *csv.Reader
	Headers                []string
	fieldInfoMap           []*fieldInfo
	MismatchedHeaders      []string
	MismatchedStructFields []string
	outType                reflect.Type
	out                    interface{}
}

// NewUnmarshaller creates an unmarshaller from a csv.Reader and a struct.
func NewUnmarshaller(reader *csv.Reader, out interface{}) (*Unmarshaller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read returns an interface{} whose runtime type is the same as the struct that
// was used to create the Unmarshaller.
func (um *Unmarshaller) Read() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadUnmatched is same as Read(), but returns a map of the columns that didn't match a field in the struct
func (um *Unmarshaller) ReadUnmatched() (interface{}, map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// validate ensures that a struct was used to create the Unmarshaller, and validates
// CSV headers against the CSV tags in the struct.
func validate(um *Unmarshaller, s interface{}, headers []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get struct info to get CSV annotations.

// Used to store the corresponding header <-> position in CSV

// unmarshalRow converts a CSV row to a struct, based on CSV struct tags.
// If unmatched is non nil, it is populated with any columns that don't map to a struct field
func (um *Unmarshaller) unmarshalRow(row []string, unmatched map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set field of struct

// RenormalizeHeaders will remap the header names based on the headerNormalizer.
// This can be used to map a CSV to a struct where the CSV header names do not match in the file but a mapping is known
func (um *Unmarshaller) RenormalizeHeaders(headerNormalizer func([]string) []string) error {
	_ = "STUB: not implemented"
	return nil
}
