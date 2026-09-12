package lock

type LockChainOperator struct {
	chainList  []Lock
	lockedList []Lock
	ctx        *LockContext
}

func NewLockChainOperator(ctx *LockContext) *LockChainOperator {
	return &LockChainOperator{ctx: ctx}
}

func (b *LockChainOperator) Add(l Lock) *LockChainOperator {
	b.chainList = append(b.chainList, l)
	return b
}

func (b *LockChainOperator) Lock() string {
	for i := 0; i < len(b.chainList); i++ {
		l := b.chainList[i]
		if !l.IsMatch(b.ctx) {
			continue
		}

		if errCode := l.Lock(b.ctx); errCode != "" {
			b.Unlock()
			return errCode
		}
		b.lockedList = append(b.lockedList, l)
	}
	return ""
}

func (b *LockChainOperator) Unlock() error {
	if b.chainList == nil || len(b.chainList) == 0 {
		return nil
	}
	for j := len(b.lockedList) - 1; j >= 0; j-- {
		b.lockedList[j].Unlock(b.ctx)
	}
	return nil
}
