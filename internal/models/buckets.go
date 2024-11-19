package models

import "time"

type Bucket struct {
	ID      uint64
	Name    string
	Balance Amount
	Created time.Time
}
