package setting

import (
	"encoding/json"
)

var EnableAIAvatar bool
var AIAVartarApi string
var AIAVartarPromoteFile string
var AIAVartarSize int
var AIAVartarAK string
var AIAVartarSK string

type PromotConfig struct {
	Prompt          []string
	HumanPrompt     []string
	AnimalPrompt    []string `json:"animal"`
	LandscapePrompt []string `json:"landscape"`
	FemalePrompt    []string `json:"female"`
	MalePrompt      []string `json:"male"`
}

var AIAvartarPromoteConfig PromotConfig = PromotConfig{}

func GetDiffusionAvatarConfig() {
	sec := Cfg.Section("picture")
	EnableAIAvatar = sec.Key("ENABLE_AI_AVATAR").MustBool(false)
	AIAVartarApi = sec.Key("AI_AVATAR_API").MustString("")
	AIAVartarAK = sec.Key("AI_AVATAR_AK").MustString("")
	AIAVartarSK = sec.Key("AI_AVATAR_SK").MustString("")

	AIAVartarPromoteFile = sec.Key("PromptFile").MustString("model/prompt.json")
	AIAVartarSize = sec.Key("AI_AVATAR_SIZE").MustInt(256)
	url := RecommentRepoAddr + AIAVartarPromoteFile
	content, err := ReadFromPromote(url)
	if err == nil {

		json.Unmarshal(content, &AIAvartarPromoteConfig)

	}
	AIAvartarPromoteConfig.HumanPrompt = make([]string, 0)
	AIAvartarPromoteConfig.HumanPrompt = append(AIAvartarPromoteConfig.HumanPrompt, AIAvartarPromoteConfig.FemalePrompt...)
	AIAvartarPromoteConfig.HumanPrompt = append(AIAvartarPromoteConfig.HumanPrompt, AIAvartarPromoteConfig.MalePrompt...)
	AIAvartarPromoteConfig.Prompt = make([]string, 0)
	AIAvartarPromoteConfig.Prompt = append(AIAvartarPromoteConfig.Prompt, AIAvartarPromoteConfig.AnimalPrompt...)
	AIAvartarPromoteConfig.Prompt = append(AIAvartarPromoteConfig.Prompt, AIAvartarPromoteConfig.LandscapePrompt...)

}
