package param

import (
	"fmt"

<<<<<<< HEAD
	"github.com/hyperledger/fabric/protos/peer"
=======
	"github.com/pkg/errors"
>>>>>>> d9270d9c7e0def8422d5975be671d71af6c232c9
	"github.com/soodakshay/cckit/convert"
	"github.com/soodakshay/cckit/router"
)

const LastPosKey = `_lastPos`

// ErrPayloadValidationError occurs when payload validation not passed
var ErrPayloadValidationError = errors.New(`payload validation`)

type (
	// Parameters list of chain code function parameters
	Parameters []Parameter

	// Parameter of chain code function
	Parameter struct {
		Name   string
		Type   interface{}
		ArgPos int
	}

	//DefinedParams

	// MiddlewareFuncMap named list of middleware functions
	MiddlewareFuncMap map[string]router.MiddlewareFunc
)

func (p Parameter) ValueFromContext(c router.Context) (arg interface{}, err error) {
	// by default args start from pos 1 , at first pos is funcName
	argsStartsFrom := 1
	//if c.Path() == router.InitFunc {
	//	argsStartsFrom = 0
	//}
	argPos := p.ArgPos
	if argPos == -1 {
		lastPos, ok := c.Param(LastPosKey).(int)
		if !ok {
			argPos = 0
		} else {
			argPos = lastPos + 1
		}
		c.SetParam(LastPosKey, argPos)
	}

	args := c.GetArgs()[argsStartsFrom:] // first arg is chaincode function name
	if argPos >= len(args) {
		return nil, fmt.Errorf(
			`method "%s", param "%s" not exists, param expected at pos : %d, stub args length: %d`,
			c.Path(), p.Name, argPos, len(args))
	}

	return convert.FromBytes(args[argPos], p.Type) //first arg is function name
}

// Add middleware function
func (pbag MiddlewareFuncMap) Add(name string, paramType interface{}) MiddlewareFuncMap {
	pbag[name] = Param(name, paramType)
	return pbag
}

// Param create middleware function for transforming stub arg to context arg
func Param(name string, paramType interface{}, argPoss ...int) router.MiddlewareFunc {
	var argPos int
	if len(argPoss) == 0 {
		argPos = -1 // use next pos
	} else {
		argPos = argPoss[0]
	}

	parameter := Parameter{name, paramType, argPos}

	return func(next router.HandlerFunc, pos ...int) router.HandlerFunc {
		return func(c router.Context) peer.Response {

			arg, err := parameter.ValueFromContext(c)
			if err != nil {
				return nil, err
			}
			c.SetParam(name, arg)
			return next(c)
		}
	}
}
