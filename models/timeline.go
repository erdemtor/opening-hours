package models

import "time"

type Timeline []OpenDuration

type OpenDuration struct {
	Open  time.Time
	Close time.Time
}
