package apiControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/helpers"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"
)

func campaignSchedulesViewallCtrl(c *gin.Context) {
	helpers.FindAllSchedules(&structs.CampaignSchedules{}, c)
}

func campaignSchedulesViewCtrl(c *gin.Context) {
	helpers.FindOne(&structs.CampaignSchedule{}, c)
}

func campaignSchedulesSaveCtrl(c *gin.Context) {

	uid, err := helpers.GetUserID(c)

	if err != nil {
		c.JSON(500, err)
		return
	}

	var campaignSchedule = &structs.CampaignSchedule{}

	err = helpers.SaveEntity(campaignSchedule, c, "Campaign Schedule has been saved", campaignSchedule.ScheduleLabel+" has been saved", "success", 1, uid)

	if err != nil {
		c.JSON(500, err)
		return
	}
}
