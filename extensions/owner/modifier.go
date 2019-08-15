package owner

import (
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
	"github.com/soodakshay/cckit/router"
)

var (
	// ErrOwnerOnly error occurs when trying to invoke chaincode func  protected by onlyOwner middleware (modifier)
	ErrOwnerOnly = errors.New(`owner only`)
)

// Only allow access from chain code owner
func Only(next router.HandlerFunc, pos ...int) router.HandlerFunc {
	return func(c router.Context) peer.Response {
		invokerIsOwner, err := IsInvoker(c)
		if invokerIsOwner && err == nil {
			return next(c)
		}
		return nil, ErrOwnerOnly
	}
}
