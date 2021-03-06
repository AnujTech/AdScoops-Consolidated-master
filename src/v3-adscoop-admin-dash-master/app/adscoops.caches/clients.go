package adscoopsCaches

import (
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"

	log "github.com/Sirupsen/logrus"
)

func UpdateClientCompactData() {
	var clients structs.Clients

	if err := clients.FindAll(); err != nil {
		log.Errorf("Cannot grab clients: %s", err)
		return
	}

	for _, c := range clients {
		if err := c.CompactOldData(); err != nil {
			log.Errorf("Cannot compact old data: %s", err)
			continue
		}
	}
}
