package entity

import (
	"time"
)

type Link struct {
	ID       int
	Origin   string
	Destiny  string
	Validate time.Time
}
