package handler

import (
	"github.com/JohnKucharsky/echo_pgx/repository"
	"github.com/JohnKucharsky/echo_pgx/serializer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) CreateOrder(c echo.Context) error {
	var orderBody serializer.OrderBody
	if err := c.Bind(&orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order, err := repository.CreateOrder(apiConfig.Pool, orderBody)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusCreated, order,
	)
}

func (apiConfig *DatabaseController) GetOrders(c echo.Context) error {

	orders, err := repository.GetOrders(apiConfig.Pool)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, orders)
}

func (apiConfig *DatabaseController) GetOneOrder(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	dbId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order, err := repository.GetOneOrder(apiConfig.Pool, uint(dbId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func (apiConfig *DatabaseController) UpdateOrder(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	dbId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var orderBody serializer.OrderBody
	if err := c.Bind(&orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order, err := repository.UpdateOrder(apiConfig.Pool, orderBody, int(dbId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, order)
}

func (apiConfig *DatabaseController) DeleteOrder(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	dbId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	deletedId, err := repository.DeleteOrder(apiConfig.Pool, int(dbId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, deletedId)
}
