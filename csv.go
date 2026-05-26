// Copyright 2014 Jonathan Picques. All rights reserved.
// Use of this source code is governed by a MIT license
// The license can be found in the LICENSE file.

// The GoCSV package aims to provide easy CSV serialization and deserialization to the golang programming language

package gocsv

import (
	"encoding/csv"
	"io"
	"mime/multipart"
	"os"
)

// FailIfUnmatchedStructTags indicates whether it is considered an error when there is an unmatched
// struct tag.
var FailIfUnmatchedStructTags = false

// FailIfDoubleHeaderNames indicates whether it is considered an error when a header name is repeated
// in the csv header.
var FailIfDoubleHeaderNames = false

// ShouldAlignDuplicateHeadersWithStructFieldOrder indicates whether we should align duplicate CSV
// headers per their alignment in the struct definition.
var ShouldAlignDuplicateHeadersWithStructFieldOrder = false

// TagName defines key in the struct field's tag to scan
var TagName = "csv"

// TagSeparator defines seperator string for multiple csv tags in struct fields
var TagSeparator = ","

// FieldSeperator defines how to combine parent struct with child struct
var FieldsCombiner = "."

// Normalizer is a function that takes and returns a string. It is applied to
// struct and header field values before they are compared. It can be used to alter
// names for comparison. For instance, you could allow case insensitive matching
// or convert '-' to '_'.
type Normalizer func(string) string

type ErrorHandler func(*csv.ParseError) bool

// normalizeName function initially set to a nop Normalizer.
var normalizeName = DefaultNameNormalizer()

// DefaultNameNormalizer is a nop Normalizer.
func DefaultNameNormalizer() Normalizer { _ = "STUB: not implemented"; return *new(Normalizer) }

// SetHeaderNormalizer sets the normalizer used to normalize struct and header field names.
func SetHeaderNormalizer(f Normalizer) {
	_ = "STUB: not implemented"

	// Need to clear the cache hen the header normalizer changes.
	return
}

// --------------------------------------------------------------------------
// CSVWriter used to format CSV

var selfCSVWriter = DefaultCSVWriter

// DefaultCSVWriter is the default SafeCSVWriter used to format CSV (cf. csv.NewWriter)
func DefaultCSVWriter(out io.Writer) *SafeCSVWriter { _ = "STUB: not implemented"; return nil }

// As only one rune can be defined as a CSV separator, we are going to trim
// the custom tag separator and use the first rune.

// SetCSVWriter sets the SafeCSVWriter used to format CSV.
func SetCSVWriter(csvWriter func(io.Writer) *SafeCSVWriter) { _ = "STUB: not implemented"; return }

func getCSVWriter(out io.Writer) *SafeCSVWriter { _ = "STUB: not implemented"; return nil }

// --------------------------------------------------------------------------
// CSVReader used to parse CSV

var selfCSVReader = DefaultCSVReader

// DefaultCSVReader is the default CSV reader used to parse CSV (cf. csv.NewReader)
func DefaultCSVReader(in io.Reader) CSVReader {
	_ = "STUB: not implemented"
	return *

	// LazyCSVReader returns a lazy CSV reader, with LazyQuotes and TrimLeadingSpace.
	new(CSVReader)
}

func LazyCSVReader(in io.Reader) CSVReader { _ = "STUB: not implemented"; return *new(CSVReader) }

// SetCSVReader sets the CSV reader used to parse CSV.
func SetCSVReader(csvReader func(io.Reader) CSVReader) { _ = "STUB: not implemented"; return }

func getCSVReader(in io.Reader) CSVReader {
	_ = "STUB: not implemented"
	return *

	// --------------------------------------------------------------------------
	// Marshal functions
	new(CSVReader)
}

// MarshalFile saves the interface as CSV in the file.
func MarshalFile(in interface{}, file *os.File) (err error) { _ = "STUB: not implemented"; return nil }

// MarshalString returns the CSV string from the interface.
func MarshalString(in interface{}) (out string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MarshalStringWithoutHeaders returns the CSV string from the interface.
func MarshalStringWithoutHeaders(in interface{}) (out string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MarshalBytes returns the CSV bytes from the interface.
func MarshalBytes(in interface{}) (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marshal returns the CSV in writer from the interface.
func Marshal(in interface{}, out io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

// MarshalWithoutHeaders returns the CSV in writer from the interface.
func MarshalWithoutHeaders(in interface{}, out io.Writer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// MarshalChan returns the CSV read from the channel.
func MarshalChan(c <-chan interface{}, out CSVWriter) error { _ = "STUB: not implemented"; return nil }

// MarshalChanWithoutHeaders returns the CSV read from the channel.
func MarshalChanWithoutHeaders(c <-chan interface{}, out CSVWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalCSV returns the CSV in writer from the interface.
func MarshalCSV(in interface{}, out CSVWriter) (err error) { _ = "STUB: not implemented"; return nil }

// MarshalCSVWithoutHeaders returns the CSV in writer from the interface.
func MarshalCSVWithoutHeaders(in interface{}, out CSVWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// --------------------------------------------------------------------------
// Unmarshal functions

// UnmarshalFile parses the CSV from the file in the interface.
func UnmarshalFile(in *os.File, out interface{}) error { _ = "STUB: not implemented"; return nil }

// UnmarshalMultipartFile parses the CSV from the multipart file in the interface.
func UnmarshalMultipartFile(in *multipart.File, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalFileWithErrorHandler parses the CSV from the file in the interface.
func UnmarshalFileWithErrorHandler(in *os.File, errHandler ErrorHandler, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalString parses the CSV from the string in the interface.
func UnmarshalString(in string, out interface{}) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBytes parses the CSV from the bytes in the interface.
func UnmarshalBytes(in []byte, out interface{}) error { _ = "STUB: not implemented"; return nil }

// Unmarshal parses the CSV from the reader in the interface.
func Unmarshal(in io.Reader, out interface{}) error { _ = "STUB: not implemented"; return nil }

// Unmarshal parses the CSV from the reader in the interface.
func UnmarshalWithErrorHandler(in io.Reader, errHandle ErrorHandler, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalWithoutHeaders parses the CSV from the reader in the interface.
func UnmarshalWithoutHeaders(in io.Reader, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalCSVWithoutHeaders parses a headerless CSV with passed in CSV reader
func UnmarshalCSVWithoutHeaders(in CSVReader, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalDecoder parses the CSV from the decoder in the interface
func UnmarshalDecoder(in Decoder, out interface{}) error { _ = "STUB: not implemented"; return nil }

// UnmarshalCSV parses the CSV from the reader in the interface.
func UnmarshalCSV(in CSVReader, out interface{}) error { _ = "STUB: not implemented"; return nil }

// UnmarshalCSVToMap parses a CSV of 2 columns into a map.
func UnmarshalCSVToMap(in CSVReader, out interface{}) error { _ = "STUB: not implemented"; return nil }

// UnmarshalToChan parses the CSV from the reader and send each value in the chan c.
// The channel must have a concrete type.
func UnmarshalToChan(in io.Reader, c interface{}) error { _ = "STUB: not implemented"; return nil }

// UnmarshalToChanWithErrorHandler parses the CSV from the reader in the interface.
func UnmarshalToChanWithErrorHandler(in io.Reader, errorHandler ErrorHandler, c interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalToChanWithoutHeaders parses the CSV from the reader and send each value in the chan c.
// The channel must have a concrete type.
func UnmarshalToChanWithoutHeaders(in io.Reader, c interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalDecoderToChan parses the CSV from the decoder and send each value in the chan c.
// The channel must have a concrete type.
func UnmarshalDecoderToChan(in SimpleDecoder, c interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalStringToChan parses the CSV from the string and send each value in the chan c.
// The channel must have a concrete type.
func UnmarshalStringToChan(in string, c interface{}) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBytesToChan parses the CSV from the bytes and send each value in the chan c.
// The channel must have a concrete type.
func UnmarshalBytesToChan(in []byte, c interface{}) error { _ = "STUB: not implemented"; return nil }

// UnmarshalToCallback parses the CSV from the reader and send each value to the given func f.
// The func must look like func(Struct).
func UnmarshalToCallback(in io.Reader, f interface{}) error { _ = "STUB: not implemented"; return nil }

// if last returned value from Call() is an error, return it

// UnmarshalDecoderToCallback parses the CSV from the decoder and send each value to the given func f.
// The func must look like func(Struct).
func UnmarshalDecoderToCallback(in SimpleDecoder, f interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalBytesToCallback parses the CSV from the bytes and send each value to the given func f.
// The func must look like func(Struct).
func UnmarshalBytesToCallback(in []byte, f interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalStringToCallback parses the CSV from the string and send each value to the given func f.
// The func must look like func(Struct).
func UnmarshalStringToCallback(in string, c interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalToCallbackWithError parses the CSV from the reader and
// send each value to the given func f.
//
// If func returns error, it will stop processing, drain the
// parser and propagate the error to caller.
//
// The func must look like func(Struct) error.
func UnmarshalToCallbackWithError(in io.Reader, f interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// callback f has already returned an error, stop processing but keep draining the chan c

// If the callback f returns an error, stores it and returns it in future.

// UnmarshalBytesToCallbackWithError parses the CSV from the bytes and
// send each value to the given func f.
//
// If func returns error, it will stop processing, drain the
// parser and propagate the error to caller.
//
// The func must look like func(Struct) error.
func UnmarshalBytesToCallbackWithError(in []byte, f interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalStringToCallbackWithError parses the CSV from the string and
// send each value to the given func f.
//
// If func returns error, it will stop processing, drain the
// parser and propagate the error to caller.
//
// The func must look like func(Struct) error.
func UnmarshalStringToCallbackWithError(in string, c interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// CSVToMap creates a simple map from a CSV of 2 columns.
func CSVToMap(in io.Reader) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

// CSVToMaps takes a reader and returns an array of dictionaries, using the header row as the keys
func CSVToMaps(reader io.Reader) ([]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CSVToChanMaps parses the CSV from the reader and send a dictionary in the chan c, using the header row as the keys.
func CSVToChanMaps(reader io.Reader, c chan<- map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}
