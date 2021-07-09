package api

import (
	"main/src/cache"
	"main/src/config"
	"main/src/database"
	"main/src/tokens"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PostTokenReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PostToken(s *gin.Context) {
	req := &PostTokenReq{}
	s.BindJSON(req)
	if req.Name == "" || req.Password == "" {
		s.JSON(http.StatusUnprocessableEntity, &Resp{
			Error:   true,
			Message: "name & password are required",
		})
		return
	}
	user, err := database.UserLogin(s, req.Name, req.Password)
	if err != nil {
		Err2Restful(s, err)
		return
	}
	cache.CacheStore.Set(user.ID, user)
	token, err := tokens.TokenStore.NewToken(user.ID)
	if err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Message: "login success",
		Data: map[string]interface{}{
			"token": &ModelToken{
				Token:   token,
				Created: time.Now(),
				TTL:     config.TokenTTL,
			},
			"user": &ModelUser{
				Created: user.Created,
				ID:      user.ID,
				Name:    user.Name,
			},
		},
	})
}

type DeleteTokenReq struct {
	Token string `json:"token"`
}

func DeleteToken(s *gin.Context) {
	req := &DeleteTokenReq{}
	s.BindJSON(req)
	if _, err := tokens.TokenStore.GetToken(req.Token); err != nil {
		Err2Restful(s, err)
		return
	}
	if err := tokens.TokenStore.DestroyToken(req.Token); err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Message: "logout success",
		Data:    nil,
	})

}

type PutTokenReq struct {
	Token string `json:"token"`
}

func PutToken(s *gin.Context) {
	req := &PutTokenReq{}
	s.BindJSON(req)
	token, err := tokens.TokenStore.ReNewToken(req.Token)
	if err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Error:   false,
		Message: "renew success",
		Data: map[string]interface{}{
			"token": &ModelToken{
				Token:   token,
				Created: time.Now(),
				TTL:     config.TokenTTL,
			},
		},
	})
}
