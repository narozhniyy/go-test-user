package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/narozhniyy/test/models"
	"github.com/narozhniyy/test/repository/connLog"
	"net/http"
	"strconv"
)

func userCheck(c echo.Context) error {
	fu, err := strconv.ParseInt(c.Param("fu"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{err.Error()})
	}

	su, err := strconv.ParseInt(c.Param("su"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{err.Error()})
	}

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