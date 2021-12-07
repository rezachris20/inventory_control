//go:build wireinject
// +build wireinject

package main

import (
	"backend-simple-pos/app"
	"backend-simple-pos/auth"
	"backend-simple-pos/handler"
	"backend-simple-pos/role"
	"backend-simple-pos/user"
	"github.com/google/wire"
)

var userSer = wire.NewSet(
	user.NewRepository,
	wire.Bind(new(user.Repository), new(*user.UserRepository)),
)

var userServiceSet = wire.NewSet(
	user.NewService,
	wire.Bind(new(user.Service), new(*user.UserService)),
)

var authSet = wire.NewSet(
	auth.NewService,
	wire.Bind(new(auth.Service), new(*auth.JwtService)))

func InitializeUserHandler() *handler.UserHandler {
	wire.Build(authSet, userSer, app.NewDB, userServiceSet, handler.NewUserHandler)
	return nil
}

var roleRepositorySet = wire.NewSet(
	role.NewRepository,
	wire.Bind(new(role.Repository), new(*role.RoleRepository)))

var roleServiceSet = wire.NewSet(
	role.NewService,
	wire.Bind(new(role.Service), new(*role.RoleService)),
)

func InitializeRoleHandler() *handler.RoleHandler {
	wire.Build(
		app.NewDB,
		roleRepositorySet,
		roleServiceSet,
		handler.NewRoleHandler,
	)
	return nil
}
