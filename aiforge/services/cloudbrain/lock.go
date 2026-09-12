package cloudbrain

import "code.gitea.io/gitea/services/lock"

func Lock4CloudbrainCreation(ctx *lock.LockContext) (*lock.LockChainOperator, string) {
	op := lock.NewLockChainOperator(ctx).Add(lock.CloudbrainUniquenessLock{}).Add(lock.CloudbrainDisplayJobNameLock{})
	errCode := op.Lock()
	if errCode != "" {
		return nil, errCode
	}
	return op, ""
}
func Lock4CloudbrainRestart(ctx *lock.LockContext) (*lock.LockChainOperator, string) {
	op := lock.NewLockChainOperator(ctx).Add(lock.CloudbrainUniquenessLock{})
	errCode := op.Lock()
	if errCode != "" {
		return nil, errCode
	}
	return op, ""
}
