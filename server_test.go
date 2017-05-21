package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestConvertImageGoodData(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/toascii",
		strings.NewReader("{\"data\": \"R0lGODdhAgACAJEAAAAAAP///wAAAAAAACH5BAkAAAIALAAAAAACAAIAAAIChFEAOw==\"}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, imageHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"data\":\"$$\\n\"}", rec.Body.String())
	}
}

func TestConvertImageNoData(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/toascii", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.Error(t, imageHandler(c)) {
		assert.Equal(t, "", rec.Body.String())
	}
}

func TestConvertImageEmptyData(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/toascii", strings.NewReader("{}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.Error(t, imageHandler(c)) {
		assert.Equal(t, "", rec.Body.String())
	}
}

func TestConvertImageBogusData(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/toascii", strings.NewReader("{\"data\": \"baddata\"}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.Error(t, imageHandler(c)) {
		assert.Equal(t, "", rec.Body.String())
	}
}
