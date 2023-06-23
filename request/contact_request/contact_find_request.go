package contact_request

type ContactFindRequest struct {
	Id string `json:"name" validate:"required"`
}
