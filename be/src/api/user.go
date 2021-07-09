package api

import (
	"main/src/cache"
	"main/src/database"
	"main/src/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserReq struct {
	Token string `json:"token"`
}

func GetUser(s *gin.Context) {
	req := &GetUserReq{}
	s.BindJSON(req)
	id, err := tokens.TokenStore.GetToken(req.Token)
	if err != nil {
		Err2Restful(s, err)
		return
	}
	var user database.User
	cache.CacheStore.Load(&user, id, database.UserInfo)
	s.JSON(http.StatusOK, &Resp{
		Message: "get info success",
		Data: map[string]interface{}{
			"user": &ModelUser{
				Created: user.Created,
				ID:      user.ID,
				Name:    user.Name,
			},
		},
	})
}

type PutUserReq struct {
	Token    string `json:"token"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PutUser(s *gin.Context) {
	req := &PutUserReq{}
	s.BindJSON(req)
	id, err := tokens.TokenStore.GetToken(req.Token)
	if err != nil {
		Err2Restful(s, err)
		return
	}
	cache.CacheStore.Del(id)
	if err := database.ChangeUserInfo(id, database.UserConfig{
		Name:     req.Name,
		Password: req.Password,
	}); err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Message: "change user info success",
	})
}
