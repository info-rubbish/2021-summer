package api

import (
	"main/src/database"
	"main/src/tokens"
	"net/http"

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
	token, err := tokens.TokenStore.NewToken(user)
	if err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Error:   false,
		Message: "login success",
		Data:    &ModelToken{Token: token},
	})
}

func DeleteToken(s *gin.Context) {
	req := &Req{}
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
		Error:   false,
		Message: "logout success",
		Data:    nil,
	})

}
func PutToken(s *gin.Context) {
	req := &Req{}
	s.BindJSON(req)
	token, err := tokens.TokenStore.ReNewToken(req.Token)
	if err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Error:   false,
		Message: "renew success",
		Data:    &ModelToken{Token: token},
	})
}
