package cache

import "main/src/config"

var CacheStore *Wrapper

func init() {
	CacheStore = NewWrapper(New(config.CacheShard2N, config.CacheLength2N))
}
