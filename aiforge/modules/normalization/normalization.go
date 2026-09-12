package normalization

import (
	"code.gitea.io/gitea/modules/setting"
)

const MAX_LINES_RECORD = 100

func Normalization(value float64, minValue float64, maxValue float64) float64 {

	min := int64(minValue * 100)
	max := int64(maxValue * 100)

	if min == max {
		return 100.0
	} else {
		return 100 * (value - minValue) / (maxValue - minValue)
	}

}

func GetRadarValue(impactValue float64, completeValue float64, livenessValue float64, projectHealthValue float64, teamHealthValue float64, growthValue float64) float64 {
	return setting.RadarMap.Impact*impactValue +
		setting.RadarMap.Completeness*completeValue +
		setting.RadarMap.Liveness*livenessValue +
		setting.RadarMap.ProjectHealth*projectHealthValue +
		setting.RadarMap.TeamHealth*teamHealthValue +
		setting.RadarMap.Growth*growthValue

}

func GetImpactInitValue(watch int64, star int64, fork int64, download int64, comments int64, browser int64) float64 {

	return setting.RadarMap.ImpactWatch*float64(watch) +
		setting.RadarMap.ImpactStar*float64(star) +
		setting.RadarMap.ImpactFork*float64(fork) +
		setting.RadarMap.ImpactCodeDownload*float64(download)*0.001 +
		setting.RadarMap.ImpactComments*float64(comments) +
		setting.RadarMap.ImpactBrowser*float64(browser)*0.001

}

func GetCompleteInitValue(issuesClosed int64, releases int64, developAge int64, dataset int64, model int64, wiki int64) float64 {

	return setting.RadarMap.CompletenessIssuesClosed*float64(issuesClosed) +
		setting.RadarMap.CompletenessReleases*float64(releases) +
		setting.RadarMap.CompletenessDevelopAge*float64(developAge) +
		setting.RadarMap.CompletenessDataset*(float64(dataset)/(1024*1024)) +
		setting.RadarMap.CompletenessModel*float64(model) +
		setting.RadarMap.CompletenessWiki*float64(wiki)

}

func GetLivenessInitValue(commits int64, issues int64, pr int64, release int64) float64 {

	return setting.RadarMap.LivenessCommit*float64(commits) +
		setting.RadarMap.LivenessIssue*float64(issues) +
		setting.RadarMap.LivenessPR*float64(pr) +
		setting.RadarMap.LivenessRelease*float64(release)

}

func GetProjectHealthInitValue(issueClosedRatio float32) float64 {

	return setting.RadarMap.ProjectHealthIssueCompleteRatio * float64(issueClosedRatio)

}

func GetTeamHealthInitValue(contributors int64, keyContributors int64, newContributors int64) float64 {

	return setting.RadarMap.TeamHealthContributors*float64(contributors) +
		setting.RadarMap.TeamHealthKeyContributors*float64(keyContributors) +
		setting.RadarMap.TeamHealthContributorsAdded*float64(newContributors)

}

func GetRepoGrowthInitValue(codeLinesGrowth int64, issueGrowth int64, commitsGrowth int64, newContributors int64, commentsGrowth int64) float64 {
	codeLinesKB := codeLinesGrowth / 1000
	if codeLinesKB > MAX_LINES_RECORD {
		codeLinesKB = MAX_LINES_RECORD
	}
	return setting.RadarMap.GrowthCodeLines*float64(codeLinesKB) +
		setting.RadarMap.GrowthIssue*float64(issueGrowth) +
		setting.RadarMap.GrowthCommit*float64(commitsGrowth) +
		setting.RadarMap.GrowthContributors*float64(newContributors) +
		setting.RadarMap.GrowthComments*float64(commentsGrowth)

}
