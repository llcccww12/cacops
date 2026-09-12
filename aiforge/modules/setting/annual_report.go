package setting

var AnnualCollection *AnnualReportCollection

type AnnualReportCollection struct {
	Enabled bool
	Year    int
}

func RefreshAnnualCollectionConfig() {
	Cfg.Reload()
	sec := Cfg.Section("annual_report_collection")
	AnnualCollection = &AnnualReportCollection{
		Enabled: sec.Key("Enabled").MustBool(false),
		Year:    sec.Key("Year").MustInt(2024),
	}
}
