package handler

import (
	"github.com/JohnKucharsky/echo_pgx/repository"
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

	product, err := repository.CreateProduct(apiConfig.Pool, productBody)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		product,
	)
}

func (apiConfig *DatabaseController) GetProducts(c echo.Context) error {
	products, err := repository.GetProduct(apiConfig.Pool)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

func (apiConfig *DatabaseController) GetOneProduct(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	dbId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := repository.GetOneProduct(apiConfig.Pool, int(dbId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (apiConfig *DatabaseController) UpdateProduct(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	dbId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var productBody serializer.ProductBody
	if err := c.Bind(&productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := repository.UpdateProduct(
		apiConfig.Pool,
		productBody,
		int(dbId),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, product)
}

func (apiConfig *DatabaseController) DeleteProduct(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	dbId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := repository.DeleteProduct(apiConfig.Pool, int(dbId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}
