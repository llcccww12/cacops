package cloudbrain

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/setting"
)

func GetAiCenterShowByAiCenterId(aiCenterId string, ctx *context.Context) string {
	aiCenterAndNameInfo := getAiCenterAndNameInfo()
	if aiCenterAndNameInfo != nil {
		if info, ok := aiCenterAndNameInfo[aiCenterId]; ok {
			if ctx.Language() == "zh-CN" {
				return info.Content
			} else {
				return info.ContentEN
			}
		}
	}
	return aiCenterId
}

func getAiCenterAndNameInfo() map[string]*setting.C2NetSequenceInfo {
	if setting.AiCenterCodeAndNameAndLocMapInfo != nil {
		return setting.AiCenterCodeAndNameAndLocMapInfo
	}
	return setting.AiCenterCodeAndNameMapInfo
}

func GetAiCenterShow(aiCenter string, ctx *context.Context) string {
	aiCenterInfo := strings.Split(aiCenter, "+")
	aiCenterAndNameInfo := getAiCenterAndNameInfo()
	if len(aiCenterInfo) == 2 {
		if aiCenterAndNameInfo != nil {
			if info, ok := aiCenterAndNameInfo[aiCenterInfo[0]]; ok {
				if ctx.Language() == "zh-CN" {
					return info.Content
				} else {
					return info.ContentEN
				}
			} else {
				return aiCenterInfo[1]
			}

		} else {
			return aiCenterInfo[1]
		}
	} else {
		if aiCenterAndNameInfo != nil {
			if info, ok := aiCenterAndNameInfo[aiCenterInfo[0]]; ok {
				if ctx.Language() == "zh-CN" {
					return info.Content
				} else {
					return info.ContentEN
				}
			}
		}
	}
	return ""
}

func GetDisplayJobName(username string) string {
	t := time.Now()
	return jobNamePrefixValid(cutString(username, 5)) + t.Format("2006010215") + strconv.Itoa(int(t.Unix()))[5:]
}

func cutString(str string, lens int) string {
	if len(str) < lens {
		return str
	}
	return str[:lens]
}

func jobNamePrefixValid(s string) string {
	lowStr := strings.ToLower(s)
	re := regexp.MustCompile(`[^a-z0-9_\\-]+`)

	removeSpecial := re.ReplaceAllString(lowStr, "")

	re = regexp.MustCompile(`^[_\\-]+`)
	return re.ReplaceAllString(removeSpecial, "")
}

func GetAiCenterInfoByCenterCode(aiCenterCode string) *setting.C2NetSequenceInfo {
	aiCenterAndNameInfo := getAiCenterAndNameInfo()
	if aiCenterAndNameInfo != nil {
		if info, ok := aiCenterAndNameInfo[aiCenterCode]; ok {
			return info
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func GetAiCenterCode(aiCenter string) string {
	aiCenterInfo := strings.Split(aiCenter, "+")
	return aiCenterInfo[0]
}

func UpdateCloudbrainAiCenter(cloudbrain *models.CloudbrainInfo) *models.CloudbrainInfo {
	if cloudbrain.Cloudbrain.Type == models.TypeCloudBrainOne {
		cloudbrain.Cloudbrain.AiCenter = models.AICenterOfCloudBrainOne
		cloudbrain.Cloudbrain.Cluster = models.OpenICluster
	}
	if cloudbrain.Cloudbrain.Type == models.TypeCloudBrainTwo {
		cloudbrain.Cloudbrain.AiCenter = models.AICenterOfCloudBrainTwo
		cloudbrain.Cloudbrain.Cluster = models.OpenICluster
	}
	if cloudbrain.Cloudbrain.Type == models.TypeCDCenter {
		cloudbrain.Cloudbrain.AiCenter = models.AICenterOfChengdu
		cloudbrain.Cloudbrain.Cluster = models.OpenICluster
	}
	if cloudbrain.Cloudbrain.Type == models.TypeC2Net {
		cloudbrain.Cloudbrain.AiCenter = GetAiCenterCode(cloudbrain.Cloudbrain.AiCenter)
		cloudbrain.Cloudbrain.Cluster = models.C2NetCluster
	}

	return cloudbrain
}

func UpdateCloudbrainComputeResource(cloudbrain *models.CloudbrainInfo) *models.CloudbrainInfo {
	if cloudbrain.Cloudbrain.ComputeResource == models.GPUResource {
		cloudbrain.Cloudbrain.ComputeResource = models.GPU
	}
	return cloudbrain
}
