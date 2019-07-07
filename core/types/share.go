package types

import (
	"github.com/sero-cash/go-czero-import/keys"
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/crypto"
)

type Lottery struct {
	ParentHash common.Hash
	PosHash    common.Hash
}

type Vote struct {
	ShareHash common.Hash
	PosHash   common.Hash
	IsPool    bool
	Sign      keys.Uint512
}

func (vote Vote) Hash() common.Hash {
	return crypto.Keccak256Hash(vote.Sign[:])
}