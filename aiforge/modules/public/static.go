// +build bindata

// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package public

import (
	"fmt"
	"io/ioutil"

	"gitea.com/macaron/macaron"
)

// Static implements the macaron static handler for serving assets.
func Static(opts *Options) macaron.Handler {
	opts.FileSystem = Assets
	// we don't need to pass the directory, because the directory var is only
	// used when in the options there is no FileSystem.
	return opts.staticHandler("")
}

func Dir(name string) ([]string, error) {

	files, err := AssetDir(name)
	if err != nil {
		return []string{}, fmt.Errorf("Failed to read embedded directory. %v", err)
	}

	return files, nil
}

func Asset(name string) ([]byte, error) {
	f, err := Assets.Open("/" + name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func AssetDir(dirName string) ([]string, error) {
	d, err := Assets.Open(dirName)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var results = make([]string, 0, len(files))
	for _, file := range files {
		results = append(results, file.Name())
	}
	return results, nil
}

func AssetNames() []string {
	realFS := Assets.(vfsgen۰FS)
	var results = make([]string, 0, len(realFS))
	for k := range realFS {
		results = append(results, k[1:])
	}
	return results
}

func AssetIsDir(name string) (bool, error) {
	if f, err := Assets.Open("/" + name); err != nil {
		return false, err
	} else {
		defer f.Close()
		if fi, err := f.Stat(); err != nil {
			return false, err
		} else {
			return fi.IsDir(), nil
		}
	}
}
