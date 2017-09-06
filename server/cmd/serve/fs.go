package serve

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

type fs string

func createFS(prefix string) (f http.FileSystem, err error) {
	prefix = strings.TrimRight(prefix, string(os.PathSeparator))

	var file *os.File
	if file, err = os.Open(prefix); err != nil {
		return
	}

	defer file.Close()

	var stat os.FileInfo
	if stat, err = file.Stat(); err != nil {
		return
	}

	if !stat.IsDir() {
		err = errors.New(fmt.Sprintf("%s is not a directory", prefix))
		return
	}

	f = fs(prefix)

	return
}

func (f fs) Open(name string) (http.File, error) {
	return os.Open(path.Join(string(f), name))
}