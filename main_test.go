package main

import (
	route "GoMon/api"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gotest.tools/assert"
)

func SetUpRouter() *gin.Engine {

	var logger = logrus.New()

	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()
	route.Setup(app, logger)

	return app
}

func TestHomepage(t *testing.T) {

	r := SetUpRouter()
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//Test for Permnent redirect
	assert.Equal(t, 301, w.Code)
}

func isJSON(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
func TestApiV1Home(t *testing.T) {

	r := SetUpRouter()
	req, _ := http.NewRequest("GET", "/api/v1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	//assert.Equal(t, mockResponse, string(responseData))
	//Test for 200
	assert.Equal(t, http.StatusOK, w.Code)

	type APIRoute struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}

	type APIRoutes struct {
		APIRoutes []APIRoute `json:"apiRoutes"`
	}

	var apiRoutesData APIRoutes
	if err := json.Unmarshal([]byte(string(responseData)), &apiRoutesData); err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
}
