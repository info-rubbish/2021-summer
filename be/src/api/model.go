package api

import (
	"errors"
	"main/src/config"
	"main/src/tokens"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var c int = http.StatusInternalServerError
	switch errors.Unwrap(e) {
	case tokens.ErrEmpty:
		c = http.StatusBadRequest
	case tokens.ErrNotFind:
		c = http.StatusUnauthorized
	case tokens.ErrNotFind:
	case gorm.ErrRecordNotFound:
		c = http.StatusNotFound
	}
	s.JSON(c, &Resp{
		Error:   true,
		Message: e.Error(),
	})
}
