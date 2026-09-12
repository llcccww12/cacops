package badge

import (
	"bytes"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/setting"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	"image"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"os"
)

type IconUploader struct {
	Config IconUploadConfig
}

type IconUploadForm struct {
	Icon *multipart.FileHeader
}

type IconUploadConfig struct {
	FileMaxSize   int64
	FileMaxWidth  int
	FileMaxHeight int
	DefaultSize   uint
	NeedResize    bool
	NeedSquare    bool
}

func NewIconUploader(config IconUploadConfig) IconUploader {
	return IconUploader{Config: config}
}

func (u IconUploader) Upload(form IconUploadForm, user *models.User) (string, error) {
	if form.Icon == nil || form.Icon.Filename == "" {
		return "", errors.New("File or fileName is empty")
	}

	fr, err := form.Icon.Open()
	if err != nil {
		return "", fmt.Errorf("Icon.Open: %v", err)
	}
	defer fr.Close()

	if form.Icon.Size > u.Config.FileMaxSize {
		return "", errors.New("File is too large")
	}

	data, err := ioutil.ReadAll(fr)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	if !base.IsImageFile(data) {
		return "", errors.New("File is not a image")
	}
	iconName, err := u.uploadIcon(data, user.ID)
	if err != nil {
		return "", fmt.Errorf("uploadIcon: %v", err)
	}
	return iconName, nil

}

func (u IconUploader) uploadIcon(data []byte, userId int64) (string, error) {
	m, err := u.prepare(data)
	if err != nil {
		return "", err
	}

	iconName := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d-%x", userId, md5.Sum(data)))))

	if err := os.MkdirAll(setting.IconUploadPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("uploadIcon. Failed to create dir %s: %v", setting.AvatarUploadPath, err)
	}

	fw, err := os.Create(models.GetCustomIconByHash(iconName))
	if err != nil {
		return "", fmt.Errorf("Create: %v", err)
	}
	defer fw.Close()

	if err = png.Encode(fw, *m); err != nil {
		return "", fmt.Errorf("Encode: %v", err)
	}

	return iconName, nil
}

func (u IconUploader) prepare(data []byte) (*image.Image, error) {
	imgCfg, _, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("DecodeConfig: %v", err)
	}
	if imgCfg.Width > u.Config.FileMaxWidth {
		return nil, fmt.Errorf("Image width is too large: %d > %d", imgCfg.Width, setting.AvatarMaxWidth)
	}
	if imgCfg.Height > u.Config.FileMaxHeight {
		return nil, fmt.Errorf("Image height is too large: %d > %d", imgCfg.Height, setting.AvatarMaxHeight)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Decode: %v", err)
	}

	if u.Config.NeedSquare {
		if imgCfg.Width != imgCfg.Height {
			var newSize, ax, ay int
			if imgCfg.Width > imgCfg.Height {
				newSize = imgCfg.Height
				ax = (imgCfg.Width - imgCfg.Height) / 2
			} else {
				newSize = imgCfg.Width
				ay = (imgCfg.Height - imgCfg.Width) / 2
			}

			img, err = cutter.Crop(img, cutter.Config{
				Width:  newSize,
				Height: newSize,
				Anchor: image.Point{ax, ay},
			})
			if err != nil {
				return nil, err
			}
		}
	}

	if u.Config.NeedResize && u.Config.DefaultSize > 0 {
		img = resize.Resize(u.Config.DefaultSize, u.Config.DefaultSize, img, resize.NearestNeighbor)
	}

	return &img, nil
}
