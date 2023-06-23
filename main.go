package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mahfuzon/test_with_uuid/controller"
	"github.com/mahfuzon/test_with_uuid/libraries"
	"github.com/mahfuzon/test_with_uuid/repository"
	"github.com/mahfuzon/test_with_uuid/service"
	"github.com/sirupsen/logrus"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db := libraries.SetDb()
	log := libraries.NewLogger()
	contactRepository := repository.NewContactRepository(db, log)
	contactService := service.NewContactService(contactRepository, log)
	contactController := controller.NewContactController(contactService, log)
	router := libraries.SetRouter()

	router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))

	api := router.Group("/api")

	apiV1 := api.Group("/v1")

	apiV1Auth := apiV1.Group("/contact")
	apiV1Auth.POST("", contactController.Create)
	apiV1Auth.PUT("/:id", contactController.Update)
	apiV1Auth.DELETE("/:id", contactController.Delete)
	apiV1Auth.GET("/:id", contactController.Find)
	apiV1Auth.GET("", contactController.All)

	router.Logger.Fatal(router.Start(":8000"))
}
