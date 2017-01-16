package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func (entry CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, e := range entry {
		data, from, err := e.ReadClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (entry CompositeEntry) String() string {
	strs := make([]string, len(entry))
	for _, e := range entry {
		strs = append(strs, e.String())
	}
	return strings.Join(strs, pathListSeparator)
}

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
