package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SirikitCloudImpl struct{}

// Configuration Resource
// (GET /configuration)
func (SirikitCloudImpl) ExtensionConfiguration(ctx echo.Context, params ExtensionConfigurationParams) error {
	config := &ExtensionConfig{
		Url: &"http://vg.no",
	}
	return ctx.JSON(http.StatusOK, config)
}

// addMedia
// (POST /intent/addMedia)
func (SirikitCloudImpl) AddMediaIntentHandling(ctx echo.Context, params AddMediaIntentHandlingParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

// playMedia
// (POST /intent/playMedia)
func (SirikitCloudImpl) PlayMediaIntentHandling(ctx echo.Context, params PlayMediaIntentHandlingParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

// updateMediaAffinity
// (POST /intent/updateMediaAffinity)
func (SirikitCloudImpl) UpdateMediaAffinityIntentHandling(ctx echo.Context, params UpdateMediaAffinityIntentHandlingParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

// playMedia
// (POST /queues/playMedia)
func (SirikitCloudImpl) PlayMediaOnQueue(ctx echo.Context, params PlayMediaOnQueueParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

// updateActivity
// (POST /queues/updateActivity)
func (SirikitCloudImpl) UpdateActivityOnQueue(ctx echo.Context, params UpdateActivityOnQueueParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

func main() {
	var myApi SirikitCloudImpl
	e := echo.New()
	RegisterHandlers(e, myApi)

	e.Logger.Fatal(e.Start(":1323"))
}
