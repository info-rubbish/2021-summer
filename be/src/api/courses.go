package api

import (
	"main/src/database"
	"main/src/tokens"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func toModelCourses(s *[]*database.Course) *[]*ModelCourse {
	r := make([]*ModelCourse, len(*s))
	for i, v := range *s {
		r[i] = &ModelCourse{
			Created:     v.Created,
			ID:          v.ID,
			Author:      v.Author,
			Title:       v.Title,
			Description: v.Description,
		}
	}
	return &r
}

func GetCoursesByUser(s *gin.Context) {
	_, err := tokens.TokenStore.GetToken(s.GetHeader("Authorization"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	offset, _ := strconv.Atoi(s.Query("offset"))
	if offset < 0 {
		offset = 0
	}
	courses, err := database.GetUserCourses(s.Param("id"), offset, s.Query("order"))
	if err != nil {
		Err2Restful(s, err)
		return
	}

	modelCourses := toModelCourses(courses)
	s.JSON(http.StatusOK, &Resp{
		Message: "get info success",
		Data: map[string]interface{}{
			"courses": modelCourses,
		},
	})
}

func GetCoursesBySearch(s *gin.Context) {
	_, err := tokens.TokenStore.GetToken(s.GetHeader("Authorization"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	offset, _ := strconv.Atoi(s.Query("offset"))
	if offset < 0 {
		offset = 0
	}
	courses, err := database.SearchCourses(s.Param("text"), offset, s.Query("order"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	modelCourses := toModelCourses(courses)
	s.JSON(http.StatusOK, &Resp{
		Message: "get info success",
		Data: map[string]interface{}{
			"courses": modelCourses,
		},
	})
}
