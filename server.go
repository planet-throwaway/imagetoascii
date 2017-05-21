package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type ImageRequest struct {
	Data string `json:"data"`
}

type AsciiResponse struct {
	Data string `json:"data"`
}

// Set by the Makefile
var version string

func main() {
	port := flag.Uint("port", 8080, "port to serve on (default 8080)")
	showVersion := flag.Bool("version", false, "displays version and exits")
	flag.Parse()

	if *showVersion {
		fmt.Printf("imagetoascii %s\n", version)
		os.Exit(0)
	}

	e := echo.New()
	e.POST("/toascii", imageHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}

func imageHandler(c echo.Context) error {
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
}
