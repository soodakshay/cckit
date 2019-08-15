package param

import (
<<<<<<< HEAD
	"github.com/hyperledger/fabric/protos/peer"
=======
>>>>>>> d9270d9c7e0def8422d5975be671d71af6c232c9
	"github.com/soodakshay/cckit/router"
)

// StrictKnown allows passing arguments to chaincode func only if parameters are defined in router
func StrictKnown(next router.HandlerFunc, pos ...int) router.HandlerFunc {
	return func(c router.Context) peer.Response {
		return next(c)
	}
}
