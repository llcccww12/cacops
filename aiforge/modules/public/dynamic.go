// +build !bindata

// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package public

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"code.gitea.io/gitea/modules/setting"
	"gitea.com/macaron/macaron"
	"github.com/unknwon/com"
)

// Static implements the macaron static handler for serving assets.
func Static(opts *Options) macaron.Handler {
	return opts.staticHandler(opts.Directory)
}

func Dir(name string) ([]string, error) {

	var (
		result []string
	)

	staticDir := path.Join(setting.StaticRootPath, "public", name)

	if com.IsDir(staticDir) {
		files, err := com.StatDir(staticDir, true)

		if err != nil {
			return []string{}, fmt.Errorf("Failed to read img directory. %v", err)
		}

		result = append(result, files...)
	}

	return result, nil
}

func Asset(name string) ([]byte, error) {

	staticPath := path.Join(setting.StaticRootPath, "public", name)

	if com.IsFile(staticPath) {
		f, err := os.Open(staticPath)
		defer f.Close()

		if err == nil {
			return ioutil.ReadAll(f)
		}
	}
	return nil, fmt.Errorf("Asset file does not exist: %s", name)
}
