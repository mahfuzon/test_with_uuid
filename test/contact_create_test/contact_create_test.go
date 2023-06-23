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

func TestCreateSuccess(t *testing.T) {
	db := libraries.SetDbTest()
	log := libraries.NewLogger()
	test.TruncateTableUsers(db)

	// buat dummy contact
	contact := models.Contact{
		Name:   "mahfuzon",
		Email:  "mahfuzon0@gmail.com",
		Phone:  "+6281278160990",
		Gender: "male",
	}
	err := db.Create(&contact).Error
	if err != nil {
		assert.NoError(t, err)
	}

	// end buat dummy contact

	requestJsonString := `{
  "name" : "mahfuzon akhiar",
"email" : "mahfuzon0@gmail1.com",
  "phone" : "+62812781609901",
  "gender" : "male"
}`

	contactController := test.SetupContactController(db, log)

	router := libraries.SetRouter()
	router.POST("api/v1/contact", contactController.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/contact", strings.NewReader(requestJsonString))
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

func TestCreateErrorUniqueEmailOrPhone(t *testing.T) {
	db := libraries.SetDbTest()
	log := libraries.NewLogger()
	test.TruncateTableUsers(db)

	// buat dummy contact
	contact := models.Contact{
		Name:   "mahfuzon",
		Email:  "mahfuzon0@gmail.com",
		Phone:  "+6281278160990",
		Gender: "male",
	}
	err := db.Create(&contact).Error
	if err != nil {
		assert.NoError(t, err)
	}

	// end buat dummy contact

	requestJsonString := `{
  "name" : "mahfuzon akhiar",
"email" : "mahfuzon0@gmail.com",
  "phone" : "+6281278160990",
  "gender" : "male"
}`

	contactController := test.SetupContactController(db, log)

	router := libraries.SetRouter()
	router.POST("api/v1/contact", contactController.Create)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/contact", strings.NewReader(requestJsonString))
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
