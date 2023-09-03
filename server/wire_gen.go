// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"context"
	"github.com/xmdhs/authlib-skin/config"
	"github.com/xmdhs/authlib-skin/server/route"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeRoute(ctx context.Context, c config.Config) (*http.Server, func(), error) {
	handler := ProvideSlog(c)
	logger := NewSlog(handler)
	db, cleanup, err := ProvideDB(c)
	if err != nil {
		return nil, nil, err
	}
	client, cleanup2 := ProvideEnt(ctx, db, c, logger)
	validate := ProvideValidate()
	node, err := ProvideSnowflake(c)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	router, err := route.NewRoute(logger, client, validate, node, c)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, cleanup3 := NewServer(c, logger, router)
	return server, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
