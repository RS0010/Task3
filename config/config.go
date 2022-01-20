package config

import "time"

const (
	AuthTokenExpireDuration    = time.Hour * 2
	RefreshTokenExpireDuration = time.Hour * 24
)

var (
	TokenSecret = []byte("West2Online")
)
