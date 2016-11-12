package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/adonnetwork"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/helpers"
)

func adonNetworkCampaignsViewAllCtrl(c *gin.Context) {
	helpers.FindAll(&adonnetwork.Campaigns{}, c)
}
