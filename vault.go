package paypal

import "fmt"

// https://developer.paypal.com/webapps/developer/docs/api/#vault

// StoreCreditCard Vault a captured payment.
func (c *Client) StoreCreditCard(captureID string, cc *CreditCard) (*Vault, error) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/vault/credit-cards", c.APIBase), c)
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
