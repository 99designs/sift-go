package sift

type CustomFields map[string]interface{}

type TypedEvent interface {
	GetType() string
}

type Event struct {
	// Required fields
	Type   string `json:"$type"`
	ApiKey string `json:"$api_key"`
	UserId string `json:"$user_id"`
	// Optional but possible on each event
	Ip   string `json:"$ip,omitempty"`
	Time int64  `json:"$time,omitempty"`
	CustomFields
}

type CreateOrderEvent struct {
	Event
	SessionId       string           `json:"$session_id,omitempty"`
	OrderId         string           `json:"$order_id,omitempty"`
	UserEmail       string           `json:"$user_email,omitempty"`
	Amount          int64            `json:"$amount,omitempty"`
	CurrencyCode    string           `json:"$currency_code,omitempty"`
	BillingAddress  *Address         `json:"$billing_address,omitempty"`
	PaymentMethods  []*PaymentMethod `json:"$payment_methods,omitempty"`
	ShippingAddress *Address         `json:"$shipping_address,omitempty"`
	Expedited       bool             `json:"$expedited_shipping,omitempty"`
	Items           []*Item          `json:"$items,omitempty"`
	Seller          string           `json:"$seller_user_id,omitempty"`
}

func NewCreateOrderEvent(api_key, user_id string) *CreateOrderEvent {
	coe := &CreateOrderEvent{}

	coe.Type = coe.GetType()
	coe.ApiKey = api_key
	coe.UserId = user_id
	coe.CustomFields = make(map[string]interface{})

	return coe
}

func (coe *CreateOrderEvent) GetType() string {
	return "$create_order"
}

func (coe *CreateOrderEvent) MarshalJSON() ([]byte, error) {
	return marshalEvent(coe, coe.CustomFields, coe.Event)
}

type TransactionEvent struct {
	Event
	TransactionId     string         `json:"$transaction_id,omitempty"`
	TransactionType   string         `json:"$transaction_type,omitempty"`
	TransactionStatus string         `json:"$transaction_status,omitempty"`
	UserEmail         string         `json:"$user_email,omitempty"`
	Amount            int64          `json:"$amount,omitempty"`
	CurrencyCode      string         `json:"$currency_code,omitempty"`
	SessionId         string         `json:"$session_id,omitempty"`
	OrderId           string         `json:"$order_id,omitempty"`
	BillingAddress    *Address       `json:"$billing_address,omitempty"`
	PaymentMethod     *PaymentMethod `json:"$payment_method,omitempty"`
	ShippingAddress   *Address       `json:"$shipping_address,omitempty"`
	Seller            string         `json:"$seller_user_id,omitempty"`
}

func NewTransactionEvent(api_key, user_id string) *TransactionEvent {
	te := &TransactionEvent{}

	te.Type = te.GetType()
	te.ApiKey = api_key
	te.UserId = user_id
	te.CustomFields = make(map[string]interface{})

	return te
}

func (te *TransactionEvent) GetType() string {
	return "$transaction"
}

func (te *TransactionEvent) MarshalJSON() ([]byte, error) {
	return marshalEvent(te, te.CustomFields, te.Event)
}

type CreateAccountEvent struct {
	Event
	SessionId      string           `json:"$session_id,omitempty"`
	UserEmail      string           `json:"$user_email,omitempty"`
	Name           string           `json:"$name,omitempty"`
	Phone          string           `json:"$phone,omitempty"`
	ReferrerUserId string           `json:"$referrer_user_id,omitempty"`
	PaymentMethods []*PaymentMethod `json:"$payment_methods,omitempty"`
	BillingAddress *Address         `json:"$billing_address,omitempty"`
	SignOnType     string           `json:"$social_sign_on_type,omitempty"`
}

func NewCreateAccountEvent(api_key, user_id string) *CreateAccountEvent {
	cae := &CreateAccountEvent{}

	cae.Type = cae.GetType()
	cae.ApiKey = api_key
	cae.UserId = user_id
	cae.CustomFields = make(map[string]interface{})

	return cae
}

func (cae *CreateAccountEvent) GetType() string {
	return "$create_account"
}

func (cae *CreateAccountEvent) MarshalJSON() ([]byte, error) {
	return marshalEvent(cae, cae.CustomFields, cae.Event)
}
