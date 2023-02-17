package cryptor

import (
	"hash"
)

type hashFun func() hash.Hash
