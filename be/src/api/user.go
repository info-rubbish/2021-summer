package api

import (
	"main/src/cache"
	"main/src/database"
	"main/src/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(s *gin.Context) {
	id, err := tokens.TokenStore.GetToken(s.GetHeader("Authorization"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	user := &database.User{}
	if err := cache.CacheStore.Load(user, id, database.UserInfoByID); err != nil {
		Err2Restful(s, err)
		return
	}
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

type PatchUserReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PatchUser(s *gin.Context) {
	id, err := tokens.TokenStore.GetToken(s.GetHeader("Authorization"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	req := &PatchUserReq{}
	if err := s.ShouldBindJSON(req); err != nil {
		s.JSON(http.StatusBadRequest, &Resp{
			Error:   true,
			Message: err.Error(),
		})
		return
	}
	if req.Name == "" && req.Password == "" {
		s.JSON(http.StatusUnprocessableEntity, &Resp{
			Error:   true,
			Message: "need name or password",
		})
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

type PostUserReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PostUser(s *gin.Context) {
	req := &PostUserReq{}
	if err := s.ShouldBindJSON(req); err != nil {
		s.JSON(http.StatusBadRequest, &Resp{
			Error:   true,
			Message: err.Error(),
		})
		return
	}
	if req.Name == "" || req.Password == "" {
		s.JSON(http.StatusUnprocessableEntity, &Resp{
			Error:   true,
			Message: "name & password are required",
		})
		return
	}
	if err := database.CreateUser(req.Name, req.Password); err != nil {
		Err2Restful(s, err)
		return
	}
	user := &database.User{}
	if err := cache.CacheStore.Load(user, req.Name, database.UserInfoByName); err != nil {
		Err2Restful(s, err)
		return
	}

	s.JSON(http.StatusOK, &Resp{
		Message: "create user success",
		Data: map[string]interface{}{
			"user": &ModelUser{
				Created: user.Created,
				ID:      user.ID,
				Name:    user.Name,
			},
		},
	})
}

func DeleteUser(s *gin.Context) {
	id, err := tokens.TokenStore.GetToken(s.GetHeader("Authorization"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	if err := database.DeleteUser(id); err != nil {
		Err2Restful(s, err)
		return
	}
	cache.CacheStore.Del(id)
	s.JSON(http.StatusOK, &Resp{
		Message: "delete account success",
	})
}
