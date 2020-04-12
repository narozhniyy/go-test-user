package router

import (
	"github.com/labstack/echo"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.GET("/:fu/:su", userCheck)
	
	return e
}