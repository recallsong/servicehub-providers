package intercepters

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/recallsong/go-utils/logs"
)

// CORS .
func CORS() interface{} {
	return middleware.CORS()
}

// Recover .
func Recover(log logs.Logger) interface{} {
	const StackSize = 4 << 10
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, StackSize)
					length := runtime.Stack(stack, true)
					log.Errorf("[PANIC RECOVER] %v %s\n", err, stack[:length])
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
