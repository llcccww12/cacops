package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"fmt"
	"time"
)

type BlockChainCommitStatus int

const (
	BlockChainCommitInit BlockChainCommitStatus = iota
	BlockChainCommitSuccess
	BlockChainCommitFailed
)

type BlockChain struct {
	ID              int64                  `xorm:"pk autoincr"`
	PrID            int64                  `xorm:"INDEX NOT NULL unique"`
	CommitID        string                 `xorm:"INDEX NOT NULL unique"`
	Contributor     string                 `xorm:"INDEX NOT NULL"`
	ContractAddress string                 `xorm:"INDEX NOT NULL"`
	Status          BlockChainCommitStatus `xorm:"INDEX NOT NULL DEFAULT 0"`
	Amount          int64                  `xorm:"INDEX"`
	UserID          int64                  `xorm:"INDEX"`
	RepoID          int64                  `xorm:"INDEX"`
	TransactionHash string                 `xorm:"INDEX"`
	CreatedUnix     timeutil.TimeStamp     `xorm:"created"`
	UpdatedUnix     timeutil.TimeStamp     `xorm:"updated"`
	DeletedAt       time.Time              `xorm:"deleted"`

	User *User       `xorm:"-"`
	Repo *Repository `xorm:"-"`
}

func getBlockChainByID(e Engine, id int64) (*BlockChain, error) {
	blockChain := new(BlockChain)
	has, err := e.ID(id).Get(blockChain)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("get block_chain by id failed(%d)", id)
	}
	return blockChain, nil
}

func GetBlockChainByID(id int64) (*BlockChain, error) {
	return getBlockChainByID(x, id)
}

func getBlockChainByPrID(e Engine, prId int64) (*BlockChain, error) {
	blockChain := new(BlockChain)
	has, err := e.Where("pr_id = ?", prId).Get(blockChain)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("get block_chain by pr_id failed(%d)", prId)
	}
	return blockChain, nil
}

func GetBlockChainByPrID(prId int64) (*BlockChain, error) {
	return getBlockChainByPrID(x, prId)
}

func getBlockChainByCommitID(e Engine, commitID string) (*BlockChain, error) {
	blockChain := new(BlockChain)
	has, err := e.Where("commit_id = ?", commitID).Get(blockChain)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("get block_chain by commitID failed(%s)", commitID)
	}
	return blockChain, nil
}

func GetBlockChainByCommitID(commitID string) (*BlockChain, error) {
	return getBlockChainByCommitID(x, commitID)
}

func updateBlockChainCols(e Engine, blockChain *BlockChain, cols ...string) error {
	_, err := e.ID(blockChain.ID).Cols(cols...).Update(blockChain)
	return err
}

func UpdateBlockChainCols(blockChain *BlockChain, cols ...string) error {
	return updateBlockChainCols(x, blockChain, cols...)
}

func GetBlockChainUnSuccessCommits() ([]*BlockChain, error) {
	blockChains := make([]*BlockChain, 0, 10)
	return blockChains, x.
		Where("status != ?", BlockChainCommitSuccess).
		Limit(100).
		Find(&blockChains)
}

func InsertBlockChain(blockChain *BlockChain) (_ *BlockChain, err error) {

	if _, err := x.Insert(blockChain); err != nil {
		return nil, err
	}

	return blockChain, nil
}
