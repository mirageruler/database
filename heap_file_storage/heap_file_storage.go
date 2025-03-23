// Copyright 2025 macbook16
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package heapfilestorage

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"

	dbbase "github.com/mirageruler/database/page"
)

type HeapFile struct {
	file *os.File
}

// NewHeapFile creates or opens a heap file at the specified path
func NewHeapFile(path string) (*HeapFile, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &HeapFile{file: f}, nil
}

// ReadPage reads a dbbase.Page from disk at the given pageID
func (hf *HeapFile) ReadPage(pageID uint32) (*dbbase.Page, error) {
	offset := int64(pageID) * dbbase.PageSize
	buf := make([]byte, dbbase.PageSize)

	n, err := hf.file.ReadAt(buf, offset)
	if err != nil {
		return nil, err
	}
	if n != dbbase.PageSize {
		return nil, errors.New("unexpected EOF")
	}

	page := &dbbase.Page{}
	err = binary.Read(
		bytes.NewReader(buf), binary.LittleEndian, page,
	)
	if err != nil {
		return nil, err
	}

	return page, nil
}

// WritePage writes a dbbase.Page to disk at the given pageID
func (hf *HeapFile) WritePage(pageID uint32, page *dbbase.Page) error {
	offset := int64(pageID) * dbbase.PageSize
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, page)
	if err != nil {
		return err
	}

	n, err := hf.file.WriteAt(buf.Bytes(), offset)
	if err != nil {
		return err
	}
	if n != dbbase.PageSize {
		return os.ErrInvalid
	}
	return nil
}

// Close closes the heap file
func (hf *HeapFile) Close() error {
	return hf.file.Close()
}
