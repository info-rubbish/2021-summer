package config

import "time"

const (
	RandomIDLength int           = 64
	TokenTTL       time.Duration = time.Hour * 23
	TokenGCT       time.Duration = time.Hour * 29
	TokenBytes     uint          = 64
)
