package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-niom/niom-sample/pkg/logger"
	helpers "github.com/go-niom/niom-sample/pkg/helpers"
	dto "github.com/go-niom/niom-sample/src/app/dto"
)

var AppController appControllerInterface = &appController{}

type appController struct {
}

type appControllerInterface interface {
	CreateApp(*fiber.Ctx) error
	RetrieveAllApp(*fiber.Ctx) error
	RetrieveApp(*fiber.Ctx) error
	UpdateApp(*fiber.Ctx) error
	DeleteApp(*fiber.Ctx) error
}

// CreateApp func create a app.
// @Description create app
// @Summary create a app
// @Tags App
// @Accept json
// @Produce json
// @Param request body CreateAppDto true "Update App Request Body"
// @Success 200 {object} helpers.ResMessage(data=App) "Ok"
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKeyApp
// @Router /v1/app [Post]
func (d *appController) CreateApp(ctx *fiber.Ctx) error {
	log := "Controller : App, Method : CreateApp"
	logger.Info(log)

	body := new(dto.CreateAppDto)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	res, err := AppService.create(body)
	if err != nil {
		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "App Created", res)
}

// GetAllApp func get all app.
// @Description get all app
// @Summary Get all app
// @Tags App
// @Accept json
// @Produce json
// @Success 200 {object} helpers.ResMessage(data=[]App) "Ok"
// {object} helpers.SuccessResponse
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKeyApp
// @Router /v1/app [Get]
func (d *appController) RetrieveAllApp(ctx *fiber.Ctx) error {
	log := "Controller : App, Method : RetrieveAllApp"
	logger.Info(log)

	query := new(dto.QueryAppDto)
	if err := ctx.QueryParser(query); err != nil {
		return err
	}
	res, err := AppService.findAll(query)
	if err != nil {
		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "All App Retrieved", res)
}

// GetApp func get a app.
// @Description get app
// @Summary get a app
// @Tags App
// @Accept json
// @Produce json
// @Param id path string true "App ID"
// @Success 200 {object} helpers.ResMessage(data=App) "Ok"
// {object} helpers.SuccessResponse
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKeyApp
// @Router /v1/app/{id} [Get]
func (d *appController) RetrieveApp(ctx *fiber.Ctx) error {
	log := "Controller : App, Method : RetrieveApp"
	logger.Info(log)

	id := ctx.Params("id")
	res, err := AppService.findOne(id)
	if err != nil {

		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "App Retrieved", res)
}

// GetAllApp func update a app.
// @Description update app
// @Summary Update App by Id
// @Tags App
// @Accept json
// @Produce json
// @Param id path string true "App ID"
// @Param request body UpdateAppDto true "Create App Request Body"
// @Success 200 {object} helpers.ResMessage(data=App) "Ok"
// {object} helpers.SuccessResponse
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKeyApp
// @Router /v1/app [Patch]
func (d *appController) UpdateApp(ctx *fiber.Ctx) error {
	log := "Controller : App, Method : UpdateApp"
	logger.Info(log)
	id := ctx.Params("id")
	body := new(dto.UpdateAppDto)
	if err := ctx.BodyParser(body); err != nil {
		logger.Error(log, err)
		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	res, err := AppService.update(id, body)
	if err != nil {

		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "App Updated", res)
}

// DeleteApp func delete a app.
// @Description delete app
// @Summary delete a app
// @Tags App
// @Accept json
// @Produce json
// @Param id path string true "App ID"
// Success 200 "Ok"
// {object} helpers.SuccessResponse
// Failure 400,404,401,500 {object} helpers.ResMessage "Error"
// @Security ApiKeyApp
// @Router /v1/app/{id} [delete]
func (d *appController) DeleteApp(ctx *fiber.Ctx) error {
	log := "Controller : App, Method : DeleteApp"
	logger.Info(log)

	id := ctx.Params("id")
	res, err := AppService.remove(id)
	if err != nil {
		return helpers.ErrorResponse(ctx, "Error Descriptions", err)
	}
	return helpers.MsgResponse(ctx, "App Deleted", res)
}

