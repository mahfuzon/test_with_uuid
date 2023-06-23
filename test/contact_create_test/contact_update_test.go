package user_register_test

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/test_with_uuid/libraries"
	"github.com/mahfuzon/test_with_uuid/models"
	"github.com/mahfuzon/test_with_uuid/response/api_response"
	"github.com/mahfuzon/test_with_uuid/test"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateSuccess(t *testing.T) {
	db := libraries.SetDbTest()
	log := libraries.NewLogger()
	test.TruncateTableUsers(db)

	// buat dummy contact
	contact1 := models.Contact{
		Name:   "mahfuzon",
		Email:  "mahfuzon0@gmail1.com",
		Phone:  "+62812781609901",
		Gender: "male",
	}
	err := db.Create(&contact1).Error
	assert.NoError(t, err)

	contact2 := models.Contact{
		Name:   "mahfuzon",
		Email:  "mahfuzon0@gmail2.com",
		Phone:  "+62812781609902",
		Gender: "male",
	}
	err = db.Create(&contact2).Error
	assert.NoError(t, err)

	// end buat dummy contact

	requestJsonString := `{
  "name" : "mahfuzon setelah di update",
"email" : "mahfuzon0@gmail1.com",
  "phone" : "+62812781609901",
  "gender" : "male"
}`

	contactController := test.SetupContactController(db, log)

	router := libraries.SetRouter()
	router.PUT("api/v1/contact/:id", contactController.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/contact/"+contact1.Id, strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var apiResponse api_response.ApiResponse

	err = json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}

func TestUpdateErrorUniqueEmailOrPhone(t *testing.T) {
	db := libraries.SetDbTest()
	log := libraries.NewLogger()
	test.TruncateTableUsers(db)

	// buat dummy contact

	contact1 := models.Contact{
		Name:   "mahfuzon",
		Email:  "mahfuzon0@gmail1.com",
		Phone:  "+62812781609901",
		Gender: "male",
	}
	err := db.Create(&contact1).Error
	assert.NoError(t, err)

	contact2 := models.Contact{
		Name:   "mahfuzon",
		Email:  "mahfuzon0@gmail2.com",
		Phone:  "+62812781609902",
		Gender: "male",
	}
	err = db.Create(&contact2).Error
	assert.NoError(t, err)
	// end buat dummy contact

	requestJsonString := `{
  "name" : "mahfuzon akhiar setelah update",
"email" : "mahfuzon0@gmail2.com",
  "phone" : "+62812781609901",
  "gender" : "male"
}`

	contactController := test.SetupContactController(db, log)

	router := libraries.SetRouter()
	router.PUT("api/v1/contact/:id", contactController.Update)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/contact/"+contact1.Id, strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, 400, result.StatusCode)

	body := result.Body

	responseBody, _ := io.ReadAll(body)
	var apiResponse api_response.ApiResponse

	err = json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}
