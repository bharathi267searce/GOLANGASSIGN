package Datastructures

import "time"

type CartReference struct {
	Referenceid string    `json:"reference_id"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
}
