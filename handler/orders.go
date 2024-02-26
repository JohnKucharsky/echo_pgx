package handler

import (
	"fmt"
	"github.com/JohnKucharsky/echo_pgx/models"
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

	return c.JSON(
		http.StatusCreated, orderBody,
	)
}

func (apiConfig *DatabaseController) GetOrders(c echo.Context) error {

	var orders []models.Order

	return c.JSON(
		http.StatusCreated, orders,
	)
}

func (apiConfig *DatabaseController) UpdateOrder(c echo.Context) error {
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

	var orderBody serializer.OrderBody
	if err := c.Bind(&orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var order = serializer.OrderBodyToOrder(orderBody)
	fmt.Println(dbId)

	return c.JSON(
		http.StatusCreated, order,
	)
}

func (apiConfig *DatabaseController) DeleteOrder(c echo.Context) error {
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

	var order models.Order
	fmt.Println(dbId)
	return c.JSON(http.StatusOK, order)
}
