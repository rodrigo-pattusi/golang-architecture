package fileWriter

import (
	"fmt"
	"io"
	"os"
)

type WriteFileError struct {
	Op  string
	err error
}

func (w WriteFileError) Error() string {
	return w.err.Error()
}

func (w WriteFileError) Unwrap() error {
	return w.err
}

type WriteFile struct {
	f   *os.File
	err error
}

func NewWriteFile(filename string) *WriteFile {
	f, err := os.Create(filename)
	if err != nil {
		return &WriteFile{
			f: nil,
			err: WriteFileError{
				Op:  "Create",
				err: fmt.Errorf("Error while creating a file: %w", err),
			},
		}
	}
	return &WriteFile{
		f: f,
		err: nil,
	}
}

func (w *WriteFile) WriteString(text string) {
	if w.err != nil {
		return
	}

	_, err := io.WriteString(w.f, text)
	if err != nil {
		w.err = WriteFileError{
			Op:  "WriteString",
			err: fmt.Errorf("Failed while writing a string: %w", err),
		}
	}
}

func (w *WriteFile) Close() {
	if w.f == nil {
		return
	}

	err := w.f.Close()
	if err != nil {
		w.err = WriteFileError{
			Op:  "Close",
			err: fmt.Errorf("Failed while closing file: %w", err),
		}
	}
}

// All errors returning from Err should be of type *WriteFileError
func (w *WriteFile) Err() error {
	return w.err
}
