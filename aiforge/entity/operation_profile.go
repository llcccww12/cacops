package entity

type OperationProfile struct {
	Events []ProfileEvent `json:"events"`
}

type ProfileEvent struct {
	Message   string `json:"message"`
	Name      string `json:"name"`
	Reason    string `json:"reason"`
	Timestamp string `json:"timestamp"`
	Action    string `json:"action"`
}

type CloudbrainOneAppExitDiagnostics struct {
	PodRoleName struct {
		Task10 string `json:"task1-0"`
	} `json:"podRoleName"`
	PodEvents struct {
		Task10 []struct {
			Uid                 string `json:"uid"`
			Reason              string `json:"reason"`
			Message             string `json:"message"`
			ReportingController string `json:"reportingController"`
			Action              string `json:"action"`
		} `json:"task1-0"`
	} `json:"podEvents"`
	Extras []struct {
		Uid                 string `json:"uid"`
		Reason              string `json:"reason"`
		Message             string `json:"message"`
		ReportingController string `json:"reportingController"`
		Action              string `json:"action"`
	} `json:"extras"`
}
