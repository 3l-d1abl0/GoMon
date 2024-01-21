package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/host"
)

func AddHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		hostInfo, _ := host.Info()
		c.Writer.Header().Set("Hostname", fmt.Sprintf("%v", hostInfo.Hostname))
		c.Writer.Header().Set("Hostid", fmt.Sprintf("%v", hostInfo.HostID))
		c.Next()
	}
}
