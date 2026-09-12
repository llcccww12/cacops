package llm_chat

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/ai_task"
	"github.com/go-resty/resty/v2"
)

const (
	GetSpeakerVoiceRoute = "/speaker/get_sample"
	SynthesizeVoiceRoute = "/tts/synthesize"
)

func GetUrlByCloudbrainId(cloudbrainId int64) string {
	tokenUrl, _ := ai_task.GetSelfEndPointUrlById(cloudbrainId)
	index := strings.Index(tokenUrl, "?token=")
	if index != -1 {
		tokenUrl = tokenUrl[:index]
	}
	log.Info("[llm_chat] Url=" + tokenUrl)

	return tokenUrl
}

func HandleRestyStatusCode(ctx *context.APIContext, resp *resty.Response) string {
	var errMsg string
	log.Info("[llm_chat] Method: %s; resp status: %s; Url: %s\n", resp.Request.Method, resp.Status(), resp.Request.URL)
	switch resp.StatusCode() {
	case 502:
		errMsg = ctx.Tr("llm_chat.loading")
	case 404:
		errMsg = ctx.Tr("llm_chat.chat_expired")
	case 200:
		contentType := resp.Header().Get("Content-Type")
		log.Info("[llm_chat] Content-Type: %s", contentType)
		if !strings.HasPrefix(contentType, "application/json") {
			return ""
		}
	default:
		errMsg = ctx.Tr("llm_chat.server_error")
		log.Error("[llm_chat] Error response RawBody(): %d", resp.RawBody())
	}
	return errMsg
}

func setResponseHeaders(ctx *context.APIContext, filename string) {
	ctx.Resp.Header().Set("Content-Type", "audio/wav")
	ctx.Resp.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
}

func GetSpeakerVoice(ctx *context.APIContext) {
	speakerId := ctx.Query("speaker_id")
	cloudbrainId := ctx.QueryInt64("task_id")

	url := GetUrlByCloudbrainId(cloudbrainId) + GetSpeakerVoiceRoute
	params := map[string]string{"speaker_id": speakerId}

	resp, err := getSpeakerVoiceResty(ctx, url, params)
	if err != nil {
		return
	}
	defer resp.RawBody().Close()

	setResponseHeaders(ctx, speakerId+".wav")
	if _, err := io.Copy(ctx.Resp, resp.RawBody()); err != nil {
		log.Error("[llm_chat] Error copying WAV data to response: %v", err)
		ctx.JSON(503, map[string]string{"error": ctx.Tr("llm_chat.server_error")})
	}
}

func getSpeakerVoiceResty(ctx *context.APIContext, url string, params map[string]string) (*resty.Response, error) {
	client := getRestyClient()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(params).
		SetDoNotParseResponse(true).
		Get(url)

	if err != nil {
		log.Error("[llm_chat] Error sending request to URL: %v", err)
		ctx.JSON(503, map[string]string{"error": ctx.Tr("llm_chat.chat_expired")})
		return nil, err
	}

	errMsg := HandleRestyStatusCode(ctx, resp)
	if errMsg != "" {
		ctx.JSON(503, map[string]string{"error": errMsg})
		return nil, fmt.Errorf(errMsg)
	}

	return resp, nil
}

type TTSConfig struct {
	Text      string `form:"text" binding:"Required"`
	Lang      string `form:"lang" binding:"Required"`
	SpeakerID string `form:"speaker_id" binding:"Required"`
	TaskID    int64  `form:"task_id" binding:"Required"`
}

func GetUploadedWavFile(ctx *context.APIContext) (io.Reader, string, error) {
	file, header, err := ctx.Req.FormFile("wav_data")
	if err != nil && err != http.ErrMissingFile {
		return nil, "", fmt.Errorf("error processing file upload")
	}

	var fileReader io.Reader
	var filename string
	if file != nil {
		defer file.Close()
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return nil, "", fmt.Errorf("error reading file")
		}
		fileReader = bytes.NewReader(fileBytes)
		if header != nil {
			filename = header.Filename
		} else {
			filename = "output.wav"
		}
	}

	return fileReader, filename, nil
}

func SynthesizeVoice(ctx *context.APIContext, config TTSConfig) {
	fileReader, filename, err := GetUploadedWavFile(ctx)
	if err != nil {
		ctx.JSON(503, map[string]string{"error": err.Error()})
		return
	}

	url := GetUrlByCloudbrainId(config.TaskID) + SynthesizeVoiceRoute
	form_data := map[string]string{
		"text":       config.Text,
		"lang":       config.Lang,
		"speaker_id": config.SpeakerID,
	}

	resp, err := synthesizeVoiceResty(ctx, url, form_data, fileReader, filename)
	if err != nil {
		return
	}
	defer resp.RawBody().Close()

	setResponseHeaders(ctx, "output.wav")
	if _, err := io.Copy(ctx.Resp, resp.RawBody()); err != nil {
		log.Error("[llm_chat] Error copying WAV data to response: %v", err)
		ctx.JSON(503, map[string]string{"error": ctx.Tr("llm_chat.server_error")})
	}
}

func synthesizeVoiceResty(ctx *context.APIContext, url string, formData map[string]string, fileReader io.Reader, filename string) (*resty.Response, error) {
	client := getRestyClient()
	req := client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFormData(formData).
		SetDoNotParseResponse(true)

	if fileReader != nil {
		req.SetFileReader("wav_data", filename, fileReader)
	}

	resp, err := req.Post(url)
	if err != nil {
		log.Error("[llm_chat] Error sending request to URL: %v", err)
		ctx.JSON(503, map[string]string{"error": ctx.Tr("llm_chat.chat_expired")})
		return nil, err
	}

	errMsg := HandleRestyStatusCode(ctx, resp)
	if errMsg != "" {
		ctx.JSON(503, map[string]string{"error": errMsg})
		return nil, fmt.Errorf(errMsg)
	}

	return resp, nil
}
