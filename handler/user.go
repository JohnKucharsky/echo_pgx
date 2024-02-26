package handler

import (
	"github.com/JohnKucharsky/echo_pgx/repository"
	"github.com/JohnKucharsky/echo_pgx/serializer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) CreateUser(c echo.Context) error {
	var userBody serializer.UserBody
	if err := c.Bind(&userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := repository.CreateUser(apiConfig.Pool, userBody)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		user,
	)
}

func (apiConfig *DatabaseController) GetUsers(c echo.Context) error {
	users, err := repository.GetUsers(apiConfig.Pool)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusOK,
		users,
	)
}

func (apiConfig *DatabaseController) GetOneUser(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	dbId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := repository.GetOneUser(apiConfig.Pool, int(dbId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusOK, user,
	)
}

func (apiConfig *DatabaseController) UpdateUser(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	var dbId int32
	res, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbId = int32(res)

	var userBody serializer.UserBody
	if err := c.Bind(&userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := repository.UpdateUser(apiConfig.Pool, userBody, int(dbId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusOK, user,
	)
}

func (apiConfig *DatabaseController) DeleteUser(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	var dbId int
	res, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbId = int(res)

	user, err := repository.DeleteUser(apiConfig.Pool, dbId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusOK, user,
	)
}
