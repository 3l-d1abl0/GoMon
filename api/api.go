package api

import (
	"log"
	"net/http"

	"GoMon/sysinfo"

	"github.com/gin-gonic/gin"
)

func Setup() {

	//Router setup
	router := gin.Default()

	//GET memory Info
	router.GET("/api/v1/resource/memory", func(c *gin.Context) {

		memInfo, err := sysinfo.GetMemInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		//fmt.Println("Total memory: ", strconv.FormatUint(memInfo.Total/(1024*1024), 10)+" MB")
		c.JSON(200, gin.H{"memoryInfo": memInfo})

	})

	//GET network Info
	router.GET("/api/v1/resource/network", func(c *gin.Context) {

		netInfo, err := sysinfo.GetNetInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(200, gin.H{"networkInfo": netInfo})

	})

	//GET CPU info
	router.GET("api/v1/resource/cpu", func(c *gin.Context) {

		cpuInfo, err := sysinfo.GetCPUInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(200, gin.H{"cpuInfo": cpuInfo})

	})

	//GET host Info
	router.GET("/api/v1/resource/host", func(c *gin.Context) {

		hostInfo, err := sysinfo.GetHostInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(200, gin.H{"hostInfo": hostInfo})

	})

	//GET System load
	router.GET("/api/v1/resource/load", func(c *gin.Context) {

		loadInfo, err := sysinfo.GetLoadInfo()
		if err != nil {
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
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//2. Network Info
		netInfo, err := sysinfo.GetNetInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//3. Cpu Info
		cpuInfo, err := sysinfo.GetCPUInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//4. host Info
		hostInfo, err := sysinfo.GetHostInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//5. loadInfo
		loadInfo, err := sysinfo.GetLoadInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		//Combine All
		c.JSON(200, gin.H{"memoryInfo": memInfo,
			"networkInfo": netInfo,
			"cpuInfo":     cpuInfo,
			"hostInfo":    hostInfo,
			"loadInfo":    loadInfo})

	})

	log.Fatal(router.Run(":3000"))

}
