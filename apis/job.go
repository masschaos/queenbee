package apis

import (
	"github.com/labstack/echo"

	"github.com/masschaos/queenbee/models"
	"github.com/masschaos/queenbee/types"
)

func CreateJob(c echo.Context) error {
	j := &models.Job{}
	err := c.Bind(j)
	if err != nil {
		return err
	}

	return j.Create()
}

func ReadJob(c echo.Context) error {
	j := &models.Job{}

	err := j.Read()
	if err != nil {
		return err
	}

	return c.JSON(200, j)
}

// TODO
func CreateJobResult(c echo.Context) error {
	jr := types.JobResult{}
	err := c.Bind(&jr)
	if err != nil {
		return err
	}

	return nil
}
