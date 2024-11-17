package models

import "time"

type Bucket struct {
	ID      uint64
	Name    string
	Balance string // todo change to some other type
	Created time.Time
}
