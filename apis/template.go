package apis

import (
	"strconv"

	"github.com/labstack/echo"

	"github.com/masschaos/queenbee/models"
)

func CreateTemplate(c echo.Context) error {
	t := &models.Template{}
	err := c.Bind(t)
	if err != nil {
		return err
	}

	return t.Create()
}

func ReadTemplate(c echo.Context) error {
	t := &models.Template{}
	id := c.Param("id")
	t.ID = id
	err := t.Read()
	if err != nil {
		return err
	}

	return c.JSON(200, t)
}

func ListTemplate(c echo.Context) error {
	pageStr := c.QueryParam("p")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return err
	}

	sizeStr := c.QueryParam("s")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return err
	}

	t := &models.Template{}
	data, err := t.List(size, page)
	if err != nil {
		return err
	}

	return c.JSON(200, data)
}

func UpdateTemplate(c echo.Context) error {
	data := map[string]interface{}{}
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	t := &models.Template{}
	id := c.Param("id")
	t.ID = id

	return t.Update(data)
}

func DeleteTemplate(c echo.Context) error {
	t := &models.Template{}
	id := c.Param("id")
	t.ID = id
	err := t.Delete()
	if err != nil {
		return err
	}

	return nil
}
