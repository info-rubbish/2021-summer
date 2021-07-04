package api

import (
	"main/src/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Error   bool                   `json:"error"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type Req struct {
	Token string      `json:"token"`
	Data  interface{} `json:"data"`
}

type ModelToken struct {
	Token   string        `json:"token"`
	Created time.Time     `json:"created"`
	TTL     time.Duration `json:"ttl"`
}

type ModelUser struct {
	Created time.Time `json:"created"`
	ID      string    `json:"id"`
	Name    string    `json:"name"`
}

// type TokenData struct {
// 	Name string `json:"name"`
// }

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
