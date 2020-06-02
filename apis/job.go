package apis

import (
	"github.com/labstack/echo"

	"github.com/masschaos/queenbee/types"
)

func GetJob(c echo.Context) error {
	t := types.Template{}
	err := c.Bind(&t)
	if err != nil {
		return err
	}

	return nil
}
