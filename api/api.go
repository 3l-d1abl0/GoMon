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

	router.GET("/api/v1/resource/memory", func(c *gin.Context) {

		memInfo, err := sysinfo.GetMemInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		//fmt.Println("Total memory: ", strconv.FormatUint(memInfo.Total/(1024*1024), 10)+" MB")
		c.JSON(200, memInfo)

	})

	log.Fatal(router.Run(":3000"))
}
