package modelapp

import (
	"bytes"
	"code.gitea.io/gitea/models"
	"crypto/tls"
	"image"
	"image/png"
	"net/http"
	"strconv"

	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/modules/base"

	"code.gitea.io/gitea/modules/context"
	"github.com/go-resty/resty/v2"
)

var restyClient *resty.Client
var tplExploreUpload base.TplName = "model/tuomin/upload"
var uploadUrl = "/extension/tuomin/upload"
var allowedContentType = []string{"image/jpeg", "image/jpg", "image/png"}

func ProcessImageUI(ctx *context.Context) {
	ctx.HTML(200, tplExploreUpload)
}

func ProcessImage(ctx *context.Context) {

	file, header, err := ctx.GetFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest,models.BaseErrorMessage(ctx.Tr("model_app.get_file_fail")))
		return
	}
	defer file.Close()

	contentType := header.Header.Get("Content-Type")

	if !isInAllowedContentType(contentType) {
		ctx.JSON(http.StatusBadRequest,models.BaseErrorMessage(ctx.Tr("model_app.content_type_unsupported")))
		return
	}

	client := getRestyClient()
	res, err := client.R().SetMultipartField(
		"file", header.Filename, contentType, file).Post(setting.ModelApp.DesensitizationUrl + "?mode=" + strconv.Itoa(ctx.QueryInt("mode")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,models.BaseErrorMessage(ctx.Tr("model_app.process_image_fail")))
		return
	}
	image, _, err := image.Decode(bytes.NewReader(res.Body()))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,models.BaseErrorMessage(ctx.Tr("model_app.process_image_fail")))
		return
	}

	png.Encode(ctx.Resp, image)
	return

}

func isInAllowedContentType(contentType string) bool {
	for _, allowType := range allowedContentType {
		if allowType == contentType {
			return true
		}
	}
	return false
}

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}
