package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"fmt"
	"time"
)

type BlockChainIssueStatus int

const (
	BlockChainIssueInit BlockChainIssueStatus = iota
	BlockChainIssueSuccess
	BlockChainIssueFailed
)

type BlockChainIssue struct {
	ID              int64                 `xorm:"pk autoincr"`
	IssueID         int64                 `xorm:"INDEX NOT NULL unique"`
	Contributor     string                `xorm:"INDEX NOT NULL"`
	ContractAddress string                `xorm:"INDEX NOT NULL"`
	Status          BlockChainIssueStatus `xorm:"INDEX NOT NULL DEFAULT 0"`
	Amount          int64                 `xorm:"INDEX"`
	UserID          int64                 `xorm:"INDEX"`
	RepoID          int64                 `xorm:"INDEX"`
	TransactionHash string                `xorm:"INDEX"`
	CreatedUnix     timeutil.TimeStamp    `xorm:"created"`
	UpdatedUnix     timeutil.TimeStamp    `xorm:"updated"`
	DeletedAt       time.Time             `xorm:"deleted"`

	User *User       `xorm:"-"`
	Repo *Repository `xorm:"-"`
}

func getBlockChainIssueByID(e Engine, id int64) (*BlockChainIssue, error) {
	blockChainIssue := new(BlockChainIssue)
	has, err := e.ID(id).Get(blockChainIssue)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("get block_chain by id failed(%d)", id)
	}
	return blockChainIssue, nil
}

func GetBlockChainIssueByID(id int64) (*BlockChainIssue, error) {
	return getBlockChainIssueByID(x, id)
}

func getBlockChainIssueByPrID(e Engine, prId int64) (*BlockChainIssue, error) {
	blockChainIssue := new(BlockChainIssue)
	has, err := e.Where("pr_id = ?", prId).Get(blockChainIssue)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("get block_chain by pr_id failed(%d)", prId)
	}
	return blockChainIssue, nil
}

func GetBlockChainIssueByPrID(prId int64) (*BlockChainIssue, error) {
	return getBlockChainIssueByPrID(x, prId)
}

func getBlockChainIssueByCommitID(e Engine, commitID string) (*BlockChainIssue, error) {
	blockChainIssue := new(BlockChainIssue)
	has, err := e.Where("commit_id = ?", commitID).Get(blockChainIssue)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("get block_chain_issue by commitID failed(%s)", commitID)
	}
	return blockChainIssue, nil
}

func GetBlockChainIssueByCommitID(commitID string) (*BlockChainIssue, error) {
	return getBlockChainIssueByCommitID(x, commitID)
}

func updateBlockChainIssueCols(e Engine, blockChainIssue *BlockChainIssue, cols ...string) error {
	_, err := e.ID(blockChainIssue.ID).Cols(cols...).Update(blockChainIssue)
	return err
}

func UpdateBlockChainIssueCols(blockChainIssue *BlockChainIssue, cols ...string) error {
	return updateBlockChainIssueCols(x, blockChainIssue, cols...)
}

func GetBlockChainUnSuccessIssues() ([]*BlockChainIssue, error) {
	blockChainIssues := make([]*BlockChainIssue, 0, 10)
	return blockChainIssues, x.
		Where("status != ?", BlockChainIssueSuccess).
		Find(&blockChainIssues)
}

func InsertBlockChainIssue(blockChainIssue *BlockChainIssue) (_ *BlockChainIssue, err error) {

	if _, err := x.Insert(blockChainIssue); err != nil {
		return nil, err
	}

	return blockChainIssue, nil
}
