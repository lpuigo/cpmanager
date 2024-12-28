package persist

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Record struct {
	id       string
	dirty    bool
	marshall func(w io.Writer) error
	idToName func(string) string
	nameToId func(string) (string, error)
}

// NewRecord create a new record with given marshalling function
func NewRecord(marshall func(w io.Writer) error) *Record {
	const defaultExt string = ".json"
	return &Record{
		marshall: marshall,
		idToName: func(id string) string {
			return id + defaultExt
		},
		nameToId: func(name string) (string, error) {
			return strings.TrimSuffix(filepath.Base(name), defaultExt), nil
		},
	}
}

// GetId returns the inner record id
func (r Record) GetId() string {
	return r.id
}

// SetId sets the inner record id
func (r *Record) SetId(id string) {
	r.id = id
}

// Dirty marks the record as dirty (need to be persisted in a file)
func (r *Record) Dirty() {
	r.dirty = true
}

// Persist writes receiver to its named file within the given path
func (r *Record) Persist(path string) error {
	file := r.GetFilePath(path)
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		return err
	}
	err = r.marshall(f)
	if err != nil {
		return fmt.Errorf("error marshalling: %v", err)
	}
	r.dirty = false
	return nil
}

// Remove moves the receiver record to deleted directory. This record won't be loaded anymore with persister init / reload
func (r Record) Remove(path string) error {
	dpath := filepath.Join(path, "deleted")
	file := r.GetFilePath(path)
	dfile := r.GetFilePath(dpath)
	return os.Rename(file, dfile)
}

// GetFilePath returns receiver's full file path name by appening given path and receiver file name
func (r Record) GetFilePath(path string) string {
	return filepath.Join(path, r.GetFileName())
}

// GetFileName returns receiver file name (zero padded id with json extension)
func (r Record) GetFileName() string {
	return r.idToName(r.GetId())
}

// GetFileName returns receiver file name (zero padded id with json extension)
func (r Record) GetIdFromFileName(file string) (string, error) {
	return r.nameToId(file)
}

// Marshall writes marshalled receiver to given writer
func (r Record) Marshall(w io.Writer) error {
	return r.marshall(w)
}

// SetIdFromFile sets receiver's id based on given file name (must be zero padded decimal digit)
func (r *Record) IdFromFile(file string) (string, error) {
	return r.nameToId(filepath.Base(file))
}
