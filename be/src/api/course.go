package api

import (
	"main/src/cache"
	"main/src/database"
	"main/src/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourse(s *gin.Context) {
	_, err := tokens.TokenStore.GetToken(s.GetHeader("Authorization"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	course := &database.Course{}
	if err := cache.CacheStore.Load(course, s.Param("id"), database.GetCourse); err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Message: "Get course successfully",
		Data: map[string]interface{}{
			"course": &ModelCourse{
				Created:     course.Created,
				ID:          course.ID,
				Author:      course.Author,
				Title:       course.Title,
				Description: course.Description,
			},
			"content": course.Content,
		},
	})
}

type PostCourseReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

func PostCourse(s *gin.Context) {
	id, err := tokens.TokenStore.GetToken(s.GetHeader("Authorization"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	req := &PostCourseReq{}
	if err := s.ShouldBindJSON(req); err != nil {
		s.JSON(http.StatusBadRequest, &Resp{
			Error:   true,
			Message: err.Error(),
		})
		return
	}
	courseID, err := database.CreateCourse(&database.CourseConfig{
		Author:      id,
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
	})
	if err != nil {
		Err2Restful(s, err)
		return
	}
	course := &database.Course{}
	if err := cache.CacheStore.Load(course, courseID, database.GetCourse); err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Message: "Created course successfully",
		Data: map[string]interface{}{
			"course": &ModelCourse{
				Created:     course.Created,
				ID:          course.ID,
				Author:      course.Author,
				Title:       course.Title,
				Description: course.Description,
			},
		},
	})
}

type PatchCourseReq struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required_without_all=Description Content"`
	Description string `json:"description" binding:"required_without_all=Title Content"`
	Content     string `json:"content" binding:"required_without_all=Title Description"`
}

func PatchCourse(s *gin.Context) {
	id, err := tokens.TokenStore.GetToken(s.GetHeader("Authorization"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	req := &PatchCourseReq{}
	if err := s.ShouldBindJSON(req); err != nil {
		s.JSON(http.StatusBadRequest, &Resp{
			Error:   true,
			Message: err.Error(),
		})
		return
	}
	if err := database.ChangeCourse(req.ID, id, &database.CourseConfig{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
	}); err != nil {
		Err2Restful(s, err)
		return
	}
	cache.CacheStore.Del(req.ID)
	s.JSON(http.StatusOK, &Resp{
		Message: "Change success",
	})

}

type DeleteCourseReq struct {
	ID string `json:"id" binding:"required"`
}

func DeleteCourse(s *gin.Context) {
	id, err := tokens.TokenStore.GetToken(s.GetHeader("Authorization"))
	if err != nil {
		Err2Restful(s, err)
		return
	}
	req := &DeleteCourseReq{}
	if err := s.ShouldBindJSON(req); err != nil {
		s.JSON(http.StatusBadRequest, &Resp{
			Error:   true,
			Message: err.Error(),
		})
		return
	}
	if err := database.DeleteCourse(req.ID, id); err != nil {
		Err2Restful(s, err)
		return
	}
	s.JSON(http.StatusOK, &Resp{
		Message: "delete successfully",
	})
}
