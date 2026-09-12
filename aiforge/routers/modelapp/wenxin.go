package modelapp

import (
	"fmt"

	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelappservice"
	"code.gitea.io/gitea/routers/ai_task"
	uuid "github.com/satori/go.uuid"
)

var modelMainTpl base.TplName = "model/index"
var modelWenXinTpl base.TplName = "model/wenxin/index"
var modelLLMChatTpl base.TplName = "model/llmchat/index"
var modelModelExperienceTpl base.TplName = "model/experience/index"
var modelSdExperienceTpl base.TplName = "model/sd/index"
var modelTtsExperienceTpl base.TplName = "model/tts/index"

const WAIT_TIME int = 7

func ModelMainPage(ctx *context.Context) {
	ctx.HTML(200, modelMainTpl)
}

func WenXinPage(ctx *context.Context) {
	ctx.HTML(200, modelWenXinTpl)
}

func ModelExperienceCreate(ctx *context.Context) {
	ctx.HTML(200, modelModelExperienceTpl)
	return
}

func ModelExperienceSd(ctx *context.Context) {
	ctx.HTML(200, modelSdExperienceTpl)
	return
}

func LLMChatPage(ctx *context.Context) {
	ctx.HTML(200, modelLLMChatTpl)
	return
}
func TtsPage(ctx *context.Context) {
	ctx.HTML(200, modelTtsExperienceTpl)
	return
}
func SdPaintNew(ctx *context.Context) {
	prompt := ctx.Query("prompt")
	negative_prompt := ctx.Query("negative_prompt")
	num_images_per_prompt := ctx.QueryInt("num_images_per_prompt")
	cloudbrainId := ctx.QueryInt64("task_id")

	seed := ctx.QueryInt("seed")
	steps := ctx.QueryInt("steps")
	height := ctx.QueryInt("height")
	width := ctx.QueryInt("width")
	scheduler_name := ctx.Query("scheduler_name")
	guidance_scale := ctx.QueryInt("guidance_scale")
	uuid := uuid.NewV4()
	id := uuid.String()
	var tokenUrl string

	task, err := models.GetCloudbrainByCloudbrainID(cloudbrainId)
	if err == nil {
		if task.Status == string(models.JobRunning) {
			tokenUrl, _ = ai_task.GetSelfEndPointUrlById(cloudbrainId)
		} else if task.Status == string(models.JobStopped) || task.Status == string(models.ModelArtsStopping) {
			tokenUrl = string(models.JobStopped)
		}
	}
	index := strings.Index(tokenUrl, "?token=")
	if index == -1 {
		log.Info("not found token=")
	} else {
		token := tokenUrl[index:]
		tokenUrl = tokenUrl[0:index] + "/text2image" + token
	}
	log.Info("sd Url=" + tokenUrl)
	modelApp := &models.ModelApp{
		ID:                    id,
		Desc:                  prompt,
		TaskId:                int(cloudbrainId),
		Negative_prompt:       negative_prompt,
		Num_images_per_prompt: num_images_per_prompt,
		Seed:                  seed,
		Steps:                 steps,
		Height:                height,
		Width:                 width,
		Scheduler_name:        scheduler_name,
		Guidance_scale:        guidance_scale,
		Count:                 1,
		Type:                  1,
		Url:                   tokenUrl,
		UserId:                ctx.User.ID,
	}
	models.SaveModelApp(modelApp)
	modelappservice.ProducerOrder(modelApp)
	ctx.JSON(200, map[string]string{
		"result": id,
		"wait":   fmt.Sprint(modelappservice.GetWaitTime()),
	})
}

func QueryAllModelExperience(ctx *context.Context) {
	//modeluuid := ctx.Query("modeluuid")
	//re, err := models.GetModelExperienceRunningAndWaitingTask(ctx.User.ID)
	re, err := models.GetModelExperienceNotFinalStatusTask(ctx.User.ID, []string{string(models.JobWaiting), string(models.JobRunning), string(models.LocalStatusPreparing), string(models.LocalStatusCreating)})
	if err != nil {
		log.Info("find result err.")
	} else {
		for _, v := range re {
			repo, err := models.GetRepositoryByID(v.RepoID)
			if err == nil {
				v.Repo = repo
			}
		}
	}
	ctx.JSON(200, re)
}

func WenXinPaintNew(ctx *context.Context) {
	textDesc := ctx.Query("textDesc")
	uuid := uuid.NewV4()
	id := uuid.String()

	count := models.QueryModelAppCount(ctx.User.ID)
	if count >= 200 {
		ctx.JSON(200, map[string]string{
			"result": "-1",
			"msg":    "Max free times exceed 200.",
		})
		return
	}
	modelApp := &models.ModelApp{
		ID:     id,
		Desc:   textDesc,
		Count:  1,
		Type:   0,
		UserId: ctx.User.ID,
	}
	models.SaveModelApp(modelApp)
	modelappservice.ProducerOrder(modelApp)
	ctx.JSON(200, map[string]string{
		"result": id,
		"wait":   fmt.Sprint(modelappservice.GetWaitTime()),
	})
}

func GetOnlineInferDebugUrl(cloudbrainId int64) string {

	url, err := ai_task.GetNoteBookUrlById(cloudbrainId)
	if err == nil {
		return url
	}

	return ""
}

func QueryWenXinPaintResult(ctx *context.Context) {
	//ctx.HTML(200, modelMainTpl)
	count := models.QueryModelAppCount(ctx.User.ID)
	ctx.JSON(200, count)
}

func QueryWenXinPaintById(ctx *context.Context) {
	result := models.QueryModelAppById(ctx.Query("id"))
	ctx.JSON(200, result)
}
