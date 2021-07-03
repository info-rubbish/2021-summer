package api

import (
	"main/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Req struct {
	Token string      `json:"token"`
	Data  interface{} `json:"data"`
}

type ModelToken struct {
	Token string `json:"token"`
}

type TokenData struct {
	Name string `json:"name"`
}

func Err2Restful(s *gin.Context, e error) {
	if err, ok := e.(*config.HttpErr); ok {
		s.JSON(err.Code, &Resp{
			Error:   true,
			Message: err.Error(),
		})
		return
	}
	s.JSON(http.StatusInternalServerError, &Resp{
		Error:   true,
		Message: e.Error(),
	})
}
