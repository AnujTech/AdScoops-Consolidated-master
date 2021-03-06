package billing

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/swhitley101/AdScoops-Consolidated/src/v3-adscoop-admin-dash-master/app/configSettting"
)

//const SECRET_KEY = "sk_live_IS0W5w62ttenjEpuQozj4VDE"

func Charge(amount uint, token string, name string) (chargeID string, err error) {
	stripe.Key = configSettting.StripeSecretKey
	chargeParams := &stripe.ChargeParams{
		Amount:   uint64(amount * 100),
		Currency: "usd",
		Customer: token,
		Desc:     "Charge for " + name,
	}

	ch, err := charge.New(chargeParams)

	if err != nil {
		return
	}

	chargeID = ch.ID
	return
}
