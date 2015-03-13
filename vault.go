package paypal

import "fmt"

// https://developer.paypal.com/webapps/developer/docs/api/#vault

// StoreCreditCard Vault a captured payment.
func (c *Client) StoreCreditCard(cc *CreditCard) (*Vault, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/vault/credit-cards", c.APIBase), struct {
		PayerID     string `json:"payer_id"`
		Number      string `json:"number"`
		Type        string `json:"type"`
		ExpireMonth string `json:"expire_month"`
		ExpireYear  string `json:"expire_year"`
		CVV2        string `json:"cvv2"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
	}{cc.PayerID, cc.Number, cc.Type, cc.ExpireMonth, cc.ExpireYear, cc.CVV2, cc.FirstName, cc.Lastname})

	if err != nil {
		return nil, err
	}

	v := &Vault{}

	err = c.SendWithAuth(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
