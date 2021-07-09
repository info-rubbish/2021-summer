package config

import "time"

const (
	RandomIDLength int           = 16
	TokenTTL       time.Duration = time.Hour * 23
	TokenGCT       time.Duration = time.Hour * 29
	TokenBytes     uint          = 64
	CacheShard2N   uint64        = 5
	CacheLength2N  uint64        = 10
)
