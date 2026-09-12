package modelappservice

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/setting"
)

var wenxinChannel = make(chan *models.ModelApp, 10000)
var isCD bool

func Init() {
	urls := setting.BaiduWenXin.ModelArtsWenXinURL
	if strings.Index(urls, "cdzs.cn") > 0 {
		isCD = true
	} else {
		isCD = false
	}
	urlarray := strings.Split(urls, ",")
	urlNums := len(urlarray)
	log.Info("url nums=" + fmt.Sprint(urlNums))
	for i := 0; i < setting.BaiduWenXin.RUN_WORKERS; i++ {
		go consumerOrder(wenxinChannel, urlarray[i%urlNums])
	}
}

func ProducerOrder(modelApp *models.ModelApp) {
	wenxinChannel <- modelApp
}

func GetWaitTime() int {
	dvid := setting.BaiduWenXin.MODEL_SERVERS
	return ((len(wenxinChannel) / dvid) + 1) * 30
}

func wenxinConsumer(modelApp *models.ModelApp, url string, goroutine_id uint64) {
	if !modelarts.SendTextReview(modelApp.Desc) {
		modelApp.Status = -1
		modelApp.Picture = "文本内容可能包括敏感关键字，请重新输入。"
		models.UpdateModelApp(modelApp)
		return
	}
	log.Info("goroutine id=" + fmt.Sprint(goroutine_id) + " wenxin text=" + modelApp.Desc)
	var result *modelarts.WenXinResult
	var err error
	if isCD {
		result, err = modelarts.CreateWenXinJobToCD(modelApp, url)
	} else {
		result, err = modelarts.CreateWenXinJob(modelApp, url)
	}
	if err == nil {
		if !modelarts.SendPictureReivew(result.Result) {
			modelApp.Status = -1
			modelApp.Picture = "生成的图像内容不合法，请重新输入关键字信息。"
			models.UpdateModelApp(modelApp)
		} else {
			modelApp.Status = 0
			modelApp.Picture = result.Result
			models.UpdateModelApp(modelApp)
		}
	} else {
		log.Info("err=" + err.Error())
		modelApp.Status = -1
		modelApp.Picture = "哇哦，服务暂时开小差了，请尝试重新生成。"
		models.UpdateModelApp(modelApp)
	}
}

func sdConsumer(modelApp *models.ModelApp, goroutine_id uint64) {
	log.Info("goroutine id=" + fmt.Sprint(goroutine_id) + " sd text=" + modelApp.Desc)
	if modelApp.Url == string(models.JobStopped) || modelApp.Url == "" {
		modelApp.Status = -1
		modelApp.Picture = "对话已过期或者任务已经被删除，请重新创建在线体验任务。"
		models.UpdateModelApp(modelApp)
		return
	}

	result, err, http_code := CreateSDJob(modelApp, modelApp.Url)
	if err == nil {
		modelApp.Status = 0
		for i := 0; i < len(result.Image_base64_list); i++ {
			modelApp.Picture += result.Image_base64_list[i]
			if i < len(result.Image_base64_list)-1 {
				modelApp.Picture += ";"
			}
		}
		models.UpdateModelApp(modelApp)
	} else {
		if http_code == 502 || http_code == 404 {
			log.Info("err=" + err.Error())
			modelApp.Status = -1
			modelApp.Picture = "模型正在加载中，请稍后再试。"
			models.UpdateModelApp(modelApp)
		} else {
			log.Info("err=" + err.Error())
			modelApp.Status = -1
			modelApp.Picture = "哇哦，服务暂时开小差了，请尝试重新生成。"
			models.UpdateModelApp(modelApp)
		}
	}
}

func consumerOrder(in <-chan *models.ModelApp, url string) {
	goroutine_id := GetGid()
	log.Info("goroutine id=" + fmt.Sprint(goroutine_id) + "consumer order start...")
	for modelApp := range in {
		if modelApp.Type == 0 {
			wenxinConsumer(modelApp, url, goroutine_id)
		} else if modelApp.Type == 1 {
			sdConsumer(modelApp, goroutine_id)
		}
	}
	log.Info("consumer order end...")
}

func GetGid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return 0
		//panic(err)
	}
	return n
}
