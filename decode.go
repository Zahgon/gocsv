package gocsv

import (
	"errors"
	"io"
	"mime/multipart"
	"reflect"
)

var (
	ErrUnmatchedStructTags = errors.New("unmatched struct tags")
	ErrDoubleHeaderNames   = errors.New("double header names")
)

// Decoder .
type Decoder interface {
	GetCSVRows() ([][]string, error)
}

// SimpleDecoder .
type SimpleDecoder interface {
	GetCSVRow() ([]string, error)
	GetCSVRows() ([][]string, error)
}

type CSVReader interface {
	Read() ([]string, error)
	ReadAll() ([][]string, error)
}

type csvDecoder struct {
	CSVReader
}

func newSimpleDecoderFromReader(r io.Reader) SimpleDecoder {
	_ = "STUB: not implemented"
	return *new(SimpleDecoder)
}

var (
	ErrEmptyCSVFile = errors.New("empty csv file given")
	ErrNoStructTags = errors.New("no csv struct tags found")
)

// NewSimpleDecoderFromCSVReader creates a SimpleDecoder, which may be passed
// to the UnmarshalDecoder* family of functions, from a CSV reader. Note that
// encoding/csv.Reader implements CSVReader, so you can pass one of those
// directly here.
func NewSimpleDecoderFromCSVReader(r CSVReader) SimpleDecoder {
	_ = "STUB: not implemented"
	return *new(SimpleDecoder)
}

func (c csvDecoder) GetCSVRows() ([][]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (c csvDecoder) GetCSVRow() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func mismatchStructFields(structInfo []fieldInfo, headers []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func mismatchHeaderFields(structInfo []fieldInfo, headers []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func maybeMissingStructFields(structInfo []fieldInfo, headers []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check that no header name is repeated twice
func maybeDoubleHeaderNames(headers []string) error { _ = "STUB: not implemented"; return nil }

// apply normalizer func to headers
func normalizeHeaders(headers []string) []string { _ = "STUB: not implemented"; return nil }

// convertTo converts multipart file to io.Reader
func convertTo(file *multipart.File) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func readTo(decoder Decoder, out interface{}) error { _ = "STUB: not implemented"; return nil }

func readToWithErrorHandler(decoder Decoder, errHandler ErrorHandler, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the concrete type (not pointer) (Slice<?> or Array<?>)

// Get the concrete inner type (not pointer) (Container<"?">)

// Get the CSV csvRows

// Ensure the container is big enough to hold the CSV content

// Get the inner struct info to get CSV annotations

// Used to store the correspondance header <-> position in CSV

//add 2 to account for the header & 0-indexing of arrays

// Position found accordingly to header name

// Set field of struct

//add 2 to account for the header & 0-indexing of arrays

func readEach(decoder SimpleDecoder, errHandler ErrorHandler, c interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the concrete type (not pointer)

// Get the concrete inner type (not pointer) (Container<"?">)

// Get the inner struct info to get CSV annotations

// Used to store the correspondance header <-> position in CSV

//add 2 to account for the header & 0-indexing of arrays

// Position found accordingly to header name

// Set field of struct

//add 2 to account for the header & 0-indexing of arrays

func readEachWithoutHeaders(decoder SimpleDecoder, c interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the concrete type (not pointer) (Slice<?> or Array<?>)

// Get the concrete inner type (not pointer) (Container<"?">)

// Get the inner struct info to get CSV annotations

// Set field of struct

//add 2 to account for the header & 0-indexing of arrays

func readToWithoutHeaders(decoder Decoder, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the concrete type (not pointer) (Slice<?> or Array<?>)

// Get the concrete inner type (not pointer) (Container<"?">)

// Get the CSV csvRows

// Ensure the container is big enough to hold the CSV content

// Get the inner struct info to get CSV annotations

// Set field of struct

// Check if the outType is an array or a slice
func ensureOutType(outType reflect.Type) error { _ = "STUB: not implemented"; return nil }

// Check if the outInnerType is of type struct
func ensureOutInnerType(outInnerType reflect.Type) error { _ = "STUB: not implemented"; return nil }

func ensureOutCapacity(out *reflect.Value, csvLen int) error { _ = "STUB: not implemented"; return nil }

// Array is not big enough to hold the CSV content (arrays are not addressable)

// Slice is not big enough tho hold the CSV content and is not addressable

// Slice is not big enough, so grows it

func getCSVFieldPosition(key string, structInfo *structInfo, curHeaderCount int) *fieldInfo {
	_ = "STUB: not implemented"
	return nil
}

func createNewOutInner(outInnerWasPointer bool, outInnerType reflect.Type) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func setInnerField(outInner *reflect.Value, outInnerWasPointer bool, index []int, value string, omitEmpty bool) error {
	_ = "STUB: not implemented"
	return nil
}

// initialize nil pointer

// grow slice when needed

// because pointers can be nil need to recurse one index at a time and perform nil check
