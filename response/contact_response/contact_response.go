package contact_response

import "github.com/mahfuzon/test_with_uuid/models"

type ContactResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ConverseToContactResponse(contact models.Contact) ContactResponse {
	contactResponse := ContactResponse{
		Id:     contact.Id,
		Name:   contact.Name,
		Email:  contact.Email,
		Phone:  contact.Phone,
		Gender: contact.Gender,
	}

	return contactResponse
}

func ConverseToListContactResponse(contacts []models.Contact) []ContactResponse {
	var listContactResponse []ContactResponse
	for _, contact := range contacts {
		contactResponse := ConverseToContactResponse(contact)
		listContactResponse = append(listContactResponse, contactResponse)
	}

	return listContactResponse
}
