package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/helpers"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"
)

func broadvidWhitelistUrlGroupsViewallCtrl(c *gin.Context) {
	helpers.FindAll(&structs.AdscoopWhitelists{}, c)
}

func broadvidWhitelistUrlGroupsViewCtrl(c *gin.Context) {
	helpers.FindOne(&structs.AdscoopWhitelist{}, c)
}

func broadvidWhitelistUrlGroupsSaveCtrl(c *gin.Context) {
	uid, err := helpers.GetUserID(c)

	if err != nil {
		c.JSON(500, err)
		return
	}

	var ad = &structs.AdscoopWhitelist{}

	err = helpers.SaveEntity(ad, c, "Ad Url Group has been saved", ad.Name+" has been saved", "success", 2, uid)

	if err != nil {
		c.JSON(500, err)
		return
	}

}
