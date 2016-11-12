package adminControllers

import (
	"bitbucket.org/broadscaler/broadscaler/app/helpers"
	"bitbucket.org/broadscaler/broadscaler/app/structs"
	"github.com/gin-gonic/gin"
)

func broadvidVideosEmbedsViewallCtrl(c *gin.Context) {
	helpers.FindAll(&structs.VideoEmbeds{}, c)
}

func broadvidVideosEmbedsViewCtrl(c *gin.Context) {
	helpers.FindOne(&structs.VideoEmbed{}, c)
}

func broadvidVideosEmbedsSaveCtrl(c *gin.Context) {
	uid, err := helpers.GetUserID(c)

	if err != nil {
		c.JSON(500, err)
		return
	}

	var ad = &structs.VideoEmbed{}

	err = helpers.SaveEntity(ad, c, "Embed has been saved", ad.Label+" has been saved", "success", 2, uid)

	if err != nil {
		c.JSON(500, err)
		return
	}

}
