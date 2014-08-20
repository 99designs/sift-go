package sift

import "fmt"

type ApiError struct {
	Status       int
	ErrorMessage string `json:"error_message"`
	Time         int64
	Request      string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%d - %s", e.Status, e.ErrorMessage)
}

type Address struct {
	Name        string `json:"$name,omitempty"`
	Phone       string `json:"$phone,omitempty"`
	Address     string `json:"$address_1,omitempty"`
	AddressCplt string `json:"$address_2,omitempty"`
	City        string `json:"$city,omitempty"`
	Region      string `json:"$region,omitempty"`
	Country     string `json:"$country,omitempty"`
	Zipcode     string `json:"$zipcode,omitempty"`
}

type PaymentMethod struct {
	Type               string `json:"$payment_type,omitempty"`
	Gateway            string `json:"$payment_gateway,omitempty"`
	CardBin            string `json:"$card_bin,omitempty"`
	CardLast           string `json:"$card_last4,omitempty"`
	AvsResult          string `json:"$avs_result_code,omitempty"`
	CvvResult          string `json:"$cvv_result_code,omitempty"`
	VerificationStatus string `json:"$verification_status,omitempty"`
	RoutingNumber      string `json:"$routing_number,omitempty"`
}

type Item struct {
	Id           string   `json:"$item_id,omitempty"`
	Title        string   `json:"$product_title,omitempty"`
	Isbn         string   `json:"$isbn,omitempty"`
	Price        int64    `json:"$price,omitempty"`
	CurrencyCode string   `json:"$currency_code,omitempty"`
	Upc          string   `json:"$upc,omitempty"`
	Sku          string   `json:"$sku,omitempty"`
	Brand        string   `json:"$brand,omitempty"`
	Manufacturer string   `json:"$manufacturer,omitempty"`
	Category     string   `json:"$category,omitempty"`
	Color        string   `json:"$color,omitempty"`
	Size         string   `json:"$size,omitempty"`
	Tags         []string `json:"$tags,omitempty"`
	Quantity     int      `json:"$quantity,omitempty"`
}
