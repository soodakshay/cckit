package defparam

import (
	"github.com/soodakshay/cckit/router"
	"github.com/soodakshay/cckit/router/param"
)

func Proto(target interface{}, argPoss ...int) router.MiddlewareFunc {
	return param.Proto(router.DefaultParam, target, argPoss...)
}
