package billing

import (
	"configSettting"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

var SECRET_KEY = configSettting.StripeSecretKey

func Charge(amount uint, token string, name string) (chargeID string, err error) {
	stripe.Key = SECRET_KEY
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
