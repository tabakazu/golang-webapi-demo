// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/tabakazu/go-webapp/application/service"
	"github.com/tabakazu/go-webapp/external/datastore"
	"github.com/tabakazu/go-webapp/external/datastore/repository"
	"github.com/tabakazu/go-webapp/interfaces/webapi"
	"github.com/tabakazu/go-webapp/interfaces/webapi/controller"
	"github.com/tabakazu/go-webapp/interfaces/webapi/generator"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeServer(dbConn *gorm.DB) *webapi.Server {
	userAccountRepository := repository.NewUserAccountRepository(dbConn)
	registerUserAccount := service.NewUserAccountRegisterService(userAccountRepository)
	userTokenGenerator := generator.NewUserTokenGenerator()
	loginUserAccount := service.NewUserAccountLoginService(userAccountRepository, userTokenGenerator)
	showUserAccount := service.NewUserAccountShowService(userAccountRepository)
	userAccountController := controller.NewUserAccountController(registerUserAccount, loginUserAccount, showUserAccount)
	server := webapi.NewServer(userAccountController)
	return server
}

// wire.go:

func InitializeDB() (*gorm.DB, func()) {
	dbConfig := datastore.NewDBConfig()
	db, dbClose := datastore.NewConnection(dbConfig)
	return db, dbClose
}
