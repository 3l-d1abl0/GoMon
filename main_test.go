package main

import (
	route "GoMon/api"
	commondata "GoMon/commonData"
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

func TestRouteMemory(t *testing.T) {

	r := SetUpRouter()
	req, _ := http.NewRequest("GET", "/api/v1/resource/memory", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	//assert.Equal(t, mockResponse, string(responseData))
	//1. Test for 200
	assert.Equal(t, http.StatusOK, w.Code)

	type MemoryInfo struct {
		Total       uint64  `json:"total"`
		Available   uint64  `json:"available"`
		Used        uint64  `json:"used"`
		UsedPercent float64 `json:"usedPercent"`
		Free        uint64  `json:"free"`
		Active      uint64  `json:"active"`
		Inactive    uint64  `json:"inactive"`
		SwapTotal   uint64  `json:"swapTotal"`
		SwapFree    uint64  `json:"swapFree"`
	}

	type MemoryInfoResponse struct {
		MemoryInfo MemoryInfo `json:"memoryInfo"`
	}

	var memoryInfoContainer MemoryInfoResponse
	err := json.Unmarshal([]byte(responseData), &memoryInfoContainer)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}
}

func TestRouteNetwork(t *testing.T) {

	r := SetUpRouter()
	req, _ := http.NewRequest("GET", "/api/v1/resource/network", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	//1. Test for 200
	assert.Equal(t, http.StatusOK, w.Code)

	var networkInfo commondata.NetworkInfo

	err := json.Unmarshal([]byte(responseData), &networkInfo)
	if err != nil {
		t.Errorf("Error decoding networkInfo JSON: %s", err)
	}
}

func TestRouteCpu(t *testing.T) {

	r := SetUpRouter()
	req, _ := http.NewRequest("GET", "/api/v1/resource/cpu", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)

	var cpuList commondata.CPUInfoList

	err := json.Unmarshal([]byte(responseData), &cpuList)
	if err != nil {
		t.Errorf("Error decoding cpuInfo JSON: %s", err)
	}

}

func TestRouteHost(t *testing.T) {

	r := SetUpRouter()
	req, _ := http.NewRequest("GET", "/api/v1/resource/host", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)

	var hostInfo commondata.HostInfo

	err := json.Unmarshal([]byte(responseData), &hostInfo)
	if err != nil {
		t.Errorf("Error decoding hostInfo JSON: %s", err)
	}

}
