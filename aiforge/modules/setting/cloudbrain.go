package setting

type CloudbrainLoginConfig struct {
	Username       string
	Password       string
	Host           string
	ImageURLPrefix string
	Expiration     string
}

var (
	Cloudbrain = CloudbrainLoginConfig{}
)

func GetCloudbrainConfig() CloudbrainLoginConfig {
	cloudbrainSec := Cfg.Section("cloudbrain")
	Cloudbrain.Username = cloudbrainSec.Key("USERNAME").MustString("")
	Cloudbrain.Password = cloudbrainSec.Key("PASSWORD").MustString("")
	Cloudbrain.Host = cloudbrainSec.Key("REST_SERVER_HOST").MustString("")
	Cloudbrain.ImageURLPrefix = cloudbrainSec.Key("IMAGE_URL_PREFIX").MustString("")
	Cloudbrain.Expiration = cloudbrainSec.Key("EXPIRATION").MustString("604800")
	return Cloudbrain
}
