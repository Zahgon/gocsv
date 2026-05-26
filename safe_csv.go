package gocsv

//Wraps around SafeCSVWriter and makes it thread safe.
import (
	"encoding/csv"
	"sync"
)

type CSVWriter interface {
	Write(row []string) error
	Flush()
	Error() error
}

type SafeCSVWriter struct {
	*csv.Writer
	m sync.Mutex
}

func NewSafeCSVWriter(original *csv.Writer) *SafeCSVWriter { _ = "STUB: not implemented"; return nil }

// Override write
func (w *SafeCSVWriter) Write(row []string) error { _ = "STUB: not implemented"; return nil }

// Override flush
func (w *SafeCSVWriter) Flush() { _ = "STUB: not implemented"; return }
