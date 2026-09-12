package slideimage

import (
	"bytes"
	"image"
	"image/png"
	"math"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/labelmsg"
	"github.com/gomodule/redigo/redis"

	"code.gitea.io/gitea/modules/public"

	"code.gitea.io/gitea/modules/log"

	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/modules/redis/redis_client"

	"gitea.com/macaron/macaron"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

type SlideImage struct {
	SubURL            string
	URLPrefix         string
	SampleImages      int
	StdWidth          int
	StdHeight         int
	MaskSize          int
	ImageY            int
	ImageXStart       int
	MinImageX         int
	Expiration        int
	Tolerance         int
	CachePrefix       string
	CacheManualPrefix string
}

type Options struct {
	// Suburl path. Default is empty.
	SubURL string
	// URL prefix of getting captcha pictures. Default is "/slideimage/".
	URLPrefix string
	//Default is 4
	SampleImages int
	//Image width default 391
	StdWidth int
	//Image Height default 196
	StdHeight int
	// default 51
	MaskSize int
	// default 125
	ImageY int
	//default 0
	ImageXStart int
	//容忍的误差 default 2px
	Tolerance int
	// default 150
	MinImageX int
	// default 600 seconds
	Expiration int
	//default slide:
	CachePrefix string
	//default mslide: 验证通过，在缓存中记录已进行过人工操作，然后在发送验证码之前再进行校验是否进行了人工操作。
	CacheManualPrefix string
}

func NewSlideImage(opt Options) *SlideImage {
	return &SlideImage{
		SubURL:       opt.SubURL,
		URLPrefix:    opt.URLPrefix,
		SampleImages: opt.SampleImages,
		StdWidth:     opt.StdWidth,
		StdHeight:    opt.StdHeight,
		MaskSize:     opt.MaskSize,
		ImageY:       opt.ImageY,
		ImageXStart:  opt.ImageXStart,
		Tolerance:    opt.Tolerance,
		MinImageX:    opt.MinImageX,
		Expiration:   opt.Expiration,
		CachePrefix:  opt.CachePrefix,
	}
}

func prepareOptions(options []Options) Options {
	var opt Options
	if len(options) > 0 {
		opt = options[0]
	}

	opt.SubURL = strings.TrimSuffix(opt.SubURL, "/")

	// Defaults.
	if len(opt.URLPrefix) == 0 {
		opt.URLPrefix = "/slideimage/"
	} else if opt.URLPrefix[len(opt.URLPrefix)-1] != '/' {
		opt.URLPrefix += "/"
	}
	if opt.SampleImages == 0 {
		opt.SampleImages = 4
	}
	if opt.StdWidth == 0 {
		opt.StdWidth = 391
	}
	if opt.StdHeight == 0 {
		opt.StdHeight = 196
	}
	if opt.MaskSize == 0 {
		opt.MaskSize = 51
	}
	if opt.ImageY == 0 {
		opt.ImageY = 75
	}
	if opt.ImageXStart == 0 {
		opt.ImageXStart = 2
	}

	if opt.Tolerance == 0 {
		opt.Tolerance = 2
	}

	if opt.MinImageX == 0 {
		opt.MinImageX = 150
	}
	if opt.Expiration == 0 {
		opt.Expiration = 600
	}

	if len(opt.CachePrefix) == 0 {
		opt.CachePrefix = "slide:"
	}
	if len(opt.CacheManualPrefix) == 0 {
		opt.CacheManualPrefix = "mslide:"
	}

	return opt
}

func (s *SlideImage) key(id string) string {
	return s.CachePrefix + id
}
func (s *SlideImage) mkey(id string) string {
	return s.CacheManualPrefix + id
}

func (s *SlideImage) VerifyManual(id string) (bool, error) {
	redisConn := labelmsg.Get()
	defer redisConn.Close()
	return redis_client.EXISTS(redisConn, s.mkey(id))
}

func (s *SlideImage) Verify(id string, x int) bool {
	redisConn := labelmsg.Get()
	defer redisConn.Close()

	v, err := redis_client.GET(redisConn, s.key(id))
	v1, err := redis.String(v, err)
	if err != nil {
		log.Warn("redis err", err)
		return false
	}
	if v1 == "" {
		return false
	}
	values := strings.Split(v1, "-")
	imageRandX, _ := strconv.Atoi(values[1])
	if int(math.Abs(float64(imageRandX-x))) <= s.Tolerance {
		redis_client.SETNX(redisConn, s.mkey(id), "1", s.Expiration)
		return true
	}
	return false

}

func (s *SlideImage) CreateCode() (string, int, int) {
	nums := rand.Intn(s.SampleImages)
	imageId := uuid.New().String()

	//获取随机x坐标
	imageRandX := rand.Intn(s.StdWidth - s.MaskSize - 3)
	if imageRandX < s.MinImageX {
		imageRandX += s.MinImageX
	}
	redis_client.Setex(s.key(imageId), strconv.Itoa(nums)+"-"+strconv.Itoa(imageRandX), time.Second*time.Duration(s.Expiration))
	return imageId, nums, imageRandX
}

func SlideImager(options ...Options) macaron.Handler {
	return func(ctx *macaron.Context) {
		slideImage := NewSlideImage(prepareOptions(options))

		if strings.HasPrefix(ctx.Req.URL.Path, slideImage.URLPrefix) {
			id := path.Base(ctx.Req.URL.Path)
			if i := strings.Index(id, "."); i > -1 {
				id = id[:i]
			}
			isScreenshot := strings.HasSuffix(id, "screenshot")
			if isScreenshot {
				id = strings.TrimSuffix(id, "screenshot")
			}

			key := slideImage.key(id)
			v, err := redis_client.Get(key)

			if err != nil || v == "" {
				ctx.Status(404)
				ctx.Write([]byte("not found"))
				//png.Encode(ctx.Resp, *setting.SlideImagesBg[0])
				return
			}

			values := strings.Split(v, "-")

			imageIndex, _ := strconv.Atoi(values[0])
			imageRandX, _ := strconv.Atoi(values[1])
			imageBg := setting.SlideImagesBg[imageIndex]

			maxPotion := image.Point{
				X: imageRandX + slideImage.MaskSize,
				Y: slideImage.ImageY + slideImage.MaskSize,
			}
			minPotion := image.Point{
				X: imageRandX,
				Y: slideImage.ImageY,
			}
			subimg := image.Rectangle{
				Max: maxPotion,
				Min: minPotion,
			}

			if isScreenshot {
				data := imaging.Crop(*imageBg, subimg)
				png.Encode(ctx.Resp, data)

			} else {
				data := imaging.Overlay(*imageBg, *setting.SlideMaskImage, minPotion, 1.0)
				png.Encode(ctx.Resp, data)
			}
			ctx.Status(200)
			return

		}
		ctx.Data["SlideImageInfo"] = slideImage
		ctx.Data["EnablePhone"] = setting.PhoneService.Enabled
		ctx.Map(slideImage)

	}
}

func InitSlideImage() {
		filenames, err := public.Dir(path.Join("img", "slide", "bg"))
		if err != nil {
			panic("Slide Image Service init failed")
		}
		maskFileName, err := public.Dir(path.Join("img", "slide", "mask"))
		if err != nil {
			panic("Slide Image Service init failed")
		}

		for _, filename := range filenames {
			if strings.HasSuffix(filename, ".png") {
				content, err := public.Asset(path.Join("img", "slide", "bg", filename))

				if err != nil {
					log.Warn("can not open "+filename, err)
					continue
				}

				m, err := png.Decode(bytes.NewReader(content))

				if err != nil {
					log.Warn("can not decode "+filename, err)
					continue
				}
				setting.SlideImagesBg = append(setting.SlideImagesBg, &m)
			}

		}
		setting.SlideImagesCount = len(setting.SlideImagesBg)
		if setting.SlideImagesCount == 0 {
			panic("Slide Image Service init failed")
		}

		maskContent, err := public.Asset(path.Join("img", "slide", "mask", maskFileName[0]))

		if err != nil {
			panic("Slide Image Service init failed")
		}

		MaskImage, err := png.Decode(bytes.NewReader(maskContent))
		if err != nil {
			panic("Slide Image Service init failed")
		}
		setting.SlideMaskImage = &MaskImage

		log.Info("Slide Image Service Enabled")

}
