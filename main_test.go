package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	params := &ExtensionConfigurationParams{XApplecloudextensionSessionId: "dummy"}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &SirikitCloudImpl{}
	// Assertions
	if assert.NoError(t, h.ExtensionConfiguration(c, *params)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "\"url\"")
	}

}
