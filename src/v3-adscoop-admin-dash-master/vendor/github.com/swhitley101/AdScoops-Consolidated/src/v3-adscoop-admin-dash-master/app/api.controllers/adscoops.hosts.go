package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"
)

func hostsViewallCtrl(c *gin.Context) {
	var hosts structs.Hosts
	err := hosts.FindAll()

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, hosts)
}
