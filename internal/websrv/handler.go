package websrv

import (
	"embed"
	"io/fs"
	"net/http"
)

func Handler(build embed.FS, dir, index string) (http.Handler, error) {
	fsys, err := fs.Sub(build, dir)
	if err != nil {
		return nil, err
	}

	if _, err = fsys.Open(index); err != nil {
		return nil, err
	}

	r := &router{fs: http.FS(fsys), index: index}

	hfs := http.FileServer(r)

	return hfs, nil
}
