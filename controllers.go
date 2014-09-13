package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
)

func pingHandler(ctx context.Context) error {
	return goweb.API.RespondWithData(ctx, objx.MSI("message", "pong!"))
}

func AowQuote(ctx context.Context) error {
	return goweb.API.RespondWithData(ctx, objx.MSI("quote", "NYI"))
}
