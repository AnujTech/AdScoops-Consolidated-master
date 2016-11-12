package helpers

import (
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/sockets"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"
)

func pushUpdate(title string, message string, messageType string, productType uint, userID uint) {

	var update structs.Updates
	update.Title = title
	update.Message = message
	update.UserID = userID
	update.Product = productType
	update.Save()

	var m structs.RealTimeUpdate
	m.Event = "message"
	m.Data.Type = messageType
	m.Data.Title = title
	m.Data.Message = message
	sockets.BroadcastMessage(1, m.JSONify(), "updates")
}
