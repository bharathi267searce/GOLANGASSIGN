package Datastructures

import (
	"time"
)

type Cart_reference struct {
	Reference_id string    `json:"reference_id"`
	Name         string    `json:"name"`
	Date         time.Time `json:"date"`
}
