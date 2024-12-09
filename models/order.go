package models

// OrderReq represents the incoming order request
type OrderReq struct {
	CouponCode string     `json:"couponCode,omitempty"`
	Items      []struct { // Inline struct definition for the items
		ProductID string `json:"productId" validate:"required"`
		Quantity  int    `json:"quantity" validate:"required"`
	} `json:"items" validate:"required,dive"` // Directly use the inline struct here
}

// Order represents the complete order response
type Order struct {
	ID    string     `json:"id"`
	Items []struct { // Inline struct for items here as well
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	} `json:"items"`
	Products []Product `json:"products"`
}
