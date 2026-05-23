package auth

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func newHasher() *hasher {
	return &hasher{
		hash: sha256.New(),
	}
}

func (h *hasher) Write(s string) (n int, err error) {
	return h.hash.Write([]byte(s))
}

func (h *hasher) GetHex() string {
	return fmt.Sprintf("%x", h.hash.Sum(nil))
}
