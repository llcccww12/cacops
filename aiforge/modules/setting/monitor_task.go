package setting



var MonitorTaskHost  string

func iniMonitorTaskConfig() {
	sec := Cfg.Section("monitor_task")
	MonitorTaskHost=sec.Key("HOST").MustString("")
	
}