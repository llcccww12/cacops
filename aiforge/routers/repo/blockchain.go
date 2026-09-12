package repo

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/blockchain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"encoding/json"
	"net/http"
	"strconv"
)

type BlockChainInitNotify struct {
	RepoId          int64  `json:"repoId"`
	ContractAddress string `json:"contractAddress"`
}

type BlockChainCommitNotify struct {
	CommitID        string `json:"commitId"`
	TransactionHash string `json:"txHash"`
}

const (
	tplBlockChainIndex base.TplName = "repo/blockchain/index"
)

func BlockChainIndex(ctx *context.Context) {
	repo := ctx.Repo.Repository
	if repo.ContractAddress == "" || ctx.User.PublicKey == "" {
		log.Error("the repo(%d) or the user(%d) has not been initialized in block_chain", repo.RepoID, ctx.User.ID, ctx.Data["msgID"])
		ctx.HTML(http.StatusInternalServerError, tplBlockChainIndex)
		return
	}

	res, err := blockchain.GetBalance(repo.ContractAddress, ctx.User.PublicKey)
	if err != nil {
		log.Error("GetBalance(%s) failed:%s", ctx.User.PublicKey, err, ctx.Data["msgID"])
		ctx.HTML(http.StatusInternalServerError, tplBlockChainIndex)
		return
	}

	ctx.Data["balance"] = res.Data
	ctx.Data["PageIsBlockChain"] = true
	ctx.HTML(200, tplBlockChainIndex)
}

func HandleBlockChainInitNotify(ctx *context.Context) {
	var req BlockChainInitNotify
	data, _ := ctx.Req.Body().Bytes()
	json.Unmarshal(data, &req)

	repo, err := models.GetRepositoryByID(req.RepoId)
	if err != nil {
		log.Error("GetRepositoryByID failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.JSON(200, map[string]string{
			"code":    "-1",
			"message": "internal error",
		})
		return
	}

	if repo.BlockChainStatus == models.RepoBlockChainSuccess && len(repo.ContractAddress) != 0 {
		log.Error("the repo has been RepoBlockChainSuccess:%d", req.RepoId, ctx.Data["msgID"])
		ctx.JSON(200, map[string]string{
			"code":    "-1",
			"message": "the repo has been RepoBlockChainSuccess",
		})
		return
	}

	repo.BlockChainStatus = models.RepoBlockChainSuccess
	repo.ContractAddress = req.ContractAddress

	if err = repo.UpdateBlockChain(); err != nil {
		log.Error("UpdateRepositoryCols failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.JSON(200, map[string]string{
			"code":    "-1",
			"message": "internal error",
		})
		return
	}

	ctx.JSON(200, map[string]string{
		"code":    "0",
		"message": "",
	})
}

func HandleBlockChainCommitNotify(ctx *context.Context) {
	var req BlockChainCommitNotify
	data, _ := ctx.Req.Body().Bytes()
	if err := json.Unmarshal(data, &req); err != nil {
		log.Error("json.Unmarshal failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.JSON(200, map[string]string{
			"code":    "-1",
			"message": "response data error",
		})
		return
	}

	blockChain, err := models.GetBlockChainByCommitID(req.CommitID)
	if err != nil {
		log.Error("GetRepositoryByID failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.JSON(200, map[string]string{
			"code":    "-1",
			"message": "internal error",
		})
		return
	}

	if blockChain.Status == models.BlockChainCommitSuccess {
		log.Error("the commit has been BlockChainCommitReady:%s", blockChain.RepoID, ctx.Data["msgID"])
		ctx.JSON(200, map[string]string{
			"code":    "-1",
			"message": "the commit has been BlockChainCommitReady",
		})
		return
	}

	blockChain.Status = models.BlockChainCommitSuccess
	blockChain.TransactionHash = req.TransactionHash

	if err = models.UpdateBlockChainCols(blockChain, "status", "transaction_hash"); err != nil {
		log.Error("UpdateBlockChainCols failed:%v", err.Error(), ctx.Data["msgID"])
		ctx.JSON(200, map[string]string{
			"code":    "-1",
			"message": "internal error",
		})
		return
	}

	ctx.JSON(200, map[string]string{
		"code":    "0",
		"message": "",
	})
}

func HandleBlockChainUnSuccessRepos() {
	repos, err := models.GetBlockChainUnSuccessRepos()
	if err != nil {
		log.Error("GetBlockChainUnSuccessRepos failed:", err.Error())
		return
	}

	for _, repo := range repos {
		err = repo.GetOwner()
		if err != nil {
			log.Error("GetOwner(%s) failed:%v", repo.Name, err)
			continue
		}
		if len(repo.Owner.PrivateKey) == 0 || len(repo.Owner.PublicKey) == 0 {
			log.Error("the user has not been init in block_chain:", repo.Owner.Name)
			continue
		}
		strRepoID := strconv.FormatInt(repo.ID, 10)
		_, err = blockchain.NewRepo(strRepoID, repo.Owner.PublicKey, repo.Name)
		if err != nil {
			log.Error("blockchain.NewRepo(%s) failed:%v", strRepoID, err)
		}
	}

	return
}

func HandleBlockChainUnSuccessCommits() {
	blockChains, err := models.GetBlockChainUnSuccessCommits()
	if err != nil {
		log.Error("GetBlockChainUnSuccessCommits failed:", err.Error())
		return
	}

	for _, block_chain := range blockChains {
		_, err = blockchain.Contribute(block_chain.ContractAddress, block_chain.Contributor, block_chain.CommitID, block_chain.Amount)
		if err != nil {
			log.Error("blockchain.Contribute(%s) failed:%v", block_chain.CommitID, err)
		}
	}

	return
}

func HandleBlockChainUnSuccessUsers() {
	users, err := models.GetBlockChainUnSuccessUsers()
	if err != nil {
		log.Error("GetBlockChainUnSuccessUsers failed:", err.Error())
		return
	}

	for _, user := range users {
		result, err := blockchain.CreateBlockchainAccount()
		if err != nil {
			log.Error("blockchain.CreateBlockchainAccount(%s) failed:%v", user.Name, err)
			continue
		}

		user.PublicKey = result.Payload["publickey"].(string)
		user.PrivateKey = result.Payload["privatekey"].(string)

		models.UpdateUser(user)
	}

	return
}

func HandleBlockChainMergedPulls() {
	prs, err := models.GetUnTransformedMergedPullRequests()
	if err != nil {
		log.Error("GetUnTransformedMergedPullRequests failed:", err.Error())
		return
	}

	for _, pr := range prs {
		_, err = models.GetBlockChainByPrID(pr.ID)
		if err == nil {
			log.Info("the pr(%s) has been transformed", pr.MergedCommitID)
			continue
		}

		err = pr.LoadIssue()
		if err != nil {
			log.Error("LoadIssue(%s) failed:%v", pr.MergedCommitID, err)
			continue
		}

		poster, err := models.GetUserByID(pr.Issue.PosterID)
		if err != nil {
			log.Error("GetUserByID(%s) failed:%v", pr.MergedCommitID, err)
			continue
		}

		if len(poster.PrivateKey) == 0 || len(poster.PublicKey) == 0 {
			log.Error("the user has not been init in block_chain:", poster.Name)
			continue
		}

		repo, err := models.GetRepositoryByID(pr.HeadRepoID)
		if err != nil {
			log.Error("GetUserByID(%s) failed:%v", pr.MergedCommitID, err)
			continue
		}

		if len(repo.ContractAddress) == 0 {
			log.Error("the repo(%s) has not been initialized in block_chain", repo.Name)
			continue
		}

		blockChain := models.BlockChain{
			Contributor:     poster.PublicKey,
			PrID:            pr.ID,
			CommitID:        pr.MergedCommitID,
			ContractAddress: repo.ContractAddress,
			Status:          models.BlockChainCommitInit,
			Amount:          int64(pr.Amount),
			UserID:          poster.ID,
			RepoID:          pr.HeadRepoID,
		}
		_, err = models.InsertBlockChain(&blockChain)
		if err != nil {
			log.Error("InsertBlockChain(%s) failed:%v", pr.MergedCommitID, err)
			continue
		}

		pr.IsTransformed = true
		pr.UpdateCols("is_transformed")

		_, err = blockchain.Contribute(repo.ContractAddress, poster.PublicKey, pr.MergedCommitID, int64(pr.Amount))
		if err != nil {
			log.Error("Contribute(%s) failed:%v", pr.MergedCommitID, err)
		}
	}

	return
}

func HandleBlockChainUnSuccessIssues() {
	issues, err := models.GetBlockChainUnSuccessCommits()
	if err != nil {
		log.Error("GetBlockChainUnSuccessIssues failed:", err.Error())
		return
	}

	for _, issue := range issues {
		_, err = blockchain.SetIssue(issue.ContractAddress, issue.Contributor, issue.ID, issue.Amount)
		if err != nil {
			log.Error("SetIssue(%s) failed:%v", issue.CommitID, err)
		}
	}

	return
}
