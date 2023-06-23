package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/test_with_uuid/helper"
	"github.com/mahfuzon/test_with_uuid/request/contact_request"
	"github.com/mahfuzon/test_with_uuid/response/api_response"
	"github.com/mahfuzon/test_with_uuid/service"
	"github.com/sirupsen/logrus"
)

type ContactController struct {
	log            *logrus.Logger
	contactService service.ContactService
}

func NewContactController(contactService service.ContactService, log *logrus.Logger) *ContactController {
	return &ContactController{contactService: contactService, log: log}
}

func (controller *ContactController) Create(ctx echo.Context) error {
	controller.log.Info("controller.Create")
	request := contact_request.ContactCreateRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		controller.log.Info(err.Error())
		errorResponse := api_response.ConverseToErrorResponse("failed create contact", err.Error())
		return ctx.JSON(500, errorResponse)
	}

	controller.log.WithFields(logrus.Fields{
		"requestBody": request,
	}).Info("request body")

	err = ctx.Validate(&request)
	if err != nil {
		controller.log.Info(err.Error())
		errorValidation := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		errorResponse := api_response.ConverseToErrorResponse("failed create contact", errorValidation)
		return ctx.JSON(422, errorResponse)
	}

	contactResponse, err := controller.contactService.Create(request)
	if err != nil {
		controller.log.Info(err.Error())
		errorResponse := api_response.ConverseToErrorResponse("failed create contact", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	successResponse := api_response.ConverseToSuccessResponse("success create contact", contactResponse)

	controller.log.Info("success controller.Create")
	return ctx.JSON(200, successResponse)
}

func (controller *ContactController) Update(ctx echo.Context) error {
	controller.log.Info("controller.Update")
	request := contact_request.ContactUpdateRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		controller.log.Info(err.Error())
		errorResponse := api_response.ConverseToErrorResponse("failed update contact", err.Error())
		return ctx.JSON(500, errorResponse)
	}

	controller.log.WithFields(logrus.Fields{
		"requestBody": request,
	}).Info("request body")

	err = ctx.Validate(&request)
	if err != nil {
		controller.log.Info(err.Error())
		errorValidation := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		errorResponse := api_response.ConverseToErrorResponse("failed update contact", errorValidation)
		return ctx.JSON(422, errorResponse)
	}

	contactResponse, err := controller.contactService.Update(request)
	if err != nil {
		controller.log.Info(err.Error())
		errorResponse := api_response.ConverseToErrorResponse("failed update contact", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	successResponse := api_response.ConverseToSuccessResponse("success update contact", contactResponse)

	controller.log.Info("success controller.Update")
	return ctx.JSON(200, successResponse)
}

func (controller *ContactController) Find(ctx echo.Context) error {
	controller.log.Info("controller.Find")
	request := contact_request.ContactFindRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		controller.log.Info(err.Error())
		errorResponse := api_response.ConverseToErrorResponse("failed find contact", err.Error())
		return ctx.JSON(500, errorResponse)
	}

	controller.log.WithFields(logrus.Fields{
		"requestBody": request,
	}).Info("request body")

	err = ctx.Validate(&request)
	if err != nil {
		controller.log.Info(err.Error())
		errorValidation := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		errorResponse := api_response.ConverseToErrorResponse("failed find contact", errorValidation)
		return ctx.JSON(422, errorResponse)
	}

	contactResponse, err := controller.contactService.Find(request)
	if err != nil {
		controller.log.Info(err.Error())
		errorResponse := api_response.ConverseToErrorResponse("failed find contact", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	successResponse := api_response.ConverseToSuccessResponse("success find contact", contactResponse)

	controller.log.Info("success controller.Find")
	return ctx.JSON(200, successResponse)
}

func (controller *ContactController) Delete(ctx echo.Context) error {
	controller.log.Info("controller.Delete")
	request := contact_request.ContactFindRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		controller.log.Info(err.Error())
		errorResponse := api_response.ConverseToErrorResponse("failed delete contact", err.Error())
		return ctx.JSON(500, errorResponse)
	}

	controller.log.WithFields(logrus.Fields{
		"requestBody": request,
	}).Info("request body")

	err = ctx.Validate(&request)
	if err != nil {
		controller.log.Info(err.Error())
		errorValidation := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		errorResponse := api_response.ConverseToErrorResponse("failed delete contact", errorValidation)
		return ctx.JSON(422, errorResponse)
	}

	err = controller.contactService.Delete(request)
	if err != nil {
		controller.log.Info(err.Error())
		errorResponse := api_response.ConverseToErrorResponse("failed delete contact", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	successResponse := api_response.ConverseToSuccessResponse("success delete contact", nil)

	controller.log.Info("success controller.Delete")
	return ctx.JSON(200, successResponse)
}

func (controller *ContactController) All(ctx echo.Context) error {
	controller.log.Info("controller.All")

	listContactResponse, err := controller.contactService.All()
	if err != nil {
		controller.log.Info(err.Error())
		errorResponse := api_response.ConverseToErrorResponse("failed get all contact", err.Error())
		return ctx.JSON(400, errorResponse)
	}

	successResponse := api_response.ConverseToSuccessResponse("success get all contact", listContactResponse)

	controller.log.Info("success controller.All")
	return ctx.JSON(200, successResponse)
}
