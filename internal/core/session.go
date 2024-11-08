package core

import "time"

type Session struct {
	Refresh_token string
	ExpiresAt     time.Time
}