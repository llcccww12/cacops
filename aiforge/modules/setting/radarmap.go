package setting

func UpdateRadarMap() {
	Cfg.DeleteSection("radar_map")
	Cfg.Reload()
	SetRadarMapConfig()
}
