package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/helpers"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"
)

func broadvidAdsViewallCtrl(c *gin.Context) {
	helpers.FindAll(&structs.Ads{}, c)
}

func broadvidAdsViewCtrl(c *gin.Context) {
	helpers.FindOne(&structs.Ad{}, c)
}

func broadvidAdsSaveCtrl(c *gin.Context) {
	uid, err := helpers.GetUserID(c)

	if err != nil {
		c.JSON(500, err)
		return
	}
	var ad = &structs.Ad{}
	err = helpers.SaveEntity(ad, c, "Ad has been saved", ad.Label+" has been saved", "success", 2, uid)

	if err != nil {
		c.JSON(500, err)
		return
	}

}
