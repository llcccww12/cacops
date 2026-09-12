package repo

import (
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
)

func SummaryStatistic() {
	log.Info("Generate summary statistic begin")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	SummaryStatisticDaily(yesterday)
	log.Info("Generate summary statistic end")
}

func SummaryStatisticDaily(date string) {
	log.Info("%s", date)
	if err := models.DeleteSummaryStatisticDaily(date); err != nil {
		log.Error("DeleteRepoStatDaily failed: %v", err.Error())
		return
	}

	//user number
	userNumber, err := models.GetUsersCount()
	if err != nil {
		log.Error("can not get user number", err)
		userNumber = 0
	}
	//organization number
	organizationNumber, err := models.GetOrganizationsCount()
	if err != nil {
		log.Error("can not get orgnazition number", err)
		organizationNumber = 0
	}
	// repository number
	repositoryNumer, err := models.GetAllRepositoriesCount()
	if err != nil {
		log.Error("can not get repository number", err)
		repositoryNumer = 0
	}

	publicRepositoryNumer, err := models.GetAllPublicRepositoriesCount()
	if err != nil {
		log.Error("can not get public repository number", err)
		publicRepositoryNumer = 0
	}

	privateRepositoryNumer := repositoryNumer - publicRepositoryNumer

	mirrorRepositoryNumber, err := models.GetAllMirrorRepositoriesCount()
	if err != nil {
		log.Error("can not get mirror repository number", err)
		mirrorRepositoryNumber = 0
	}
	forkRepositoryNumber, err := models.GetAllForkRepositoriesCount()
	if err != nil {
		log.Error("can not get fork mirror repository number", err)
		forkRepositoryNumber = 0
	}
	selfRepositoryNumber := repositoryNumer - mirrorRepositoryNumber - forkRepositoryNumber

	organizationRepoNumber, err := models.GetAllOrgRepositoriesCount()
	if err != nil {
		log.Error("can not get org repository number", err)
		organizationRepoNumber = 0
	}

	//repository size
	repositorySize, err := models.GetAllRepositoriesSize()
	if err != nil {
		log.Error("can not get repository size", err)
		repositorySize = 0
	}
	// dataset size
	allDatasetSize, err := models.GetAllAttachmentSize()
	if err != nil {
		log.Error("can not get dataset size", err)
		allDatasetSize = 0
	}
	//topic repo number
	topics, err := models.GetAllUsedTopics()
	if err != nil {
		log.Error("can not get topics", err)
	}
	var topicsCount [11]int
	for _, topic := range topics {

		index, exists := models.DomainMap[topic.Name]
		if exists {
			topicsCount[index] = topic.RepoCount
		}

	}

	summaryStat := models.SummaryStatistic{
		Date:               date,
		NumUsers:           userNumber,
		RepoSize:           repositorySize,
		DatasetSize:        allDatasetSize,
		NumOrganizations:   organizationNumber,
		NumRepos:           repositoryNumer,
		NumRepoFork:        forkRepositoryNumber,
		NumRepoMirror:      mirrorRepositoryNumber,
		NumRepoPrivate:     privateRepositoryNumer,
		NumRepoPublic:      publicRepositoryNumer,
		NumRepoSelf:        selfRepositoryNumber,
		NumRepoOrg:         organizationRepoNumber,
		NumRepoBigModel:    topicsCount[0],
		NumRepoAI:          topicsCount[1],
		NumRepoVision:      topicsCount[2],
		NumRepoNLP:         topicsCount[3],
		NumRepoML:          topicsCount[4],
		NumRepoNN:          topicsCount[5],
		NumRepoAutoDrive:   topicsCount[6],
		NumRepoRobot:       topicsCount[7],
		NumRepoLeagueLearn: topicsCount[8],
		NumRepoDataMining:  topicsCount[9],
		NumRepoRISC:        topicsCount[10],
	}

	if _, err = models.InsertSummaryStatistic(&summaryStat); err != nil {
		log.Error("Insert summary Stat failed: %v", err.Error())
	}

	log.Info("finish summary statistic")
}
