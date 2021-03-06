package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/helpers"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"
)

func broadvidVideosThemesViewallCtrl(c *gin.Context) {
	helpers.FindAll(&structs.VideoThemes{}, c)
}

func broadvidVideosThemesViewCtrl(c *gin.Context) {
	helpers.FindOne(&structs.VideoTheme{}, c)
}

func broadvidVideosThemesSaveCtrl(c *gin.Context) {
	uid, err := helpers.GetUserID(c)

	if err != nil {
		c.JSON(500, err)
		return
	}

	var ad = &structs.VideoTheme{}

	err = helpers.SaveEntity(ad, c, "Theme has been saved", ad.Name+" has been saved", "success", 2, uid)

	if err != nil {
		c.JSON(500, err)
		return
	}

}
