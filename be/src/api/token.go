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
				TTL:     config.TokenTTL / time.Millisecond,
			},
			"user": &ModelUser{
				Created: user.Created,
				ID:      user.ID,
				Name:    user.Name,
			},
		},
	})
}

func DeleteToken(s *gin.Context) {
	if err := tokens.TokenStore.DestroyToken(s.GetHeader("Authorization")); err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Message: "logout success",
		Data:    nil,
	})

}

func PutToken(s *gin.Context) {
	token, err := tokens.TokenStore.ReNewToken(s.GetHeader("Authorization"))
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
