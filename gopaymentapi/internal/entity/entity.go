package entity

import "errors"

type OrderRequest struct {
	OrderID  string  `json:"order_id"`
	CardHash string  `json:"card_hash"`
	Total    float64 `json:"total"`
}

func NewOrderRequest(orderID string, cardHash string, total float64) *OrderRequest {
	return &OrderRequest{
		OrderID:  orderID,
		CardHash: cardHash,
		Total:    total,
	}
}

func (orderRequest *OrderRequest) Validate() error {
	if orderRequest.OrderID == "" {
		return errors.New("order_id is required")
	}

	if orderRequest.CardHash == "" {
		return errors.New("card_hash is required")
	}

	if orderRequest.Total <= 0 {
		return errors.New("total must be greater than 0")
	}

	return nil
}

func (orderRequest *OrderRequest) Process() (*OrderResponse, error) {
	if err := orderRequest.Validate(); err != nil {
		return nil, err
	}

	orderResponse := NewOrderResponse(orderRequest.OrderID, "failed")

	if orderRequest.Total < 100.00 {
		orderResponse.Status = "paid"
	}

	return orderResponse, nil
}

type OrderResponse struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

func NewOrderResponse(orderID string, status string) *OrderResponse {
	return &OrderResponse{
		OrderID: orderID,
		Status:  status,
	}
}
