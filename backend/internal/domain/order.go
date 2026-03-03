package domain

import "time"

type OrderStatus string

const (
	StatusPending  OrderStatus = "PENDING"
	StatusApproved OrderStatus = "APPROVED"
	StatusFailed   OrderStatus = "FAILED"
)

type Order struct {
	ID        string
	Amount    float64
	Status    OrderStatus
	CreatedAt time.Time
}