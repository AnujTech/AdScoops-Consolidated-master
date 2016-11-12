package adscoopsCaches

import (
	"encoding/json"
	"fmt"

	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/structs"
)

func LoadAdscoopHostById(id string) (host structs.Host) {
	AdscoopsDB.Find(&host, id)

	b, err := json.Marshal(host)

	if err == nil {
		rp := RedisPool.Get()
		defer rp.Close()

		rp.Do("SET", fmt.Sprintf(HOSTS_BY_ID_KEY, id), b)
	}

	return
}

func LoadAdscoopHostByHost(id string) (host structs.Host) {
	AdscoopsDB.Where("host = ?", id).Find(&host)

	b, err := json.Marshal(host)

	if err == nil {
		rp := RedisPool.Get()
		defer rp.Close()

		rp.Do("SET", fmt.Sprintf(HOSTS_BY_HOST_KEY, id), b)
	}

	return
}
