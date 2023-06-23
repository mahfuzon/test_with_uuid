package test

import (
	"github.com/mahfuzon/test_with_uuid/controller"
	"github.com/mahfuzon/test_with_uuid/repository"
	"github.com/mahfuzon/test_with_uuid/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupContactController(db *gorm.DB, log *logrus.Logger) *controller.ContactController {
	contactRepository := repository.NewContactRepository(db, log)
	contactService := service.NewContactService(contactRepository, log)
	contactController := controller.NewContactController(contactService, log)
	return contactController
}
