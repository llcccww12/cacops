package common

import (
	"archive/zip"
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"io"
	"net/http"
	"net/url"
)

func WriteDownloadContent2Resp(ctx *context.Context, res *entity.FileDownloadInfo) error {
	defer func() {
		res.Close()
	}()
	resp := ctx.Resp

	//优先重定向到下载链接
	if res.DownloadUrl != "" {
		ctx.Resp.Header().Set("Cache-Control", "max-age=0")
		http.Redirect(ctx.Resp, ctx.Req.Request, res.DownloadUrl, http.StatusTemporaryRedirect)
		return nil
	}

	//没有下载链接则直接返回文件流
	resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(res.ResultFileName))
	resp.Header().Set("Content-Type", "application/octet-stream")
	var reader io.Reader

	switch res.ResultType {
	case entity.FileTypeTXT:
		for _, f := range res.Readers {
			reader = f.Reader
			io.Copy(resp, reader)
		}
	case entity.FileTypeZIP:
		w := zip.NewWriter(resp)
		defer w.Close()
		for _, f := range res.Readers {
			fDest, err := w.Create(f.Name)
			if err != nil {
				log.Error("GetAITaskLog error.%v", err)
				return err
			}
			p := make([]byte, 1024)
			var readErr error
			var readCount int
			// 读取对象内容
			for {
				readCount, readErr = f.Reader.Read(p)
				if readCount > 0 {
					fDest.Write(p[:readCount])
				}
				if readErr != nil {
					break
				}
			}
		}
	}
	return nil
}
