package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path) - 1]
	compositeEntry := []Entry{}

	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(strings.ToLower(path), ".jar") || strings.HasSuffix(strings.ToLower(path), ".zip") {
			compositeEntry = append(compositeEntry, newZipEntry(path))
		}
		return nil
	})

	return compositeEntry
}
