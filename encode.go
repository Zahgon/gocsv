package gocsv

import (
	"errors"
	"io"
	"reflect"
)

var (
	ErrChannelIsClosed = errors.New("channel is closed")
)

type encoder struct {
	out io.Writer
}

func newEncoder(out io.Writer) *encoder { _ = "STUB: not implemented"; return nil }

func writeFromChan(writer CSVWriter, c <-chan interface{}, omitHeaders bool) error {
	_ = "STUB: not implemented"
	// Get the first value. It wil determine the header structure.
	return nil
}

// Get the concrete type

// Get the inner struct info to get CSV annotations

// Used to write the header (first line) in CSV

// Get the correct field header <-> position

// Get the concrete type (not pointer) (Slice<?> or Array<?>)

func writeTo(writer CSVWriter, in interface{}, omitHeaders bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the concrete type (not pointer) (Slice<?> or Array<?>)

// Get the concrete inner type (not pointer) (Container<"?">)

// Get the inner struct info to get CSV annotations

// Used to write the header (first line) in CSV

// Iterate over container rows

// Get the correct field header <-> position

func ensureStructOrPtr(t reflect.Type) error { _ = "STUB: not implemented"; return nil }

// Check if the inType is an array or a slice
func ensureInType(outType reflect.Type) error { _ = "STUB: not implemented"; return nil }

// Check if the inInnerType is of type struct
func ensureInInnerType(outInnerType reflect.Type) error { _ = "STUB: not implemented"; return nil }

func getInnerField(outInner reflect.Value, outInnerWasPointer bool, index []int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// because pointers can be nil need to recurse one index at a time and perform nil check
