package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
	"math/rand"
)

func randInt(max int) int {
	return rand.Intn(max)
}

func pingHandler(ctx context.Context) error {
	return goweb.API.RespondWithData(ctx, objx.MSI("message", "pong!"))
}

func AowQuote(ctx context.Context) error {
	index := randInt(len(quotes))
	quote := quotes[index]
	return goweb.API.RespondWithData(ctx, objx.MSI("quote", quote))
}
