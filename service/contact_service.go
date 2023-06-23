package service

import (
	"errors"
	"github.com/mahfuzon/test_with_uuid/models"
	"github.com/mahfuzon/test_with_uuid/repository"
	"github.com/mahfuzon/test_with_uuid/request/contact_request"
	"github.com/mahfuzon/test_with_uuid/response/contact_response"
	"github.com/sirupsen/logrus"
)

type ContactService interface {
	Create(request contact_request.ContactCreateRequest) (contact_response.ContactResponse, error)
	Find(request contact_request.ContactFindRequest) (contact_response.ContactResponse, error)
	Update(request contact_request.ContactUpdateRequest) (contact_response.ContactResponse, error)
	Delete(request contact_request.ContactFindRequest) error
	All() ([]contact_response.ContactResponse, error)
}

type contactService struct {
	log               *logrus.Logger
	contactRepository repository.ContactRepository
}

func NewContactService(contactRepository repository.ContactRepository, log *logrus.Logger) *contactService {
	return &contactService{contactRepository: contactRepository, log: log}
}

func (service *contactService) All() ([]contact_response.ContactResponse, error) {
	service.log.Info("contactService.All")
	contacts, err := service.contactRepository.All()
	if err != nil {
		return []contact_response.ContactResponse{}, err
	}

	listContactResponse := contact_response.ConverseToListContactResponse(contacts)

	service.log.Info("success contactService.All")
	return listContactResponse, nil
}

func (service *contactService) Delete(request contact_request.ContactFindRequest) error {
	service.log.Info("contactService.Delete")
	contact, err := service.contactRepository.Find(request.Id)
	if err != nil {
		return err
	}

	err = service.contactRepository.Delete(&contact)
	if err != nil {
		return err
	}

	service.log.Info("success contactService.Delete")
	return nil
}

func (service *contactService) Update(request contact_request.ContactUpdateRequest) (contact_response.ContactResponse, error) {
	service.log.Info("contactService.Update")
	// find by id
	contact, err := service.contactRepository.Find(request.Id)
	if contact.Id == "" {
		return contact_response.ContactResponse{}, errors.New("contact not found")
	}

	if err != nil {
		return contact_response.ContactResponse{}, err
	}

	checkPhoneAndEmailAlreadyExists, err := service.contactRepository.CheckEmailAndPhoneAlreadyExists(request.Phone, request.Email)

	if len(checkPhoneAndEmailAlreadyExists) > 0 {
		for _, exist := range checkPhoneAndEmailAlreadyExists {
			if exist.Id != contact.Id {
				if exist.Phone == request.Phone || exist.Email == request.Email {
					return contact_response.ContactResponse{}, errors.New("email or phone already exists")
				}
			}

		}
	}

	contact.Name = request.Name
	contact.Email = request.Email
	contact.Phone = request.Phone
	contact.Gender = request.Gender

	err = service.contactRepository.Update(&contact)
	if err != nil {
		return contact_response.ContactResponse{}, err
	}

	contactResponse := contact_response.ConverseToContactResponse(contact)

	service.log.Info("success contactService.Update")
	return contactResponse, nil
}

func (service *contactService) Create(request contact_request.ContactCreateRequest) (contact_response.ContactResponse, error) {
	service.log.Info("contactService.Create")

	checkExists, err := service.contactRepository.CheckEmailAndPhoneAlreadyExists(request.Phone, request.Email)
	if len(checkExists) > 0 {
		return contact_response.ContactResponse{}, errors.New("email or phone already exists")
	}

	if err != nil {
		return contact_response.ContactResponse{}, err
	}

	contact := models.Contact{
		Name:   request.Name,
		Email:  request.Email,
		Phone:  request.Phone,
		Gender: request.Gender,
	}

	err = service.contactRepository.Create(&contact)
	if err != nil {
		return contact_response.ContactResponse{}, err
	}

	contactResponse := contact_response.ConverseToContactResponse(contact)

	service.log.Info("success contactService.Create")

	return contactResponse, nil
}

func (service *contactService) Find(request contact_request.ContactFindRequest) (contact_response.ContactResponse, error) {
	service.log.Info("contactService.Find")
	contact, err := service.contactRepository.Find(request.Id)
	if err != nil {
		return contact_response.ContactResponse{}, err
	}

	contactResponse := contact_response.ConverseToContactResponse(contact)

	service.log.Info("success contactService.Find")
	return contactResponse, nil
}
