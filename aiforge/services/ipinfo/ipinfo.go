package ipinfo

import (
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/ipinfo"
	"code.gitea.io/gitea/modules/log"
)

func GetLocationByIp(ip string) *models.IPLocation {

	ipLocation, err := models.GetIpLocation(ip)

	if err != nil {
		if !models.IsErrRecordNotExist(err) {
			log.Error("can not get ip locatin from db", err)
		}

		ipInfoRes, err := ipinfo.GetLocationByIp(ip)

		ipLocationTemp := &models.IPLocation{
			IpAddr: ip,
		}
		if err != nil || ipInfoRes.Bogon || ipInfoRes.Loc == "" {
			return ipLocationTemp
		}
		latLongArray := strings.Split(ipInfoRes.Loc, ",")
		ipLocationTemp.Latitude = latLongArray[0]
		ipLocationTemp.Longitude = latLongArray[1]
		models.CreateIPLocation(ipLocationTemp)
		return ipLocationTemp

	} else {
		return ipLocation
	}

}
