package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/helpers"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"
)

func broadvidVideosRedirectsViewallCtrl(c *gin.Context) {
	helpers.FindAll(&structs.VideoRedirectors{}, c)
}

func broadvidVideosRedirectsViewCtrl(c *gin.Context) {
	helpers.FindOne(&structs.VideoRedirector{}, c)
}

func broadvidVideosRedirectsSaveCtrl(c *gin.Context) {
	uid, err := helpers.GetUserID(c)

	if err != nil {
		c.JSON(500, err)
		return
	}

	var ad = &structs.VideoRedirector{}

	err = helpers.SaveEntity(ad, c, "Redirect has been saved", ad.Key+" has been saved", "success", 2, uid)

	if err != nil {
		c.JSON(500, err)
		return
	}

}
