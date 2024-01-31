package api

import (
	"encoding/json"
	"net/http"

	"GoMon/sysinfo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Setup(router *gin.Engine, logger *logrus.Logger) {

	//GET memory Info
	router.GET("/api/v1/resource/memory", func(c *gin.Context) {

		memInfo, err := sysinfo.GetMemInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		//fmt.Println("Total memory: ", strconv.FormatUint(memInfo.Total/(1024*1024), 10)+" MB")
		c.JSON(200, gin.H{"memoryInfo": memInfo})

	})

	//GET network Info
	router.GET("/api/v1/resource/network", func(c *gin.Context) {

		netInfo, err := sysinfo.GetNetInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(200, gin.H{"networkInfo": netInfo})

	})

	//GET CPU info
	router.GET("api/v1/resource/cpu", func(c *gin.Context) {

		cpuInfo, err := sysinfo.GetCPUInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(200, gin.H{"cpuInfo": cpuInfo})

	})

	//GET host Info
	router.GET("/api/v1/resource/host", func(c *gin.Context) {

		hostInfo, err := sysinfo.GetHostInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(200, gin.H{"hostInfo": hostInfo})

	})

	//GET System load
	router.GET("/api/v1/resource/load", func(c *gin.Context) {

		loadInfo, err := sysinfo.GetLoadInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(200, gin.H{"loadInfo": loadInfo})
	})

	//GET All resources
	router.GET("/api/v1/resource/all", func(c *gin.Context) {

		c.Redirect(http.StatusMovedPermanently, "/api/v1/resource/")

	})

	router.GET("/api/v1/resource/", func(c *gin.Context) {

		//1. Memory info
		memInfo, err := sysinfo.GetMemInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//2. Network Info
		netInfo, err := sysinfo.GetNetInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//3. Cpu Info
		cpuInfo, err := sysinfo.GetCPUInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//4. host Info
		hostInfo, err := sysinfo.GetHostInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//5. loadInfo
		loadInfo, err := sysinfo.GetLoadInfo()
		if err != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//Combine All
		c.JSON(200, gin.H{"memoryInfo": memInfo,
			"networkInfo": netInfo,
			"cpuInfo":     cpuInfo,
			"hostInfo":    hostInfo,
			"loadInfo":    loadInfo})

	})

	//GET all routes
	router.GET("/api/v1", func(c *gin.Context) {

		type Routes struct {
			Method string `json:"method"`
			Path   string `json:"path"`
		}

		var routesInfo []Routes
		for _, item := range router.Routes() {
			routesInfo = append(routesInfo, Routes{Method: item.Method, Path: item.Path})
		}

		routesInfoJson, errMarshal := json.Marshal(routesInfo)
		if errMarshal != nil {
			logger.Error("Path: ", c.Request.URL.Path, "Error: ", errMarshal.Error())
			c.JSON(http.StatusInternalServerError, "Error Processing routes")
		}

		//routesInfoJson - bytes
		c.JSON(200, gin.H{"apiRoutes": string(routesInfoJson)})

		//c.JSON(200, gin.H{"apiRoutes": json.RawMessage(string(routesInfoJson))})
	})

	//REDIRECT to /api/v1
	router.GET("/", func(c *gin.Context) {

		c.Redirect(http.StatusMovedPermanently, "/api/v1")

	})

}
