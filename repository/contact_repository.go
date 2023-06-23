package repository

import (
	"github.com/mahfuzon/test_with_uuid/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *models.Contact) error
	Find(id string) (models.Contact, error)
	CheckEmailAndPhoneAlreadyExists(phone, email string) ([]models.Contact, error)
	Update(contact *models.Contact) error
	Delete(contact *models.Contact) error
	All() ([]models.Contact, error)
}

type contactRepository struct {
	log *logrus.Logger
	db  *gorm.DB
}

func NewContactRepository(db *gorm.DB, log *logrus.Logger) *contactRepository {
	return &contactRepository{
		log: log,
		db:  db,
	}
}

func (repository *contactRepository) All() ([]models.Contact, error) {
	repository.log.Info("masuk contact repository All")
	var contacts []models.Contact
	err := repository.db.Find(&contacts).Error
	if err != nil {
		return []models.Contact{}, err
	}

	repository.log.Info("selesai contact repository Delete")
	return contacts, nil
}

func (repository *contactRepository) Delete(contact *models.Contact) error {
	repository.log.Info("masuk contact repository Delete")
	err := repository.db.Delete(contact).Error
	if err != nil {
		return err
	}

	repository.log.Info("selesai contact repository Delete")
	return nil
}

func (repository *contactRepository) Update(contact *models.Contact) error {
	repository.log.Info("masuk contact repository Update")
	err := repository.db.Save(contact).Error
	if err != nil {
		return err
	}

	repository.log.Info("selesai contact repository Update")
	return nil
}

func (repository *contactRepository) Create(contact *models.Contact) error {
	repository.log.Info("masuk ke contact repository create")
	err := repository.db.Create(contact).Error
	if err != nil {
		return err
	}

	repository.log.Info("success contactRepository.Create")
	return nil
}

func (repository *contactRepository) Find(id string) (models.Contact, error) {
	repository.log.Info("masuk ke contact repository Find")
	contact := models.Contact{}
	err := repository.db.First(&contact, "id = ?", id).Error
	if err != nil {
		return contact, err
	}

	repository.log.Info("success find by id")
	return contact, nil
}

func (repository *contactRepository) CheckEmailAndPhoneAlreadyExists(phone, email string) ([]models.Contact, error) {
	repository.log.Info("contactRepository.CheckEmailAndPhoneAlreadyExists")
	var contact []models.Contact
	err := repository.db.Where("email = ?", email).Or("phone = ?", phone).Find(&contact).Error
	if err != nil {
		return []models.Contact{}, err
	}

	repository.log.Info("success contactRepository.CheckEmailAndPhoneAlreadyExists")
	return contact, nil
}
