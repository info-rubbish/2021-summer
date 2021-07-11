package api

import (
	"fmt"
	"main/src/database"
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

func Err2Restful(s *gin.Context, e error) {
	fmt.Printf("%T\n", e)
	var c int = http.StatusInternalServerError
	switch e {
	case tokens.ErrEmpty:
		c = http.StatusBadRequest
	case tokens.ErrNotFind:
	case database.ErrPasswordNotMatch:
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
