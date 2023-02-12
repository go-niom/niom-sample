package app

import (
	"fmt"

	"github.com/go-niom/niom-sample/pkg/logger"
)

var AppService appServiceInterface = &appService{}

type appService struct {
}

type appServiceInterface interface {
	create(interface{}) (string, interface{})
	findAll(interface{}) (string, interface{})
	findOne(string) (string, interface{})
	update(string, interface{}) (string, interface{})
	remove(string) (string, interface{})
}

func (a *appService) create(data interface{}) (string, interface{}) {
	logger.Info("Service : App, Method : create")
	return "This action adds a new app", nil
}

func (a *appService) findAll(data interface{}) (string, interface{}) {
	logger.Info("Service : App, Method : findAll")
	return "This action find all app", nil
}

func (a *appService) findOne(id string) (string, interface{}) {
	logger.Info("Service : App, Method : findOne")
	return fmt.Sprintf("This action returns a #%s app", id), nil
}

func (a *appService) update(id string, data interface{}) (string, interface{}) {
	logger.Info("Service : App, Method : update")
	return fmt.Sprintf("This action updates a #%s user", id), nil
}

func (a *appService) remove(id string) (string, interface{}) {
	logger.Info("Service : App, Method : remove")
	return fmt.Sprintf("This action removes a #%s user", id), nil
}

