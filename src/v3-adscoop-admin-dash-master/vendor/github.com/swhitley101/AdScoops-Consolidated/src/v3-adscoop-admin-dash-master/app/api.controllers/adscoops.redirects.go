package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/adscoops.caches"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/helpers"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"
)

func redirectsViewallCtrl(c *gin.Context) {
	helpers.FindAll(&structs.Redirects{}, c)
}

func redirectsViewCtrl(c *gin.Context) {
	helpers.FindOne(&structs.Redirect{}, c)
}

func redirectsSaveCtrl(c *gin.Context) {

	uid, err := helpers.GetUserID(c)

	if err != nil {
		c.JSON(500, err)
		return
	}

	var redirect = &structs.Redirect{}

	err = helpers.SaveEntity(redirect, c, "Redirect has been saved", redirect.Name+" has been saved", "success", 1, uid)

	if err != nil {
		c.JSON(500, err)
		return
	}

	adscoopsCaches.UpdateRedirect(redirect)
}
