package handler

import (
	"fmt"
	"github.com/JohnKucharsky/echo_pgx/models"
	"github.com/JohnKucharsky/echo_pgx/serializer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) CreateProduct(c echo.Context) error {
	var productBody serializer.ProductBody
	if err := c.Bind(&productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		productBody,
	)
}

func (apiConfig *DatabaseController) GetProducts(c echo.Context) error {

	return c.NoContent(
		http.StatusOK,
	)
}

func (apiConfig *DatabaseController) GetOneProduct(c echo.Context) error {
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

	var product models.Product

	fmt.Println(dbId)
	return c.JSON(
		http.StatusOK, product,
	)
}

func (apiConfig *DatabaseController) UpdateProduct(c echo.Context) error {
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

	var productBody serializer.ProductBody
	if err := c.Bind(&productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var product = serializer.ProductBodyToProduct(productBody)

	fmt.Println(dbId)
	return c.JSON(
		http.StatusCreated, product,
	)
}

func (apiConfig *DatabaseController) DeleteProduct(c echo.Context) error {
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

	var product models.Product

	fmt.Println(dbId)

	return c.JSON(http.StatusOK, product)
}
