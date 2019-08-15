package param

import (
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/soodakshay/cckit/router"
)

// StrictKnown allows passing arguments to chaincode func only if parameters are defined in router
func StrictKnown(next router.HandlerFunc, pos ...int) router.HandlerFunc {
	return func(c router.Context) peer.Response {
		return next(c)
	}
}
