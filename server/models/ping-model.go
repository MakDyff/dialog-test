package models

import "time"

type PingModel struct {
	Url        string
	CreatedOn  time.Time
	PingTime   time.Duration
	Status     string
	StatusCode int
	Err        error
}
