package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/narozhniyy/test/models"
	"github.com/narozhniyy/test/repository/connLog"
	"net/http"
)

func userCheck(c echo.Context) error {
	fu := c.Param("fu")
	su := c.Param("su")

	var dps models.Result
	repo := connLogRepository.InitConnLogRepository()

	ipOccurrence, err := repo.Compare(fu, su)

	if err != nil {
		log.Fatal(err)
	}

	if len(ipOccurrence) >= 2 {
		dps.Duples = true
	} else {
		dps.Duples = false
	}

	return c.JSON(http.StatusOK, dps)
}