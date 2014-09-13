package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
	"math/rand"
)

func isKeyValid(key string) bool {
	var valid bool

	switch key {
	case config.AndroidKey, config.DevKey, config.FFOSKey, config.TizenKey, config.WebKey:
		valid = true
	default:
		valid = false
	}

	return valid
}

func randInt(max int) int {
	return rand.Intn(max)
}

func pingHandler(ctx context.Context) error {
	return goweb.API.RespondWithData(ctx, objx.MSI("message", "pong!"))
}

func AowQuote(ctx context.Context) error {
	index := randInt(len(quotes))
	quote := quotes[index]
	key := ctx.PathParams().Get("apiKey").Str()

	if valid := isKeyValid(key); valid == true {
		return goweb.API.RespondWithData(ctx, objx.MSI("quote", quote))
	}
	return goweb.API.RespondWithError(ctx, 400, "Access Denied: Invalid API Key")
}
