package websrv

import (
	"errors"
	"io/fs"
	"net/http"
)

type router struct {
	fs    http.FileSystem
	index string
}

func (r *router) Open(name string) (http.File, error) {
	file, err := r.fs.Open(name)

	if err == nil {
		return file, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		file, err = r.fs.Open(r.index)
		return file, err
	}

	return nil, err
}
