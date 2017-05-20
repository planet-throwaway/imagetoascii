package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type ImageRequest struct {
	Data string `json:"data"`
}

type AsciiResponse struct {
	Data string `json:"data"`
}

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		imageRequest := new(ImageRequest)
		if err := c.Bind(imageRequest); err != nil {
			return err
		}
		if imageRequest.Data == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Expected a 'data' element")
		}
		err, converted := Convert(imageRequest.Data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Problem with the image data provided")
		} else {
			response := AsciiResponse{Data: converted}
			return c.JSON(http.StatusOK, response)
		}
	})
	e.Logger.Fatal(e.Start(":8080"))
}
