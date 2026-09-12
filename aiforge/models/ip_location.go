package models

type IPLocation struct {
	ID        int64  `xorm:"pk autoincr"`
	IpAddr    string `xorm:"unique"`
	Longitude string
	Latitude  string
}

func CreateIPLocation(ipLocation *IPLocation) (err error) {
	_, err = x.Insert(ipLocation)
	return err

}

func GetIpLocation(ip string) (*IPLocation, error) {

	ipLocation := &IPLocation{IpAddr: ip}
	has, err := x.Get(ipLocation)
	if err != nil {
		return nil, err
	}

	if has {
		return ipLocation, nil
	} else {
		return nil, ErrRecordNotExist{}
	}

}
