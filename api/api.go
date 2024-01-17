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

	//Get memory Info
	router.GET("/api/v1/resource/memory", func(c *gin.Context) {

		memInfo, err := sysinfo.GetMemInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		//fmt.Println("Total memory: ", strconv.FormatUint(memInfo.Total/(1024*1024), 10)+" MB")
		c.JSON(200, gin.H{"memoryInfo": memInfo})

	})

	router.GET("/api/v1/resource/network", func(c *gin.Context) {

		netInfo, err := sysinfo.GetNetInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(200, gin.H{"networkInfo": netInfo})

	})

	log.Fatal(router.Run(":3000"))
}
