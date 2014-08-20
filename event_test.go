package sift

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	coe := NewCreateOrderEvent("api_key", "user_id")

	coe.SessionId = "SESSION"
	coe.UserEmail = "mail@example.com"

	payment := &PaymentMethod{
		Type:     "$credit_card",
		Gateway:  "$braintree",
		CardBin:  "542486",
		CardLast: "4444",
	}

	coe.PaymentMethods = []*PaymentMethod{payment}

	coe.CustomFields["testing"] = "value"

	b, err := json.Marshal(coe)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s", string(b))
}
