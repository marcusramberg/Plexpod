package main

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

type sirikitCloudImpl struct{}

// Configuration Resource
// (GET /configuration)
func (sirikitCloudImpl) ExtensionConfiguration(ctx echo.Context, params ExtensionConfigurationParams) error {
	config := &ExtensionConfig{
		Url: &ctx.Request().RequestURI,
	}
	ctx.Response().Header().Set("x-applecloudextension-session-id", ctx.Request().Header.Get("x-applecloudextension-session-id"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"alg": "ES256",
		"kid": "secret",
	})

	secret := []byte("Secret")
	signature, err := token.SignedString(secret)
	if err != nil {
		return err
	}
	ctx.Response().Header().Set("token", signature)
	return ctx.JSON(http.StatusOK, config)
}

// addMedia
// (POST /intent/addMedia)
func (sirikitCloudImpl) AddMediaIntentHandling(ctx echo.Context, params AddMediaIntentHandlingParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

// playMedia
// (POST /intent/playMedia)
func (sirikitCloudImpl) PlayMediaIntentHandling(ctx echo.Context, params PlayMediaIntentHandlingParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

// updateMediaAffinity
// (POST /intent/updateMediaAffinity)
func (sirikitCloudImpl) UpdateMediaAffinityIntentHandling(ctx echo.Context, params UpdateMediaAffinityIntentHandlingParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

// playMedia
// (POST /queues/playMedia)
func (sirikitCloudImpl) PlayMediaOnQueue(ctx echo.Context, params PlayMediaOnQueueParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

// updateActivity
// (POST /queues/updateActivity)
func (sirikitCloudImpl) UpdateActivityOnQueue(ctx echo.Context, params UpdateActivityOnQueueParams) error {
	return ctx.JSON(http.StatusOK, "OK")
}

func main() {
	var myAPI sirikitCloudImpl
	e := echo.New()
	e.Debug = true
	p := prometheus.NewPrometheus("plexpod", nil)
	p.Use(e)

	RegisterHandlers(e, myAPI)

	e.Logger.Fatal(e.Start(":1323"))
}
