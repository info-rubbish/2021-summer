package tokens

import "main/src/config"

var TokenStore *Store

func init() {
	TokenStore = NewStore(config.TokenTTL, config.TokenGCT, config.TokenBytes)
}
