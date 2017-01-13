package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath       string
	zipReadCloser *zip.ReadCloser
	zipError      error
}

func (entry *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	if entry.zipError != nil {
		return nil, nil, entry.zipError
	}
	for _, f := range entry.zipReadCloser.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			data, err := ioutil.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, nil, err
			}
			return data, entry, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (entry *ZipEntry) String() string {
	return entry.absPath
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	r, err := zip.OpenReader(absPath)
	return &ZipEntry{
		absPath:       absPath,
		zipReadCloser: r,
		zipError:      err,
	}
}

func (entry *ZipEntry) Close() {
	if entry.zipReadCloser != nil {
		entry.zipReadCloser.Close()
	}
}
