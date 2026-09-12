package cloudbrain

import (
	"net/http"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/services/ipinfo"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
)

type AiCenterLocationInfo struct {
	Name      string `json:"name"`
	Value     int    `json:"value"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type CloudBrainLocationInfo struct {
	DisplayJobName string `json:"job_name"`
	FromLongitude  string `json:"from_longitude"`
	FromLatitude   string `json:"from_latitude"`
	ToLongitude    string `json:"to_longitude"`
	ToLatitude     string `json:"to_latitude"`
}

func GetCloudbrainLocationInfo(ctx *context.Context) {
	var locationInfos = make([]*CloudBrainLocationInfo, 0)
	cloudbrains, err := models.GetLastestNCloudbrain(20)
	if err != nil {
		log.Error("can not get cloudbrain info", err)
		ctx.JSON(http.StatusOK, locationInfos)
		return
	}
	var userCenterMap = make(map[string]string, 0)

	for _, cloudbrain := range cloudbrains {
		key := strconv.FormatInt(cloudbrain.UserID, 10) + "_" + strconv.Itoa(cloudbrain.Type) + "_" + cloudbrain.AiCenter

		if _, ok := userCenterMap[key]; ok {
			continue
		} else {
			userCenterMap[key] = ""
		}

		ip := models.GetIpByUID(cloudbrain.UserID)
		if ip != "" {
			ipLocation := ipinfo.GetLocationByIp(ip)
			aicenter := cloudbrain.GetAiCenter()

			if ipLocation.Latitude != "" && aicenter != "" {
				if value, ok := setting.AiCenterCodeAndNameAndLocMapInfo[aicenter]; ok {
					longLat := strings.Split(value.Loc, ",")
					cloudBrainLocationInfo := &CloudBrainLocationInfo{
						DisplayJobName: cloudbrain.DisplayJobName,
						FromLongitude:  ipLocation.Longitude,
						FromLatitude:   ipLocation.Latitude,
						ToLatitude:     longLat[1],
						ToLongitude:    longLat[0],
					}
					if len(locationInfos) == 10 {
						break
					}
					locationInfos = append(locationInfos, cloudBrainLocationInfo)
				}

			}

		}

	}

	ctx.JSON(http.StatusOK, locationInfos)

}
